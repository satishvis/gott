#!/usr/bin/env make


tidy:
	go mod tidy

get:
	go get -v

build:
	go build
