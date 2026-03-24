//go:build ignore

// Section 2, Topic 11: const keyword, untyped constants
//
// Constants are immutable values known at compile time.
// Untyped constants have higher precision and adapt to context.
//
// GOTCHA: Constants cannot be declared with :=.
// GOTCHA: Untyped constants can be used with any compatible type.
// GOTCHA: Constants must be compile-time evaluable (no function calls).
//
// Run: go run examples/s02_constants.go

package main

import (
	"fmt"
	"math"
)

// Package-level constants
const Pi = 3.14159265358979323846
const AppName = "LearnGo"

// Grouped constants
const (
	MaxRetries = 3
	Timeout    = 30 // seconds
)

// Typed vs untyped
const typedInt int = 42
const untypedInt = 42 // untyped: adapts to context

func main() {
	fmt.Println("=== Constants ===")
	fmt.Println()

	// 1. Basic usage
	fmt.Println("App:", AppName)
	fmt.Println("Pi:", Pi)
	fmt.Println("MaxRetries:", MaxRetries)

	// 2. Untyped constant adapts
	fmt.Println("\n-- Untyped constants --")
	var f32 float32 = untypedInt // works: untyped adapts to float32
	var f64 float64 = untypedInt // works: untyped adapts to float64
	var i64 int64 = untypedInt   // works: untyped adapts to int64
	fmt.Printf("float32: %f, float64: %f, int64: %d\n", f32, f64, i64)

	// Typed constant is strict:
	// var f float64 = typedInt // ERROR: cannot use int as float64
	_ = typedInt

	// 3. Constant expressions
	fmt.Println("\n-- Constant expressions --")
	const x = 10
	const y = 20
	const sum = x + y
	const product = x * y
	fmt.Printf("sum=%d, product=%d\n", sum, product)

	// 4. High-precision untyped constants
	fmt.Println("\n-- High precision --")
	const huge = 1e1000 // this is fine as a constant
	// var v float64 = huge // ERROR: overflows float64
	const small = huge / 1e999
	fmt.Println("huge/1e999 =", small) // 10

	// 5. Math constants
	fmt.Println("\n-- Math --")
	fmt.Println("math.Pi:", math.Pi)
	fmt.Println("math.E:", math.E)
	fmt.Println("math.MaxInt64:", math.MaxInt64)
}
