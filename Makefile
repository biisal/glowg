run:
	go run ./cmd/debug/main.go
test:
	go test ./...
format:
	gofmt -w .
lint:
	golangci-lint run

.PHONY: run