//go:build ignore

// Section 8, Topic 61: strings.Builder — Efficient Concatenation
//
// String concatenation with + creates a new string each time (O(n²) for loops).
// strings.Builder accumulates bytes efficiently, like Java's StringBuilder.
//
// GOTCHA: Don't copy a Builder after writing to it. Use pointer if passing around.
// GOTCHA: Builder.String() does NOT copy — it returns a string referencing the
//         internal buffer. Writing to the builder after String() is still safe
//         but the returned string won't reflect new writes.
//
// Run: go run examples/s08_string_builder.go

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("=== strings.Builder ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. Basic usage
	// ─────────────────────────────────────────────
	var b strings.Builder
	b.WriteString("Hello")
	b.WriteString(", ")
	b.WriteString("World")
	b.WriteByte('!')
	fmt.Println(b.String()) // "Hello, World!"

	// ─────────────────────────────────────────────
	// 2. WriteRune for Unicode
	// ─────────────────────────────────────────────
	fmt.Println("\n-- WriteRune --")
	var b2 strings.Builder
	b2.WriteString("Go ")
	b2.WriteRune('🚀')
	b2.WriteRune(' ')
	b2.WriteRune('世')
	b2.WriteRune('界')
	fmt.Println(b2.String())

	// ─────────────────────────────────────────────
	// 3. Grow (pre-allocate)
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Grow --")
	var b3 strings.Builder
	b3.Grow(100) // pre-allocate 100 bytes
	for i := 0; i < 10; i++ {
		fmt.Fprintf(&b3, "item %d, ", i) // can use as io.Writer!
	}
	fmt.Println(b3.String())

	// ─────────────────────────────────────────────
	// 4. Reset (reuse builder)
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Reset --")
	var b4 strings.Builder
	b4.WriteString("first")
	fmt.Printf("Before reset: %q\n", b4.String())
	b4.Reset()
	b4.WriteString("second")
	fmt.Printf("After reset: %q\n", b4.String())

	// ─────────────────────────────────────────────
	// 5. Len and Cap
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Len --")
	var b5 strings.Builder
	b5.WriteString("hello")
	fmt.Printf("Len: %d\n", b5.Len()) // 5

	// ─────────────────────────────────────────────
	// 6. Performance comparison
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Performance example --")

	// BAD: O(n²) — each + creates a new string
	// s := ""
	// for i := 0; i < 1000; i++ { s += "x" }

	// GOOD: O(n) — Builder accumulates
	var fast strings.Builder
	for i := 0; i < 1000; i++ {
		fast.WriteByte('x')
	}
	result := fast.String()
	fmt.Printf("Built string of length %d\n", len(result))

	// Also GOOD: strings.Join for slices
	parts := []string{"a", "b", "c", "d"}
	joined := strings.Join(parts, "-")
	fmt.Printf("Joined: %s\n", joined)

	// Also GOOD: fmt.Sprintf for formatting
	formatted := fmt.Sprintf("Name: %s, Age: %d", "Alice", 30)
	fmt.Printf("Sprintf: %s\n", formatted)
}
