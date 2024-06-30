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

# Docker Compose commands
.PHONY: up
up:
	@echo "Starting all services with Docker Compose..."
	docker-compose up -d

.PHONY: down
down:
	@echo "Stopping and removing all services with Docker Compose..."
	docker-compose down

.PHONY: ps
ps:
	@echo "Listing all running Docker services..."
	docker-compose ps

.PHONY: logs
logs:
	@echo "Displaying logs from Docker services..."
	docker-compose logs

# Display help
.PHONY: help
help:
	@echo "Makefile for managing the Docker image for the Go service and Docker Compose environments"
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
	@echo "  up         Start all services in detached mode using Docker Compose"
	@echo "  down       Stop all services using Docker Compose"
	@echo "  ps         List Docker Compose services"
	@echo "  logs       Display logs for Docker Compose services"
	@echo "  help       Display this help message"
