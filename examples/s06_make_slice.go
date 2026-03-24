//go:build ignore

// Section 6, Topic 44: make() for Slices
//
// `make` creates slices (and maps and channels) with a specified length
// and optional capacity. Elements are zero-valued.
//
// Syntax: make([]T, length)  or  make([]T, length, capacity)
//
// GOTCHA: make([]int, 5) creates [0,0,0,0,0] with len=5.
//         If you then append, you get [0,0,0,0,0,newValue] — len=6!
// GOTCHA: make([]int, 0, 5) creates empty slice with cap=5.
//         append works as expected: [newValue] — len=1.
//
// Run: go run examples/s06_make_slice.go

package main

import "fmt"

func main() {
	fmt.Println("=== make() for Slices ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. make with length only
	// ─────────────────────────────────────────────
	s1 := make([]int, 5) // len=5, cap=5, all zeros
	fmt.Printf("make([]int, 5): %v (len=%d, cap=%d)\n", s1, len(s1), cap(s1))

	// ─────────────────────────────────────────────
	// 2. make with length and capacity
	// ─────────────────────────────────────────────
	s2 := make([]int, 3, 10) // len=3, cap=10
	fmt.Printf("make([]int, 3, 10): %v (len=%d, cap=%d)\n", s2, len(s2), cap(s2))

	// ─────────────────────────────────────────────
	// 3. make with length=0 (for append pattern)
	// ─────────────────────────────────────────────
	s3 := make([]int, 0, 10) // len=0, cap=10
	fmt.Printf("make([]int, 0, 10): %v (len=%d, cap=%d)\n", s3, len(s3), cap(s3))

	// ─────────────────────────────────────────────
	// 4. GOTCHA: make(len=5) + append = surprise!
	// ─────────────────────────────────────────────
	fmt.Println("\n-- GOTCHA: make length + append --")
	s := make([]int, 5) // [0 0 0 0 0]
	s = append(s, 99)   // [0 0 0 0 0 99] — appended AFTER the zeros!
	fmt.Printf("make(5) + append(99): %v (len=%d)\n", s, len(s))

	// If you want to append without pre-filled zeros:
	s = make([]int, 0, 5) // [] with capacity 5
	s = append(s, 99)     // [99]
	fmt.Printf("make(0,5) + append(99): %v (len=%d)\n", s, len(s))

	// ─────────────────────────────────────────────
	// 5. Pre-allocating for known size (performance)
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Pre-allocating --")
	// When you know the size, pre-allocate to avoid resizing:
	result := make([]int, 0, 100)
	for i := 0; i < 100; i++ {
		result = append(result, i*i)
	}
	fmt.Printf("Pre-allocated: len=%d, cap=%d\n", len(result), cap(result))

	// Or fill directly via index:
	result2 := make([]int, 100)
	for i := 0; i < 100; i++ {
		result2[i] = i * i
	}
	fmt.Printf("Index-filled: len=%d, cap=%d\n", len(result2), cap(result2))

	// ─────────────────────────────────────────────
	// 6. make vs literal vs new
	// ─────────────────────────────────────────────
	fmt.Println("\n-- make vs literal vs new --")
	a := make([]int, 0) // empty, not nil
	b := []int{}        // empty, not nil
	var c []int         // nil
	d := new([]int)     // *[]int pointing to nil slice

	fmt.Printf("make:    nil=%t, len=%d\n", a == nil, len(a))
	fmt.Printf("literal: nil=%t, len=%d\n", b == nil, len(b))
	fmt.Printf("var:     nil=%t, len=%d\n", c == nil, len(c))
	fmt.Printf("new:     nil=%t, len=%d\n", *d == nil, len(*d))
}
