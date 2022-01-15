package main

import (
	"archive/zip"
	"context"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
)

// Notes:
// This pipeline concept is usable but has some issues
// - Its difficult to propagate state changes to previous stages ( file closing )
// - Download retries are imposible
// - It was planed to be flexible, in reality the stages are tightly coupled
//
// TODO:
// - Should be simplified by extracting the working bits to some "provider" services
// and write "simple" glue function/functions which will handle all the joining bits
//
// - because of ruian lack of transfer speed, its desirable to keep preloading flow from Downloader
// - downloading, caching to disk, parsing and cleanup can be effectively merged to single component

type HTTPDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

type ZipFile interface {
	fs.File
	io.ReaderAt
}

type DownloadedFile struct {
	ContentType string
	FileName    string
	Content     io.ReadCloser
}

type CachedFile struct {
	ContentType string
	FileName    string
	Content     ZipFile
}

type FileContent struct {
	ContentType string
	FileName    string
	Content     io.ReadCloser
}

type Downloader func(ctx context.Context) error

func NewDownloader(doer HTTPDoer, links []string) (<-chan DownloadedFile, Downloader) {
	ch := make(chan DownloadedFile, 1)

	return ch, func(ctx context.Context) error {
		for i, url := range links {
			log.Printf("Downloading [%d/%d] %s", i, len(links), url)

			file, err := downloadFile(ctx, url, doer)
			if err != nil {
				return err
			}

			ch <- file
		}

		log.Println("Downloader finished")
		close(ch)

		return nil
	}
}

func downloadFile(ctx context.Context, url string, doer HTTPDoer) (DownloadedFile, error) {
	file := DownloadedFile{}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return file, fmt.Errorf("downloader: new request err: %w", err)
	}

	response, err := doer.Do(req) //nolint: bodyclose // closed in different part of pipeline
	if err != nil {
		return file, fmt.Errorf("downloader: do request err: %w", err)
	}

	if response.StatusCode != http.StatusOK {
		return file, fmt.Errorf("downloader: bad status code: %d", response.StatusCode)
	}

	file.Content = response.Body

	file.ContentType = response.Header.Get("Content-Type")
	if file.ContentType == "" {
		return file, fmt.Errorf("downloader: no content type")
	}

	file.FileName = path.Base(url)

	return file, nil
}

func NewFileCacher(in <-chan DownloadedFile) (<-chan CachedFile, func(context.Context) error) {
	out := make(chan CachedFile, 1)

	return out, func(ctx context.Context) error {
		for {
			select {
			case input, more := <-in:
				if !more {
					log.Println("FileCacher finished")
					return nil
				}

				log.Println("Caching to disk", input.FileName)

				outFile, err := cacheFile(input.Content)
				if err != nil {
					return err
				}

				out <- CachedFile{
					ContentType: input.ContentType,
					FileName:    input.FileName,
					Content: selfDeletingFile{
						ZipFile:  outFile,
						fileName: outFile.Name(),
					},
				}

			case <-ctx.Done():
				return ctx.Err()
			}
		}
	}
}

func cacheFile(input io.ReadCloser) (*os.File, error) {
	defer input.Close()

	outFile, err := ioutil.TempFile("", "ruian_*.zip")
	if err != nil {
		return nil, fmt.Errorf("file cacher: create temp file err: %w", err)
	}

	// todo: retry download on failure
	_, err = io.Copy(outFile, input)
	if err != nil {
		return nil, fmt.Errorf("file cacher: cant copy file to disk: %w", err)
	}

	_, err = outFile.Seek(0, io.SeekStart)
	if err != nil {
		return nil, fmt.Errorf("file cacher: cant reset file position: %w", err)
	}

	return outFile, nil
}

type Decompressor func(ctx context.Context) error

func NewDecompressor(in <-chan CachedFile) (<-chan FileContent, Decompressor) {
	ch := make(chan FileContent, 1)

	return ch, func(ctx context.Context) error {
		defer close(ch)

		for {
			select {
			case inputFile, more := <-in:
				if !more {
					log.Println("Decompressor finished")
					return nil
				}

				log.Println("Decompressing", inputFile.FileName)

				stats, err := inputFile.Content.Stat()
				if err != nil {
					return fmt.Errorf("decompressor: input file stat err: %w", err)
				}

				readerAt, ok := inputFile.Content.(io.ReaderAt)
				if !ok {
					return fmt.Errorf("decompressor: input file is not readerAt")
				}

				reader, err := zip.NewReader(readerAt, stats.Size())
				if err != nil {
					return fmt.Errorf("decompressor: zip reader err: %w", err)
				}

				if len(reader.File) != 1 {
					return fmt.Errorf("decompressor: bad zip file, more than 1 embedded file")
				}

				file := reader.File[0]

				contentReader, err := file.Open()
				if err != nil {
					return fmt.Errorf("decompressor: open embedded file: %w", err)
				}

				ch <- FileContent{
					ContentType: "application/xml",
					FileName:    inputFile.FileName,
					Content: &CloserSpy{
						r: contentReader,
						closeFunc: func() {
							inputFile.Content.Close()
						},
					},
				}
			case <-ctx.Done():
				return ctx.Err()
			}
		}
	}
}

type Mapper func(ctx context.Context) error

func NewMapper(in <-chan FileContent) (<-chan interface{}, Mapper) {
	out := make(chan interface{}, 1)

	return out, func(ctx context.Context) error {
		defer close(out)

		for {
			select {
			case inputFile, more := <-in:
				if !more {
					log.Println("Mapper finished")
					return nil
				}

				err := mapFile(inputFile, out)
				if err != nil {
					return err
				}
			case <-ctx.Done():
				return ctx.Err()
			}
		}
	}
}

func mapFile(inputFile FileContent, out chan interface{}) error {
	defer inputFile.Content.Close()

	log.Println("Mapping", inputFile.FileName)
	decoder := NewDecoder(inputFile.Content)

	for {
		data, err := decoder.Next()
		if errors.Is(err, io.EOF) {
			return nil
		}

		if err != nil {
			return fmt.Errorf("mapper: cant decode item: %w", err)
		}

		out <- data
	}
}

type selfDeletingFile struct {
	ZipFile
	fileName string
}

func (s selfDeletingFile) Close() error {
	log.Println(s.fileName)

	defer os.Remove(s.fileName)

	return s.ZipFile.Close()
}

type CloserSpy struct {
	r         io.ReadCloser
	closeFunc func()
}

func (spy CloserSpy) Close() error {
	defer spy.closeFunc()

	return spy.r.Close()
}

func (spy CloserSpy) Read(b []byte) (int, error) {
	return spy.r.Read(b)
}
