# Variables
BINARY_NAME=iac-ttl
CDK_DIR=cdk
CLI_DIR=cli

# Build the CLI
.PHONY: build
build:
	@echo "Building CLI..."
	cd $(CLI_DIR) && go build -o ../$(BINARY_NAME)

# Install the CLI
.PHONY: install
install: build
	@echo "Installing CLI..."
	mv $(BINARY_NAME) $(HOME)/bin/

# Clean build artifacts
.PHONY: clean
clean:
	@echo "Cleaning..."
	rm -rf $(BINARY_NAME)
	rm -rf $(CDK_DIR)/cdk.out
	rm -rf $(CDK_DIR)/node_modules

# Install CDK dependencies
.PHONY: install-cdk
install-cdk:
	@echo "Installing CDK dependencies..."
	cd $(CDK_DIR) && npm install

# Build CDK
.PHONY: build-cdk
build-cdk: install-cdk
	@echo "Building CDK..."
	cd $(CDK_DIR) && npm run build

# Deploy CDK
.PHONY: cdk-deploy
cdk-deploy: build-cdk
	@echo "Deploying CDK..."
	cd $(CDK_DIR) && npm run cdk deploy

# Destroy CDK
.PHONY: cdk-destroy
cdk-destroy:
	@echo "Destroying CDK..."
	cd $(CDK_DIR) && npm run cdk destroy

# Run tests
.PHONY: test
test:
	@echo "Running tests..."
	cd $(CLI_DIR) && go test -v ./...

# Help
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  build        - Build the CLI"
	@echo "  install      - Install the CLI to $(HOME)/bin"
	@echo "  clean        - Clean build artifacts"
	@echo "  install-cdk  - Install CDK dependencies"
	@echo "  build-cdk    - Build CDK"
	@echo "  cdk-deploy   - Deploy CDK"
	@echo "  cdk-destroy  - Destroy CDK"
	@echo "  test         - Run tests"
	@echo "  help         - Show this help message"