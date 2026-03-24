//go:build ignore

// Section 5, Topic 35: if/else — Initializer Statement, No Ternary
//
// Go's if/else is similar to other C-family languages but with key differences:
//   - No parentheses around the condition
//   - Braces { } are REQUIRED (even for single statements)
//   - Supports an initializer statement: if x := expr; condition { }
//   - NO ternary operator (no a ? b : c)
//
// GOTCHA: Opening brace MUST be on the same line (implicit semicolons).
// GOTCHA: The initializer variable is scoped to the entire if/else chain.
// GOTCHA: Conditions must be bool — no truthy/falsy values.
//
// Run: go run examples/s05_if_else.go

package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("=== if/else ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. Basic if/else
	// ─────────────────────────────────────────────
	x := 42
	if x > 0 {
		fmt.Println("Positive")
	} else if x < 0 {
		fmt.Println("Negative")
	} else {
		fmt.Println("Zero")
	}

	// ─────────────────────────────────────────────
	// 2. If with initializer statement
	// ─────────────────────────────────────────────
	fmt.Println("\n-- if with initializer --")
	// The variable `n` is scoped to the if/else block
	if n, err := strconv.Atoi("42"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Parsed:", n)
	}
	// fmt.Println(n)  // ERROR: n not declared in this scope

	// Common pattern: error checking
	if err := doSomething(); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Success!")
	}

	// ─────────────────────────────────────────────
	// 3. GOTCHA: No ternary operator
	// ─────────────────────────────────────────────
	fmt.Println("\n-- No ternary --")
	// Python: result = "yes" if True else "no"
	// Rust:   let result = if true { "yes" } else { "no" };
	// Go:     NO ternary. Use if/else:
	age := 20
	var status string
	if age >= 18 {
		status = "adult"
	} else {
		status = "minor"
	}
	fmt.Println("Status:", status)

	// Some people use a helper, but most Go code just uses if/else:
	fmt.Println("Status:", ternary(age >= 18, "adult", "minor"))

	// ─────────────────────────────────────────────
	// 4. GOTCHA: Conditions must be bool
	// ─────────────────────────────────────────────
	// if 1 { }        // ERROR: non-bool 1 (type int) used as condition
	// if "hello" { }  // ERROR: non-bool "hello" used as condition
	// if nil { }      // ERROR: non-bool nil used as condition
	// Unlike Python/JS/C, Go has NO truthy/falsy values.

	// ─────────────────────────────────────────────
	// 5. GOTCHA: Brace position
	// ─────────────────────────────────────────────
	// WRONG:
	// if true
	// {              // ERROR: unexpected semicolon before {
	// }
	//
	// CORRECT:
	// if true {      // brace on same line
	// }

	// ─────────────────────────────────────────────
	// 6. if as a statement (not an expression)
	// ─────────────────────────────────────────────
	// In Rust, if is an expression: let x = if true { 1 } else { 2 };
	// In Go, if is a statement: you can't assign its result directly.
	// You must declare the variable first, then assign inside the blocks.

	// ─────────────────────────────────────────────
	// 7. Nested if with initializers
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Nested initializers --")
	if a, err := strconv.Atoi("10"); err == nil {
		if b, err := strconv.Atoi("20"); err == nil {
			fmt.Printf("a=%d, b=%d, sum=%d\n", a, b, a+b)
		} else {
			fmt.Println("Error parsing b:", err)
		}
	}
}

func doSomething() error { return nil }

// Ternary helper (not idiomatic, but sometimes used)
func ternary(cond bool, a, b string) string {
	if cond {
		return a
	}
	return b
}
