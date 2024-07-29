MODULE := github.com/0x0FACED/go-collections

.PHONY: test-list test-stack test-queue test-all test-rbt

test-list:
	go test ./list/

test-stack:
	go test ./stack/

test-queue:
	go test ./queue/

test-all:
	go test -race -v -timeout 600s ./...

test-list-race:
	go test -race -v -timeout 120s ./list/

test-rbt:
	go test ./trees/