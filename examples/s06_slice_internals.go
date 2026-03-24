//go:build ignore

// Section 6, Topic 48: Slice Internals - Header, Shared Backing Array
//
// A slice is a 3-word struct: {pointer, length, capacity}
// Multiple slices can share the same backing array.
//
// GOTCHA: Modifying one slice can affect another if they share a backing array.
// GOTCHA: append may or may not create a new backing array.
// GOTCHA: Re-slicing doesn't copy data.
//
// Run: go run examples/s06_slice_internals.go

package main

import "fmt"

func main() {
	fmt.Println("=== Slice Internals ===")
	fmt.Println()

	// 1. Slice header
	fmt.Println("-- Slice header --")
	s := make([]int, 3, 5)
	fmt.Printf("s: len=%d, cap=%d, data=%v\n", len(s), cap(s), s)

	// 2. Shared backing array
	fmt.Println("\n-- Shared backing array --")
	original := []int{10, 20, 30, 40, 50}
	slice1 := original[1:3] // [20, 30]
	slice2 := original[2:5] // [30, 40, 50]

	fmt.Println("original:", original)
	fmt.Println("slice1 [1:3]:", slice1)
	fmt.Println("slice2 [2:5]:", slice2)

	// Modify through slice1 - affects original and slice2!
	slice1[1] = 999 // changes original[2]
	fmt.Println("\nAfter slice1[1] = 999:")
	fmt.Println("original:", original) // [10, 20, 999, 40, 50]
	fmt.Println("slice1:", slice1)     // [20, 999]
	fmt.Println("slice2:", slice2)     // [999, 40, 50]

	// 3. Append may create new backing array
	fmt.Println("\n-- Append and capacity --")
	a := make([]int, 3, 3)
	a[0], a[1], a[2] = 1, 2, 3
	b := a[:2] // shares backing array

	fmt.Printf("Before: a=%v (cap=%d), b=%v (cap=%d)\n",
		a, cap(a), b, cap(b))

	b = append(b, 99) // cap exceeded -> new backing array
	fmt.Println("After append to b:")
	fmt.Println("  a:", a) // unchanged
	fmt.Println("  b:", b) // new backing array

	// 4. Full slice expression (prevent shared mutation)
	fmt.Println("\n-- Full slice expression [low:high:max] --")
	src := []int{1, 2, 3, 4, 5}
	safe := src[1:3:3] // cap limited to 2 (3-1)
	fmt.Printf("safe: %v, len=%d, cap=%d\n", safe, len(safe), cap(safe))
	safe = append(safe, 99) // forced to allocate new array
	fmt.Println("src after safe append:", src) // unchanged
}
