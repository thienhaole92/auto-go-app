SERVICE_NAME = auto-go-app

.PHONY: pre-commit
pre-commit:
	go mod tidy
	go vet ./...
	go fmt ./...

.PHONY: build
build:
	env GOOS=$(TARGETOS) GOARCH=$(TARGETARCH) go build -o $(SERVICE_NAME)

.PHONY:pre-gci
pre-gci:
	go install github.com/daixiang0/gci@latest

.PHONY:gci
gci: pre-gci
	gci write -s standard -s default -s "prefix($$(go list -m))" .

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

.PHONY:pre-test-go
pre-test-go:
	go install -mod=mod github.com/onsi/ginkgo/v2/ginkgo
	go install -mod=mod github.com/boumenot/gocover-cobertura

.PHONY:test-go
test-go: pre-test-go
	ginkgo run -r --race --keep-going --junit-report report.xml --cover --coverprofile cover.out
