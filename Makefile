MODULE := github.com/0x0FACED/go-collections

.PHONY: test-array-list test-all

test-list:
	go test ./list/

test-all:
	go test ./...