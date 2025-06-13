# RFC 0001: Split Between API and Infrastructure Repositories

- **Title:** Split Between API and Infrastructure Repositories
- **Status:** Accepted
- **Date:** 2025-06-13
- **Author:** Gabriel de Espindula
- **Supersedes:** None

---

## Objective

To establish a clean separation between the Go API codebase and the infrastructure that runs it, by splitting them into two distinct and purpose-driven repositories. This separation aims to simplify deployments, encourage modular development, and facilitate infrastructure automation using Docker, Kubernetes, and Terraform.

---

## Motivation

- Avoid mixing business logic (Go code) with infrastructure concerns (Terraform, Kubernetes, etc).
- Enable clear Docker image versioning and publishing strategies.
- Simplify CI/CD pipelines by separating responsibilities.
- Improve the scalability and maintainability of both application and infra.

---

## Proposal

We will create two separate repositories:

### 1. `goapi` – Application Repository

A self-contained Go project, responsible for building and publishing the Docker image of the API.

**Structure:**
```

goapi/
├── Dockerfile              # Builds the Go binary
├── docker-compose.yml      # Optional: for local development
├── go.mod / go.sum
├── api/
│   └── main.go
├── db/
├── handlers/
├── models/
├── router/
└── README.md

```

- All code related to the Go service lives here.
- The Dockerfile uses multi-stage builds to produce a lightweight final image.
- Docker image is tagged using immutable versions (e.g. `v1.2.3`).

### 2. `goapi-infra` – Infrastructure Repository

Contains all infrastructure as code (IaC) necessary to deploy the Go API to cloud environments using Kubernetes and Terraform.

**Structure:**
```

goapi-infra/
├── docker-compose.yml         # For local/staging use, pulls published image
├── k8s/
│   ├── deployment.yaml        # References published image tag
│   ├── service.yaml
│   └── postgres-deploy.yaml
├── terraform/
│   ├── main.tf                # Infrastructure provisioning
│   └── variables.tf
└── README.md

````

---

## Technical Details

- **Docker Image Hosting:** We will use [GitHub Container Registry (GHCR)](https://ghcr.io).
  - Example tag: `ghcr.io/username/goapi:v1.2.3`
- **Kubernetes:** The `deployment.yaml` in `goapi-infra/k8s/` references a specific image tag. No `latest` tag will be used to ensure reproducibility and rollback capability.
- **Terraform:** Manages infrastructure resources like clusters, databases, and networking.

---

## Deployment Strategy

1. **From the `goapi` repo:**
   - On push to `main` or a release tag:
     - Build and push Docker image to GHCR with semantic version tag (`vX.Y.Z`)
     - Optionally update a changelog or GitHub Release

2. **From the `goapi-infra` repo:**
   - Manually or via CI, update `deployment.yaml` with the new image tag.
   - Apply changes using:
     ```bash
     kubectl apply -f k8s/
     terraform apply
     ```

---

## Rollbacks

Because we use immutable image tags, rollbacks can be done simply by reapplying the previous version tag in `deployment.yaml`. This provides a reliable and predictable way to recover from failed deployments.

---

## Alternatives Considered

- Keeping a monorepo: rejected due to increased complexity and tight coupling between layers.
- Using `latest` image tag: rejected due to lack of traceability and rollback safety.

---

## Next Steps

- Create and initialize both repositories (`goapi`, `goapi-infra`)
- Set up GitHub Actions to build and publish images on `goapi`
- Define initial deployment in `goapi-infra/k8s/`
- Document image update procedure in `goapi-infra/README.md`

---
