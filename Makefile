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
	go tool cover -html=cover.txt -o cover.html

clean:
	go clean
	rm -rf cover.*

.PHONY: vet build test cover clean
