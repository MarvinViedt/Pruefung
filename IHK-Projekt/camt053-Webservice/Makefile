GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
BINARY_NAME=camt053-Webservice
BINARY_UNIX=$(BINARY_NAME)_UNIX

all: build
build:
	$(GOBUILD) -o $(BINARY_NAME) -v 
#-race
#test:
#	$(GOTEST) -v ./...
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)
														
