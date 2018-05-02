package main

import (
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
)

func ReadDir(pattern string) ([]io.ReadCloser, error) {
	glob, err := filepath.Glob(pattern)
	if err != nil {
		return nil, err
	}

	files := make([]io.ReadCloser, 0, len(glob))

	for _, filename := range glob {
		f, err := openFile(filename)
		if err != nil {
			_ = closeAllFiles(files)
			return nil, err
		}

		files = append(files, f)
	}

	return files, nil
}

func closeAllFiles(files []io.ReadCloser) error {
	for _, file := range files {
		if err := file.Close(); err != nil {
			return err
		}
	}

	return nil
}

func openFile(filename string) (io.ReadCloser, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	return decompressIfNeeded(filename, f)
}

func decompressIfNeeded(filename string, f *os.File) (io.ReadCloser, error) {
	if filepath.Ext(filename) == ".gz" {
		gzipReader, err := gzip.NewReader(f)
		if err != nil {
			return nil, err
		}

		return splitReadCloser{closer: f, reader: gzipReader}, nil
	}

	return f, nil
}

type splitReadCloser struct {
	reader io.Reader
	closer io.Closer
}

func (rc splitReadCloser) Close() error {
	return rc.closer.Close()
}

func (rc splitReadCloser) Read(p []byte) (n int, err error) {
	return rc.reader.Read(p)
}
