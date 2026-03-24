//go:build ignore

// Section 6, Topic 45: append() — Behavior and Growth Strategy
//
// `append` adds elements to a slice and returns a new slice.
// If the backing array is full, append allocates a new, larger array.
//
// GOTCHA: append may or may not allocate a new backing array. ALWAYS use
//         the return value: s = append(s, val)
// GOTCHA: Growth factor is roughly 2x for small slices, less for large ones.
// GOTCHA: When capacity grows, the new slice no longer shares the old array.
//
// Run: go run examples/s06_append.go

package main

import "fmt"

func main() {
	fmt.Println("=== append() ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. Basic append
	// ─────────────────────────────────────────────
	s := []int{1, 2, 3}
	fmt.Printf("Before: %v (len=%d, cap=%d)\n", s, len(s), cap(s))
	s = append(s, 4) // MUST reassign!
	fmt.Printf("After:  %v (len=%d, cap=%d)\n", s, len(s), cap(s))

	// Append multiple:
	s = append(s, 5, 6, 7)
	fmt.Printf("Multi:  %v (len=%d, cap=%d)\n", s, len(s), cap(s))

	// Append another slice:
	more := []int{8, 9, 10}
	s = append(s, more...) // spread operator
	fmt.Printf("Slice:  %v (len=%d, cap=%d)\n", s, len(s), cap(s))

	// ─────────────────────────────────────────────
	// 2. Growth strategy observation
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Growth strategy --")
	var g []int
	prevCap := 0
	for i := 0; i < 20; i++ {
		g = append(g, i)
		if cap(g) != prevCap {
			fmt.Printf("  len=%2d, cap=%2d (grew from %d)\n", len(g), cap(g), prevCap)
			prevCap = cap(g)
		}
	}
	// Typical pattern: 0→1→2→4→8→16→32 (roughly doubles)

	// ─────────────────────────────────────────────
	// 3. GOTCHA: Append to nil slice works fine
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Append to nil --")
	var nilSlice []int
	nilSlice = append(nilSlice, 1, 2, 3)
	fmt.Printf("Nil → %v\n", nilSlice)

	// ─────────────────────────────────────────────
	// 4. GOTCHA: Append may or may not reallocate
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Reallocation check --")
	a := make([]int, 3, 5) // has room (cap=5, len=3)
	a[0], a[1], a[2] = 1, 2, 3
	b := append(a, 4) // fits in existing capacity
	fmt.Printf("a=%v, b=%v\n", a, b)
	b[0] = 999
	fmt.Printf("After b[0]=999: a=%v (shared!)\n", a) // a[0] is also 999!

	// But if we append beyond capacity:
	c := make([]int, 3) //nolint:gosimple // at full capacity (explicit cap=len)
	c[0], c[1], c[2] = 1, 2, 3
	d := append(c, 4) // must reallocate
	d[0] = 999
	fmt.Printf("After realloc: c=%v, d=%v (independent)\n", c, d) // c unchanged

	// ─────────────────────────────────────────────
	// 5. GOTCHA: Always reassign result
	// ─────────────────────────────────────────────
	// BAD:  append(s, val)    // return value discarded!
	// GOOD: s = append(s, val)
}
