run:
	go run -gcflags=-G=3 src/map.go 

build:
	go build -gcflags=-G=3 src/* 

test:
	go test -gcflags=-G=3 ./...
