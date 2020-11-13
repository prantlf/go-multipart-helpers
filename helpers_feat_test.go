package helpers_test

import (
	"bytes"
	"mime/multipart"
	"testing"

	helpers "github.com/prantlf/go-multipart-helpers"
)

func TestComposer_WriteFile_missing(t *testing.T) {
	message := &bytes.Buffer{}
	writer := multipart.NewWriter(message)
	if err := helpers.WriteFile(writer, "file", "missing.txt"); err == nil {
		t.Error("helpers: invalid file added")
	}
}
