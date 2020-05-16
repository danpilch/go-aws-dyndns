#basic build goapp
# Go parameters
    GOCMD=go
    GOBUILD=$(GOCMD) build
    GOCLEAN=$(GOCMD) clean
    GOTEST=$(GOCMD) test
    GOGET=$(GOCMD) get
	BINARY_PATH=./build
    BINARY_NAME=$(BINARY_PATH)/go-aws-dyndns
    
    all: clean test build-linux-amd64 build-linux-arm5 build-linux-arm6 build-windows-amd64
    test: 
		$(GOBUILD) fmt
		$(GOTEST) -v ./...
    clean: 
		$(GOCLEAN)
		rm -rf $(BINARY_PATH)
    run:
		$(GOBUILD) fmt
		$(GOBUILD) -o $(BINARY_NAME) -v ./...
		./$(BINARY_NAME)
    
    # Cross compilation
    build-linux-amd64:
		CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME)-linux-amd64 -v
    build-linux-arm5:
		CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=5 $(GOBUILD) -o $(BINARY_NAME)-linux-arm5 -v
    build-linux-arm6:
		CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=6 $(GOBUILD) -o $(BINARY_NAME)-linux-arm6 -v
    build-windows-amd64:
		CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME)-win-amd64 -v

#    docker-build:
#            docker run --rm -it -v "$(GOPATH)":/go -w /go/src/bitbucket.org/rsohlich/makepost golang:latest go build -o "$(BINARY_UNIX)" -v

