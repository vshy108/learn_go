//go:build ignore

// Section 9, Topic 71: Anonymous Structs
//
// Anonymous structs have no named type — defined and used inline.
// Useful for one-off data structures, test data, and JSON parsing.
//
// Run: go run examples/s09_anonymous_structs.go

package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("=== Anonymous Structs ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. Inline variable
	// ─────────────────────────────────────────────
	point := struct {
		X, Y int
	}{X: 10, Y: 20}
	fmt.Printf("Point: %+v\n", point)

	// ─────────────────────────────────────────────
	// 2. Grouping related config
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Config --")
	config := struct {
		Host  string
		Port  int
		Debug bool
	}{
		Host:  "localhost",
		Port:  8080,
		Debug: true,
	}
	fmt.Printf("Config: %+v\n", config)

	// ─────────────────────────────────────────────
	// 3. JSON parsing without defining a type
	// ─────────────────────────────────────────────
	fmt.Println("\n-- JSON parsing --")
	data := `{"name": "Alice", "scores": [95, 87, 92]}`
	var result struct {
		Name string `json:"name"`

		Scores []int `json:"scores"`
	}
	_ = json.Unmarshal([]byte(data), &result)
	fmt.Printf("Parsed: %+v\n", result)
	// ─────────────────────────────────────────────
	// 4. Table-driven tests pattern
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Table-driven test data --")
	tests := []struct {
		input int

		expected int
	}{
		{1, 1},

		{2, 4},

		{3, 9},

		{4, 16},
	}
	for _, tt := range tests {
		got := tt.input * tt.input

		status := "PASS"

		if got != tt.expected {

			status = "FAIL"

		}

		fmt.Printf("  square(%d) = %d [%s]\n", tt.input, got, status)

	}
	// ─────────────────────────────────────────────
	// 5. Anonymous struct comparison
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Comparison --")
	a := struct{ X, Y int }{1, 2}
	b := struct{ X, Y int }{1, 2}
	c := struct{ X, Y int }{3, 4}
	fmt.Printf("a == b: %t\n", a == b) // true (same type, same values)
	fmt.Printf("a == c: %t\n", a == c) // false
}
