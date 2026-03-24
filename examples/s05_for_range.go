//go:build ignore

// Section 5, Topic 37: for-range — Iterating Over Collections
//
// `for range` iterates over slices, arrays, maps, strings, and channels.
// It yields index/key and value pairs.
//
// Syntax:
//   for index, value := range collection { }
//   for index := range collection { }         // value omitted
//   for _, value := range collection { }      // index discarded
//   for range collection { }                  // both omitted (Go 1.22+)
//
// GOTCHA: `range` over a string iterates RUNES, not bytes.
// GOTCHA: `range` over a map has RANDOM order (not insertion order).
// GOTCHA: The `value` in range is a COPY — modifying it doesn't affect the original.
//
// Run: go run examples/s05_for_range.go

package main

import "fmt"

func main() {
	fmt.Println("=== for-range ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. Slice/array range
	// ─────────────────────────────────────────────
	fmt.Println("-- Slice range --")
	fruits := []string{"apple", "banana", "cherry"}
	for i, fruit := range fruits {
		fmt.Printf("  [%d] %s\n", i, fruit)
	}

	// Index only:
	fmt.Print("Indices: ")
	for i := range fruits {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// Value only:
	fmt.Print("Values: ")
	for _, fruit := range fruits {
		fmt.Printf("%s ", fruit)
	}
	fmt.Println()

	// ─────────────────────────────────────────────
	// 2. Map range (random order!)
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Map range (random order!) --")
	ages := map[string]int{"Alice": 30, "Bob": 25, "Charlie": 35}
	for name, age := range ages {
		fmt.Printf("  %s: %d\n", name, age)
	}
	// Run multiple times — order will vary!

	// Keys only:
	fmt.Print("Keys: ")
	for name := range ages {
		fmt.Printf("%s ", name)
	}
	fmt.Println()

	// ─────────────────────────────────────────────
	// 3. String range (iterates RUNES, not bytes)
	// ─────────────────────────────────────────────
	fmt.Println("\n-- String range (rune iteration) --")
	s := "Go 世界"
	for i, r := range s {
		fmt.Printf("  byte[%d] = '%c' (U+%04X)\n", i, r, r)
	}
	// Note: byte offsets jump by 3 for CJK characters

	// ─────────────────────────────────────────────
	// 4. Channel range
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Channel range --")
	ch := make(chan int, 5)
	for i := 0; i < 5; i++ {
		ch <- i * 10
	}
	close(ch) // must close for range to terminate

	for v := range ch {
		fmt.Printf("%d ", v)
	}
	fmt.Println()

	// ─────────────────────────────────────────────
	// 5. GOTCHA: range value is a COPY
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Range value is a copy --")
	nums := []int{1, 2, 3}
	for _, v := range nums {
		v *= 10 // modifies the COPY, not the slice
	}
	fmt.Println("After range modify value:", nums) // [1 2 3] — unchanged!

	// To modify: use the index
	for i := range nums {
		nums[i] *= 10
	}
	fmt.Println("After index modify:", nums) // [10 20 30]

	// ─────────────────────────────────────────────
	// 6. GOTCHA: range over nil slice/map is safe
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Range over nil --")
	var nilSlice []int
	_ = nilSlice
	// for i, v := range nilSlice { ... } // runs zero times, no panic
	fmt.Println("Range over nil slice: safe (zero iterations, no panic)")

	// ─────────────────────────────────────────────
	// 7. Go 1.22+: range over integer
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Range over integer (Go 1.22+) --")
	for i := range 5 {
		fmt.Printf("%d ", i) // 0 1 2 3 4
	}
	fmt.Println()
}
