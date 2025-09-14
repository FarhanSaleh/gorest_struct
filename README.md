# Go REST API Project Folder Structure

This repository contains the **folder layout** for a **Golang REST API** project.  
The structure is designed to be modular, scalable, and easy to maintain.

## Folder Structure

```
.
├── cmd
├── config
├── constants
├── database
├── internal/
│   ├── app /
│   │   ├── module/
│   │   │   ├── service.go
│   │   │   ├── handler.go
│   │   │   └── repository.go  
│   │   └── app.go
│   ├── cli
│   ├── domain
│   └── middleware 
└── pkg/
    └── helper
```

## Directory Overview

| Folder / File             | Description                                                                 |
|---------------------------|-----------------------------------------------------------------------------|
| **cmd/**                  | Application entry point, e.g., `main.go` or other executables.             |
| **config/**               | Configuration files (YAML/JSON/env loader, etc.).                          |
| **constants/**            | Global constants.                                                           |
| **database/**             | Database connection setup, migrations, and related files.                   |
| **internal/**             | Core application code, private to this module.                             |
| `internal/app/`           | Main business logic of the application.                                     |
| `internal/app/module/`    | Example feature module (service, handler, repository).                      |
| `internal/app/app.go`     | Application initialization (router, server setup).                          |
| `internal/cli/`           | Command-line utilities if needed.                                           |
| `internal/domain/`        | Domain/entity definitions.                                                  |
| `internal/middleware/`    | HTTP middleware (authentication, logging, etc.).                            |
| **pkg/**                  | Reusable packages that can be imported by external projects if desired.     |
| `pkg/helper/`             | Helper utilities (formatters, general-purpose functions).                   |

## Notes

- This structure follows the common **Go Project Layout** pattern, adapted for a REST API.
- The `internal` directory ensures that code cannot be imported from outside the module, preserving encapsulation.
