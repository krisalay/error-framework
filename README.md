# Go Error Framework

Production‑grade centralized error management framework for Go services with built‑in support for:

* Echo framework
* Zap structured logging
* PostgreSQL (pgx)
* validator/v10
* Trace ID propagation
* Stack traces
* Panic recovery
* Sensitive error protection
* Fully modular and extendible architecture

This framework is designed using SOLID principles and enterprise‑grade design patterns.

---

# Table of Contents

* Overview
* Why this framework
* Architecture
* Installation
* Quick Start
* Configuration
* Core Concepts
* Error Object Structure
* Error Levels
* Error Codes
* Framework Helpers
* Echo Integration
* Validator Integration
* Database Integration (pgx)
* Error Wrapping
* Panic Recovery
* Trace ID
* Logging
* Response Format
* Best Practices
* Extending the Framework
* Example Usage
* FAQ

---

# Overview

This framework provides centralized, structured, and safe error management for Go microservices and backend systems.

It solves common problems like:

* inconsistent error formats
* leaking sensitive internal errors
* poor logging
* lack of traceability
* missing stack traces
* duplicated error handling logic

---

# Why this Framework

Traditional Go error handling:

```go
return err
```

Problems:

* no error codes
* no structured logging
* no stack trace
* no trace ID
* unsafe client exposure

Framework provides:

```go
return framework.DB(err)
return framework.Validation(err)
return framework.Wrap(err, "failed to fetch user")
```

With:

* structured error object
* centralized logging
* automatic trace IDs
* safe client responses

---

# Architecture

```
errorframework/
│
├── core/            # Core error types and manager
├── framework/       # Public helper functions
├── adapters/        # Framework integrations
│   ├── echo/
│   ├── pgx/
│   └── validator/
├── logging/         # Zap logger integration
├── utils/           # Stacktrace and trace provider
├── config/          # Configuration
```

Design Patterns Used:

* Builder Pattern
* Adapter Pattern
* Facade Pattern
* Strategy Pattern
* Dependency Injection

---

# Installation

```bash
go get github.com/krisalay/error-framework
```

---

# Quick Start

Initialize framework:

```go
cfg := config.Config{
    Logger: config.LoggerConfig{
        ConsoleEnabled: true,
        Level: "debug",
        Encoding: "console",
    },
    Trace: config.TraceConfig{ Enabled: true },
    StackTrace: config.StackTraceConfig{ Enabled: true },
    Database: config.DatabaseConfig{ Type: "pgx" },
    Validator: config.ValidatorConfig{ Enabled: true },
}

manager, err := framework.InitFromConfig(cfg)
```

Echo integration:

```go
e := echo.New()

e.Use(echoadapter.TraceMiddleware(traceProvider))
e.Use(echoadapter.PanicMiddleware(manager))
e.HTTPErrorHandler = echoadapter.NewHandler(manager).Handle
```

---

# Configuration

```go
type Config struct {
    Logger LoggerConfig
    Trace TraceConfig
    StackTrace StackTraceConfig
    Database DatabaseConfig
    Validator ValidatorConfig
}
```

LoggerConfig:

```go
type LoggerConfig struct {
    ConsoleEnabled bool
    FileEnabled bool
    FilePath string
    Level string
    Encoding string
}
```

---

# Core Concepts

## AppError

Central error structure:

```go
type AppError struct {
    Message string
    Code string
    Status int
    Details map[string]any

    Level ErrorLevel
    Err error
    IsSensitive bool

    Timestamp time.Time
    StackTrace string
    TraceID string
}
```

---

# Error Levels

```go
LevelDebug
LevelInfo
LevelWarn
LevelError
LevelFatal
```

Used for logging severity.

---

# Error Codes

Examples:

```
INTERNAL_ERROR
VALIDATION_ERROR
DB_ERROR
NOT_FOUND
ALREADY_EXISTS
```

Benefits:

* machine readable
* monitoring
* analytics

---

# Framework Helpers

## Database Error

```go
return framework.DB(err)
```

## Validation Error

```go
return framework.Validation(err)
```

## Internal Error

```go
return framework.Internal(err)
```

## Wrap Error

Sensitive:

```go
return framework.Wrap(err, "failed to fetch user")
```

Safe:

```go
return framework.WrapSafe(err, "user service unavailable")
```

## Custom Errors

```go
return framework.NotFound("user not found")
return framework.AlreadyExists("user exists")
```

---

# Echo Integration

Middleware:

```go
e.Use(echoadapter.TraceMiddleware(traceProvider))
e.Use(echoadapter.PanicMiddleware(manager))
e.HTTPErrorHandler = echoadapter.NewHandler(manager).Handle
```

Handlers:

```go
func handler(c echo.Context) error {
    return framework.NotFound("user not found")
}
```

---

# Validator Integration

```go
if err := validator.Struct(req); err != nil {
    return framework.Validation(err)
}
```

Response:

```json
{
  "code": "VALIDATION_ERROR",
  "details": {
    "email": "must be valid"
  }
}
```

---

# Database Integration (pgx)

```go
if err != nil {
    return framework.DB(err)
}
```

Handles:

* duplicate keys
* foreign keys
* connection errors

---

# Error Wrapping

Add context:

```go
return framework.Wrap(err, "failed to create user")
```

Preserves:

* original error
* stack trace
* trace ID

---

# Panic Recovery

Echo panic recovery built‑in.

Goroutine recovery:

```go
go func() {
    defer framework.Recover(ctx)
    panic("failure")
}()
```

---

# Trace ID

Automatically:

* generated
* propagated
* returned in response

Header supported:

```
X-Trace-ID
X-Request-ID
traceparent
```

---

# Logging

Uses Uber Zap.

Structured logs:

```json
{
  "message": "Database error",
  "code": "DB_ERROR",
  "trace_id": "abc"
}
```

---

# Response Format

Standard:

```json
{
  "message": "Validation failed",
  "code": "VALIDATION_ERROR",
  "status": 400,
  "trace_id": "abc",
  "details": {}
}
```

Sensitive error:

```json
{
  "message": "Internal server error",
  "code": "INTERNAL_ERROR"
}
```

---

# Best Practices

Use framework helpers:

```
framework.DB()
framework.Validation()
framework.Wrap()
```

Avoid:

```
return err
```

---

# Extending Framework

Add new DB adapter:

```
adapters/mysql/
```

Implement:

```
FromError(err error)
```

---

# Example Usage

Service:

```go
if err != nil {
    return framework.DB(err)
}
```

Handler:

```go
return err
```

Framework handles rest.

---

# FAQ

## Does it leak sensitive errors?

No.

## Is it production ready?

Yes.

## Can I use with other frameworks?

Yes. Echo adapter is included, others can be added.

---

# Conclusion

This framework provides:

* centralized error handling
* structured logging
* safe client responses
* traceability
* production‑grade architecture

Recommended for all Go backend services.

---

---

# Sequence Diagrams of Error Flow

## HTTP Request Error Flow

```
Client
  |
  | HTTP Request
  v
Echo Router
  |
  v
Handler
  |
  | returns error
  v
Framework Helper (framework.DB / Validation / Wrap)
  |
  v
AppError
  |
  v
Error Manager
  |  ├─ attach trace_id
  |  ├─ attach stacktrace
  |  └─ log via zap
  v
Echo Error Handler
  |
  v
Client Response (safe)
```

---

## Panic Flow

```
Handler / Goroutine
  |
  | panic()
  v
Recover Middleware / framework.Recover()
  |
  v
Manager.HandlePanic()
  |
  v
Logger
  |
  v
Safe Response
```

---

# Advanced Examples

---

## Microservice Example

Service Layer:

```go
func GetUser(ctx context.Context, id string) (*User, error) {

    user, err := repo.GetUser(ctx, id)

    if err != nil {
        return nil, framework.DB(err)
    }

    if user == nil {
        return nil, framework.NotFound("user not found")
    }

    return user, nil
}
```

Handler:

```go
func GetUserHandler(c echo.Context) error {

    user, err := service.GetUser(c.Request().Context(), "123")

    if err != nil {
        return err
    }

    return c.JSON(200, user)
}
```

---

## Worker Example

```go
func StartWorker(ctx context.Context) {

    defer framework.Recover(ctx)

    err := processJob()

    if err != nil {
        framework.Wrap(err, "worker failed")
    }
}
```

---

## Cron Job Example

```go
func DailyCron(ctx context.Context) {

    defer framework.Recover(ctx)

    err := runBilling()

    if err != nil {
        framework.Wrap(err, "billing cron failed")
    }
}
```

---

# OpenTelemetry Integration Guide

Framework supports OpenTelemetry via TraceProvider.

Example custom provider:

```go
import "go.opentelemetry.io/otel/trace"

func NewOTelTraceProvider() core.TraceProvider {

    return &OTelProvider{}
}

type OTelProvider struct{}

func (p *OTelProvider) GetTraceID(ctx context.Context) string {

    span := trace.SpanFromContext(ctx)

    return span.SpanContext().TraceID().String()
}
```

Initialize:

```go
manager := core.NewManager(core.ManagerConfig{
    Logger: logger,
    TraceProvider: NewOTelTraceProvider(),
})
```

Now logs automatically include OpenTelemetry trace IDs.

---

# Packaging Instructions (Publishing as Go Module)

## Step 1: Initialize Module

```bash
go mod init github.com/yourorg/errorframework
```

---

## Step 2: Add Dependencies

```bash
go mod tidy
```

---

## Step 3: Version the Module

```bash
git init
git add .
git commit -m "initial release"
git tag v1.0.0
```

---

## Step 4: Push to GitHub

```bash
git remote add origin https://github.com/krisalay/error-framework.git
git push origin main
git push origin v1.0.0
```

---

## Step 5: Use in Another Project

```bash
go get github.com/krisalay/error-framework@v1.0.0
```

---

# Recommended Versioning Strategy

Semantic Versioning:

```
v1.0.0  initial release
v1.1.0  new features
v1.1.1  bug fixes
v2.0.0  breaking changes
```

---

# Recommended Production Setup

```
errorframework/
  core/
  adapters/
  logging/
  framework/
  config/
  utils/
  README.md
```

---

# Final Summary

Framework provides enterprise-grade error management with:

* centralized handling
* structured logging
* stack traces
* trace propagation
* panic recovery
* modular architecture

Suitable for:

* microservices
* monoliths
* workers
* cron jobs
* distributed systems

---
