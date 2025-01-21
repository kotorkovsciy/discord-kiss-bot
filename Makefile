BINARY_NAME=Bot
SRC_DIR=cmd/bot
GOFLAGS=
SOURCES=$(wildcard $(SRC_DIR)/*.go)
BUILD_DIR=build

.PHONY: build run clean deps all

build: $(BUILD_DIR)/$(BINARY_NAME)

$(BUILD_DIR)/$(BINARY_NAME): $(SOURCES)
	@echo "Build project"
	go build $(GOFLAGS) -o $@ $(SOURCES)

run: build
	@echo "Run project"
	./$(BUILD_DIR)/$(BINARY_NAME)

clean:
	@echo "Clean project"
	rm -f $(BUILD_DIR)/$(BINARY_NAME)

deps:
	@echo "Install deps"
	go mod tidy

all: build
