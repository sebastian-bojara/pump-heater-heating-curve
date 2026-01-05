.PHONY: build run clean help

IMAGE_NAME=heating-curve
BINARY_NAME=heating-curve

help: ## Display help
	@echo "Available commands:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  %-15s %s\n", $$1, $$2}'

build: ## Build the Docker image
	docker build -t $(IMAGE_NAME) .

run: build ## Run the application in Docker (use ARGS="..." to pass arguments, e.g., make run ARGS="-min-ext -25")
	docker run --rm $(IMAGE_NAME) $(ARGS)

clean: ## Remove Docker image
	docker rmi $(IMAGE_NAME)
