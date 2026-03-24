//go:build ignore

// Section 6, Topic 42: Arrays — Fixed Size, Value Type
//
// Arrays in Go have a FIXED size that is part of the type:
//   [5]int and [10]int are different types!
//
// Arrays are VALUE TYPES — assignment copies the entire array.
// This is different from most languages where arrays are reference types.
//
// GOTCHA: Arrays are rarely used directly. Slices are preferred.
// GOTCHA: Array size is part of the type — [3]int != [4]int.
// GOTCHA: Passing arrays to functions copies the ENTIRE array.
//
// Run: go run examples/s06_arrays.go

package main

import "fmt"

func main() {
	fmt.Println("=== Arrays ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. Declaration and initialization
	// ─────────────────────────────────────────────
	var a [5]int // zero-valued: [0 0 0 0 0]
	fmt.Printf("Zero-valued: %v\n", a)

	b := [3]string{"Go", "Rust", "Python"}
	fmt.Printf("Initialized: %v\n", b)

	// Compiler counts the size with ...:
	c := [...]int{10, 20, 30, 40, 50}
	fmt.Printf("Auto-sized: %v (len=%d)\n", c, len(c))

	// Indexed initialization:
	d := [5]int{0: 10, 4: 50} // only set indices 0 and 4
	fmt.Printf("Indexed init: %v\n", d)

	// ─────────────────────────────────────────────
	// 2. Access and modify
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Access and modify --")
	a[0] = 100
	a[4] = 500
	fmt.Printf("a[0]=%d, a[4]=%d\n", a[0], a[4])
	fmt.Printf("len(a)=%d\n", len(a))

	// Out of bounds: compile-time error for constants, runtime panic for variables
	// a[5] = 1  // compile error: index 5 out of bounds
	// i := 5; a[i] = 1  // runtime panic: index out of range

	// ─────────────────────────────────────────────
	// 3. Array is a VALUE TYPE (copy on assignment)
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Value type (copy on assignment) --")
	original := [3]int{1, 2, 3}
	copied := original // FULL COPY
	copied[0] = 999
	fmt.Printf("Original: %v\n", original) // [1 2 3] — unchanged
	fmt.Printf("Copied:   %v\n", copied)   // [999 2 3]

	// ─────────────────────────────────────────────
	// 4. Array comparison
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Array comparison --")
	x := [3]int{1, 2, 3}
	y := [3]int{1, 2, 3}
	z := [3]int{1, 2, 4}
	fmt.Printf("[1,2,3] == [1,2,3]: %t\n", x == y)
	fmt.Printf("[1,2,3] == [1,2,4]: %t\n", x == z)

	// Different sizes can't be compared:
	// [3]int == [4]int  // compile error: mismatched types

	// ─────────────────────────────────────────────
	// 5. Passing arrays (copies the whole thing!)
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Passing arrays (copied) --")
	arr := [3]int{1, 2, 3}
	modifyArray(arr) // passes a COPY
	fmt.Printf("After modifyArray: %v (unchanged!)\n", arr)

	modifyArrayPtr(&arr) // passes a pointer
	fmt.Printf("After modifyArrayPtr: %v (modified)\n", arr)

	// ─────────────────────────────────────────────
	// 6. Iterating
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Iteration --")
	for i, v := range c {
		fmt.Printf("  [%d]=%d\n", i, v)
	}

	// ─────────────────────────────────────────────
	// When to use arrays vs slices
	// ─────────────────────────────────────────────
	// Arrays: fixed-size data, value semantics needed (crypto hashes, IP addresses)
	// Slices: everything else (dynamic size, passed by reference-like header)
	// In practice, slices are used 99% of the time.
}

func modifyArray(a [3]int) {
	a[0] = 999 // modifies the copy
}

func modifyArrayPtr(a *[3]int) {
	a[0] = 999 // modifies the original
}
