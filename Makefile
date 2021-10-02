.PHONY: \
	build

all:

build:
	@go build -o dist/rss main.go

dev: \
	build
	@./dist/rss

start:
	@./dist/rss
