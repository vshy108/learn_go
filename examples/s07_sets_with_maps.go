//go:build ignore

// Section 7, Topic 57: Implementing Sets with Maps
//
// Go has no built-in set type. Use map[T]bool or map[T]struct{} instead.
//
// map[T]bool:     easier to use: if visited[key] { ... }
// map[T]struct{}: zero-size values, slightly more memory-efficient
//
// GOTCHA: map[T]bool — a missing key returns false, which is usually what you want.
// GOTCHA: map[T]struct{} — must use `_, ok := s[key]` or `s[key] = struct{}{}`
//
// Run: go run examples/s07_sets_with_maps.go

package main

import "fmt"

func main() {
	fmt.Println("=== Sets with Maps ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. Set using map[string]bool
	// ─────────────────────────────────────────────
	fmt.Println("-- map[string]bool --")
	visited := map[string]bool{
		"Apple":  true,
		"Banana": true,
	}

	// Check membership (idiomatic — false for absent keys):
	fmt.Printf("Apple in set: %t\n", visited["Apple"])    // true
	fmt.Printf("Cherry in set: %t\n", visited["Cherry"])  // false

	// Add:





























































































































}	return result	}		}			result[k] = struct{}{}		if _, ok := b[k]; !ok {	for k := range a {	result := newSet()func difference(a, b StringSet) StringSet {}	return result	}		}			result[k] = struct{}{}		if _, ok := b[k]; ok {	for k := range a {	result := newSet()func intersection(a, b StringSet) StringSet {}	return result	}		result[k] = struct{}{}	for k := range b {	}		result[k] = struct{}{}	for k := range a {	result := newSet()func union(a, b StringSet) StringSet {}	return result	}		result = append(result, k)	for k := range s {	result := make([]string, 0, len(s))func setToSlice(s StringSet) []string {}	return s	}		s[item] = struct{}{}	for _, item := range items {	s := make(StringSet, len(items))func newSet(items ...string) StringSet {type StringSet map[string]struct{}// Helper type and functions}	fmt.Printf("Unique: %v\n", unique)	}		}			unique = append(unique, item)			seen[item] = true		if !seen[item] {	for _, item := range items {	var unique []string	seen := make(map[string]bool)	items := []string{"go", "rust", "go", "python", "rust", "go"}	fmt.Println("\n-- Deduplication --")	// ─────────────────────────────────────────────	// 4. Deduplication using set	// ─────────────────────────────────────────────	fmt.Printf("Difference:   %v\n", setToSlice(diff))	diff := difference(a, b)	// Difference (A - B):	fmt.Printf("Intersection: %v\n", setToSlice(inter))	inter := intersection(a, b)	// Intersection:	fmt.Printf("Union:        %v\n", setToSlice(u))	u := union(a, b)	// Union:	fmt.Printf("B: %v\n", setToSlice(b))	fmt.Printf("A: %v\n", setToSlice(a))	b := newSet("Go", "Java", "TypeScript", "Kotlin")	a := newSet("Go", "Rust", "Python", "TypeScript")	fmt.Println("\n-- Set operations --")	// ─────────────────────────────────────────────	// 3. Set operations	// ─────────────────────────────────────────────	fmt.Printf("Size: %d\n", len(set))	// Size:	delete(set, "Python")	// Remove:	set["TypeScript"] = struct{}{}	// Add:	}		fmt.Println("Java is NOT in the set")	if _, ok := set["Java"]; !ok {	}		fmt.Println("Go is in the set")	if _, ok := set["Go"]; ok {	// Check membership:	}		"Python": {},		"Rust":   {},		"Go":     {},	set := map[string]struct{}{	fmt.Println("\n-- map[string]struct{} --")	// ─────────────────────────────────────────────	// 2. Set using map[string]struct{} (memory efficient)	// ─────────────────────────────────────────────	fmt.Println()	}		fmt.Printf("%s ", k)	for k := range visited {	fmt.Printf("Set: ")	delete(visited, "Banana")	// Remove:	visited["Cherry"] = true