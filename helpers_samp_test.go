package helpers_test

import (
	"bytes"
	"fmt"
	"log"
	"mime/multipart"
	"regexp"
	"strings"

	helpers "github.com/prantlf/go-multipart-helpers"
)

func Example() {
	// Create a new buffer for the message content.
	message := &bytes.Buffer{}
	// Create a new multipart message writrer with a random boundary.
	writer := multipart.NewWriter(message)
	// Write a textual field.
	if err := writer.WriteField("comment", "a comment"); err != nil {
		log.Fatal(err)
	}
	// Write a file.
	if err := helpers.WriteFile(writer, "file", "test.txt"); err != nil {
		log.Fatal(err)
	}
	// Finalize rhe message by appending the trailing boundary separatore.
	if err := writer.Close(); err != nil {
		log.Fatal(err)
	}

	// Make a network request with the composed content type and request body.
	printRequest(writer.FormDataContentType(), message)
	// Output:
	// Content-Type: multipart/form-data; boundary=1879bcd06ac39a4d8fa5
	// Content-Length: 383
	//
	// --1879bcd06ac39a4d8fa5
	// Content-Disposition: form-data; name="comment"
	//
	// a comment
	// --1879bcd06ac39a4d8fa5
	// Content-Disposition: form-data; name="file"; filename="test.txt"
	// Content-Type: text/plain; charset=utf-8
	//
	// text file content
	// --1879bcd06ac39a4d8fa5--
}

func ExampleCreateFilePart() {
	message := &bytes.Buffer{}
	writer := multipart.NewWriter(message)

	// Create an empty file part.
	if _, err := helpers.CreateFilePart(writer, "file", "test.txt"); err != nil {
		log.Fatal(err)
	}

	writer.Close()
	printRequest(writer.FormDataContentType(), message)
	// Output:
	// Content-Type: multipart/form-data; boundary=1879bcd06ac39a4d8fa5
	// Content-Length: 241
	//
	// --1879bcd06ac39a4d8fa5
	// Content-Disposition: form-data; name="file"; filename="test.txt"
	// Content-Type: text/plain; charset=utf-8
	//
	// --1879bcd06ac39a4d8fa5--
}

func ExampleWriteFile() {
	message := &bytes.Buffer{}
	writer := multipart.NewWriter(message)

	// Write a file.
	if err := helpers.WriteFile(writer, "file", "test.bin"); err != nil {
		log.Fatal(err)
	}

	writer.Close()
	printRequest(writer.FormDataContentType(), message)
	// Output:
	// Content-Type: multipart/form-data; boundary=1879bcd06ac39a4d8fa5
	// Content-Length: 259
	//
	// --1879bcd06ac39a4d8fa5
	// Content-Disposition: form-data; name="file"; filename="test.bin"
	// Content-Type: application/octet-stream
	//
	// binary file content
	// --1879bcd06ac39a4d8fa5--
}

func ExampleWriteFileReader() {
	message := &bytes.Buffer{}
	writer := multipart.NewWriter(message)

	// Write a file.
	if err := helpers.WriteFileReader(writer, "file", "test.txt",
		strings.NewReader("test")); err != nil {
		log.Fatal(err)
	}

	writer.Close()
	printRequest(writer.FormDataContentType(), message)
	// Output:
	// Content-Type: multipart/form-data; boundary=1879bcd06ac39a4d8fa5
	// Content-Length: 245
	//
	// --1879bcd06ac39a4d8fa5
	// Content-Disposition: form-data; name="file"; filename="test.txt"
	// Content-Type: text/plain; charset=utf-8
	//
	// test
	// --1879bcd06ac39a4d8fa5--
}

const commonBoundary = "1879bcd06ac39a4d8fa5"

var contentTypeBoundary = regexp.MustCompile("boundary=.+")
var requestBodyBoundary = regexp.MustCompile("--[0-9a-z]+")

func printRequest(contentType string, reqBody *bytes.Buffer) {
	printContentType(contentType)
	printContentLength(reqBody)
	fmt.Println()
	printRequestBody(reqBody)
}

func printContentType(contentType string) {
	fmt.Printf("Content-Type: %s\n", contentTypeBoundary.ReplaceAllLiteralString(
		contentType, "boundary="+commonBoundary))
}

func printContentLength(reqBody *bytes.Buffer) {
	fmt.Printf("Content-Length: %d\n", reqBody.Len())
}

func printRequestBody(reqBody *bytes.Buffer) {
	fmt.Println(requestBodyBoundary.ReplaceAllLiteralString(
		stringifyBuffer(reqBody), "--"+commonBoundary))
}

func stringifyBuffer(reqBody *bytes.Buffer) string {
	return strings.ReplaceAll(
		strings.ReplaceAll(reqBody.String(), "\r\n", "\n"), "\n\n\n", "\n\n")
}
