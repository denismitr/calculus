GO_VERSION ?= 1.14

.PHONY: test

test:
	go test ./...
