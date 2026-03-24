//go:build ignore

// Section 8, Topic 64: String Mutability
//
// Go strings are IMMUTABLE. You cannot modify individual bytes or characters.
// To mutate, convert to []byte or []rune, modify, convert back.
//
// GOTCHA: string → []byte → string creates TWO copies!
// GOTCHA: Use []byte for ASCII manipulation, []rune for Unicode.
// GOTCHA: strings.Builder is preferred for building strings incrementally.
//
// Run: go run examples/s08_string_mutability.go

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("=== String Mutability ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. Strings are immutable
	// ─────────────────────────────────────────────
	s := "hello"
	// s[0] = 'H'  // COMPILE ERROR: cannot assign to s[0]
	fmt.Printf("Original: %s\n", s)

	// ─────────────────────────────────────────────
	// 2. Mutation via []byte (for ASCII/single-byte)
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Mutation via []byte --")
	b := []byte(s) // copy
	b[0] = 'H'
	s2 := string(b)                  // another copy
	fmt.Printf("Modified: %s\n", s2) // "Hello"
	fmt.Printf("Original unchanged: %s\n", s)

	// ─────────────────────────────────────────────
	// 3. Mutation via []rune (for Unicode)
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Mutation via []rune --")
	s3 := "café"
	runes := []rune(s3)
	runes[3] = 'E' // change é to E
	s4 := string(runes)
	fmt.Printf("Modified: %s\n", s4) // "cafE"

	// ─────────────────────────────────────────────
	// 4. Concatenation creates new strings
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Concatenation --")
	a := "Hello"
	b2 := a + ", World!" // new string allocated
	fmt.Printf("a: %s (unchanged)\n", a)
	fmt.Printf("b: %s (new string)\n", b2)

	// ─────────────────────────────────────────────
	// 5. strings.Builder for efficient building
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Builder --")
	var builder strings.Builder
	builder.WriteString("Hello")
	builder.WriteString(", ")
	builder.WriteString("World!")
	result := builder.String()
	fmt.Printf("Built: %s\n", result)

	// ─────────────────────────────────────────────
	// 6. Replace creates new string
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Replace --")
	original := "hello world"
	replaced := strings.Replace(original, "world", "Go", 1)
	fmt.Printf("Original: %s (unchanged)\n", original)
	fmt.Printf("Replaced: %s (new string)\n", replaced)

	// ─────────────────────────────────────────────
	// 7. Common patterns for "modified" strings
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Common patterns --")

	// Uppercase first letter:
	name := "alice"
	capitalized := strings.ToUpper(name[:1]) + name[1:]
	fmt.Printf("Capitalized: %s\n", capitalized)

	// Reverse string:
	reversed := reverseString("Hello, 世界")
	fmt.Printf("Reversed: %s\n", reversed)
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
