//go:build ignore

// Section 6, Topic 43: Slice Basics — Creation, len, cap
//
// Slices are Go's primary sequence type — dynamic-length views into arrays.
// A slice has three components: pointer, length, and capacity.
//
// Think of a slice as: struct { ptr *T; len int; cap int }
//
// GOTCHA: A slice is NOT an array. It's a reference to an underlying array.
// GOTCHA: Multiple slices can share the same backing array.
// GOTCHA: len() is how many elements; cap() is the total available space.
//
// Run: go run examples/s06_slice_basics.go

package main

import "fmt"

func main() {
	fmt.Println("=== Slice Basics ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. Creating slices
	// ─────────────────────────────────────────────
	fmt.Println("-- Creating slices --")

	// Literal:
	s1 := []int{1, 2, 3, 4, 5}
	fmt.Printf("Literal:   %v (len=%d, cap=%d)\n", s1, len(s1), cap(s1))

	// From array:
	arr := [5]int{10, 20, 30, 40, 50}
	s2 := arr[1:4] // elements at index 1, 2, 3
	fmt.Printf("From arr:  %v (len=%d, cap=%d)\n", s2, len(s2), cap(s2))

	// Empty slice:
	s3 := []int{}
	fmt.Printf("Empty:     %v (len=%d, cap=%d, nil=%t)\n", s3, len(s3), cap(s3), s3 == nil)

	// Nil slice:
	var s4 []int
	fmt.Printf("Nil:       %v (len=%d, cap=%d, nil=%t)\n", s4, len(s4), cap(s4), s4 == nil)

	// ─────────────────────────────────────────────
	// 2. Slicing syntax
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Slicing syntax --")
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("s[2:5] = %v (indices 2,3,4)\n", s[2:5])
	fmt.Printf("s[:3]  = %v (first 3)\n", s[:3])
	fmt.Printf("s[7:]  = %v (from index 7)\n", s[7:])
	fmt.Printf("s[:]   = %v (full copy reference)\n", s[:])

	// ─────────────────────────────────────────────
	// 3. len vs cap
	// ─────────────────────────────────────────────
	fmt.Println("\n-- len vs cap --")
	base := make([]int, 3, 10) // len=3, cap=10
	fmt.Printf("make([]int, 3, 10): len=%d, cap=%d, val=%v\n", len(base), cap(base), base)

	// len: number of elements currently in the slice
	// cap: number of elements the backing array can hold starting from the slice's pointer
	sub := s[2:5]
	fmt.Printf("s[2:5]: len=%d, cap=%d\n", len(sub), cap(sub))
	// cap = len(s) - 2 = 8 (can extend rightward to the end of s)

	// ─────────────────────────────────────────────
	// 4. GOTCHA: Slices share backing arrays
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Shared backing array --")
	original := []int{1, 2, 3, 4, 5}
	slice := original[1:3] // [2, 3]
	slice[0] = 999
	fmt.Printf("Slice: %v\n", slice)       // [999 3]
	fmt.Printf("Original: %v\n", original) // [1 999 3 4 5] — modified!

	// ─────────────────────────────────────────────
	// 5. Three-index slice (controls capacity)
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Three-index slice [low:high:max] --")
	data := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s5 := data[2:5:7] // elements 2,3,4; cap limited to 7-2=5
	fmt.Printf("data[2:5:7]: %v (len=%d, cap=%d)\n", s5, len(s5), cap(s5))
	// This prevents the slice from seeing elements beyond index 7.

	// ─────────────────────────────────────────────
	// Comparison: Go slices vs Rust slices
	// ─────────────────────────────────────────────
	// Go:   slice = {ptr, len, cap} — can grow via append
	// Rust: &[T] = {ptr, len}       — fixed view, can't grow
	// Rust: Vec<T> = {ptr, len, cap} — similar to Go slice + append
}
