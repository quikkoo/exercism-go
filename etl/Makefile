SHELL = /bin/sh

export GOPATH = $(HOME):$(CURDIR)

MODULE = etl

GOGET   = go get -u
GOVET   = go vet -v
GOLINT  = golint
GOTEST  = go test -v -cover -covermode=count -coverprofile=coverage/coverage.cov
GOCOVER = go tool cover
MKDIR   = mkdir -p 
RM      = rm -rf

all: test

check:
	$(GOGET) github.com/golang/lint/golint
	$(GOVET) $(MODULE)
	$(GOLINT) $(MODULE)

test: check
	$(MKDIR) coverage
	$(GOTEST) $(MODULE) 
	$(GOCOVER) -func=coverage/coverage.cov
	$(GOCOVER) -html=coverage/coverage.cov -o coverage/coverage.html

clean:
	$(RM) coverage
