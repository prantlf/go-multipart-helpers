// Package helpers helps writing file parts to MIME multipart messages while
// preserving the original content type inferred from the file extension.
//
// CreateFormFile from multipart Writer sets the content type always to
// `application/octet-stream`. If you need to preserve the content type
// of the file decided by its file extension, thi spackage will help you.
//
// Files can be written using their path by a convenience method:
//
//     message := &bytes.Buffer{}
//     writer := multipart.NewWriter(message)
//     err := helpers.WriteFile(writer, "file", "test.bin")
//     err := writer.Close()
package helpers

import (
	"fmt"
	"io"
	"mime"
	"mime/multipart"
	"net/textproto"
	"os"
	"path/filepath"
	"strings"
)

// CreateFilePart creates a new form-data header with the provided field
// name, file name and its content type inferred from its file extension.
// It calls CreatePart to create a new multipart section with the provided
// header. The content of the file should be written to the returned Writer.
func CreateFilePart(writer *multipart.Writer, fieldName, fileName string) (io.Writer, error) {
	header := make(textproto.MIMEHeader)
	header.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`,
		escapeQuotes(fieldName), escapeQuotes(fileName)))
	contentType := mime.TypeByExtension(filepath.Ext(fileName))
	if contentType == "" {
		contentType = "application/octet-stream"
	}
	header.Set("Content-Type", contentType)
	return writer.CreatePart(header)
}

// WriteFile calls CreateFilePart and then writes the file content
// to the part writer.
func WriteFile(writer *multipart.Writer, fieldName, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	return WriteFileReader(writer, fieldName, filepath.Base(filePath), file)
}

// WriteFileReader calls CreateFilePart and then writes the reader content
// to the part writer.
func WriteFileReader(writer *multipart.Writer, fieldName, fileName string, reader io.Reader) error {
	part, err := CreateFilePart(writer, fieldName, filepath.Base(fileName))
	if err != nil {
		return err
	}
	_, err = io.Copy(part, reader)
	return err
}

var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

func escapeQuotes(value string) string {
	return quoteEscaper.Replace(value)
}
