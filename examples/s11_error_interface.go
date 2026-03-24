//go:build ignore

// Section 11, Topic 87: error Interface
//
// The built-in error interface:
//   type error interface {
//       Error() string
//   }
//
// Any type with an Error() string method is an error.
// This is the foundation of Go's error handling.
//
// Run: go run examples/s11_error_interface.go

package main

import (
	"errors"
	"fmt"
)

// ─────────────────────────────────────────────
// Custom error type
// ─────────────────────────────────────────────
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation: %s — %s", e.Field, e.Message)
}

func validateAge(age int) error {
	if age < 0 {
		return &ValidationError{Field: "age", Message: "cannot be negative"}
	}
	if age > 150 {
		return &ValidationError{Field: "age", Message: "unrealistically high"}
	}
	return nil
}

func main() {
	fmt.Println("=== error Interface ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. errors.New — simplest error
	// ─────────────────────────────────────────────
	err := errors.New("something went wrong")
	fmt.Println("Error:", err)
	fmt.Printf("Type: %T\n", err)

	// ─────────────────────────────────────────────
	// 2. fmt.Errorf — formatted error
	// ─────────────────────────────────────────────
	fmt.Println("\n-- fmt.Errorf --")
	name := "file.txt"
	err = fmt.Errorf("failed to open %s: permission denied", name)
	fmt.Println(err)

	// ─────────────────────────────────────────────
	// 3. Custom error types
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Custom error --")
	if err := validateAge(-5); err != nil {
		fmt.Println("Error:", err)

		// Type assert to get details:
		if ve, ok := err.(*ValidationError); ok {
			fmt.Printf("  Field: %s\n", ve.Field)
			fmt.Printf("  Message: %s\n", ve.Message)
		}
	}

	if err := validateAge(30); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Age 30 is valid")
	}

	// ─────────────────────────────────────────────
	// 4. nil error means success
	// ─────────────────────────────────────────────
	fmt.Println("\n-- nil = success --")
	var e error // nil
	fmt.Printf("nil error: %v (is nil: %t)\n", e, e == nil)

	// ─────────────────────────────────────────────
	// 5. error is an interface
	// ─────────────────────────────────────────────
	// Any type with Error() string satisfies error.
	// This means you can carry arbitrary data in errors.
}
