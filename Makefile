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

# Supported platforms
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
	@for os in $(PLATFORMS); do \
		echo "Building for $$os..."; \
		GOOS=$$os GOARCH=$(GOARCH) $(GOBUILD) \
			-ldflags="-X main.Version=$(VERSION)" \
			-o $(BUILD_DIR)/$$os/$(BINARY_NAME)$$(test $$os = windows && echo ".exe") \
			$(MAIN_DIR); \
		mkdir -p $(DIST_DIR); \
		if [ "$$os" = "windows" ]; then \
			zip -j $(DIST_DIR)/kuncy-$(VERSION)-$$os.zip $(BUILD_DIR)/$$os/$(BINARY_NAME).exe; \
		else \
			tar czf $(DIST_DIR)/kuncy-$(VERSION)-$$os.tar.gz -C $(BUILD_DIR)/$$os $(BINARY_NAME); \
		fi; \
	done

	@echo "Removing build directory..."
	@rm -rf $(BUILD_DIR)

	@echo "Build complete! Artifacts available in $(DIST_DIR)/"
	@echo "Built version: $(VERSION)"

# Show current version
version:
	@echo $(VERSION)

.PHONY: init clean build version