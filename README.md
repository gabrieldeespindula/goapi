# goapi

> A simple CRUD API in Go, built for learning Docker, Kubernetes, Terraform, and CI/CD.

---

## ✨ Purpose

This project serves as a practical learning base for:

- Go (building a RESTful API)
- Docker (containerization for development and production)
- Kubernetes (deployment, services, ingress, volumes)
- Terraform (infrastructure as code for provisioning)
- CI/CD with GitHub Actions

---

## 📁 Project Structure

```

goapi/
├── api/                  # Go API source code
│   └── main.go
├── Dockerfile            # Builds the Go application
├── docker-compose.yml    # Local environment with API + DB
├── k8s/                  # Kubernetes manifests
│   ├── deployment.yaml
│   ├── service.yaml
│   ├── postgres-deploy.yaml
│   └── ingress.yaml
├── terraform/            # Terraform infrastructure config
│   ├── main.tf
│   └── variables.tf
└── .github/workflows/    # CI/CD workflows
└── deploy.yml

```

---

## 🛠 Technologies

- [Go](https://golang.org)
- [Docker](https://www.docker.com/)
- [Kubernetes](https://kubernetes.io/)
- [Terraform](https://www.terraform.io/)
- [PostgreSQL](https://www.postgresql.org/)
- [GitHub Actions](https://github.com/features/actions)

---

## 🚀 Running Locally

```bash
  docker compose up --build
```


The API will be available at: [http://localhost:8080](http://localhost:8080)

---

## ✅ Planned Features

* [ ] Create, retrieve, update, and delete records (`/records`)
* [ ] PostgreSQL database with persistence
* [ ] Local and cloud Kubernetes deployment
* [ ] Terraform-managed infrastructure
* [ ] CI/CD pipeline with GitHub Actions

---

## 📚 In Progress

This repository is a personal learning project and will be continuously updated and improved.
