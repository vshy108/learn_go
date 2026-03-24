//go:build ignore

// Section 7, Topic 52: Accessing Values — Comma-Ok Idiom
//
// Accessing a missing key returns the zero value for the value type.
// Use the "comma ok" idiom to distinguish "key exists with zero value"
// from "key doesn't exist".
//
//   val, ok := m[key]
//   val — the value (or zero value if missing)
//   ok  — true if key exists, false otherwise
//
// GOTCHA: Without comma-ok, you can't tell if 0 means "value is 0"
//         or "key not found".
//
// Run: go run examples/s07_access_comma_ok.go

package main

import "fmt"

func main() {
	fmt.Println("=== Map Access & Comma-Ok ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. Basic access
	// ─────────────────────────────────────────────
	m := map[string]int{
		"Alice": 90,
		"Bob":   0, // intentionally zero
	}

	fmt.Println("Alice:", m["Alice"]) // 90
	fmt.Println("Bob:", m["Bob"])     // 0
	fmt.Println("Eve:", m["Eve"])     // 0 (missing key → zero value)

	// Problem: Bob=0 and Eve=0 are indistinguishable!

	// ─────────────────────────────────────────────
	// 2. Comma-ok idiom
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Comma-ok --")
	if val, ok := m["Bob"]; ok {
		fmt.Printf("Bob exists: %d\n", val)
	} else {
		fmt.Println("Bob not found")
	}

	if val, ok := m["Eve"]; ok {
		fmt.Printf("Eve exists: %d\n", val)
	} else {
		fmt.Println("Eve not found") // this branch
	}

	// ─────────────────────────────────────────────
	// 3. Pattern: check-then-act
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Check-then-act --")
	word := "hello"
	freq := map[rune]int{}
	for _, ch := range word {
		freq[ch]++ // missing key returns 0, so 0+1=1 on first hit
	}
	fmt.Println("Frequency:", freq)

	// ─────────────────────────────────────────────
	// 4. Pattern: default value
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Default value pattern --")
	config := map[string]string{
		"host": "localhost",
	}
	host := getOrDefault(config, "host", "0.0.0.0")
	port := getOrDefault(config, "port", "8080")
	fmt.Printf("host=%s, port=%s\n", host, port)

	// ─────────────────────────────────────────────
	// 5. GOTCHA: Zero values for different types
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Zero values for missing keys --")
	intMap := map[string]int{}
	strMap := map[string]string{}
	boolMap := map[string]bool{}
	sliceMap := map[string][]int{}

	fmt.Printf("int:    %d\n", intMap["x"])  // 0
	fmt.Printf("string: %q\n", strMap["x"])  // ""
	fmt.Printf("bool:   %t\n", boolMap["x"]) // false
	fmt.Printf("slice:  %v (nil=%t)\n", sliceMap["x"], sliceMap["x"] == nil)

	// All zero values — always use comma-ok when ambiguity matters!
}

func getOrDefault(m map[string]string, key, defaultVal string) string {
	if v, ok := m[key]; ok {
		return v
	}
	return defaultVal
}
