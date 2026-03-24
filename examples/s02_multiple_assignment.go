//go:build ignore

// Section 2, Topic 13: Multiple Variable Assignment and Swapping
//
// Go supports multiple assignment in a single statement.
// This enables clean swaps without temp variables.
//
// Run: go run examples/s02_multiple_assignment.go

package main

import "fmt"

func main() {
	fmt.Println("=== Multiple Assignment ===")
	fmt.Println()

	// 1. Multiple declaration
	var a, b, c int = 1, 2, 3
	fmt.Printf("a=%d, b=%d, c=%d\n", a, b, c)

	// 2. Mixed types with var
	var (
		name = "Alice"
		age  = 30
		gpa  = 3.9
	)
	fmt.Printf("name=%s, age=%d, gpa=%.1f\n", name, age, gpa)

	// 3. Short declaration
	x, y, z := 10, 20, 30
	fmt.Printf("x=%d, y=%d, z=%d\n", x, y, z)

	// 4. Swap (no temp variable needed!)
	fmt.Println("\n-- Swap --")
	fmt.Printf("Before: x=%d, y=%d\n", x, y)
	x, y = y, x
	fmt.Printf("After:  x=%d, y=%d\n", x, y)

	// 5. Multiple return values
	fmt.Println("\n-- Multiple returns --")
	quotient, remainder := divmod(17, 5)
	fmt.Printf("17 divmod 5 = %d remainder %d\n", quotient, remainder)

	// 6. Ignore a return value with _
	q, _ := divmod(10, 3)
	fmt.Printf("10 / 3 = %d (remainder ignored)\n", q)
}

func divmod(a, b int) (int, int) {
	return a / b, a % b
}
