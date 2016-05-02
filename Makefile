.PHONY: all clean deps lint test

PACKAGES = $(shell go list ./... | grep -v /vendor/)

all: deps test

clean:
	go clean -i ./...

deps:
	go get -u github.com/govend/govend
	govend -v

lint:
	go fmt $(PACKAGES)
	go vet $(PACKAGES)

test:
	@for PKG in $(PACKAGES); do go test -cover -coverprofile $$GOPATH/src/$$PKG/coverage.out $$PKG || exit 1; done;
