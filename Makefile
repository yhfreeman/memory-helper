GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=memoryhelper
BINARY_UNIX=$(BINARY_NAME)

all: deps build
build:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GOBUILD) -o ${BINARY_NAME}

deps:
	$(GOGET) github.com/gin-gonic/gin
	$(GOGET) github.com/garyburd/redigo/redis

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v
