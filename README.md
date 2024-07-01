# Service-Catalog Project

## Project Overview

The Service-Catalog project is a backend application designed to serve data about services and their versions. It utilizes a RESTful API built in Go and uses PostgreSQL as its database to store service data.

## Features

- **User Authentication**: Secure login process and JWT-based authentication for API access. Current implementation is mock     authetnication and doesnt integrate with external integration.
- **Service Versions**: Ability to handle multiple versions per service.
- **Docker Integration**: Containerized application and database for easy deployment and scaling.

## Assumptions
- Authentication and Authorization is mocked and not using any external integration.
- PostgreSQL is used, however only GET APIs are implemented.
- Application startup will initialize DB with mock data for 100 services.
- No implementation of filter, sort or pagination for now.
- Logging has been implemented but not in exhaustive manner.

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

### Clone the Repository

Start by cloning the repository to your local machine. To do this, open a terminal and run the following command:

```bash
git clone https://github.com/rakshpatel/services.git
cd services
```

### Using Docker Compose

This section details how to use Docker Compose to build and start the services defined in the `docker-compose.yml` file:

#### Build and Start the Services

To build and start your services in detached mode, run the following command in your terminal:

```bash
# Start application with make
make up

# Stat application with docker-compose tool
docker-compose up -d
```

#### To examine the logs 
```bash
# Logs with make
make logs

# Logs with docker-compose tool
docker-compose logs
```

#### To Stop the Services
```bash
# Stop the application with make
make down

# Stop the application with docker-compose
docker-compose down

```

## Accessing the API

Once the application is running, you can access the API at `http://localhost:8080/v1/`. Below are some examples of how you can interact with the API using `curl`, a command-line tool that can be used to send requests to the server.

### Authenticate and Receive a JWT

Use the `/login` endpoint to authenticate and receive a JWT. You need to replace `'your_username'` and `'your_password'` with your actual login credentials. For mock you can user username as `user` and password as `password`.

```bash
curl -X POST http://localhost:8080/v1/login \
-H "Content-Type: application/json" \
-d '{"username": "your_username", "password": "your_password"}'
```

### List All Services
```bash
curl -X GET http://localhost:8080/v1/services \
-H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### Retrieve a specific service
```bash
# This retrieves the details for service with ID 1, replace 1 with any value between 1 to 100.
curl -X GET http://localhost:8080/v1/services/1 \
-H "Authorization: Bearer YOUR_JWT_TOKEN"

```

### Retrieve versions for a service
```bash
curl -X GET http://localhost:8080/v1/services/1/versions \
-H "Authorization: Bearer YOUR_JWT_TOKEN"
```

## Future Improvements
- Update middleware and integrate with actual system for Authn/Authz
- Exhaustive logging
- Implement iterface for DB interaction and add test with mocking API calls