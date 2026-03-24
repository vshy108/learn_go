//go:build ignore

// Section 18, Topic 133: strings and strconv
//
// strings: String manipulation functions.
// strconv: String ↔ number conversions.
//
// Run: go run examples/s18_strings_strconv.go

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("=== strings and strconv ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. strings package
	// ─────────────────────────────────────────────
	s := "Hello, World!"

	fmt.Println("Contains:", strings.Contains(s, "World"))
	fmt.Println("HasPrefix:", strings.HasPrefix(s, "Hello"))
	fmt.Println("HasSuffix:", strings.HasSuffix(s, "!"))
	fmt.Println("Index:", strings.Index(s, "World"))
	fmt.Println("ToUpper:", strings.ToUpper(s))
	fmt.Println("ToLower:", strings.ToLower(s))
	fmt.Println("TrimSpace:", strings.TrimSpace("  hello  "))
	fmt.Println("Replace:", strings.Replace(s, "World", "Go", 1))
	fmt.Println("ReplaceAll:", strings.ReplaceAll("aabaa", "a", "x"))
	fmt.Println("Count:", strings.Count("aabaa", "a"))
	fmt.Println("Repeat:", strings.Repeat("Go! ", 3))

	// ─────────────────────────────────────────────
	// 2. Split and Join
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Split/Join --")
	parts := strings.Split("a,b,c,d", ",")
	fmt.Println("Split:", parts)
	fmt.Println("Join:", strings.Join(parts, " | "))

	fields := strings.Fields("  hello   world  ") // splits on whitespace
	fmt.Println("Fields:", fields)

	// ─────────────────────────────────────────────
	// 3. strings.Builder (efficient concatenation)
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Builder --")
	var sb strings.Builder
	for i := 0; i < 5; i++ {
		sb.WriteString(fmt.Sprintf("item%d ", i))
	}
	fmt.Println("Built:", sb.String())

	// ─────────────────────────────────────────────
	// 4. strconv: string ↔ number
	// ─────────────────────────────────────────────
	fmt.Println("\n-- strconv --")

	// String → int:
	n, err := strconv.Atoi("42")
	fmt.Printf("Atoi(\"42\"): %d, err=%v\n", n, err)

	n2, err := strconv.Atoi("abc")
	fmt.Printf("Atoi(\"abc\"): %d, err=%v\n", n2, err)

	// Int → string:
	s2 := strconv.Itoa(42)
	fmt.Printf("Itoa(42): %q\n", s2)

	// Float parsing:
	f, _ := strconv.ParseFloat("3.14", 64)
	fmt.Printf("ParseFloat: %f\n", f)

	// Bool parsing:
	b, _ := strconv.ParseBool("true")
	fmt.Printf("ParseBool: %t\n", b)

	// Format float:
	fs := strconv.FormatFloat(3.14159, 'f', 2, 64)
	fmt.Printf("FormatFloat: %q\n", fs)

	// ─────────────────────────────────────────────
	// 5. strings.NewReader
	// ─────────────────────────────────────────────
	fmt.Println("\n-- NewReader --")
	r := strings.NewReader("hello")
	fmt.Printf("Reader: len=%d\n", r.Len())
	// Implements io.Reader — can be used with io.Copy, json.Decoder, etc.
}
