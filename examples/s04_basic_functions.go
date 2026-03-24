//go:build ignore

// Section 4, Topic 26: Basic Functions — Declaration, Parameters, Return Values
//
// Functions in Go are declared with the `func` keyword. Parameters are typed
// after the name (like Pascal), not before (like C).
//
// Syntax: func name(param1 type1, param2 type2) returnType { ... }
//
// GOTCHA: All parameters are PASSED BY VALUE (copies). To modify the original,
//         pass a pointer. Slices, maps, and channels are "reference-like" because
//         their values contain internal pointers, but they're still passed by value.
// GOTCHA: Go has no default parameter values. Use variadic or functional options instead.
// GOTCHA: Go has no function overloading. Each function name must be unique in a package.
//
// Run: go run examples/s04_basic_functions.go

package main

import "fmt"

// ─────────────────────────────────────────────
// 1. Basic function with parameters and return
// ─────────────────────────────────────────────
func add(a int, b int) int {
	return a + b
}

// Consecutive parameters of the same type can share the type:
func multiply(a, b int) int {
	return a * b
}

// ─────────────────────────────────────────────
// 2. No return value
// ─────────────────────────────────────────────
func greet(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

// ─────────────────────────────────────────────
// 3. Pass by value demonstration
// ─────────────────────────────────────────────
func tryToModify(x int) {
	x = 999 // modifies the COPY, not the original
}

func modifyWithPointer(x *int) {
	*x = 999 // modifies the original through the pointer
}

// ─────────────────────────────────────────────
// 4. Slices are "reference-like" (header is copied, but points to same array)
// ─────────────────────────────────────────────
func modifySlice(s []int) {
	if len(s) > 0 {
		s[0] = 999 // modifies the underlying array!
	}
}

func main() {
	fmt.Println("=== Basic Functions ===")
	fmt.Println()

	// Basic calls
	fmt.Println("add(2, 3) =", add(2, 3))
	fmt.Println("multiply(4, 5) =", multiply(4, 5))
	greet("Gopher")

	// Pass by value
	fmt.Println("\n-- Pass by value --")
	n := 42
	tryToModify(n)
	fmt.Println("After tryToModify:", n) // still 42

	modifyWithPointer(&n)
	fmt.Println("After modifyWithPointer:", n) // 999

	// Slices appear to pass by reference (but it's a value copy of the header)
	fmt.Println("\n-- Slice modification --")
	nums := []int{1, 2, 3}
	modifySlice(nums)
	fmt.Println("After modifySlice:", nums) // [999 2 3]

	// ─────────────────────────────────────────────
	// 5. Functions must be called with exact number of args
	// ─────────────────────────────────────────────
	// add(1)       // ERROR: not enough arguments
	// add(1, 2, 3) // ERROR: too many arguments

	// ─────────────────────────────────────────────
	// 6. No default parameters
	// ─────────────────────────────────────────────
	// func greet(name string, greeting string = "Hello") // ERROR: syntax error
	// Workaround: use variadic functions or functional options pattern.

	// ─────────────────────────────────────────────
	// 7. No function overloading
	// ─────────────────────────────────────────────
	// func add(a, b float64) float64 { ... } // ERROR: add already declared
	// Workaround: use different names (addInt, addFloat) or generics.

	// ─────────────────────────────────────────────
	// Comparison: Go vs Rust
	// ─────────────────────────────────────────────
	// Go:   func add(a, b int) int { return a + b }
	// Rust: fn add(a: i32, b: i32) -> i32 { a + b }
	// Go:   parameters after name, return type at end
	// Rust: parameters after name, return type after ->
	// Go:   all pass by value (use *T for references)
	// Rust: pass by value (move) or borrow (&T, &mut T)
}
