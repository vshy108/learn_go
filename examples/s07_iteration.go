//go:build ignore

// Section 7, Topic 54: Iterating Maps — Random Order!
//
// Maps are iterated using for-range. The iteration order is
// NOT guaranteed and is intentionally randomized by the Go runtime.
//
// GOTCHA: Map iteration order changes between runs!
//         Never rely on map ordering for deterministic output.
// GOTCHA: It's safe to delete map entries during iteration.
// GOTCHA: Adding entries during iteration may or may not be visited.
//
// Run: go run examples/s07_iteration.go

package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("=== Map Iteration ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. Basic iteration (random order)
	// ─────────────────────────────────────────────
	m := map[string]int{
		"Alice": 90,
		"Bob":   85,
		"Carol": 92,
		"Dave":  78,
	}
	fmt.Println("-- Random order (run twice to see difference) --")
	for k, v := range m {
		fmt.Printf("  %s: %d\n", k, v)
	}

	// ─────────────────────────────────────────────
	// 2. Keys only
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Keys only --")
	for k := range m {
		fmt.Printf("  key: %s\n", k)
	}

	// ─────────────────────────────────────────────
	// 3. Sorted iteration (collect keys, sort, iterate)
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Sorted by key --")
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("  %s: %d\n", k, m[k])
	}

	// ─────────────────────────────────────────────
	// 4. Delete during iteration (safe!)
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Delete during iteration --")
	data := map[string]int{
		"keep1":   1,
		"remove1": 2,
		"keep2":   3,
		"remove2": 4,
	}
	for k := range data {
		if len(k) >= 6 && k[:6] == "remove" {
			delete(data, k)
		}
	}
	fmt.Println("After delete:", data)

	// ─────────────────────────────────────────────
	// 5. Counting frequency
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Word frequency --")
	words := []string{"go", "is", "fun", "go", "is", "great", "go"}
	freq := make(map[string]int)
	for _, w := range words {
		freq[w]++
	}
	// Print sorted:
	wordKeys := make([]string, 0, len(freq))
	for k := range freq {
		wordKeys = append(wordKeys, k)
	}
	sort.Strings(wordKeys)
	for _, k := range wordKeys {
		fmt.Printf("  %s: %d\n", k, freq[k])
	}

	// ─────────────────────────────────────────────
	// 6. GOTCHA: Why random order?
	// ─────────────────────────────────────────────
	// Go intentionally randomizes map iteration to prevent developers
	// from relying on a specific order. Before Go 1.12, iteration was
	// more predictable (but not guaranteed), leading to subtle bugs.
	//
	// If you need ordered maps, use a separate sorted key slice
	// or a third-party ordered map library.
}
