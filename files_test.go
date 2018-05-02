package main

import (
	"github.com/matryer/is"
	"io"
	"testing"
)

func TestReadDir(t *testing.T) {
	var is = is.New(t)

	files, err := ReadDir("./*.go")
	is.NoErr(err) // ReadDir should not return an error

	is.True(len(files) > 0) // ReadDir should return at least one file
}

func TestDecodeGzipFile(t *testing.T) {
	var is = is.New(t)

	files, err := ReadDir("./testdata/sample.gz")
	is.NoErr(err) // ReadDir should not return an error

	is.True(len(files) == 1) // ReadDir should return exactly one file

	file := files[0]
	defer file.Close()

	content, err := io.ReadAll(file)
	is.NoErr(err) // file should be readable

	is.True(string(content) == "__sample_file__\n") // file should contain the expected content
}
