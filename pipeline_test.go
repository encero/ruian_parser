package main

import (
	"bytes"
	"context"
	"embed"
	_ "embed"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"testing"
	"time"

	is_ "github.com/matryer/is"
)

const defaultWaitTime = time.Millisecond * 10

//go:embed testdata/*
var testData embed.FS

func TestDownloader(t *testing.T) {
	is := is_.New(t)
	links := []string{
		"https://vdp.cuzk.cz/vymenny_format/soucasna/20211031_OB_500011_UKSH.xml.zip",
		"https://vdp.cuzk.cz/vymenny_format/soucasna/20211031_OB_500011_UKSH.xml.zip",
	}

	ctx, cancel := context.WithTimeout(context.Background(), defaultWaitTime)
	defer cancel()

	called := 0
	doer := httpDoer(func(req *http.Request) (*http.Response, error) {
		is.Equal(req.URL.String(), links[0]) // check request URL
		called++

		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(bytes.NewBufferString("data")),
			Header: map[string][]string{
				"Content-Type": {"application/zip"},
			},
		}, nil
	})

	downCh, downloader := NewDownloader(doer, links)

	go func() {
		err := downloader(ctx)
		is.NoErr(err)
	}()

	for i := range links {
		select {
		case file := <-downCh:
			is.True(called >= i+1) // check if doer was called

			content, err := io.ReadAll(file.Content)
			is.NoErr(err) // read file content

			is.Equal(file.FileName, "20211031_OB_500011_UKSH.xml.zip") // DownloadedFile.FileName
			is.Equal(file.ContentType, "application/zip")              // DownloadedFile.ContentType
			is.Equal(string(content), "data")                          // DownloadedFile.Content
		case <-ctx.Done():
			is.Fail() // no file received on channel
		}
	}

	select {
	case _, ok := <-downCh:
		is.True(!ok) // channel should be closed
	case <-ctx.Done():
		is.Fail() // channel not closed in time
	}
}

type TestCloserSpy struct {
	io.Reader
	closed bool
}

func (c *TestCloserSpy) Close() error {
	c.closed = true

	return nil
}

func TestNewFileCacher(t *testing.T) {
	is := is_.New(t)

	ctx, cancel := context.WithTimeout(context.Background(), defaultWaitTime)
	defer cancel()

	wait := make(chan struct{})
	defer close(wait)

	inFile := &TestCloserSpy{
		Reader: bytes.NewBufferString("the data"),
	}

	in := make(chan DownloadedFile, 1)
	in <- DownloadedFile{
		ContentType: "application/zip",
		FileName:    "",
		Content:     inFile,
	}

	close(in)

	out, filecacher := NewFileCacher(in)

	go func() {
		err := filecacher(ctx)
		is.NoErr(err)

		wait <- struct{}{}
	}()

	select {
	case file, ok := <-out:
		is.True(ok) // channel should be open

		defer file.Content.Close()
		is.Equal(file.ContentType, "application/zip")

		_, ok = file.Content.(io.ReaderAt)
		is.True(ok) // cached file must be io.ReaderAt

		data, err := io.ReadAll(file.Content)
		is.NoErr(err)

		is.True(inFile.closed) // input file should be closed

		is.Equal(string(data), "the data")
	case <-ctx.Done():
		is.Fail() // no file received on channel
	}

	waitForChannel(ctx, is, wait)
}

func TestDecompressor(t *testing.T) {
	is := is_.New(t)

	ctx, cancel := context.WithTimeout(context.Background(), defaultWaitTime)
	defer cancel()

	testFile, err := testData.Open("testdata/zip_file.zip")
	is.NoErr(err)

	in := make(chan CachedFile, 1)
	in <- CachedFile{
		ContentType: "application/zip",
		Content:     readAtWrapper{File: testFile},
	}

	close(in)

	wait := make(chan struct{})
	defer close(wait)

	out, decompressor := NewDecompressor(in)

	go func() {
		err := decompressor(ctx)
		is.NoErr(err)

		wait <- struct{}{}
	}()

	select {
	case file, ok := <-out:
		is.True(ok)                                   // channel should be open
		is.Equal(file.ContentType, "application/xml") // DownloadedFile.ContentType

		content, err := io.ReadAll(file.Content)
		is.NoErr(err)

		is.Equal(string(content), "the zip data\n") // DownloadedFile.Content
	case <-ctx.Done():
		is.Fail() // no file received on channel
	}

	select {
	case _, ok := <-out:
		is.True(!ok) // channel should be closed
	case <-ctx.Done():
		is.Fail() // channel not closed in time
	}

	waitForChannel(ctx, is, wait)

	select {
	case _, ok := <-out:
		is.True(!ok) // channel should be closed
	case <-ctx.Done():
		is.Fail() // channel not closed in time
	}
}

//go:embed testdata/sample.xml
var xmlSample []byte

func TestMapper(t *testing.T) {
	is := is_.New(t)

	ctx, cancel := context.WithTimeout(context.Background(), defaultWaitTime)
	defer cancel()

	wait := make(chan struct{})
	defer close(wait)

	inFile := &TestCloserSpy{
		Reader: bytes.NewBuffer(xmlSample),
	}
	in := make(chan FileContent, 1)
	in <- FileContent{
		ContentType: "application/xml",
		Content:     inFile,
	}
	close(in)

	out, mapper := NewMapper(in)

	go func() {
		err := mapper(ctx)
		is.NoErr(err)
		wait <- struct{}{}
	}()

	items := make([]interface{}, 0)

loop:
	for {
		select {
		case item, ok := <-out:
			if !ok {
				break loop
			}

			items = append(items, item)
		case <-ctx.Done():
			is.Fail() // channel not closed in time
		}
	}

	expected := []interface{}{
		sampleAddressPlace, sampleCity, sampleStreet,
	}

	is.Equal(len(items), len(expected)) // unexpected number of items

	for i, exp := range expected {
		is.Equal(items[i], exp)
	}

	waitForChannel(ctx, is, wait)
}

func waitForChannel(ctx context.Context, is *is_.I, wait chan struct{}) {
	select {
	case <-wait:
	case <-ctx.Done():
		is.Fail() // Context exceeded when waiting for channel
	}
}

type readAtWrapper struct {
	fs.File
}

func (f readAtWrapper) ReadAt(p []byte, off int64) (n int, err error) {
	seeker, ok := f.File.(io.Seeker)
	if !ok {
		return 0, fmt.Errorf("not seeker")
	}

	_, err = seeker.Seek(off, io.SeekStart)
	if err != nil {
		return 0, err
	}

	return f.File.Read(p)
}
