.PHONY: \
	build

all:

build:
	@gofmt -w .
	@go build -o dist/rss main.go

build-docker:
	@docker build -f ./build/Dockerfile  -t impossible98/rss .

dev: \
	build
	@./dist/rss

start:
	@./dist/rss
