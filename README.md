# Service-Catalog Project

## Project Overview

The Service-Catalog project is a backend application designed to serve data about services and their versions. It utilizes a RESTful API built in Go and uses PostgreSQL as its database to store service data.

## Features

- **User Authentication**: Secure login process and JWT-based authentication for API access. Current implementation is mock     authetnication and doesnt integrate with external application.
- **Service Versions**: Ability to handle multiple versions per service.
- **Docker Integration**: Containerized application and database for easy deployment and scaling.

## Project Structure

```plaintext
/services
|-- /auth
|   |-- auth.go                # Authentication middleware
|-- /backend
|   |-- db.go                  # Database connection setup
|   |-- servics.go             # Database interaction
|-- /config
|   |-- config.go              # Configuration setup
|-- /datamodels
|   |-- services.go            # Data models
|-- /handlers
|   |-- services.go            # API endpoint handlers
|-- /sqls
|   |-- init.sql               # SQL script for initializing the database, creating tables, insert mock data for 100 services
|-- docker-compose.yml         # Docker Compose to orchestrate the service and PostgreSQL
|-- Dockerfile                 # Dockerfile for building the service
|-- main.go                    # Entry point of the application
|-- Makefile                   # Makefile for simplifying build and deploy commands
|-- README.md                  # This file
```

## Prerequisites

Before you begin, ensure you have met the following requirements:

- **Docker and Docker Compose**: You will need Docker installed on your machine to build and run the containers defined in the `docker-compose.yml`. Install Docker from [Docker's official website](https://www.docker.com/get-started).

- **Go Programming Language**: Required if you intend to run or test the application locally without Docker. Install Go from the [official Go website](https://golang.org/dl/).

- **Make**: While optional, having `make` installed will allow you to use the convenient commands in the `Makefile` that simplify tasks like building, running, and cleaning up Docker containers and images. Install `make` from [GNU Make](https://www.gnu.org/software/make/) or use a package manager specific to your operating system.


## Getting Started

To get the application running locally, follow these steps:

### 1. Clone the Repository

Start by cloning the repository to your local machine. To do this, open a terminal and run the following command:

```bash
git clone https://github.com/rakshpatel/services.git
cd service-catalog
```

### Using Docker Compose

This section details how to use Docker Compose to build and start the services defined in the `docker-compose.yml` file:


### 2. Using Docker Compose

Docker Compose is used to define and run multi-container Docker applications. With Docker Compose, you can manage the lifecycle of your application and its components through a single command. Here's how to use it:

#### Build and Start the Services

To build and start your services in detached mode, run the following command in your terminal:

```bash
make up
```