# Contributing to errorframework

First of all, thank you for your interest in contributing to
**errorframework**. This project aims to provide an enterprise-grade,
scalable, and extensible error management framework for Go services.

This guide explains how to contribute in a way that maintains the
quality, stability, and reliability of the project.

------------------------------------------------------------------------

# Table of Contents

-   Code of Conduct
-   Contribution Philosophy
-   Ways to Contribute
-   Development Environment Setup
-   Project Architecture Overview
-   Branching Strategy
-   Coding Standards
-   Error Framework Design Principles
-   Writing Tests
-   Test Coverage Requirements
-   Running Tests
-   Running Linters
-   Benchmarking
-   Commit Message Standards
-   Pull Request Process
-   CI/CD Pipeline Requirements
-   Documentation Standards
-   Adding New Features
-   Adding New Adapters
-   Security Guidelines
-   Reporting Bugs
-   Suggesting Enhancements
-   Release Process
-   Maintainer Guidelines

------------------------------------------------------------------------

# Code of Conduct

This project adheres to professional and respectful collaboration.

Please:

-   Be respectful
-   Be constructive
-   Avoid personal criticism
-   Focus on technical improvements

------------------------------------------------------------------------

# Contribution Philosophy

We follow these principles:

-   Reliability first
-   Backward compatibility
-   High test coverage
-   Clear documentation
-   Modular design
-   SOLID principles

------------------------------------------------------------------------

# Ways to Contribute

You can contribute by:

-   Fixing bugs
-   Improving documentation
-   Writing tests
-   Adding adapters (database, framework, logging)
-   Improving performance
-   Refactoring code
-   Adding examples
-   Improving CI/CD

------------------------------------------------------------------------

# Development Environment Setup

## Requirements

-   Go 1.22 or newer
-   Git
-   Linux, macOS, or Windows

Verify Go version:

    go version

------------------------------------------------------------------------

## Clone Repository

    git clone https://github.com/krisalay/errorframework.git
    cd errorframework

------------------------------------------------------------------------

## Install dependencies

    go mod tidy

------------------------------------------------------------------------

# Project Architecture Overview

Project structure:

    core/        Core error types and manager
    framework/   Public helper APIs
    adapters/    Integrations (Echo, pgx, validator)
    logging/     Logger implementations
    utils/       Utilities
    config/      Configuration
    examples/    Usage examples
    docs/        Documentation

------------------------------------------------------------------------

# Branching Strategy

We use GitHub Flow:

main branch:

-   always stable
-   always releasable

Feature branch example:

    git checkout -b feat/add-mysql-adapter

Bugfix branch example:

    git checkout -b fix/validator-panic

------------------------------------------------------------------------

# Coding Standards

Follow official Go standards:

    go fmt ./...
    go vet ./...

Use golangci-lint:

    golangci-lint run

Code must:

-   compile without warnings
-   pass all tests
-   pass lint checks

------------------------------------------------------------------------

# Error Framework Design Principles

Contributions must respect:

-   centralized error management
-   immutable error objects
-   structured logging
-   safe client error messages
-   extensibility
-   framework independence

------------------------------------------------------------------------

# Writing Tests

All contributions must include tests.

Test file naming:

    *_test.go

Use table-driven tests:

Example:

    func TestWrapSafe(t *testing.T) {
        tests := []struct {
            name string
            err error
        }{
            {"nil error", nil},
            {"normal error", errors.New("test")},
        }

        for _, tt := range tests {
            t.Run(tt.name, func(t *testing.T) {
                result := WrapSafe(tt.err, "message")
                if tt.err != nil && result == nil {
                    t.Fatal("expected error")
                }
            })
        }
    }

------------------------------------------------------------------------

# Test Coverage Requirements

Minimum coverage:

    80%

Target coverage:

    90%+

Check coverage:

    go test ./... -coverprofile=coverage.out
    go tool cover -func=coverage.out

------------------------------------------------------------------------

# Running Tests

Run all tests:

    go test ./...

Run with race detection:

    go test ./... -race

------------------------------------------------------------------------

# Benchmarking

Run benchmarks:

    go test -bench=.

------------------------------------------------------------------------

# Commit Message Standards

Use Conventional Commits:

Examples:

    feat: add mysql adapter
    fix: resolve pgx duplicate key handling
    docs: update README
    test: add validation adapter tests
    refactor: improve error manager

Breaking change:

    feat!: redesign error builder

------------------------------------------------------------------------

# Pull Request Process

Step 1: Fork repository

Step 2: Create branch

    git checkout -b feat/new-feature

Step 3: Make changes

Step 4: Add tests

Step 5: Run checks

    go test ./...
    golangci-lint run

Step 6: Commit

    git commit -m "feat: add new feature"

Step 7: Push

    git push origin feat/new-feature

Step 8: Open Pull Request

Include:

-   description
-   motivation
-   test coverage

------------------------------------------------------------------------

# CI/CD Requirements

All PRs must pass:

-   build
-   tests
-   coverage
-   lint

GitHub Actions enforces this automatically.

------------------------------------------------------------------------

# Documentation Standards

All exported functions must include GoDoc comments:

Example:

    // WrapSafe wraps an error with a safe message
    func WrapSafe(err error, message string) *AppError

------------------------------------------------------------------------

# Adding New Features

Steps:

-   discuss via issue
-   implement feature
-   add tests
-   update documentation

------------------------------------------------------------------------

# Adding New Adapters

Adapters must:

-   be isolated
-   include tests
-   follow adapter pattern

Example:

    adapters/mysql/

------------------------------------------------------------------------

# Security Guidelines

Do NOT expose:

-   internal errors
-   stack traces to clients
-   database sensitive details

------------------------------------------------------------------------

# Reporting Bugs

Open GitHub issue and include:

-   Go version
-   OS
-   reproduction steps
-   expected behavior
-   actual behavior

------------------------------------------------------------------------

# Suggesting Enhancements

Open GitHub issue with:

-   use case
-   problem description
-   proposed solution

------------------------------------------------------------------------

# Release Process

Maintainers create releases:

    v1.0.0
    v1.1.0
    v1.1.1

------------------------------------------------------------------------

# Maintainer Guidelines

Maintainers must:

-   review PRs
-   ensure quality
-   maintain documentation
-   maintain releases

------------------------------------------------------------------------

Thank you for contributing!
