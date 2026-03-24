//go:build ignore

// Section 2, Topic 10: Type Inference Rules
//
// Go infers types from the right-hand side of := and var x = expr.
//
// GOTCHA: Numeric literals default to int (not int32/int64).
// GOTCHA: Float literals default to float64 (not float32).
// GOTCHA: Untyped constants have a "default type" used during inference.
//
// Run: go run examples/s02_type_inference.go

package main

import "fmt"

func main() {
	fmt.Println("=== Type Inference ===")
	fmt.Println()

	// 1. Integer literals -> int
	a := 42
	fmt.Printf("42 -> %T\n", a) // int

	// 2. Float literals -> float64
	b := 3.14
	fmt.Printf("3.14 -> %T\n", b) // float64

	// 3. String literal -> string
	c := "hello"
	fmt.Printf("\"hello\" -> %T\n", c) // string

	// 4. Bool literal -> bool
	d := true
	fmt.Printf("true -> %T\n", d) // bool

	// 5. Rune literal -> int32
	e := 'A'
	fmt.Printf("'A' -> %T (rune = int32)\n", e)

	// 6. Complex literal -> complex128
	f := 1 + 2i
	fmt.Printf("1+2i -> %T\n", f)

	// 7. From function return type
	g := sum(3, 4)
	fmt.Printf("sum(3,4) -> %T\n", g)

	// 8. From expression
	h := a * 2 // int * untyped int -> int
	fmt.Printf("a * 2 -> %T\n", h)

	// 9. Mixed: float64 wins over int
	i := float64(a) + b
	fmt.Printf("float64(a) + b -> %T\n", i)
}

func sum(a, b int) int {
	return a + b
}
