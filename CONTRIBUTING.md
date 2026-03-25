# Contributing to Hexagonal Architecture Templates

Thank you for your interest in contributing! This project follows the Hexagonal Architecture pattern (Ports & Adapters).

## Architecture Overview

```
┌─────────────────────────────────────────────────────────┐
│                   Infrastructure Layer                    │
│  ┌──────────────┐ ┌──────────────┐ ┌──────────────┐  │
│  │  Persistence │ │     API     │ │  Messaging   │  │
│  └──────┬───────┘ └──────┬───────┘ └──────┬───────┘  │
└─────────┼────────────────┼────────────────┼───────────┘
          │                │                │
          ▼                ▼                ▼
┌─────────────────────────────────────────────────────┐
│                   Application Layer                    │
│  ┌──────────────┐ ┌──────────────┐ ┌────────────┐  │
│  │    Ports     │ │  Use Cases  │ │    DTOs   │  │
│  └──────┬───────┘ └──────┬───────┘ └────────────┘  │
└─────────┼────────────────┼───────────────────────────┘
          │                │
          ▼                ▼
┌─────────────────────────────────────────────────────┐
│                     Domain Layer                      │
│  ┌──────────────┐ ┌──────────────┐ ┌────────────┐  │
│  │  Entities   │ │   Services  │ │   Events   │  │
│  └─────────────┘ └─────────────┘ └────────────┘  │
└─────────────────────────────────────────────────────┘
```

## Methodologies

- **Hexagonal Architecture** (Ports & Adapters)
- **Clean Architecture**
- **SOLID Principles**
- **KISS, DRY, YAGNI**
- **TDD/BDD**
- **Domain-Driven Design (DDD)**

## Pull Request Guidelines

1. **Fork** the repository
2. **Create** a feature branch: `feat/your-feature`
3. **Write tests first** (TDD approach)
4. **Follow** the hexagonal structure
5. **Ensure** all tests pass
6. **Update** documentation as needed
7. **Commit** using conventional commits
8. **Push** and open a Pull Request

## Code Style

- Follow language-specific idioms
- Use meaningful names
- Keep functions small (Single Responsibility)
- Dependency injection for adapters

## License

MIT - See [LICENSE](LICENSE)
