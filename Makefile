.PHONY: test

test:
	@go test -v ./...

fmt:
	@go fmt ./...

generate:
	@go generate ./...
