# PRD — hexagon-go

## Overview

`hexagon-go` is a production-ready Go project template following hexagonal architecture (Ports and Adapters), Clean Architecture, SOLID principles, and xDD methodologies (TDD, BDD, DDD, CDD, ADD). It serves as the canonical starting point for all Go services in the Phenotype ecosystem.

## Goals

- Eliminate boilerplate for new Go services by providing a fully structured template.
- Enforce hexagonal layer boundaries at the code level with dependency guard tooling.
- Demonstrate best-practice patterns for domain modeling, port definition, and adapter implementation in Go.

## Epics

### E1 — Project Structure
- E1.1 Canonical directory layout: domain/, application/, ports/, adapters/, infrastructure/.
- E1.2 Go module with idiomatic naming and versioning.
- E1.3 Makefile / Taskfile with standard targets: build, test, lint, generate.

### E2 — Domain Layer
- E2.1 Example entity, value object, and aggregate.
- E2.2 Domain events with typed payloads.
- E2.3 Domain service interface and implementation.

### E3 — Application Layer (CQRS)
- E3.1 Command and query types with handlers.
- E3.2 Application service orchestrating domain operations.
- E3.3 Unit tests for all handlers using mock adapters.

### E4 — Ports and Adapters
- E4.1 Inbound ports: HTTP handler (chi/gin), gRPC server.
- E4.2 Outbound ports: repository, event publisher.
- E4.3 Adapter stubs ready for real implementations.

### E5 — Quality Gates
- E5.1 golangci-lint configuration with strict rules.
- E5.2 depguard rules enforcing hexagonal layer imports.
- E5.3 Test coverage gate at 80%.

## Acceptance Criteria

- A new service scaffolded from this template compiles and passes all tests out of the box.
- `make lint` passes with zero errors on the template code.
- Domain layer has no imports from adapters or infrastructure.
