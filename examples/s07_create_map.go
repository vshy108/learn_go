//go:build ignore

// Section 7, Topic 51: Creating Maps
//
// Maps are hash tables mapping keys to values: map[KeyType]ValueType
// Keys must be comparable (==, !=): strings, ints, bools, structs (if all
// fields are comparable). Slices, maps, and functions CANNOT be keys.
//
// GOTCHA: You MUST initialize a map before writing to it.
//         var m map[string]int → nil map → panic on write!
// GOTCHA: Map access for missing key returns zero value (no error).
// GOTCHA: Maps are NOT safe for concurrent access.
//
// Run: go run examples/s07_create_map.go

package main

import "fmt"

func main() {
	fmt.Println("=== Creating Maps ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. Map literal
	// ─────────────────────────────────────────────
	fmt.Println("-- Map literal --")
	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
		"blue":  "#0000FF",
	} // trailing comma required on last line!
	fmt.Println(colors)

	// ─────────────────────────────────────────────
	// 2. Using make
	// ─────────────────────────────────────────────
	fmt.Println("\n-- make --")
	ages := make(map[string]int)
	ages["Alice"] = 30
	ages["Bob"] = 25
	fmt.Println(ages)

	// With initial capacity hint (optimization, not limit):
	scores := make(map[string]int, 100)
	scores["test"] = 99
	fmt.Printf("scores len=%d\n", len(scores))

	// ─────────────────────────────────────────────
	// 3. Empty map vs nil map
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Empty vs nil --")
	empty := map[string]int{}    // empty, not nil
	made := make(map[string]int) // empty, not nil
	var nilMap map[string]int    // nil!

	fmt.Printf("empty == nil: %t\n", empty == nil)   // false
	fmt.Printf("made == nil:  %t\n", made == nil)    // false
	fmt.Printf("nilMap == nil: %t\n", nilMap == nil) // true

	// Reading from nil is OK:
	fmt.Printf("nilMap[\"x\"] = %d (zero value, no panic)\n", nilMap["x"])

	// Writing to nil map panics:
	// nilMap["x"] = 1  // PANIC: assignment to entry in nil map

	// ─────────────────────────────────────────────
	// 4. Complex value types
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Complex values --")
	// Map of slices:
	graph := map[string][]string{
		"A": {"B", "C"},
		"B": {"A", "D"},
	}
	fmt.Println("Graph:", graph)

	// Map of maps:
	db := map[string]map[string]int{
		"Alice": {"math": 95, "science": 88},
		"Bob":   {"math": 72, "science": 91},
	}
	fmt.Println("DB:", db)

	// ─────────────────────────────────────────────
	// 5. Valid key types
	// ─────────────────────────────────────────────
	// Comparable types (can be keys): int, string, bool, float,
	//   complex, pointers, channels, arrays, structs with comparable fields
	// NOT comparable (cannot be keys): slices, maps, functions

	type Point struct{ X, Y int }
	pointMap := map[Point]string{
		{0, 0}: "origin",
		{1, 2}: "point A",
	}
	fmt.Println("\nStruct keys:", pointMap)

	// Array keys:
	arrayMap := map[[2]int]string{
		{0, 0}: "origin",
	}
	fmt.Println("Array keys:", arrayMap)
}
