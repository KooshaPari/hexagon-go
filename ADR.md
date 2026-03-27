# Architecture Decision Records — hexagon-go

## ADR-001 — Chi as Default HTTP Router

**Status:** Accepted  
**Date:** 2026-03-27

### Context
The template needs a default HTTP router. Chi is lightweight, stdlib-compatible, and well-maintained. Gin is heavier but more popular.

### Decision
Use `go-chi/chi` v5 as the default HTTP router. Adapters for gin can be added as an alternative.

### Consequences
- Chi middleware (logging, recovery, CORS) is included in the HTTP adapter.
- No gin dependency in the base template.

---

## ADR-002 — depguard for Layer Boundary Enforcement

**Status:** Accepted  
**Date:** 2026-03-27

### Context
Go has no built-in package boundary enforcement. Without tooling, developers accidentally import infrastructure packages from domain code.

### Decision
Configure `depguard` in `.golangci.yml` to deny infrastructure and adapter imports in the domain and application packages.

### Consequences
- `make lint` fails if a domain package imports an adapter or infrastructure package.
- Layer violations are caught at CI time, not code review time.

---

## ADR-003 — In-Memory Adapters as Default Test Doubles

**Status:** Accepted  
**Date:** 2026-03-27

### Context
Unit tests for application services require adapter implementations. Mocking frameworks add complexity; simple in-memory implementations are easier to reason about.

### Decision
Provide concrete in-memory adapter implementations (not mocks) for all outbound ports. These live in the `adapters/memory/` package.

### Consequences
- Tests use real adapter implementations, reducing mock-related false positives.
- In-memory adapters can also serve as development stubs.
