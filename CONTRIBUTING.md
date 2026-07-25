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

This project adheres to the [Contributor Covenant Code of Conduct](CODE_OF_CONDUCT.md).
By participating, you are expected to uphold this code.

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

Follow the patterns already used under [`es/`](es/). When in doubt, copy an existing similar query or aggregation file.

### Go Version

This project uses **Go 1.18** as the minimum version to ensure broad compatibility. Do not use features from newer Go versions. Generics (introduced in Go 1.18) are used throughout the API.

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

### Package Layout

| Path | Package | Role |
|------|---------|------|
| `es/` | `es` | Main DSL: queries, aggregations, sort, highlight, etc. (flat package) |
| `es/condition/` | `condition` | Conditional helpers (`If`, `IfElse`, `ElseIf`, `Else`) |
| `es/enums/<kebab-case>/` | combined lowercase (e.g. `operator`, `validationmethod`) | Elasticsearch string constants |

Do not introduce new top-level packages under `es/` unless they match these existing roles.

### File Naming

| Kind | Pattern | Example |
|------|---------|---------|
| Query | `{snake}_query.go` + `_test.go` | `match_none_query.go` |
| Aggregation | `aggregation_{name}.go` + `_test.go` | `aggregation_avg.go` |
| Shared helpers | descriptive snake_case | `types.go`, `base_query.go`, `generic_put_in_the_field.go` |
| Enum | `es/enums/{kebab-case}/{snake}.go` | `es/enums/operator/operator.go` |
| White-box tests | `*_private_test.go` | `generic_put_in_the_field_private_test.go` |

Some non-query builders omit `_query` in the filename when that matches existing style (`range.go`, `sort.go`, `script.go`, `highlight.go`, `inner_hits.go`). Prefer `{name}_query.go` for new query types.

### Code Style Guidelines

#### Line Length
- Maximum **140 characters** per line
- Break long lines at logical points

#### Function Length
- Maximum **100 lines** per function
- If a function exceeds this, consider refactoring into smaller functions

#### Type Model

Query and aggregation types are **aliases of `Object`**, not structs:

```go
type Object map[string]any
type Array []any

type termType Object      // unexported query type
type matchNoneType Object // unexported query type
```

Do not introduce new struct types for query builders.

#### Naming Conventions

**Types:**
```go
// Most query/aggregation types are unexported: camelCase + Type suffix
type termType Object
type matchNoneType Object
type avgAggType Object

// Exported exceptions used by the public API
type Object map[string]any
type Array []any
type BoolType Object
type FilterType Array
type MustType Array
type MustNotType Array
type ShouldType Array

// Good: descriptive, matches Elasticsearch concept
type termType Object
// Bad: abbreviations or fake exported query types
type TQ Object
type TermQueryType Object
```

**Functions:**
```go
// Constructors are exported PascalCase; return the unexported (or Bool) type
func Bool() BoolType { ... }
func Term[T any](key string, value T) termType { ... }
func MatchNone() matchNoneType { ... }

// Option methods are verbs / Elasticsearch parameter names; return the same type
func (t termType) Boost(boost float64) termType { ... }
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

All exported functions and types must have documentation comments with **Example usage**, **Parameters** (when useful), and **Returns**. Refer to types by their real names (`es.termType`, `es.BoolType`):

```go
// Term creates a new es.termType object with the specified key-value pair.
//
// Example usage:
//
//	t := es.Term("category", "books")
//	// t now contains an es.termType object with a term query for the "category" field.
//
// Parameters:
//   - key: A string representing the field name for the term query.
//   - value: The value to be searched for in the specified field. The type is generic.
//
// Returns:
//
//	An es.termType object containing the specified term query.
func Term[T any](key string, value T) termType {
	return termType{
		"term": Object{
			key: Object{
				"value": value,
			},
		},
	}
}
```

Unexported helpers may use a short comment or none (follow neighboring code).

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

#### Option Fields (`putInTheField`)

Option methods write into the nested Elasticsearch object via a private `putInTheField` helper that delegates to generics in `generic_put_in_the_field.go`:

- `genericPutInTheField` — key under a named parent (e.g. `"match_none"`)
- `genericPutInTheFieldOfFirstChild` — key on the first child object (e.g. term field object)
- `genericPutInTheFieldOfFirstObject` — key on the first object in the root (e.g. sort, query_string)

```go
func (m matchNoneType) Boost(boost float64) matchNoneType {
	return m.putInTheField("boost", boost)
}

func (m matchNoneType) putInTheField(key string, value any) matchNoneType {
	return genericPutInTheField(m, "match_none", key, value)
}
```

#### Conditional Constructors (`XFunc` / `XIf`)

Many queries expose `{Name}Func` and `{Name}If` constructors that return `nil` when the condition is false. Nil / typed-nil clauses are filtered by `correctType` (no errors):

```go
func TermIf[T any](key string, value T, condition bool) termType {
	if !condition {
		return nil
	}
	return Term(key, value)
}

func TermFunc[T any](key string, value T, f func(key string, value T) bool) termType {
	if !f(key, value) {
		return nil
	}
	return Term(key, value)
}
```

When adding a new field-based query, include `Func` / `If` variants if sibling queries already have them.

#### Error Handling

Query builders in `es/` **do not return `error` and must not panic**. Invalid or nil clauses are skipped (via `correctType`) or result in an empty/partial query object. Contract correctness is enforced by tests that assert exact JSON output.

#### Enums

Place Elasticsearch string constants under `es/enums/<kebab-case>/`:

```go
package operator

type Operator string

const (
	Or  Operator = "or"
	And Operator = "and"
)

func (operator Operator) String() string {
	return string(operator)
}
```

Import with an alias equal to the exported type name:

```go
import (
	Operator "github.com/Trendyol/es-query-builder/es/enums/operator"
)

// usage: Operator.And
```

#### Condition Package

Use `es/condition` for conditional clause selection without branching in call sites:

- `condition.If(item, condition)`
- `condition.IfElse(condition, item, branches...)`
- `condition.ElseIf(condition, item)` / `condition.Else(item)`

#### Adding a New Query

Use existing files as templates:

| Complexity | Template |
|------------|----------|
| Simple (no field) | `es/match_none_query.go` + `_test.go` |
| Field + options | `es/term_query.go` + `_test.go` |
| Variadic children | `es/dis_max_query.go` + `_test.go` |

Checklist:

1. Add `es/{name}_query.go` with `package es`, `type {name}Type Object`, exported constructor, option methods, and `putInTheField`.
2. Add `{Name}Func` / `{Name}If` when appropriate.
3. Add `es/{name}_query_test.go` (`package es_test`): exist, create type, JSON asserts, method coverage.
4. Add enums under `es/enums/...` if new string constants are needed; import with a type-name alias.
5. Do not add external dependencies.

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

- Black-box API tests: same directory, `_test.go` suffix, **`package es_test`**
- White-box tests for unexported helpers: `*_private_test.go`, **`package es`**

```go
package es_test

import (
	"testing"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/Trendyol/es-query-builder/test/assert"
)
```

Group related tests with a section header comment:

```go
////   Term   ////
```

### Test Naming Convention

Use descriptive names with underscores:

```go
// Patterns used in es/:
//   Test_<Symbol>_should_<behavior>
//   Test_<Symbol>_<Method>_should_<behavior>
//   Test_<Symbol>_method_should_create_<typeName>

func Test_Term_should_exist_on_es_package(t *testing.T) { ... }
func Test_Term_method_should_create_termType(t *testing.T) { ... }
func Test_Term_CaseInsensitive_should_create_json_with_case_insensitive_field_inside_term(t *testing.T) { ... }
func Test_MatchNone_method_should_create_matchNoneType(t *testing.T) { ... }
```

### Test Structure

Use the Given-When-Then pattern. Assert real type names (`es.termType`, `es.BoolType`, `es.matchNoneType`):

```go
func Test_Term_method_should_create_termType(t *testing.T) {
	t.Parallel()
	// Given
	b := es.Term("key", "value")

	// Then
	assert.NotNil(t, b)
	assert.IsTypeString(t, "es.termType", b)
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

Always verify JSON serialization for query builders with an exact string:

```go
func Test_Term_should_create_json_with_term_field_inside_query(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.Term("key", "value"),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"term\":{\"key\":{\"value\":\"value\"}}}}", bodyJSON)
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
