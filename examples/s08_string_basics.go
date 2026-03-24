//go:build ignore

// Section 8, Topic 58: Strings — Immutable UTF-8 Byte Slices
//
// Go strings are immutable sequences of bytes, conventionally holding UTF-8 text.
// A string is essentially a read-only []byte with syntactic sugar.
//
// Internally: struct { ptr *byte; len int }  (16 bytes on 64-bit)
//
// GOTCHA: len("hello") returns BYTE count, not CHARACTER count.
// GOTCHA: Indexing s[i] returns a BYTE, not a rune (character).
// GOTCHA: Strings are immutable — you cannot modify individual bytes.
// GOTCHA: Strings can hold arbitrary bytes, not just valid UTF-8.
//
// Run: go run examples/s08_string_basics.go

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println("=== String Basics ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. String declaration
	// ─────────────────────────────────────────────
	s := "Hello, World!"
	fmt.Println(s)
	fmt.Printf("Type: %T\n", s)

	// ─────────────────────────────────────────────
	// 2. len() returns BYTE count
	// ─────────────────────────────────────────────
	fmt.Println("\n-- len() = byte count --")
	ascii := "hello"
	emoji := "Hello, 世界! 🌍"
	fmt.Printf("'%s': len=%d bytes\n", ascii, len(ascii))                // 5
	fmt.Printf("'%s': len=%d bytes\n", emoji, len(emoji))                // 20
	fmt.Printf("'%s': runes=%d\n", emoji, utf8.RuneCountInString(emoji)) // 12

	// ─────────────────────────────────────────────
	// 3. Indexing returns bytes, not characters
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Indexing --")
	s2 := "café"
	fmt.Printf("len('%s') = %d (byte count)\n", s2, len(s2)) // 5 (é = 2 bytes)
	fmt.Printf("s[0] = %d (%c) — byte!\n", s2[0], s2[0])
	// s[3] returns 0xC3 (first byte of é), NOT 'é'!
	fmt.Printf("s[3] = 0x%X (first byte of é)\n", s2[3])

	// ─────────────────────────────────────────────
	// 4. String comparison
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Comparison --")
	s3, s4 := "abc", "abc"
	fmt.Printf("\"abc\" == \"abc\": %t\n", s3 == s4)
	fmt.Printf("\"abc\" < \"abd\": %t\n", "abc" < "abd") // lexicographic
	fmt.Printf("\"abc\" < \"b\":   %t\n", "abc" < "b")

	// ─────────────────────────────────────────────
	// 5. String concatenation
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Concatenation --")
	a := "Hello"
	b := "World"
	c := a + ", " + b + "!"
	fmt.Println(c)
	// For many concatenations, use strings.Builder (see s08_string_builder)

	// ─────────────────────────────────────────────
	// 6. Multi-line strings
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Multi-line --")
	multiLine := "line 1\n" +
		"line 2\n" +
		"line 3"
	fmt.Println(multiLine)

	// ─────────────────────────────────────────────
	// 7. Zero value
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Zero value --")
	var empty string
	fmt.Printf("Zero value: %q (len=%d)\n", empty, len(empty)) // ""

	// ─────────────────────────────────────────────
	// 8. Strings can hold any bytes
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Arbitrary bytes --")
	binary := "\x00\x01\x02\xff"
	fmt.Printf("Binary string: len=%d, bytes=%v\n", len(binary), []byte(binary))

	// ─────────────────────────────────────────────
	// Comparison: Go strings vs Rust
	// ─────────────────────────────────────────────
	// Go:   string = immutable []byte, typically UTF-8
	// Rust: String = owned, growable, guaranteed UTF-8
	// Rust: &str   = immutable reference to UTF-8 data
	// Go is more permissive — strings can contain invalid UTF-8.
}
