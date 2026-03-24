//go:build ignore

// Section 6, Topic 46: copy() Function
//
// `copy` copies elements from a source slice to a destination slice.
// It returns the number of elements copied (min of len(dst), len(src)).
//
// GOTCHA: copy() does NOT allocate memory. dst must already have enough length.
// GOTCHA: copy() copies min(len(dst), len(src)) elements.
//
// Run: go run examples/s06_copy.go

package main

import "fmt"

func main() {
	fmt.Println("=== copy() Function ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. Basic copy
	// ─────────────────────────────────────────────
	src := []int{1, 2, 3, 4, 5}
	dst := make([]int, len(src))
	n := copy(dst, src)
	fmt.Printf("Copied %d elements: %v\n", n, dst)

	// Modify dst — src is NOT affected (independent):
	dst[0] = 999
	fmt.Printf("src=%v, dst=%v\n", src, dst)

	// ─────────────────────────────────────────────
	// 2. Copy to smaller destination
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Smaller destination --")
	small := make([]int, 3)
	n = copy(small, src) // copies only 3 elements
	fmt.Printf("Copied %d: %v\n", n, small)

	// ─────────────────────────────────────────────
	// 3. Copy from larger source
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Partial copy --")
	partial := make([]int, 2)
	copy(partial, src[2:]) // copies src[2:] = [3,4,5] → only [3,4]
	fmt.Printf("Partial: %v\n", partial)

	// ─────────────────────────────────────────────
	// 4. Overlapping copy (same slice)
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Overlapping copy --")
	s := []int{0, 1, 2, 3, 4}
	copy(s[1:], s[:4]) // shift right by 1
	fmt.Printf("Shift right: %v\n", s)

	// ─────────────────────────────────────────────
	// 5. Copy strings (string → []byte)
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Copy string to byte slice --")
	bytes := make([]byte, 5)
	n = copy(bytes, "Hello, World!") // copies first 5 bytes
	fmt.Printf("Copied %d bytes: %s\n", n, bytes)

	// ─────────────────────────────────────────────
	// 6. Deep copy pattern (create independent copy)
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Deep copy pattern --")
	original := []int{1, 2, 3, 4, 5}
	clone := make([]int, len(original))
	copy(clone, original)
	clone[0] = 999
	fmt.Printf("Original: %v (unchanged)\n", original)
	fmt.Printf("Clone:    %v\n", clone)

	// Alternative using append:
	clone2 := append([]int(nil), original...) // also creates independent copy
	clone2[0] = 888
	fmt.Printf("Clone2:   %v\n", clone2)
	fmt.Printf("Original: %v (still unchanged)\n", original)
}
