//go:build ignore

// Section 2, Topic 7: var keyword - explicit types and zero values
//
// var declares variables with explicit types. Uninitialized variables
// get the type's ZERO VALUE (not undefined/null like other languages).
//
// GOTCHA: var at package level is global - use sparingly.
// GOTCHA: Unused local variables are compile errors in Go.
//
// Run: go run examples/s02_var_declaration.go

package main

import "fmt"

// Package-level var (no := allowed here)
var packageVar = "I'm package-level"

func main() {
	fmt.Println("=== var Declaration ===")
	fmt.Println()

	// 1. Explicit type
	var name string = "Alice"
	var age int = 30
	var height float64 = 5.9
	fmt.Printf("name=%s, age=%d, height=%.1f\n", name, age, height)

	// 2. Type inferred from value
	var city = "Tokyo" // inferred as string
	var count = 42     // inferred as int
	fmt.Printf("city=%s (%T), count=%d (%T)\n", city, city, count, count)

	// 3. Zero values (no initialization)
	var zeroInt int
	var zeroFloat float64
	var zeroString string
	var zeroBool bool
	var zeroPtr *int
	fmt.Println("\n-- Zero values --")
	fmt.Printf("int: %d, float64: %f, string: %q, bool: %t, *int: %v\n",
		zeroInt, zeroFloat, zeroString, zeroBool, zeroPtr)

	// 4. Grouped declaration
	var (
		firstName = "Bob"
		lastName  = "Smith"
		score     = 95
	)
	fmt.Printf("\nGrouped: %s %s, score=%d\n", firstName, lastName, score)

	// 5. Multiple same-type on one line
	var x, y, z int = 1, 2, 3
	fmt.Printf("x=%d, y=%d, z=%d\n", x, y, z)

	// 6. Package-level
	fmt.Println("\nPackage var:", packageVar)
}
