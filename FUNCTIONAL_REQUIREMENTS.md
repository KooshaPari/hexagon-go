# Functional Requirements — hexagon-go

## FR-STRUCT — Project Structure

| ID | Requirement |
|----|-------------|
| FR-STRUCT-001 | The template SHALL provide a canonical Go directory layout with domain, application, ports, adapters, and infrastructure packages. |
| FR-STRUCT-002 | The template SHALL include a go.mod with idiomatic module naming. |
| FR-STRUCT-003 | The template SHALL include a Taskfile with build, test, lint, and generate targets. |

## FR-DOMAIN — Domain Layer

| ID | Requirement |
|----|-------------|
| FR-DOMAIN-001 | The template SHALL include an example entity, value object, and aggregate. |
| FR-DOMAIN-002 | The template SHALL include typed domain events. |
| FR-DOMAIN-003 | The domain package SHALL have zero imports from adapters or infrastructure packages. |

## FR-APP — Application Layer

| ID | Requirement |
|----|-------------|
| FR-APP-001 | The template SHALL include command and query types with handlers following CQRS. |
| FR-APP-002 | The template SHALL include an application service that orchestrates domain operations. |
| FR-APP-003 | All command and query handlers SHALL have unit tests using mock adapters. |

## FR-PORTS — Ports and Adapters

| ID | Requirement |
|----|-------------|
| FR-PORTS-001 | The template SHALL include inbound port adapters for HTTP and gRPC. |
| FR-PORTS-002 | The template SHALL include outbound port interfaces for repository and event publishing. |
| FR-PORTS-003 | The template SHALL include in-memory adapter stubs implementing all outbound ports. |

## FR-QUALITY — Quality Gates

| ID | Requirement |
|----|-------------|
| FR-QUALITY-001 | The template SHALL include a golangci-lint configuration with strict rules enabled. |
| FR-QUALITY-002 | The template SHALL include depguard rules enforcing hexagonal layer import restrictions. |
| FR-QUALITY-003 | The template SHALL include a test coverage gate requiring 80% minimum coverage. |
