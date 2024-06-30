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
