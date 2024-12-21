# Binary name and versioning
BINARY_NAME=kuncy
VERSION=$(shell git describe --tags --always)

# Build configuration
DIST_DIR=dist
MAIN_DIR=.

# Go build settings
GOBUILD=go build
GOARCH=amd64

# Supported platforms
PLATFORMS=linux windows darwin

# Default target
.DEFAULT_GOAL := build

# Initialize build directories
init:
	@echo "Creating build directories..."
	@mkdir -p $(DIST_DIR)

# Clean previous builds
clean:
	@echo "Cleaning up previous builds..."
	@rm -rf $(DIST_DIR)

# Build for all platforms
build: clean init
	@echo "Building version: $(VERSION)"
	@for os in $(PLATFORMS); do \
		echo "Building for $$os..."; \
		mkdir -p $(DIST_DIR)/$$os; \
		if [ "$$os" = "windows" ]; then \
			GOOS=$$os GOARCH=$(GOARCH) $(GOBUILD) \
				-ldflags="-X main.Version=$(VERSION)" \
				-o $(DIST_DIR)/$$os/$(BINARY_NAME).exe \
				$(MAIN_DIR); \
		else \
			GOOS=$$os GOARCH=$(GOARCH) $(GOBUILD) \
				-ldflags="-X main.Version=$(VERSION)" \
				-o $(DIST_DIR)/$$os/$(BINARY_NAME) \
				$(MAIN_DIR); \
		fi; \
	done

	@echo "Build complete! Artifacts available in $(DIST_DIR)/"
	@echo "Built version: $(VERSION)"

# Show current version
version:
	@echo $(VERSION)

.PHONY: init clean build version