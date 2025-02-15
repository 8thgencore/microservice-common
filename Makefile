install-deps:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/gojuno/minimock/v3/cmd/minimock@latest
	go install mvdan.cc/gofumpt@latest
	
# Linting
lint:
	golangci-lint run ./... --config .golangci.yaml

# Formating
format:
	gofumpt -l -w .

generate-mocks:
	go generate ./pkg/db
