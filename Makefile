NAME := plop

test: unit acceptance lint

acceptance: build
	bats test

build: dependencies
	go build -o bin/$(NAME)

unit: dependencies
	go test ./...

lint: dependencies
	@which gometalinter || (go get -u github.com/alecthomas/gometalinter && gometalinter --install --update)
	gometalinter --deadline 1m ./...

dependencies:
	go get -t

.PHONY: acceptance build dependencies lint test unit
