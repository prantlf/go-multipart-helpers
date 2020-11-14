# go-multipart-helpers

Helps writing file parts to MIME multipart messages according to [RFC7578] while preserving the original content type inferred from the file extension. See the [documentation] for more information.

`CreateFormFile` from [multipart writer] sets the content type always to `application/octet-stream`. If you need to preserve the content type of the file decided by its file extension, choose a function among `CreateFilePart`,` WriteFile` or `WriteFileReader` from this package.

## Installation

Add this package to `go.mod` and `go.sub` in your Go project:

    go get github.com/prantlf/go-multipart-helpers

## Usage

Upload a file with comment:

```go
import (
  "net/http"
  helpers "github.com/prantlf/go-multipart-helpers"
)
// compose a multipart form-data content
message := &bytes.Buffer{}
writer := multipart.NewWriter(message)
comp.AddField("comment", "a comment")
err := helpers.WriteFile(writer, "file", "test.txt")
// post a request with the generated content type and body
resp, err := http.DefaultClient.Post("http://host.com/upload",
  writer.FormDataContentType(), message)
```

See the [documentation] for the full interface.

[documentation]: https://pkg.go.dev/github.com/prantlf/go-multipart-helpers
[RFC7578]: https://tools.ietf.org/html/rfc7578
[multipart writer]: https://golang.org/pkg/mime/multipart/#Writer
