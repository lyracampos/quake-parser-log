GOLANGCI_LINT := go run github.com/golangci/golangci-lint/cmd/golangci-lint@v1.58.2

lint:
	$(GOLANGCI_LINT) run --fix
	
build:
	go build cmd/main.go

run:
	go run cmd/main.go -h ./static/html

test:
	go test ./... -cover