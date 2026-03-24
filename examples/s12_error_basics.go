//go:build ignore

// Section 12, Topic 90: Error Basics and Idiomatic Error Handling
//
// Go handles errors explicitly — no exceptions, no try/catch.
// Functions return error as the last return value.
// The caller MUST check it.
//
// Pattern:
//   result, err := doSomething()
//   if err != nil {
//       // handle error
//   }
//
// GOTCHA: Ignoring errors is a common source of bugs.
// GOTCHA: Always check errors before using the result.
//
// Run: go run examples/s12_error_basics.go

package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("=== Error Basics ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. Basic error handling pattern
	// ─────────────────────────────────────────────
	n, err := strconv.Atoi("42")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Parsed:", n)
	}

	n2, err := strconv.Atoi("not-a-number")
	if err != nil {
		fmt.Println("Error:", err) // Error: strconv.Atoi: parsing "not-a-number": invalid syntax
	} else {
		fmt.Println("Parsed:", n2)
	}

	// ─────────────────────────────────────────────
	// 2. errors.New — simplest error
	// ─────────────────────────────────────────────
	fmt.Println("\n-- errors.New --")
	err = errors.New("something went wrong")
	fmt.Println(err)

	// ─────────────────────────────────────────────
	// 3. fmt.Errorf — formatted error
	// ─────────────────────────────────────────────
	fmt.Println("\n-- fmt.Errorf --")
	filename := "data.csv"
	err = fmt.Errorf("cannot open %s: file not found", filename)
	fmt.Println(err)

	// ─────────────────────────────────────────────
	// 4. Returning errors from functions
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Function errors --")
	result, err := divide(10, 3)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("10 / 3 = %.2f\n", result)
	}

	result, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err) // division by zero
	} else {
		fmt.Printf("Result: %.2f\n", result)
	}

	// ─────────────────────────────────────────────
	// 5. Error is nil on success
	// ─────────────────────────────────────────────
	fmt.Println("\n-- nil means success --")
	if err := validate("hello"); err != nil {
		fmt.Println("Invalid:", err)
	} else {
		fmt.Println("Valid!")
	}

	if err := validate(""); err != nil {
		fmt.Println("Invalid:", err)
	}

	// ─────────────────────────────────────────────
	// Comparison: Go vs Rust
	// ─────────────────────────────────────────────
	// Go:   value, err := f(); if err != nil { ... }
	// Rust: match f() { Ok(v) => ..., Err(e) => ... }
	//       or: let v = f()?;  (? operator)
	// Both make error handling explicit. Rust's ? is more concise.
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func validate(s string) error {
	if s == "" {
		return errors.New("must not be empty")
	}
	return nil
}
