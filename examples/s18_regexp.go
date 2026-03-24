//go:build ignore

// Section 18, Topic 137: regexp Package
//
// Go's regexp package implements RE2 syntax.
// RE2 guarantees linear time matching (no catastrophic backtracking).
//
// GOTCHA: Go regex does NOT support lookahead/lookbehind.
// GOTCHA: Use raw strings `...` for patterns to avoid escaping.
// GOTCHA: Compile() returns error; MustCompile() panics on bad pattern.
//
// Run: go run examples/s18_regexp.go

package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println("=== regexp Package ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. Basic matching
	// ─────────────────────────────────────────────
	matched, _ := regexp.MatchString(`\d+`, "abc123def")
	fmt.Println("Has digits:", matched)

	// ─────────────────────────────────────────────
	// 2. Compiled regex (preferred for reuse)
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Compiled --")
	re := regexp.MustCompile(`\b\w+@\w+\.\w+\b`)

	text := "Contact alice@example.com or bob@test.org"
	fmt.Println("Match:", re.MatchString(text))
	fmt.Println("Find:", re.FindString(text))
	fmt.Println("FindAll:", re.FindAllString(text, -1))

	// ─────────────────────────────────────────────
	// 3. Submatches (capture groups)
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Capture groups --")
	dateRe := regexp.MustCompile(`(\d{4})-(\d{2})-(\d{2})`)
	match := dateRe.FindStringSubmatch("Today is 2024-03-15")
	if match != nil {
		fmt.Println("Full match:", match[0])
		fmt.Println("Year:", match[1])
		fmt.Println("Month:", match[2])
		fmt.Println("Day:", match[3])
	}

	// ─────────────────────────────────────────────
	// 4. Named capture groups
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Named groups --")
	namedRe := regexp.MustCompile(`(?P<name>\w+):(?P<value>\d+)`)
	match2 := namedRe.FindStringSubmatch("port:8080")
	for i, name := range namedRe.SubexpNames() {
		if name != "" {
			fmt.Printf("  %s: %s\n", name, match2[i])
		}
	}

	// ─────────────────────────────────────────────
	// 5. Replace
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Replace --")
	re2 := regexp.MustCompile(`\d+`)
	result := re2.ReplaceAllString("abc123def456", "NUM")
	fmt.Println("Replace:", result)

	// Replace with function:
	result2 := re2.ReplaceAllStringFunc("abc123def456", func(s string) string {
		return "[" + s + "]"
	})
	fmt.Println("ReplaceFunc:", result2)

	// ─────────────────────────────────────────────
	// 6. Split
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Split --")
	splitRe := regexp.MustCompile(`[,;\s]+`)
	parts := splitRe.Split("a, b; c  d", -1)
	fmt.Println("Split:", parts)

	// ─────────────────────────────────────────────
	// 7. RE2 limitations
	// ─────────────────────────────────────────────
	// NOT supported:
	// - Lookahead: (?=...) (?!...)
	// - Lookbehind: (?<=...) (?<!...)
	// - Backreferences: \1
	// - Possessive quantifiers: a++
	// These guarantee linear time O(n) matching.
}
