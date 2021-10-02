.PHONY: \
	build

all:

build:
	@gofmt -w .
	@go build -o dist/rss main.go

dev: \
	build
	@./dist/rss

start:
	@./dist/rss
