//go:build ignore

// Section 2, Topic 13: Multiple Assignment and Swapping
//
// Go supports assigning multiple variables in a single statement.
// The classic use case is swapping values without a temporary variable.
//
// GOTCHA: All right-hand side expressions are evaluated BEFORE any assignment.
//         This is what makes swap work without a temp variable.
//
// Run: go run examples/s02_multiple_assignment.go

package main

import "fmt"

func main() {
	fmt.Println("=== Multiple Assignment ===")
	fmt.Println()

	// ─────────────────────────────────────────────







































































}	return a / b, a % bfunc divide(a, b int) (int, int) {}	// Rust:   std::mem::swap(&mut a, &mut b)  or  (a, b) = (b, a); [unstable]	// Python: a, b = b, a       (tuple unpacking)	// Go:     a, b = b, a       (built-in tuple assignment)	// ─────────────────────────────────────────────	// Comparison: Go vs Python vs Rust	// ─────────────────────────────────────────────	fmt.Printf("python: val=%d, ok=%t\n", val, ok) // val=0 (zero value), ok=false	val, ok = m["python"]	fmt.Printf("go: val=%d, ok=%t\n", val, ok)	val, ok := m["go"]	// Two-value assignment: value + ok boolean	m := map[string]int{"go": 2009, "rust": 2010}	fmt.Println("\n-- Comma-ok idiom --")	// ─────────────────────────────────────────────	// 5. Common pattern: comma-ok idiom	// ─────────────────────────────────────────────	fmt.Printf("arr=%v, i=%d\n", arr, i) // arr=[11 20 30], i=1	// Step 2: assign:       arr[0]=11, i=1	// Step 1: evaluate RHS: arr[0]+1=11, 0+1=1	arr[i], i = arr[i]+1, i+1	arr := [3]int{10, 20, 30}	// All RHS values are computed before any LHS assignment:	i := 0	fmt.Println("\n-- Evaluation order --")	// ─────────────────────────────────────────────	// 4. GOTCHA: Evaluation order	// ─────────────────────────────────────────────	fmt.Printf("Just the remainder: %d\n", remainder)	_, remainder := divide(17, 5)	// Discarding a return value with _:	fmt.Printf("17 / 5 = %d remainder %d\n", quot, rem)	quot, rem := divide(17, 5)	fmt.Println("\n-- Multiple returns --")	// ─────────────────────────────────────────────	// 3. Multiple return values from functions	// ─────────────────────────────────────────────	fmt.Printf("After:  p=%d, q=%d, r=%d\n", p, q, r)	p, q, r = q, r, p	fmt.Printf("Before: p=%d, q=%d, r=%d\n", p, q, r)	p, q, r := 1, 2, 3	// Three-way rotation:	fmt.Printf("After:  a=%d, b=%d\n", a, b)	a, b = b, a // all RHS evaluated first, then assigned	fmt.Printf("Before: a=%d, b=%d\n", a, b)	a, b := 10, 20	fmt.Println("\n-- Swap --")	// ─────────────────────────────────────────────	// 2. Swapping — no temp variable needed	// ─────────────────────────────────────────────	fmt.Printf("%s is %d with score %.1f\n", name, age, score)	name, age, score := "Alice", 30, 95.5	// Mixed types:	fmt.Printf("x=%d, y=%d, z=%d\n", x, y, z)	x, y, z := 1, 2, 3	// ─────────────────────────────────────────────	// 1. Multiple assignment