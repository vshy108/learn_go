//go:build ignore

// Section 10, Topic 77: No Pointer Arithmetic in Go
//
// Unlike C/C++, Go does NOT allow pointer arithmetic.
// You cannot do p++, p+1, or p[i] on a pointer.
//
// This is a deliberate safety choice — prevents buffer overflows
// and makes garbage collection possible.
//
// unsafe.Pointer can bypass this, but is strongly discouraged.
//
// Run: go run examples/s10_no_pointer_arithmetic.go

package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("=== No Pointer Arithmetic ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. These are ILLEGAL in Go:
	// ─────────────────────────────────────────────
	x := 42
	p := &x
	fmt.Printf("p = %p, *p = %d\n", p, *p)

	// p++        // ERROR: invalid operation
	// p + 1      // ERROR: invalid operation
	// p[0]       // ERROR: invalid operation
	// p - q      // ERROR: invalid operation

	// ─────────────────────────────────────────────
	// 2. Use slices instead of pointer arithmetic
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Use slices instead --")
	arr := [5]int{10, 20, 30, 40, 50}
	// In C: int *p = arr; *(p+2) gives 30
	// In Go: use a slice
	s := arr[:]
	fmt.Printf("s[2] = %d (no pointer arithmetic needed)\n", s[2])

	// ─────────────────────────────────────────────
	// 3. unsafe.Pointer (escape hatch — avoid!)
	// ─────────────────────────────────────────────
	fmt.Println("\n-- unsafe.Pointer (DO NOT use in normal code) --")
	// This is how you CAN do pointer arithmetic, but shouldn't:
	data := [3]int{100, 200, 300}
	ptr := unsafe.Pointer(&data[0])
	intSize := unsafe.Sizeof(data[0])

	// Access data[1] via pointer arithmetic:
	ptr1 := unsafe.Pointer(uintptr(ptr) + intSize)
	val := *(*int)(ptr1)
	fmt.Printf("data[1] via unsafe: %d\n", val)

	// This is UNSAFE because:
	// - No bounds checking
	// - Can break with different architectures
	// - GC may move memory (uintptr is not tracked by GC)
	// - Violates Go's memory safety guarantees

	// ─────────────────────────────────────────────
	// Comparison: Go vs C vs Rust
	// ─────────────────────────────────────────────
	// C:    Full pointer arithmetic (dangerous, powerful)
	// Go:   No pointer arithmetic (safe, use slices)
	// Rust: No pointer arithmetic in safe code (raw pointers in unsafe blocks)
}
