# Go parameters
GOCMD=go
GOFMT=gofmt
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get
GOTEST=$(GOCMD) test -cover -v
GOGET=$(GOCMD) get 
BINARY_NAME=libgame
PORT?=80

.PHONY: test

all: setup build test run

setup:
		$(GOGET) -u github.com/stretchr/testify/assert github.com/go-playground/overalls
build: 
		$(GOFMT) -s -w .
		$(GOGET) ./
		$(GOBUILD) -o $(BINARY_NAME) -v *.go
test: 
		${GOPATH}/bin/overalls -project=github.com/nlittlepoole/cookiecutter-monorepo -covermode=count -debug
clean: 
		$(GOCLEAN)
		rm -f $(BINARY_NAME)
run:
		./$(BINARY_NAME) -port=$(PORT)