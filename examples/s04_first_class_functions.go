//go:build ignore

// Section 4, Topic 30: Functions as Values and Types
//
// In Go, functions are first-class citizens:
//   - Assign to variables
//   - Pass as arguments
//   - Return from functions
//   - Store in data structures
//
// Function types describe the signature: func(int, int) int
//
// GOTCHA: You cannot compare functions (except to nil).
// GOTCHA: Function types must match exactly (parameter and return types).
//
// Run: go run examples/s04_first_class_functions.go

package main

import (
	"fmt"
	"strings"
)

// ─────────────────────────────────────────────
// 1. Function type definition
// ─────────────────────────────────────────────
type MathFunc func(int, int) int
type Transformer func(string) string

func add(a, b int) int      { return a + b }
func subtract(a, b int) int { return a - b }
func multiply(a, b int) int { return a * b }

func main() {
	fmt.Println("=== First-Class Functions ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. Assign function to variable
	// ─────────────────────────────────────────────
	fmt.Println("-- Function assigned to variable --")
	var op MathFunc = add
	fmt.Printf("op(3, 5) = %d (using add)\n", op(3, 5))

	op = multiply
	fmt.Printf("op(3, 5) = %d (using multiply)\n", op(3, 5))

	// Without type alias:
	fn := subtract // inferred as func(int, int) int
	fmt.Printf("fn(10, 3) = %d\n", fn(10, 3))

	// ─────────────────────────────────────────────
	// 2. Functions as parameters
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Functions as parameters --")
	fmt.Println("apply(add, 3, 5) =", apply(add, 3, 5))
	fmt.Println("apply(multiply, 3, 5) =", apply(multiply, 3, 5))

	// ─────────────────────────────────────────────
	// 3. Functions as return values
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Functions as return values --")
	adder := makeAdder(10)
	fmt.Println("makeAdder(10)(5) =", adder(5))   // 15
	fmt.Println("makeAdder(10)(20) =", adder(20)) // 30

	// ─────────────────────────────────────────────
	// 4. Slice of functions
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Slice of functions --")
	operations := []MathFunc{add, subtract, multiply}
	names := []string{"add", "subtract", "multiply"}
	for i, op := range operations {
		fmt.Printf("  %s(10, 3) = %d\n", names[i], op(10, 3))
	}

	// ─────────────────────────────────────────────
	// 5. Map of functions
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Map of functions --")
	ops := map[string]MathFunc{
		"+": add,
		"-": subtract,
		"*": multiply,
	}
	for symbol, fn := range ops {
		fmt.Printf("  10 %s 3 = %d\n", symbol, fn(10, 3))
	}

	// ─────────────────────────────────────────────
	// 6. Standard library function values
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Standard library function values --")
	transform := []Transformer{
		strings.ToUpper,
		strings.ToLower,
		strings.Title,
	}
	for _, t := range transform {
		fmt.Printf("  %q\n", t("hello world"))
	}

	// ─────────────────────────────────────────────
	// 7. Nil function check
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Nil function check --")
	var nilFn MathFunc
	fmt.Printf("nilFn == nil: %t\n", nilFn == nil)
	// nilFn(1, 2)  // PANIC: nil function call
}

func apply(fn MathFunc, a, b int) int {
	return fn(a, b)
}

func makeAdder(base int) func(int) int {
	return func(x int) int {
		return base + x // closes over `base`
	}
}
