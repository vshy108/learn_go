//go:build ignore

// Section 2, Topic 15: The Blank Identifier (_)
//
// _ discards values. Used when you must accept a value but don't need it.
//
// Common uses:
//   - Ignore return values: _, err := func()
//   - Import for side effects: import _ "pkg"
//   - Verify interface implementation: var _ Interface = (*Type)(nil)
//   - Ignore loop index: for _, v := range slice
//
// GOTCHA: Ignoring errors with _ is a code smell.
//
// Run: go run examples/s02_blank_identifier.go

package main

import "fmt"

func main() {
	fmt.Println("=== Blank Identifier _ ===")
	fmt.Println()

	// 1. Ignore return value
	_, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	}

	result, _ := divide(10, 3) // ignore error (not recommended)
	fmt.Println("Result:", result)

	// 2. Ignore loop index
	fmt.Println("\n-- Ignore index --")
	fruits := []string{"apple", "banana", "cherry"}
	for _, fruit := range fruits {
		fmt.Println(" ", fruit)
	}

	// 3. Ignore loop value
	fmt.Println("\n-- Ignore value --")
	for i := range fruits {
		fmt.Printf("  index %d\n", i)
	}

	// 4. Multiple ignored values
	fmt.Println("\n-- Multiple ignores --")
	_, b, _ := threeValues()
	fmt.Println("Middle value:", b)

	// 5. Interface compliance check (compile-time)
	// var _ fmt.Stringer = (*MyType)(nil)
	fmt.Println("\n-- Side-effect import (concept) --")
	fmt.Println("import _ \"image/png\" // registers PNG decoder")
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

func threeValues() (int, string, bool) {
	return 1, "hello", true
}
