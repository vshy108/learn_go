//go:build ignore

// Section 10, Topic 74: Pointer Basics — & and * Operators
//
// A pointer holds the memory address of a value.
//   &x  — address of x (creates a pointer)
//   *p  — dereference p (access the value at address)
//
// GOTCHA: Go has pointers but NO pointer arithmetic (unlike C).
// GOTCHA: Zero value of a pointer is nil.
// GOTCHA: Go's garbage collector handles memory — no free/delete needed.
//
// Run: go run examples/s10_pointer_basics.go

package main

import "fmt"

func main() {
	fmt.Println("=== Pointer Basics ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. Creating pointers
	// ─────────────────────────────────────────────
	x := 42
	p := &x // p is *int, holds address of x
	fmt.Printf("x  = %d\n", x)
	fmt.Printf("&x = %p (address)\n", &x)
	fmt.Printf("p  = %p (same address)\n", p)
	fmt.Printf("*p = %d (dereferenced value)\n", *p)

	// ─────────────────────────────────────────────
	// 2. Modifying through pointer
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Modify via pointer --")
	*p = 100
	fmt.Printf("x after *p=100: %d\n", x) // 100 — modified through pointer

	// ─────────────────────────────────────────────
	// 3. Pointer type
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Pointer types --")
	var ip *int
	var sp *string
	var fp *float64
	fmt.Printf("*int:     %v (nil=%t)\n", ip, ip == nil)
	fmt.Printf("*string:  %v (nil=%t)\n", sp, sp == nil)
	fmt.Printf("*float64: %v (nil=%t)\n", fp, fp == nil)

	// ─────────────────────────────────────────────
	// 4. Passing pointers to functions
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Function parameters --")
	n := 10
	fmt.Printf("Before: %d\n", n)
	increment(&n)
	fmt.Printf("After increment: %d\n", n) // 11

	// Compare with value parameter:
	tryIncrement(n)
	fmt.Printf("After tryIncrement: %d (unchanged)\n", n) // still 11

	// ─────────────────────────────────────────────
	// 5. Returning pointers (safe in Go!)
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Returning pointers --")
	p2 := createInt(42)
	fmt.Printf("Returned pointer: %d\n", *p2)
	// In C, returning &localVar is dangerous (stack reference).
	// In Go, the compiler detects this and allocates on the heap ("escape analysis").

	// ─────────────────────────────────────────────
	// 6. Pointer to pointer
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Pointer to pointer --")
	val := 42
	p1 := &val
	pp := &p1 // **int
	fmt.Printf("val=%d, *p1=%d, **pp=%d\n", val, *p1, **pp)
}

func increment(n *int) {
	*n++
}

func tryIncrement(n int) {
	n++ // modifies local copy
}

func createInt(val int) *int {
	return &val // safe! Go heap-allocates when needed
}
