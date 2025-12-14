# Testing

Standard library package for writing tests. Test functions start with `Test` and take `*testing.T` parameter. Use `t.Error()`, `t.Fatal()` for failures. Test files end with `_test.go`. Run with 'go test'. Supports benchmarks and examples

## Table-driven Tests

Table-driven tests use slices of test cases to test multiple scenarios with the same logic. Each case contains inputs and expected outputs. Makes adding test cases easy and provides comprehensive coverage with minimal code duplication.

## Mocks and Stubs

Replace dependencies with controlled implementations for isolated testing. Stubs provide predefined responses while mocks verify method calls. Go's interfaces make mocking natural. Essential for testing without external dependencies.

## httptest for HTTP Tests

The `httptest` package provides utilities for testing HTTP servers and clients without network connections. Includes `httptest.Server`, `ResponseRecorder`, and helpers for creating test requests. Essential for testing handlers, middleware, and HTTP services.

## Benchmarks

Measure code performance by timing repeated executions. Functions start with `Benchmark` and use `*testing.B` parameter. Run with `go test -bench=.` to identify bottleneck, compare implementations, and track performance changes over time.
