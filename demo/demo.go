package demo

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
)

const commonBoundary = "1879bcd06ac39a4d8fa5"

var contentTypeBoundary = regexp.MustCompile("boundary=.+")
var requestBodyBoundary = regexp.MustCompile("--[0-9a-z]+")

func PrintRequest(contentType string, reqBody *bytes.Buffer) {
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
