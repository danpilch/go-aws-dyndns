#basic build goapp
# Go parameters
    GOCMD=go
    GOBUILD=$(GOCMD) build
    GOCLEAN=$(GOCMD) clean
    GOTEST=$(GOCMD) test
    GOGET=$(GOCMD) get
	BINARY_PATH=./build
    BINARY_NAME=$(BINARY_PATH)/go-aws-dyndns
	BINARY_UNIX=$(BINARY_PATH)/$(BINARY_NAME)_unix
    
    all: test build
    build: 
		$(GOBUILD) -o $(BINARY_NAME) -v
    test: 
		$(GOTEST) -v ./...
    clean: 
		$(GOCLEAN)
		rm -rf $(BINARY_PATH)
    run:
		$(GOBUILD) -o $(BINARY_NAME) -v ./...
		./$(BINARY_NAME)
    deps:
		$(GOGET) github.com/markbates/goth
		$(GOGET) github.com/markbates/pop
    
    # Cross compilation
    build-linux:
		CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v
#    docker-build:
#            docker run --rm -it -v "$(GOPATH)":/go -w /go/src/bitbucket.org/rsohlich/makepost golang:latest go build -o "$(BINARY_UNIX)" -v

