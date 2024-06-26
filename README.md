# Simple Bank Project

This repository contains the source code for the Simple Bank project, a backend application for managing bank accounts and transactions.
The project is in development and currently uses Golang, Docker, and AWS, with a focus on best practices for software development, testing, and deployment.

## Currently Implemented Features

### Authentication and Authorization
- Implemented basic login and sign-up journeys.
- Utilized JWT and PASETO tokens for secure authentication.
- Integrated authentication middleware in Gin framework.

### Money Transfer
- Users can transfer money between accounts.
- Developed transfer money API with custom parameter validation.

### Database Management
- Handled concurrent transactions to ensure data integrity.
- Implemented mechanisms to handle and prevent DB deadlocks.
- Used PostgreSQL for database operations.

### Testing and CI/CD
- Achieved 100% test coverage by mocking database interactions.
- Set up GitHub Actions for automated testing and CI/CD pipeline.
- Wrote unit tests for CRUD operations and API endpoints.

### Configuration Management
- Managed application configurations using Viper.
- Secured passwords with Bcrypt hashing.

## In Pipeline

- **Create a production DB on AWS RDS**
- **Store & retrieve production secrets with AWS Secrets Manager**
- **Understanding Kubernetes architecture & create an EKS cluster on AWS**

## Techstack used

- **Golang**: Primary programming language
- **PostgreSQL**: Database management system
- **Docker**: Containerization platform
- **AWS**: Cloud service provider
- **Gin**: HTTP web framework for Golang
- **Viper**: Configuration management
- **Gomock**: Mocking framework for testing
- **PASETO**: Token-based authentication
- **GitHub Actions**: CI/CD pipeline**
