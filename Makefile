ROOT_DIR:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

all:
	go get ./...
	go build -o ${ROOT_DIR}

test:
	go test -v ./...