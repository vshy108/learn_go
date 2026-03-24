//go:build ignore

// Section 7, Topic 53: Adding and Deleting Map Entries
//
// Adding: m[key] = value  (inserts or overwrites)
// Deleting: delete(m, key) (no-op if key doesn't exist)
//
// GOTCHA: delete() on a nil map is a no-op (doesn't panic).
// GOTCHA: delete() on a missing key is a no-op (doesn't panic).
// GOTCHA: You can't take the address of a map value: &m[key] is illegal.
//
// Run: go run examples/s07_add_delete.go

package main

import "fmt"

func main() {
	fmt.Println("=== Map Add & Delete ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. Adding entries
	// ─────────────────────────────────────────────
	m := make(map[string]int)
	m["one"] = 1
	m["two"] = 2
	m["three"] = 3
	fmt.Println("After adds:", m)

	// Overwrite:
	m["one"] = 100
	fmt.Println("After overwrite:", m)

	// ─────────────────────────────────────────────
	// 2. Deleting entries
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Delete --")
	delete(m, "two")
	fmt.Println("After delete 'two':", m)

	// Delete non-existent key — no panic:
	delete(m, "nonexistent")
	fmt.Println("After deleting missing key:", m) // unchanged

	// Delete from nil map — also no panic:
	var nilMap map[string]int
	delete(nilMap, "key") // no-op
	fmt.Println("Delete from nil map: ok")

	// ─────────────────────────────────────────────
	// 3. len() for maps
	// ─────────────────────────────────────────────
	fmt.Println("\n-- len --")
	fmt.Printf("len(m) = %d\n", len(m))
	// There is NO cap() for maps

	// ─────────────────────────────────────────────
	// 4. GOTCHA: Can't address map values
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Can't address map values --")
	// m2 := map[string]int{"a": 1}
	// p := &m2["a"]  // COMPILE ERROR: cannot take address of m2["a"]

	// Why? Maps may relocate values during growth. A pointer could dangle.
	// Workaround: use a map of pointers:
	type Data struct{ Val int }
	m3 := map[string]*Data{
		"a": {Val: 1},
	}
	m3["a"].Val = 999 // OK — pointer is stable
	fmt.Printf("m3[\"a\"].Val = %d\n", m3["a"].Val)

	// ─────────────────────────────────────────────
	// 5. GOTCHA: Can't modify struct value directly in map
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Can't modify struct field in map --")
	type Person struct{ Age int }
	people := map[string]Person{
		"Alice": {Age: 30},
	}
	// people["Alice"].Age = 31  // COMPILE ERROR: cannot assign to struct field

	// Workaround 1: reassign the whole struct:
	p := people["Alice"]
	p.Age = 31
	people["Alice"] = p
	fmt.Println("Alice:", people["Alice"])

	// Workaround 2: use map of pointers:
	people2 := map[string]*Person{
		"Alice": {Age: 30},
	}
	people2["Alice"].Age = 31 // OK
	fmt.Println("Alice (ptr):", people2["Alice"])

	// ─────────────────────────────────────────────
	// 6. Clear all entries (Go 1.21+)
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Clear map --")
	data := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Printf("Before clear: %v (len=%d)\n", data, len(data))

	// Pre-1.21: iterate and delete
	for k := range data {
		delete(data, k)
	}
	fmt.Printf("After clear:  %v (len=%d)\n", data, len(data))

	// Go 1.21+: clear(data) — built-in function
}
