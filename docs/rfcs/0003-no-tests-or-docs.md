# RFC 0003: No Automated Tests or Formal API Documentation
## Context

The main goal of this project is to explore and practice modern infrastructure concepts using:

- **Docker** for packaging
- **Kubernetes** for container orchestration
- **Terraform** for infrastructure as code
- **CI/CD** for deployment automation

The Go application is a simple CRUD service with no complex business logic.

## Decision

It was decided **not to include automated tests or formal API documentation (e.g., Swagger)**.

### Justifications

- The project’s core focus is on **infrastructure**, not building a production-grade API.
- The time and effort required for writing tests and generating documentation would not provide meaningful value aligned with the project's goals.
- The API is minimal, with straightforward endpoints that are easy to understand by reading the source code.
- Validation and testing will be done manually using tools like `curl`, `httpie`, or REST clients like Insomnia/Postman.
- Tests and documentation can be introduced later if the scope of the project expands.

## Considered Alternatives

- Writing integration tests using `httptest` with `Gin` — rejected to maintain focus and simplicity.
- Generating Swagger docs using `swaggo/swag` — also rejected as it's out of scope.

## Consequences

- Lower test coverage, but increased focus and speed toward the main learning objectives: Docker, Kubernetes, Terraform, and CI/CD.
- Simplicity of the code is intentional for educational purposes.
