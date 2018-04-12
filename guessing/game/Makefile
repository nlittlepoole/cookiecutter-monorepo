# Go parameters
GOCMD=go
GOFMT=gofmt
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get
GOTEST=$(GOCMD) test -cover -v
GOGET=$(GOCMD) get 
GOVET=$(GOCMD) vet
BINARY_NAME=libgame
PORT?=80

.PHONY: test

all: clean setup build test run

travis: | setup build test codecov

setup:
		$(GOGET) -u github.com/stretchr/testify/assert github.com/go-playground/overalls
		$(GOGET) -u golang.org/x/lint/golint
		$(GOGET) -u github.com/fzipp/gocyclo
build: 
		$(GOFMT) -s -w .
		$(GOGET) ./
		$(GOBUILD) -o $(BINARY_NAME) -v *.go
test: 
		${GOPATH}/bin/overalls -project=github.com/nlittlepoole/cookiecutter-monorepo/guessing/game -covermode=count -debug 
		mv overalls.coverprofile coverage.txt
		$(GOVET)
		${GOPATH}/bin/golint ./...
		${GOPATH}/bin/gocyclo -over 10 *.go */*.go
codecov:
		curl -s https://codecov.io/bash | bash -

clean: 
		$(GOCLEAN)
		rm -f $(BINARY_NAME)
		rm -f *coverage.txt
		rm -f *.coverprofile */*.coverprofile
run:
		./$(BINARY_NAME) -port=$(PORT)