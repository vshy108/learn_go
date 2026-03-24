//go:build ignore

// Section 6, Topic 49: Common Slice Tricks — Delete, Insert, Filter
//
// Go doesn't have built-in slice manipulation functions (pre-1.21).
// Here are common patterns. Since Go 1.21, the `slices` package provides
// many of these operations.
//
// GOTCHA: Most slice tricks modify the underlying array in-place.
//         Make a copy first if you need to preserve the original.
// GOTCHA: Deleting from a slice may leave old pointers, causing memory leaks
//         with pointer-containing slices. Zero out removed elements.
//
// Run: go run examples/s06_slice_tricks.go

package main

import "fmt"

func main() {
	fmt.Println("=== Slice Tricks ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. Delete element at index (preserving order)
	// ─────────────────────────────────────────────
	fmt.Println("-- Delete (preserve order) --")
	s := []int{1, 2, 3, 4, 5}
	i := 2 // delete index 2 (value 3)
	s = append(s[:i], s[i+1:]...)
	fmt.Printf("After deleting index 2: %v\n", s) // [1 2 4 5]

	// ─────────────────────────────────────────────
	// 2. Delete element (NOT preserving order — faster)
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Delete (swap with last — O(1)) --")
	s = []int{1, 2, 3, 4, 5}
	i = 2
	s[i] = s[len(s)-1]                               // replace with last element
	s = s[:len(s)-1]                                 // shrink
	fmt.Printf("After swap-delete index 2: %v\n", s) // [1 2 5 4]

	// ─────────────────────────────────────────────
	// 3. Insert element at index
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Insert at index --")
	s = []int{1, 2, 4, 5}
	i = 2
	val := 3
	s = append(s[:i+1], s[i:]...)
	s[i] = val
	fmt.Printf("After inserting 3 at index 2: %v\n", s) // [1 2 3 4 5]

	// ─────────────────────────────────────────────
	// 4. Filter (create new slice with matching elements)
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Filter --")
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evens := filter(nums, func(n int) bool { return n%2 == 0 })
	fmt.Printf("Evens: %v\n", evens)

	// ─────────────────────────────────────────────
	// 5. Filter in-place (reuse backing array)
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Filter in-place --")
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	n := 0
	for _, v := range data {
		if v%2 == 0 {
			data[n] = v
			n++
		}
	}
	data = data[:n]
	fmt.Printf("In-place evens: %v\n", data)

	// ─────────────────────────────────────────────
	// 6. Reverse
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Reverse --")
	r := []int{1, 2, 3, 4, 5}
	for left, right := 0, len(r)-1; left < right; left, right = left+1, right-1 {
		r[left], r[right] = r[right], r[left]
	}
	fmt.Printf("Reversed: %v\n", r)

	// ─────────────────────────────────────────────
	// 7. Deduplicate (sorted slice)
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Deduplicate (sorted input) --")
	sorted := []int{1, 1, 2, 3, 3, 3, 4, 5, 5}
	deduped := dedupSorted(sorted)
	fmt.Printf("Deduped: %v\n", deduped)

	// ─────────────────────────────────────────────
	// 8. Pop (remove last element)
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Pop --")
	stack := []int{1, 2, 3, 4, 5}
	top := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	fmt.Printf("Popped %d, remaining: %v\n", top, stack)

	// ─────────────────────────────────────────────
	// 9. Shift (remove first element)
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Shift --")
	queue := []int{1, 2, 3, 4, 5}
	front := queue[0]
	queue = queue[1:]
	fmt.Printf("Shifted %d, remaining: %v\n", front, queue)
}

func filter(s []int, pred func(int) bool) []int {
	var result []int
	for _, v := range s {
		if pred(v) {
			result = append(result, v)
		}
	}
	return result
}

func dedupSorted(s []int) []int {
	if len(s) == 0 {
		return s
	}
	j := 0
	for i := 1; i < len(s); i++ {
		if s[i] != s[j] {
			j++
			s[j] = s[i]
		}
	}
	return s[:j+1]
}
