.PHONY: all clean lint test

PACKAGES = $(shell go list ./... | grep -v /vendor/)

all: test

clean:
	go clean -i ./...

lint:
	go fmt $(PACKAGES)
	go vet $(PACKAGES)

test:
	echo "mode: count" > coverage-all.out
	$(foreach pkg,$(PACKAGES),\
		go test -coverprofile=coverage.out -covermode=count $(pkg);\
		tail -n +2 coverage.out >> coverage-all.out;)
	go tool cover -html=coverage-all.out
