export GO111MODULE=on
BIN := s2test
SOURCES ?= $(shell find . -name "*.go" -type f)

.PHONY: build
build: $(BIN)

.PHONY: test
test:
	go test ./...

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: vet
vet:
	go vet ./...

.PHONY: lint
lint: vet fmt

$(BIN): $(SOURCES)
	go build -o $@
