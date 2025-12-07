SRC_DIR = .
BIN_DIR = $(SRC_DIR)/bin
CMD_DIR = $(SRC_DIR)/cmd
DOC_DIR = $(SRC_DIR)/docs

# Source files are all files with ".go" extension except test files ("_test.go").
SRC_FILES = $(shell find $(SRC_DIR) -type "f" -name "*.go" ! -name "*_test.go")

# golangci-lint is a linter for Go.
GOLANGCI = $(BIN_DIR)/golangci-lint

ifndef VERBOSE
# --silent drops the need to prepend `@` to suppress command output.
MAKEFLAGS += --silent
endif

.DEFAULT_GOAL := all

## all: Performs a complete health check of the project.
.PHONY: all
all: lint test

## check: Performs a complete health check of the project.
.PHONY: check
check: all

## clean: Cleans up everything.
.PHONY: clean
clean:
	rm -rf $(BIN_DIR)

## help: Creates this help message.
.PHONY: help
help:
	awk 'BEGIN { printf "\n\033[36m\033[0mUsage:\n  make\n" } \
	/^##/ { sub(/^## /, ""); split($$0, a, ": "); printf "  \033[36m%-27s\033[0m %s\n", a[1], a[2] } \
	END { printf "\033[0m" }' $(MAKEFILE_LIST)

## lint: Performs a check for broken conventions and code quality violations.
.PHONY: lint
lint: | $(GOLANGCI)
	go fmt ./...
	go vet ./...
	$(GOLANGCI) run

## test: Executes a full (unit and integration) test suite.
.PHONY: test
test:
	go test $(shell go list ./...) -count=1 -short

## run: Runs the CLI
.PHONY: run
run:
	go run cmd/aoc2025/main.go

# Ensures $BIN_DIR exists.
$(BIN_DIR):
	mkdir -p $@

# Ensures $GOLANGCI is installed.
$(GOLANGCI): | $(BIN_DIR)
	curl -sSfL \
		https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh \
	| sh -s -- -b $(BIN_DIR) v2.7.1
	golangci-lint --version
