//go:build ignore

// Section 1, Topic 6: Implicit Semicolons and Brace Placement Rules
//
// Go uses implicit semicolons — the lexer inserts a semicolon at the end of
// a line if the last token is:
//   - An identifier (x, foo, int, etc.)
//   - A literal (42, "hello", true, etc.)
//   - One of: break, continue, fallthrough, return
//   - One of: ++, --, ), ], }
//
// This is WHY Go enforces the "opening brace on the same line" style.
// It's not just a style preference — putting { on the next line is a COMPILE ERROR.
//
// GOTCHA: This is the #1 surprise for developers coming from C/C++/C#/Java
//         where opening brace placement is a style choice, not a language rule.
//
// Run: go run examples/s01_semicolons_braces.go

package main

import "fmt"

func main() {
	fmt.Println("=== Implicit Semicolons and Brace Placement ===")
	fmt.Println()

	// ─────────────────────────────────────
	// 1. Correct brace placement
	// ─────────────────────────────────────
	// Opening brace MUST be on the same line as the statement.
	if true {
		fmt.Println("✓ Brace on same line — this compiles")
	}

	// WRONG (doesn't compile):
	// if true
	// {
	//     fmt.Println("✗ This is a syntax error!")
	// }
	// The lexer sees: if true; { ... }
	// Which is: if true; followed by a block — syntax error!

	// ─────────────────────────────────────
	// 2. How semicolons are inserted
	// ─────────────────────────────────────
	// These are logically equivalent:
	x := 42        // → x := 42;
	fmt.Println(x) // → fmt.Println(x);

	// The semicolons are invisible but real. You CAN write them explicitly:
	y := 10
	z := 20
	fmt.Println("y+z =", y+z)
	// But this is non-idiomatic. gofmt will leave single-line multiples alone,
	// but it's rarely used outside of for-loop initializers.

	// ─────────────────────────────────────
	// 3. Multi-line expressions
	// ─────────────────────────────────────
	// To break a long expression across lines, the line break must come
	// AFTER an operator or comma (so no semicolon is inserted):

	result := 1 +
		2 +
		3 // ✓ OK: line ends with +, so no semicolon after 1 or 2
	fmt.Println("Multi-line sum:", result)

	// WRONG:
	// result := 1
	//         + 2    // ERROR: 1; +2 — semicolon after 1!

	// Same for function calls:
	fmt.Println(
		"This",
		"works",
		"fine",
	) // ✓ The trailing comma after "fine", is REQUIRED for multi-line

	// GOTCHA: In multi-line composite literals and function calls,
	// the last element MUST have a trailing comma:
	names := []string{
		"Alice",
		"Bob",
		"Charlie", // ← trailing comma required! Without it: syntax error
	}
	fmt.Println("Names:", names)

	// ─────────────────────────────────────
	// 4. For loops and semicolons
	// ─────────────────────────────────────
	// The for loop uses explicit semicolons to separate clauses:
	for i := 0; i < 3; i++ {
		fmt.Printf("i=%d ", i)
	}
	fmt.Println()

	// ─────────────────────────────────────
	// 5. return and implicit semicolons
	// ─────────────────────────────────────
	// A bare `return` gets a semicolon, so expressions after it
	// on the next line are unreachable:
	//
	//   return      // → return;
	//   x + y       // unreachable — this is a new statement!
	//
	// For multi-line returns, keep the expression on the same line
	// or use parentheses (rare in Go).

	// ─────────────────────────────────────
	// 6. Comparison with other languages
	// ─────────────────────────────────────
	// JavaScript: also has ASI (Automatic Semicolon Insertion) but with
	//             different rules and many more edge cases/gotchas
	// Python:     uses indentation, no semicolons needed
	// Rust:       explicit semicolons required (expressions vs statements)
	// Go:         implicit semicolons with simple, predictable rules

	fmt.Println("\nGo's semicolon rules are simple: if a line ends with an")
	fmt.Println("identifier, literal, or certain keywords/punctuation, a")
	fmt.Println("semicolon is inserted. This forces consistent brace style.")
}
