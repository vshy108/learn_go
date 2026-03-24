//go:build ignore

// Section 8, Topic 60: strings Package
//
// The strings package provides many useful string manipulation functions.
// These are analogous to Python's str methods or Rust's &str methods.
//
// Run: go run examples/s08_strings_package.go

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("=== strings Package ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. Searching
	// ─────────────────────────────────────────────
	fmt.Println("-- Searching --")
	s := "Hello, World! Hello, Go!"
	fmt.Printf("Contains 'Go':    %t\n", strings.Contains(s, "Go"))
	fmt.Printf("ContainsRune 'W': %t\n", strings.ContainsRune(s, 'W'))
	fmt.Printf("ContainsAny 'xyz': %t\n", strings.ContainsAny(s, "xyz"))
	fmt.Printf("HasPrefix 'Hello': %t\n", strings.HasPrefix(s, "Hello"))
	fmt.Printf("HasSuffix 'Go!':   %t\n", strings.HasSuffix(s, "Go!"))
	fmt.Printf("Index 'World':     %d\n", strings.Index(s, "World"))
	fmt.Printf("LastIndex 'Hello': %d\n", strings.LastIndex(s, "Hello"))
	fmt.Printf("Count 'Hello':     %d\n", strings.Count(s, "Hello"))

	// ─────────────────────────────────────────────
	// 2. Splitting and joining
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Split and Join --")
	csv := "apple,banana,cherry,date"
	parts := strings.Split(csv, ",")
	fmt.Printf("Split: %v\n", parts)

	joined := strings.Join(parts, " | ")
	fmt.Printf("Join:  %s\n", joined)

	// SplitN — limit splits:
	limited := strings.SplitN(csv, ",", 2)
	fmt.Printf("SplitN(2): %v\n", limited) // ["apple", "banana,cherry,date"]

	// Fields — split by whitespace:
	text := "  hello   world   go  "
	fields := strings.Fields(text)
	fmt.Printf("Fields: %v\n", fields) // ["hello", "world", "go"]

	// ─────────────────────────────────────────────
	// 3. Replacing
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Replace --")
	original := "foo bar foo baz foo"
	replaced := strings.Replace(original, "foo", "qux", 2) // replace first 2
	fmt.Printf("Replace(2): %s\n", replaced)

	allReplaced := strings.ReplaceAll(original, "foo", "qux")
	fmt.Printf("ReplaceAll: %s\n", allReplaced)

	// ─────────────────────────────────────────────
	// 4. Case conversion
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Case --")
	fmt.Printf("ToUpper:  %s\n", strings.ToUpper("hello"))
	fmt.Printf("ToLower:  %s\n", strings.ToLower("HELLO"))
	fmt.Printf("ToTitle:  %s\n", strings.ToTitle("hello world"))
	fmt.Printf("Title:    %s\n", strings.Title("hello world")) //nolint // deprecated but shown for reference

	// ─────────────────────────────────────────────
	// 5. Trimming
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Trim --")
	padded := "  \t hello \n "
	fmt.Printf("TrimSpace:  %q\n", strings.TrimSpace(padded))
	fmt.Printf("Trim:       %q\n", strings.Trim("***hello***", "*"))
	fmt.Printf("TrimLeft:   %q\n", strings.TrimLeft("***hello***", "*"))
	fmt.Printf("TrimRight:  %q\n", strings.TrimRight("***hello***", "*"))
	fmt.Printf("TrimPrefix: %q\n", strings.TrimPrefix("Hello, World", "Hello, "))
	fmt.Printf("TrimSuffix: %q\n", strings.TrimSuffix("file.go", ".go"))

	// ─────────────────────────────────────────────
	// 6. Repeat
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Repeat --")
	fmt.Println(strings.Repeat("Go! ", 3))
	fmt.Println(strings.Repeat("-", 40))

	// ─────────────────────────────────────────────
	// 7. Map (transform each rune)
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Map --")
	rot13 := strings.Map(func(r rune) rune {
		switch {
		case r >= 'a' && r <= 'z':
			return 'a' + (r-'a'+13)%26
		case r >= 'A' && r <= 'Z':
			return 'A' + (r-'A'+13)%26
		}
		return r
	}, "Hello, World!")
	fmt.Printf("ROT13: %s\n", rot13)

	// ─────────────────────────────────────────────
	// 8. EqualFold (case-insensitive comparison)
	// ─────────────────────────────────────────────
	fmt.Println("\n-- EqualFold --")
	fmt.Printf("EqualFold('Go', 'go'): %t\n", strings.EqualFold("Go", "go"))
	fmt.Printf("EqualFold('Go', 'GO'): %t\n", strings.EqualFold("Go", "GO"))

	// ─────────────────────────────────────────────
	// 9. NewReader (io.Reader from string)
	// ─────────────────────────────────────────────
	fmt.Println("\n-- NewReader --")
	reader := strings.NewReader("Hello")
	fmt.Printf("Reader len: %d, size: %d\n", reader.Len(), reader.Size())
}
