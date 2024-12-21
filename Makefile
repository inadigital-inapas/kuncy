# Binary name and versioning
BINARY_NAME=kuncy
VERSION=$(shell git describe --tags --always)

# Build configuration
BUILD_DIR=build
DIST_DIR=dist
MAIN_DIR=.

# Go build settings
GOBUILD=go build
GOARCH=amd64

# Our supported platforms
PLATFORMS=linux windows darwin

# Default target
.DEFAULT_GOAL := build

# Initialize build directories
init:
	@echo "Creating build directories..."
	@mkdir -p $(BUILD_DIR)
	@mkdir -p $(DIST_DIR)

# Clean previous builds
clean:
	@echo "Cleaning up previous builds..."
	@rm -rf $(BUILD_DIR)
	@rm -rf $(DIST_DIR)

# Build for all platforms
build: clean init
	@echo "Building version: $(VERSION)"
	@# Build for Linux
	@echo "Building for Linux..."
	@GOOS=linux GOARCH=$(GOARCH) $(GOBUILD) \
		-ldflags="-X main.Version=$(VERSION)" \
		-o $(BUILD_DIR)/$(BINARY_NAME)-$(VERSION)-linux-$(GOARCH) \
		$(MAIN_DIR)
	
	@# Build for Windows
	@echo "Building for Windows..."
	@GOOS=windows GOARCH=$(GOARCH) $(GOBUILD) \
		-ldflags="-X main.Version=$(VERSION)" \
		-o $(BUILD_DIR)/$(BINARY_NAME)-$(VERSION)-windows-$(GOARCH).exe \
		$(MAIN_DIR)
	
	@# Build for macOS
	@echo "Building for Darwin/macOS..."
	@GOOS=darwin GOARCH=$(GOARCH) $(GOBUILD) \
		-ldflags="-X main.Version=$(VERSION)" \
		-o $(BUILD_DIR)/$(BINARY_NAME)-$(VERSION)-darwin-$(GOARCH) \
		$(MAIN_DIR)
	
	@# Package the builds
	@echo "Packaging builds..."
	@cd $(BUILD_DIR) && \
		tar czf ../$(DIST_DIR)/$(BINARY_NAME)-$(VERSION)-linux-$(GOARCH).tar.gz \
			$(BINARY_NAME)-$(VERSION)-linux-$(GOARCH) && \
		zip ../$(DIST_DIR)/$(BINARY_NAME)-$(VERSION)-windows-$(GOARCH).zip \
			$(BINARY_NAME)-$(VERSION)-windows-$(GOARCH).exe && \
		tar czf ../$(DIST_DIR)/$(BINARY_NAME)-$(VERSION)-darwin-$(GOARCH).tar.gz \
			$(BINARY_NAME)-$(VERSION)-darwin-$(GOARCH)
	
	@echo "Build complete! Artifacts available in $(DIST_DIR)/"
	@echo "Built version: $(VERSION)"

# Show current version
version:
	@echo $(VERSION)

.PHONY: init clean build version