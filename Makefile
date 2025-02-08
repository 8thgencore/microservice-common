LOCAL_BIN:=$(CURDIR)/bin

install-deps:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	GOBIN=$(LOCAL_BIN) go install github.com/gojuno/minimock/v3/cmd/minimock@latest
	GOBIN=$(LOCAL_BIN) go install mvdan.cc/gofumpt@latest

# Linting
lint:
	GOBIN=$(LOCAL_BIN) bin/golangci-lint run ./... --config .golangci.yaml

# Formating
format:
	GOBIN=$(LOCAL_BIN) bin/gofumpt -l -w .

generate-mocks:
	go generate ./pkg/db
