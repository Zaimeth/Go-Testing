#!bin/sh
define USAGE
Commands:
    run-rest			Run Rest API main.go
	  build-rest		Build Rest API  binary
    start-rest    Run Rest API built app
endef

run-rest:
	go mod vendor
	go run cmd/rest/main.go

build-rest:
	go mod vendor
	go build -o dist/omdb-service-rest cmd/rest/main.go

start-rest:
	./dist/omdb-service