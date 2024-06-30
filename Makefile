# Name of the Docker image
IMAGE_NAME = service-catalog
IMAGE_TAG = v0.0.1

# Dockerfile location
DOCKERFILE = Dockerfile

# Default target
.PHONY: all
all: build

# Build the Docker image
.PHONY: build
build:
	@echo "Building the Docker image..."
	docker build -t $(IMAGE_NAME):$(IMAGE_TAG) -f $(DOCKERFILE) .

# Run the Docker container
.PHONY: run
run:
	@echo "Running the Docker container..."
	docker run -d -p 8080:8080 $(IMAGE_NAME):$(IMAGE_TAG)

# Stop and remove the running container
.PHONY: stop
stop:
	@echo "Stopping the Docker container..."
	docker ps -q --filter "ancestor=$(IMAGE_NAME):$(IMAGE_TAG)" | xargs -r docker stop
	@echo "Removing the Docker container..."
	docker ps -a -q --filter "ancestor=$(IMAGE_NAME):$(IMAGE_TAG)" | xargs -r docker rm

# Remove the Docker image
.PHONY: clean
clean: stop
	@echo "Removing the Docker image..."
	docker rmi $(IMAGE_NAME):$(IMAGE_TAG)

# Rebuild the Docker image
.PHONY: rebuild
rebuild: clean build

# Display help
.PHONY: help
help:
	@echo "Makefile for managing the Docker image for the Go service"
	@echo ""
	@echo "Usage:"
	@echo "  make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  all        Default target, builds the Docker image"
	@echo "  build      Build the Docker image"
	@echo "  run        Run the Docker container"
	@echo "  stop       Stop and remove the running container"
	@echo "  clean      Remove the Docker image and container"
	@echo "  rebuild    Rebuild the Docker image"
	@echo "  help       Display this help message"

