SHELL := /bin/bash

.PHONY: build run test

build:
	@go build -o bin/ccwc main.go

run: build
	@./bin/ccwc

test:
	@go test -v ./... 
