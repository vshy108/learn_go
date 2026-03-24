//go:build ignore

// Section 5, Topic 39: Type Switch — Switching on Interface Types
//
// A type switch lets you branch based on the dynamic type of an interface value.
// Syntax: switch v := x.(type) { case int: ... case string: ... }
//
// This is Go's way of "pattern matching" on types, used heavily with
// interface{} (any) values.
//
// GOTCHA: The .(type) syntax ONLY works inside a switch statement.
// GOTCHA: nil is a valid case in a type switch.
//
// Run: go run examples/s05_type_switch.go

package main

import "fmt"

func main() {
	fmt.Println("=== Type Switch ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. Basic type switch
	// ─────────────────────────────────────────────
	fmt.Println("-- Basic type switch --")
	printType(42)
	printType("hello")
	printType(3.14)
	printType(true)
	printType(nil)
	printType([]int{1, 2, 3})

	// ─────────────────────────────────────────────
	// 2. Type switch with assignment
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Type switch with assignment --")
	describe(42)
	describe("hello")
	describe([]int{1, 2, 3})

	// ─────────────────────────────────────────────
	// 3. Multiple types in one case
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Multiple types per case --")
	isNumeric(42)
	isNumeric(3.14)
	isNumeric("hello")

	// ─────────────────────────────────────────────
	// 4. Practical: processing heterogeneous data
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Processing mixed data --")
	data := []any{42, "hello", 3.14, true, nil, []int{1, 2}}
	for _, item := range data {
		fmt.Printf("  %v → %s\n", item, classify(item))
	}
}

func printType(x any) {
	switch x.(type) {
	case int:
		fmt.Println("  int")
	case string:
		fmt.Println("  string")
	case float64:
		fmt.Println("  float64")
	case bool:
		fmt.Println("  bool")
	case nil:
		fmt.Println("  nil")
	default:
		fmt.Printf("  unknown: %T\n", x)
	}
}

func describe(x any) {
	// The `v` variable gets the concrete type in each case:
	switch v := x.(type) {
	case int:
		fmt.Printf("  int: %d (doubled: %d)\n", v, v*2)
	case string:
		fmt.Printf("  string: %q (length: %d)\n", v, len(v))
	case []int:
		fmt.Printf("  []int: %v (length: %d)\n", v, len(v))
	default:
		fmt.Printf("  unknown type: %T\n", v)
	}
}

func isNumeric(x any) {
	switch x.(type) {
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64,
		float32, float64:
		fmt.Printf("  %v (%T) is numeric\n", x, x)
	default:
		fmt.Printf("  %v (%T) is NOT numeric\n", x, x)
	}
}

func classify(x any) string {
	switch x.(type) {
	case nil:
		return "nil"
	case int, float64:
		return "number"
	case string:
		return "text"
	case bool:
		return "boolean"
	default:
		return fmt.Sprintf("other (%T)", x)
	}
}
