# Contributing to es-query-builder

Thank you for your interest in contributing to es-query-builder! This document provides guidelines and requirements for contributing to the project.

## Table of Contents

- [Code of Conduct](#code-of-conduct)
- [Getting Started](#getting-started)
- [Development Setup](#development-setup)
- [Coding Standards](#coding-standards)
- [Testing Requirements](#testing-requirements)
- [Pull Request Process](#pull-request-process)
- [Reporting Issues](#reporting-issues)

## Code of Conduct

By participating in this project, you agree to maintain a respectful and inclusive environment. Please be considerate of others and focus on constructive collaboration.

## Getting Started

1. Fork the repository
2. Clone your fork locally
3. Create a new branch for your feature or bug fix
4. Make your changes
5. Submit a pull request

## Development Setup

### Prerequisites

- Go 1.18 or later
- Make (optional, but recommended)

### Installing Development Tools

Run the following command to install all required development tools:

```bash
make init
```

This installs:
- `golangci-lint` - Linter for Go code
- `fieldalignment` - Struct field alignment optimizer
- `gotestfmt` - Test output formatter
- `go-run-bench` - Benchmark runner
- `go-carpet` - Code coverage tool

### Building the Project

```bash
go build -v ./...
```

### Running Tests

```bash
# Run all tests with race detection
make run-test

# Run tests with pretty output
make unit-test-pretty

# Generate coverage report
make coverage-html
```

## Coding Standards

### Go Version

This project uses **Go 1.18** as the minimum version to ensure broad compatibility. Do not use features from newer Go versions.

### Zero External Dependencies

This project maintains a **zero external dependencies** policy for the main module. All functionality must be implemented using only the Go standard library. This ensures:
- Minimal attack surface
- No dependency vulnerabilities
- Easy integration for users

### Linting Rules

We use `golangci-lint` with a strict configuration. Before submitting, run:

```bash
make linter
```

#### Enabled Linters

| Linter | Purpose |
|--------|---------|
| `bodyclose` | Checks whether HTTP response body is closed |
| `errcheck` | Checks for unchecked errors |
| `exhaustive` | Checks exhaustiveness of enum switch statements |
| `funlen` | Limits function length (max 100 lines) |
| `goconst` | Finds repeated strings that could be constants |
| `gocritic` | Provides various code analysis checks |
| `gocyclo` | Checks cyclomatic complexity |
| `gosimple` | Suggests code simplifications |
| `govet` | Reports suspicious constructs |
| `gosec` | Security-focused linting |
| `ineffassign` | Detects ineffective assignments |
| `lll` | Limits line length (max 140 characters) |
| `misspell` | Finds commonly misspelled words |
| `nakedret` | Checks for naked returns |
| `gofumpt` | Stricter gofmt |
| `staticcheck` | Advanced static analysis |
| `stylecheck` | Style checks |
| `typecheck` | Type checking |
| `unconvert` | Removes unnecessary conversions |
| `unparam` | Reports unused function parameters |
| `unused` | Checks for unused code |
| `whitespace` | Checks for unnecessary whitespace |

### Code Style Guidelines

#### Line Length
- Maximum **140 characters** per line
- Break long lines at logical points

#### Function Length
- Maximum **100 lines** per function
- If a function exceeds this, consider refactoring into smaller functions

#### Naming Conventions

**Types:**
```go
// Use PascalCase for exported types
type BoolType Object
type FilterType Array

// Use descriptive names that indicate purpose
type TermQueryType Object  // Good
type TQ Object              // Bad
```

**Functions:**
```go
// Constructor functions should be named after what they create
func Bool() BoolType { ... }
func Term(field string, value any) TermType { ... }

// Method names should be verbs or descriptive actions
func (b BoolType) Filter(items ...any) BoolType { ... }
func (b BoolType) MinimumShouldMatch(value any) BoolType { ... }
```

**Variables:**
```go
// Use camelCase for local variables
minimumShouldMatch := 2
adjustPureNegative := true

// Use short names for loop variables
for i := 0; i < len(items); i++ { ... }
```

#### Documentation

All exported functions and types must have documentation comments:

```go
// Bool creates and returns an empty BoolType object.
//
// This function is typically used to initialize an es.BoolType, which can be
// populated later with the appropriate boolean query conditions.
//
// Example usage:
//
//	b := es.Bool()
//	// b is now an empty es.BoolType object that can be used in a query.
//
// Returns:
//
//	An empty es.BoolType object.
func Bool() BoolType {
    return BoolType{}
}
```

Documentation should include:
- A brief description of what the function does
- Example usage (when helpful)
- Parameter descriptions (for complex functions)
- Return value description

#### Fluent API Pattern

This library uses a fluent/builder pattern. Methods should:
- Return the modified type to allow chaining
- Accept variadic parameters where appropriate

```go
// Good: Returns the modified type for chaining
func (b BoolType) Filter(items ...any) BoolType {
    // ... implementation
    return b
}

// Usage: Allows method chaining
query := es.Bool().
    Must(es.Term("field", "value")).
    Filter(es.Exists("field"))
```

### Struct Field Alignment

Optimize struct field alignment for memory efficiency:

```bash
make fixfieldalignment
```

## Testing Requirements

### Test Coverage

- All new code must have tests
- Aim for high test coverage (check with `make coverage`)
- Tests must pass with race detection enabled

### Test File Organization

- Test files must be in the same package with `_test` suffix
- Use the `es_test` package name for black-box testing

```go
package es_test

import (
    "testing"
    
    "github.com/Trendyol/es-query-builder/es"
    "github.com/Trendyol/es-query-builder/test/assert"
)
```

### Test Naming Convention

Use descriptive test names with underscores:

```go
// Pattern: Test_<Type>_<Method>_should_<expected_behavior>

func Test_Bool_should_exist_on_es_package(t *testing.T) { ... }
func Test_Bool_method_should_create_boolType(t *testing.T) { ... }
func Test_Bool_MinimumShouldMatch_should_create_json_with_int_minimum_should_match_field_inside_bool(t *testing.T) { ... }
```

### Test Structure

Use the Given-When-Then pattern:

```go
func Test_Bool_method_should_create_boolType(t *testing.T) {
    t.Parallel()
    // Given
    b := es.Bool()

    // Then
    assert.NotNil(t, b)
    assert.IsTypeString(t, "es.BoolType", b)
}
```

### Parallel Tests

All tests should run in parallel when possible:

```go
func Test_Example(t *testing.T) {
    t.Parallel()
    // test code
}
```

### Custom Assert Package

Use the project's custom assert package instead of external testing libraries:

```go
import "github.com/Trendyol/es-query-builder/test/assert"

// Available assertions:
assert.Equal(t, expected, actual)
assert.NotNil(t, value)
assert.Nil(t, value)
assert.True(t, condition)
assert.False(t, condition)
assert.IsType(t, expected, actual)
assert.IsTypeString(t, "expectedType", actual)
assert.MarshalWithoutError(t, body)
```

### JSON Output Verification

Always verify JSON serialization for query builders:

```go
func Test_Bool_Filter_should_create_correct_json(t *testing.T) {
    t.Parallel()
    // Given
    query := es.NewQuery(
        es.Bool().
            Filter(es.Term("id", 12345)),
    )

    // When Then
    assert.NotNil(t, query)
    bodyJSON := assert.MarshalWithoutError(t, query)
    assert.Equal(t, "{\"query\":{\"bool\":{\"filter\":[{\"term\":{\"id\":{\"value\":12345}}}]}}}", bodyJSON)
}
```

## Pull Request Process

### Before Submitting

1. **Run linter**: `make linter`
2. **Run tests**: `make run-test`
3. **Check coverage**: `make coverage`
4. **Update documentation** if needed

### PR Requirements

- [ ] Code follows the coding standards
- [ ] All tests pass
- [ ] New code has appropriate test coverage
- [ ] Documentation is updated (if applicable)
- [ ] Linter passes without errors
- [ ] No external dependencies added
- [ ] Commit messages are clear and descriptive

### Commit Message Guidelines

Write clear, concise commit messages:

```
Add MinimumShouldMatch method to BoolType

- Implement MinimumShouldMatch for boolean queries
- Add comprehensive tests for int and string values
- Update documentation with usage examples
```

### Review Process

1. Create a pull request against the `main` branch
2. Ensure all CI checks pass
3. Wait for maintainer review
4. Address any feedback
5. Once approved, a maintainer will merge your PR

## Reporting Issues

### Bug Reports

When reporting a bug, please include:

1. Go version (`go version`)
2. es-query-builder version
3. Minimal code to reproduce the issue
4. Expected behavior
5. Actual behavior
6. Any error messages

### Feature Requests

For feature requests, please:

1. Check if the feature already exists or is planned
2. Describe the use case
3. Provide example usage if possible
4. Explain why this would benefit other users

### Security Vulnerabilities

For security issues, please follow our [Security Policy](SECURITY.md).

## Questions?

If you have questions about contributing, feel free to open an issue for discussion.

Thank you for contributing to es-query-builder!
