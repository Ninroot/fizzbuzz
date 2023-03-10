ROOT_DIR:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

all:
	go build -o ${ROOT_DIR}

test:
	go test -v ./...