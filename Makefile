.PHONY: all clean lint test

PACKAGES = $(shell go list ./... | grep -v /vendor/)

all: test

clean:
	go clean -i ./...

lint:
	go fmt $(PACKAGES)
	go vet $(PACKAGES)

test:
	@for PKG in $(PACKAGES); do go test -cover -coverprofile $$GOPATH/src/$$PKG/coverage.out $$PKG || exit 1; done;
