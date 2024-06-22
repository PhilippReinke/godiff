GIT_VERSION := $(shell git rev-parse HEAD)

test:
	go test ./...

install:
	go install -ldflags "-X main.GitCommit=$(GIT_VERSION)" .

build:
	go build -ldflags "-X main.GitCommit=$(GIT_VERSION)"

clean:
	go clean -testcache
