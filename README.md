# error-framework

<p align="center">
  <img src="https://raw.githubusercontent.com/krisalay/error-framework/main/docs/banner.png" width="640" alt="error-framework banner">
</p>

<p align="center">
  <b>Enterprise-grade centralized error management framework for Go.</b>
</p>

<p align="center">

[![Go Reference](https://pkg.go.dev/badge/github.com/krisalay/error-framework.svg)](https://pkg.go.dev/github.com/krisalay/error-framework)
[![CI](https://github.com/krisalay/error-framework/actions/workflows/ci.yml/badge.svg)](https://github.com/krisalay/error-framework/actions)
[![codecov](https://codecov.io/gh/krisalay/error-framework/branch/main/graph/badge.svg)](https://codecov.io/gh/krisalay/error-framework)
[![Go Report Card](https://goreportcard.com/badge/github.com/krisalay/error-framework)](https://goreportcard.com/report/github.com/krisalay/error-framework)
[![License](https://img.shields.io/github/license/krisalay/error-framework)](LICENSE)

</p>

---

# Overview

Go's built-in error handling is simple but insufficient for modern distributed systems.

Production services require:

- Structured error objects  
- Centralized error management  
- Consistent client responses  
- Safe handling of sensitive errors  
- Database error normalization  
- Validation error normalization  
- Structured logging  
- Trace ID correlation  
- Stack trace support  

**error-framework provides all of these in a clean, extendible, and production-grade architecture.**

---

# Features

- Centralized error management
- Structured error model
- Safe error handling for clients
- Error wrapping and chaining
- PostgreSQL (pgx) integration
- validator/v10 integration
- Echo framework middleware
- Zap logger integration
- Trace ID propagation
- Stack trace support
- Extendible adapter architecture
- High performance
- Production ready

---

# Installation

```bash
go get github.com/krisalay/error-framework@latest
```

# Quick Start
## Basic Usage

```go
package main

import (
    "errors"
    "github.com/krisalay/error-framework/framework"
)

func GetUser(id string) error {

    err := errors.New("database connection failed")

    return framework.Wrap(err, "failed to fetch user")
}
```
Client receives safe response:
```json
{
  "message": "Internal server error",
  "code": "INTERNAL_ERROR",
  "status": 500
}
```
Full error is logged internally with stack trace and trace ID.

# Error Model

Primary error structure:

```go
type AppError struct {
    Message    string
    Code       string
    Status     int
    Details    map[string]any
    TraceID    string
    Level      ErrorLevel
    StackTrace string
}
```
This ensures consistent error handling across your entire application.

# Database Error Handling (PostgreSQL pgx)
Automatically converts database errors into structured errors:

```go
err := db.Query(...)

return framework.DB(err)
```

Duplicate key example response:

```json
{
  "code": "DB_DUPLICATE_KEY",
  "status": 409,
  "message": "resource already exists"
}
```

# Validation Error Handling

Works with validator/v10:

```go
if err := validate.Struct(req); err != nil {
    return framework.Validation(err)
}
```

Response:
```json
{
  "code": "VALIDATION_ERROR",
  "status": 400,
  "details": {
    "email": "must be a valid email"
  }
}
```

# Echo Framework Integration
```go
e := echo.New()

e.Use(echoAdapter.Middleware(manager))
```

All errors are automatically handled, logged, and returned safely.

# Error Wrapping

Sensitive error (hidden from client):
```go
return framework.Wrap(err, "internal failure")
```

Safe error (visible to client):
```go
return framework.WrapSafe(err, "user not found")
```

# Logging Integration

Zap logger integration:

```go
logger := logging.NewZapLogger(config)
manager := core.NewManager(core.ManagerConfig{
    Logger: logger,
})
```

Logs include:
- stack trace
- trace ID
- error code
- error level
- timestamp

# Architecture
```
Application
    ↓
framework.Wrap()
    ↓
core.Manager
    ↓
Logger + StackTrace + TraceID
    ↓
Safe Client Response
```

# Comparison
|Feature|	error-framework | native errors | pkg/errors |
|---|---|---|---|
|Structured errors|	✓ | ✗ |	partial |
|Safe client responses|	✓ |	✗ |	✗ |
|DB error normalization| ✓ | ✗ | ✗ |
|Validation integration| ✓ | ✗ | ✗ |
|Stack traces|	✓|	✓|	✓|
|Logger integration|	✓|	✗|	✗|
|Trace ID support|	✓|	✗|	✗|

# Examples
See:
```
examples/
```
Includes:
- Echo REST API
- Validation example
- Database example


# Versioning

Uses Semantic Versioning:
```
v1.0.0
v1.1.0
v2.0.0
```

# Roadmap
Upcoming features:
- Gin adapter
- Fiber adapter
- gRPC adapter
- MongoDB adapter
- Redis adapter
- OpenTelemetry integration

# Contributing
See `CONTRIBUTING.md`

Pull requests welcome.

# License

MIT License

See LICENSE file.

# Production Ready
error-framework is designed for:
- microservices
- distributed systems
- REST APIs
- backend platforms
- enterprise Go services

# Star the repo

If you find this useful, please consider starring the repository.