# Learn Go

A structured, example-driven Go learning repository — one concept per file,
with detailed comments covering usage, edge cases, and gotchas.

## Quick Start

```sh
# Run any example
go run examples/s01_hello_world.go

# Run with race detector (for concurrency examples)
go run -race examples/s13_goroutine_basics.go

# Vet a file for common mistakes
go vet examples/s01_hello_world.go
```

## Structure

See [PLAN.md](PLAN.md) for the full topic roadmap.

Each file in `examples/` is a standalone `package main` with its own `main()` function.
Files use `//go:build ignore` so they don't conflict when the directory is compiled as a package.
Files are prefixed with `s{NN}_` to group them by section.

## Sections

| Section | Topics |
|---------|--------|
| 01 | Basics |
| 02 | Variables and Constants |
| 03 | Data Types |
| 04 | Functions |
| 05 | Control Flow |
| 06 | Arrays and Slices |
| 07 | Maps |
| 08 | Strings |
| 09 | Structs |
| 10 | Pointers |
| 11 | Interfaces |
| 12 | Error Handling |
| 13 | Goroutines |
| 14 | Channels |
| 15 | Packages and Modules |
| 16 | Generics (Go 1.18+) |
| 17 | Testing |
| 18 | Standard Library Highlights |
| 19 | Advanced Patterns |
| 20 | Reflection and Unsafe |
