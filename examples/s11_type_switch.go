//go:build ignore

// Section 11, Topic 84: Type Switch
//
// Type switch is a special form of switch that compares types instead of values.
// Cleaner than chaining type assertions.
//
// Syntax:
//   switch v := i.(type) {
//   case int:    ...
//   case string: ...
//   default:     ...
//   }
//
// GOTCHA: `.(type)` can only be used inside a switch statement.
// GOTCHA: `v` in each case has the concrete type (not interface).
//
// Run: go run examples/s11_type_switch.go

package main

import "fmt"

func main() {
	fmt.Println("=== Type Switch ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. Basic type switch
	// ─────────────────────────────────────────────
	describe(42)
	describe("hello")
	describe(3.14)
	describe(true)
	describe(nil)
	describe([]int{1, 2, 3})

	// ─────────────────────────────────────────────
	// 2. Multiple types in one case
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Multiple types per case --")
	values := []any{42, int64(100), "hi", 3.14, float32(1.5)}
	for _, v := range values {
		switch v.(type) {
		case int, int64:
			fmt.Printf("  %v is an integer type\n", v)
		case float32, float64:
			fmt.Printf("  %v is a float type\n", v)
		case string:
			fmt.Printf("  %v is a string\n", v)
		}
	}

	// ─────────────────────────────────────────────
	// 3. Type switch with interface check
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Interface check --")
	type Stringer interface{ String() string }

	type MyError struct{ msg string }
	things := []any{
		42,
		&MyError{msg: "oops"}, // won't match Stringer since no String() method
		"hello",
	}
	for _, t := range things {
		switch v := t.(type) {
		case Stringer:
			fmt.Printf("  Stringer: %s\n", v.String())
		case string:
			fmt.Printf("  string: %q\n", v)
		default:
			fmt.Printf("  other: %v (%T)\n", v, v)
		}
	}
}

func describe(i any) {
	switch v := i.(type) {
	case int:
		fmt.Printf("  int: %d (doubled: %d)\n", v, v*2)
	case string:
		fmt.Printf("  string: %q (len: %d)\n", v, len(v))
	case float64:
		fmt.Printf("  float64: %.2f\n", v)
	case bool:
		fmt.Printf("  bool: %t\n", v)
	case nil:
		fmt.Println("  nil!")
	default:
		fmt.Printf("  unknown: %v (%T)\n", v, v)
	}
}
