//go:build ignore

// Section 18, Topic 136: sort Package
//
// sort package provides sorting for slices and user-defined collections.
//
// Since Go 1.21: slices.Sort is preferred for simple cases.
//
// Run: go run examples/s18_sort.go

package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("=== sort Package ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. Sort basic types
	// ─────────────────────────────────────────────
	ints := []int{5, 2, 8, 1, 9, 3}
	sort.Ints(ints)
	fmt.Println("Sorted ints:", ints)

	strs := []string{"banana", "apple", "cherry", "date"}
	sort.Strings(strs)
	fmt.Println("Sorted strings:", strs)

	floats := []float64{3.14, 1.41, 2.72, 0.57}
	sort.Float64s(floats)
	fmt.Println("Sorted floats:", floats)

	// ─────────────────────────────────────────────
	// 2. sort.Slice — custom comparator
	// ─────────────────────────────────────────────
	fmt.Println("\n-- sort.Slice --")
	people := []struct {
		Name string
		Age  int
	}{
		{"Charlie", 30},
		{"Alice", 25},
		{"Bob", 35},
		{"Diana", 28},
	}

	// Sort by age:
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println("By age:", people)
	// Sort by name:
	sort.Slice(people, func(i, j int) bool {
		return people[i].Name < people[j].Name

	})
	fmt.Println("By name:", people)
	// ─────────────────────────────────────────────
	// 3. Stable sort (preserves order of equal elements)
	// ─────────────────────────────────────────────
	fmt.Println("\n-- StableSort --")
	items := []struct {
		Name string

		Priority int
	}{
		{"A", 2}, {"B", 1}, {"C", 2}, {"D", 1}, {"E", 3},
	}
	sort.SliceStable(items, func(i, j int) bool {
		return items[i].Priority < items[j].Priority

	})
	fmt.Println("Stable sort:", items)
	// B and D (priority 1) maintain relative order
	// A and C (priority 2) maintain relative order
	// ─────────────────────────────────────────────
	// 4. Reverse sort
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Reverse --")
	nums := []int{1, 2, 3, 4, 5}
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	fmt.Println("Reversed:", nums)
	// Or with sort.Slice:
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]

	})
	fmt.Println("Descending:", nums)
	// ─────────────────────────────────────────────
	// 5. Binary search (slice must be sorted!)
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Search --")
	sorted := []int{1, 3, 5, 7, 9, 11}
	idx := sort.SearchInts(sorted, 7)
	fmt.Printf("Index of 7: %d\n", idx) // 3
	idx = sort.SearchInts(sorted, 6)
	fmt.Printf("Search for 6 (not found): insertion point %d\n", idx) // 3
	// ─────────────────────────────────────────────
	// 6. sort.IsSorted
	// ─────────────────────────────────────────────
	fmt.Println("\n-- IsSorted --")
	a := []int{1, 2, 3, 4}
	b := []int{4, 2, 3, 1}
	fmt.Println("a sorted:", sort.IntsAreSorted(a)) // true
	fmt.Println("b sorted:", sort.IntsAreSorted(b)) // false
}
