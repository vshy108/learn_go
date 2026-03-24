//go:build ignore

// Section 3, Topic 20: byte vs rune — Understanding Go's Character Types
//
// Go has two character-related type aliases:
//   - byte = uint8 (a single byte, 0-255)
//   - rune = int32 (a Unicode code point, can represent any character)
//
// Go strings are sequences of BYTES (not runes). This is critical:
//   - ASCII characters: 1 byte each
//   - Many non-ASCII characters: 2-4 bytes each (UTF-8 encoding)
//
// GOTCHA: len("世界") returns 6 (bytes), not 2 (characters).
// GOTCHA: Indexing a string gives a byte, not a rune.
// GOTCHA: for-range over a string iterates runes; for i:=0;i<len gives bytes.
//
// Run: go run examples/s03_byte_and_rune.go

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println("=== byte vs rune ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. byte — alias for uint8
	// ─────────────────────────────────────────────
	var b byte = 'A' // A single ASCII character
	fmt.Printf("byte 'A': value=%d, char=%c, type=%T\n", b, b, b)

	// byte can only hold 0-255:
	var maxByte byte = 255
	fmt.Printf("max byte: %d\n", maxByte)
	// var overflow byte = 256  // ERROR: constant 256 overflows byte

	// ─────────────────────────────────────────────
	// 2. rune — alias for int32
	// ─────────────────────────────────────────────
	var r rune = '世' // A Chinese character (Unicode U+4E16)
	fmt.Printf("rune '世': value=%d (U+%04X), char=%c, type=%T\n", r, r, r, r)

	var emoji rune = '🚀' // Rocket emoji (U+1F680)
	fmt.Printf("rune '🚀': value=%d (U+%04X), char=%c, type=%T\n", emoji, emoji, emoji, emoji)

	// ─────────────────────────────────────────────
	// 3. Strings: bytes vs runes
	// ─────────────────────────────────────────────
	s := "Hello, 世界"
	fmt.Println("\n-- String analysis --")
	fmt.Printf("String:       %s\n", s)
	fmt.Printf("len():        %d bytes\n", len(s))                    // 13
	fmt.Printf("RuneCount:    %d runes\n", utf8.RuneCountInString(s)) // 9
	fmt.Printf("[]rune len:   %d\n", len([]rune(s)))                  // 9

	// ─────────────────────────────────────────────
	// 4. GOTCHA: Indexing gives bytes, not runes
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Indexing gotcha --")
	fmt.Printf("s[0] = %d (%c) — a byte, not a rune!\n", s[0], s[0]) // 72 (H)
	fmt.Printf("s[7] = %d — NOT '世', it's a UTF-8 byte!\n", s[7])

	// To get the rune at position, convert to []rune first:
	runes := []rune(s)
	fmt.Printf("[]rune(s)[7] = %c — now it's '世'\n", runes[7])

	// ─────────────────────────────────────────────
	// 5. for-range iterates RUNES (not bytes)
	// ─────────────────────────────────────────────
	fmt.Println("\n-- for-range (rune iteration) --")
	for i, r := range s {
		fmt.Printf("  byte offset %2d: rune %c (U+%04X)\n", i, r, r)
	}
	// Note: byte offsets jump by 3 for CJK characters

	// ─────────────────────────────────────────────
	// 6. for-i iterates BYTES
	// ─────────────────────────────────────────────
	fmt.Println("\n-- for-i (byte iteration) --")
	for i := 0; i < len(s); i++ {
		fmt.Printf("  byte[%2d] = 0x%02X\n", i, s[i])
	}

	// ─────────────────────────────────────────────
	// 7. []byte vs []rune conversion
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Conversions --")
	bytes := []byte(s)
	fmt.Printf("[]byte: %v\n", bytes)
	fmt.Printf("[]rune: %v\n", runes)

	// Back to string:
	fmt.Printf("string([]byte): %s\n", string(bytes))
	fmt.Printf("string([]rune): %s\n", string(runes))

	// Single rune/byte to string:
	fmt.Printf("string(65):   %s\n", string(rune(65))) // "A"
	fmt.Printf("string('世'): %s\n", string('世'))       // "世"

	// ─────────────────────────────────────────────
	// 8. UTF-8 encoding details
	// ─────────────────────────────────────────────
	fmt.Println("\n-- UTF-8 encoding --")
	// UTF-8 uses 1-4 bytes per character:
	// U+0000–U+007F:   1 byte  (ASCII)
	// U+0080–U+07FF:   2 bytes (Latin, Greek, Cyrillic, etc.)
	// U+0800–UFFFF:    3 bytes (CJK, most scripts)
	// U+10000–U+10FFFF: 4 bytes (emojis, historic scripts)
	chars := []rune{'A', 'é', '世', '🚀'}
	for _, c := range chars {
		buf := make([]byte, 4)
		n := utf8.EncodeRune(buf, c)
		fmt.Printf("  %c (U+%04X): %d bytes → %v\n", c, c, n, buf[:n])
	}

	// ─────────────────────────────────────────────
	// Comparison: Go vs Rust
	// ─────────────────────────────────────────────
	// Go:   byte=uint8, rune=int32, string=[]byte (UTF-8)
	// Rust: u8,         char (4 bytes, Unicode scalar), String/&str (UTF-8)
	// Both: strings are UTF-8, indexing gives bytes, iteration gives chars
	// Rust forbids s[0] on &str entirely; Go allows it but gives a byte.
}
