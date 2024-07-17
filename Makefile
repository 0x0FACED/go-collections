MODULE := github.com/0x0FACED/go-collections

.PHONY: test-list test-stack test-all

test-list:
	go test ./list/

test-stack:
	go test ./stack/

test-all:
	go test ./...