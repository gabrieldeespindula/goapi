# goapi

A simple Go API aimed at educational purposes and experimentation with modern DevOps practices, including:

- Docker image build and versioning
- Deployment via Kubernetes
- Infrastructure management with Terraform
- Automation with GitHub Actions

## 🎯 Purpose

This project serves as a foundation to learn and demonstrate best practices for:

- Organizing Go projects
- Packaging with Docker
- Deploying in orchestrated environments (like Kubernetes)
- Separating application code from infrastructure

The application itself is a basic CRUD API that connects to a PostgreSQL database to manage simple records.

## 📦 Project Structure

```

goapi/
├── Dockerfile             # Defines the API Docker image
├── go.mod / go.sum        # Go dependencies
├── api/                   # Application entrypoint
│   └── main.go
├── db/                    # Database connection
├── handlers/              # Route handlers logic
├── models/                # Data structures
└── router/                # API routes definitions

````

## 🚀 Running Locally

### With Docker Compose

```bash
docker compose up --build
````

API available at `http://localhost:8080`.

### Without Docker

```bash
go run ./api
```

(Requires a local PostgreSQL database running)

## 🐳 Docker

### Manual Image Build

```bash
docker build -t ghcr.io/your-username/goapi:<tag> .
```

### Publishing (CI-driven)

When creating a new Git tag (e.g. `v1.2.3`), the CI pipeline will automatically:

* Build the Docker image
* Push it to: `ghcr.io/your-username/goapi:<tag>`
* Use immutable tags for traceability and safe rollback

## 🧪 Tests

```bash
go test ./...
```

## 📘 RFCs

All architecture, organization, and deployment decisions are documented under the [`docs/rfcs/`](./docs/rfcs/) folder.

Initial document: [`docs/rfcs/0001-split-between-API-and-infra.md`](./docs/rfcs/0001-split-between-API-and-infra.md)

## 🗂️ Related Repositories

This repository contains **only the API code**.

Infrastructure (Kubernetes, Terraform, database, etc.) is located in the repository:

👉 [`goapi-infra`](https://github.com/your-username/goapi-infra)
