SERVICE_NAME = auto-go-app

.PHONY: pre-commit
pre-commit:
	go mod tidy
	go vet ./...
	go fmt ./...

.PHONY: build
build:
	env GOOS=$(TARGETOS) GOARCH=$(TARGETARCH) go build -o $(SERVICE_NAME)

.PHONY:pre-lint
pre-lint:
	go install mvdan.cc/gofumpt@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

.PHONY:lint
lint: pre-lint
	go mod tidy
	gofumpt -l -w .
	go vet ./...
	golangci-lint run
