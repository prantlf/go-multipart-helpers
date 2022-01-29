.DEFAULT_GOAL = test

all: vet build test

vet:
	go vet

build:
	go build

test:
	go test -v

cover:
	go test -coverprofile cover.txt
	go tool cover -html=coverage.txt -o coverage.html

clean:
	go clean
	rm -rf coverage.*

publish:
	GOPROXY=proxy.golang.org go list -m 'github.com/prantlf/go-multipart-helpers@v$(VERSION)'

.PHONY: vet build test cover clean
