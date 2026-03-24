//go:build ignore

// Section 7, Topic 56: Maps with Slice Values
//
// Maps can hold slices as values: map[string][]int
// Useful for grouping, multi-value lookups, adjacency lists, etc.
//
// GOTCHA: Accessing a missing key returns nil slice (can append to nil slice,
//         but remember to store it back in the map).
//
// Run: go run examples/s07_map_of_slices.go

package main

import "fmt"

func main() {
	fmt.Println("=== Map of Slices ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. Grouping values
	// ─────────────────────────────────────────────
	fmt.Println("-- Grouping --")
	students := map[string][]string{
		"math":    {"Alice", "Bob"},
		"science": {"Bob", "Eve"},
	}
	students["math"] = append(students["math"], "Charlie")
	for subject, names := range students {
		fmt.Printf("  %s: %v\n", subject, names)
	}

	// ─────────────────────────────────────────────
	// 2. GOTCHA: Must store back after append
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Append to missing key --")
	graph := make(map[string][]string)

	// Append to missing key — append(nil, x) works:
	graph["A"] = append(graph["A"], "B") // graph["A"] was nil
	graph["A"] = append(graph["A"], "C")
	graph["B"] = append(graph["B"], "A")
	fmt.Printf("Graph: %v\n", graph)

	// ─────────────────────────────────────────────
	// 3. Word index (inverted index)
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Inverted index --")
	docs := map[int]string{
		1: "the quick brown fox",
		2: "the lazy brown dog",
		3: "quick fox jumps",
	}
	index := make(map[string][]int)
	for id, text := range docs {
		for _, word := range split(text) {
			index[word] = append(index[word], id)
		}
	}
	fmt.Printf("'brown' in docs: %v\n", index["brown"])
	fmt.Printf("'quick' in docs: %v\n", index["quick"])
	fmt.Printf("'fox' in docs:   %v\n", index["fox"])

	// ─────────────────────────────────────────────
	// 4. Adjacency list
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Adjacency list --")
	adj := map[int][]int{
		1: {2, 3},
		2: {1, 4},
		3: {1},
		4: {2},
	}
	for node, neighbors := range adj {
		fmt.Printf("  Node %d → %v\n", node, neighbors)
	}
}

// Simple split by space (no strings package import needed for demo)
func split(s string) []string {
	var result []string
	word := ""
	for _, r := range s {
		if r == ' ' {
			if word != "" {
				result = append(result, word)
				word = ""
			}
		} else {
			word += string(r)
		}
	}
	if word != "" {
		result = append(result, word)
	}
	return result
}
