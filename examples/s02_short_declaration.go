//go:build ignore

// Section 2, Topic 8: Short variable declaration (:=)
//
// := declares AND initializes a variable with type inference.
// Only works inside functions (not at package level).
//
// GOTCHA: := requires at least one NEW variable on the left side.
// GOTCHA: := in a new scope creates a new variable (shadowing).
//
// Run: go run examples/s02_short_declaration.go

package main

import "fmt"

func main() {
	fmt.Println("=== Short Declaration := ===")
	fmt.Println()

	// 1. Basic usage
	name := "Alice"
	age := 30
	height := 5.9
	fmt.Printf("name=%s (%T), age=%d (%T), height=%.1f (%T)\n",
		name, name, age, age, height, height)

	// 2. Multiple variables
	x, y := 10, 20
	fmt.Printf("x=%d, y=%d\n", x, y)

	// 3. Redeclaration: at least one new variable required
	x, z := 100, 300 // x is reassigned, z is new
	fmt.Printf("x=%d, z=%d\n", x, z)

	// 4. Shadowing in inner scope
	fmt.Println("\n-- Shadowing --")
	val := "outer"
	fmt.Println("Before block:", val)
	{
		val := "inner" // new variable, shadows outer
		fmt.Println("Inside block:", val)
	}
	fmt.Println("After block:", val) // still "outer"

	// 5. Cannot use := at package level
	// shortVar := "error" // would not compile outside func

	// 6. Common pattern with error
	result, err := divide(10, 3)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("10 / 3 = %.2f\n", result)
	}

	// Reuse err (only result2 is new):
	result2, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("10 / 0 = %.2f\n", result2)
	}
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}
