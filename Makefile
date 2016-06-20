.PHONY: all clean deps lint test docs

PACKAGES = $(shell go list ./... | grep -v /vendor/)
HUGO ?= 0.15_linux_amd64

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

docs:
	curl -L -o /tmp/hugo_$(HUGO).tar.gz https://github.com/spf13/hugo/releases/download/v0.15/hugo_$(HUGO).tar.gz
	tar xvf /tmp/hugo_$(HUGO).tar.gz -C /tmp/
	mv /tmp/hugo_$(HUGO)/hugo_$(HUGO) /tmp/hugo
	/tmp/hugo -s docs/
