//go:build ignore

// Section 8, Topic 59: Runes vs Bytes
//
// A rune is an alias for int32 — represents a Unicode code point.
// A byte is an alias for uint8.
//
// Iterating with for-range gives RUNES (characters).
// Iterating with index gives BYTES.
//
// GOTCHA: len("café") = 5 (bytes), but range gives 4 runes.
// GOTCHA: Some Unicode characters are multiple runes (combining marks, emoji sequences).
//
// Run: go run examples/s08_runes_vs_bytes.go

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println("=== Runes vs Bytes ===")
	fmt.Println()

	s := "Hello, 世界! 🌍"

	// ─────────────────────────────────────────────
	// 1. Byte iteration (index loop)
	// ─────────────────────────────────────────────
	fmt.Println("-- Byte iteration --")
	fmt.Printf("String: %s (len=%d bytes)\n", s, len(s))
	for i := 0; i < len(s); i++ {
		fmt.Printf("  [%2d] byte=0x%02X\n", i, s[i])
	}

	// ─────────────────────────────────────────────
	// 2. Rune iteration (for-range)
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Rune iteration (for range) --")
	for i, r := range s {
		fmt.Printf("  [%2d] rune=U+%04X char=%c bytes=%d\n", i, r, r, utf8.RuneLen(r))
	}
	// Notice: index jumps (not +1 for multi-byte runes)

	// ─────────────────────────────────────────────
	// 3. Rune count vs byte count
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Counting --")
	text := "café"
	fmt.Printf("'%s': bytes=%d, runes=%d\n", text, len(text), utf8.RuneCountInString(text))
	// bytes=5 (é is 2 bytes), runes=4

	// ─────────────────────────────────────────────
	// 4. Converting to []rune and []byte
	// ─────────────────────────────────────────────
	fmt.Println("\n-- []rune and []byte --")
	runes := []rune(text)
	bytes := []byte(text)
	fmt.Printf("[]rune: %v (len=%d)\n", runes, len(runes))
	fmt.Printf("[]byte: %v (len=%d)\n", bytes, len(bytes))

	// Access by character index using []rune:
	fmt.Printf("runes[3] = '%c' (U+%04X)\n", runes[3], runes[3]) // é

	// ─────────────────────────────────────────────
	// 5. Rune literal
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Rune literals --")
	var r rune = '世'
	fmt.Printf("Rune literal: '%c' = U+%04X = %d\n", r, r, r)
	fmt.Printf("Type: %T, size: %d bytes (int32)\n", r, 4)

	// ASCII rune:
	var a rune = 'A'
	fmt.Printf("'A' = %d (fits in one byte, but rune is int32)\n", a)

	// ─────────────────────────────────────────────
	// 6. GOTCHA: Substrings are byte-based
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Substring gotcha --")
	s2 := "café"
	// s2[:4] cuts in the middle of 'é' → invalid string display
	sub := s2[:4]
	fmt.Printf("s2[:4] = %q (might be invalid!)\n", sub)
	fmt.Printf("Valid UTF-8: %t\n", utf8.ValidString(sub))

	// Safe substring by runes:
	runeSlice := []rune(s2)
	safeSub := string(runeSlice[:4])
	fmt.Printf("rune-safe[:4] = %q\n", safeSub) // "café"

	// ─────────────────────────────────────────────
	// 7. Decode rune by rune
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Manual decoding --")
	remaining := "世界"
	for len(remaining) > 0 {
		r, size := utf8.DecodeRuneInString(remaining)
		fmt.Printf("  '%c' (U+%04X, %d bytes)\n", r, r, size)
		remaining = remaining[size:]
	}
}
