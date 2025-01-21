BINARY_NAME=Bot
SRC_DIR=cmd/bot
GOFLAGS=
SOURCES=$(wildcard $(SRC_DIR)/*.go)
GOFILES=$(shell go list -f '{{range .GoFiles}}{{$$.Dir}}/{{.}}{{"\n"}}{{end}}' ./...)
BUILD_DIR=build

.PHONY: build run clean deps all

build: $(BUILD_DIR)/$(BINARY_NAME)

$(BUILD_DIR)/$(BINARY_NAME): $(GOFILES)
	@echo Build project
	@go build $(GOFLAGS) -o $@ $(SOURCES)

run: build
	@echo Run project
	@./$(BUILD_DIR)/$(BINARY_NAME)

clean:
	@echo Clean project
ifeq ($(OS),Windows_NT)
	@powershell -Command "Remove-Item -Recurse -Force $(BUILD_DIR)"
else
	@rm -rf $(BUILD_DIR)
endif

deps:
	@echo Install deps
	@go mod tidy

all: build
