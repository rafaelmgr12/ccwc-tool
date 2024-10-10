SHELL := /bin/bash

build:
	@go build -o bin/ccwc main.go

run: build
	@./bin/ccwc

test:
	@go test -v ./... 