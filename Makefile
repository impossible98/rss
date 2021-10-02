.PHONY: \
	build

all:

build:
	@gofmt -w .
	@go build -o dist/rss main.go

build-docker:
	@docker build -f ./build/Dockerfile  -t impossible98/rss .

build-frontend:
	cd frontend; \
	npm install; \
	./node_modules/.bin/vue-cli-service build; \
	cp -r dist/ ../web/public/

dev: \
	build
	@./dist/rss

start:
	@./dist/rss
