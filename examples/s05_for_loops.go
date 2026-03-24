//go:build ignore

// Section 5, Topic 36: for Loops — Go's Only Loop Construct
//
// Go has ONLY `for` — no while, no do-while. The `for` keyword handles all:
//   1. Three-component for:  for init; cond; post { }
//   2. While-style for:      for condition { }
//   3. Infinite for:          for { }
//
// GOTCHA: No parentheses around the for clauses.
// GOTCHA: `break` breaks the innermost loop; use labels for outer loops.
// GOTCHA: Variables declared in the init statement are scoped to the loop.
//
// Run: go run examples/s05_for_loops.go

package main

import "fmt"

func main() {
	fmt.Println("=== for Loops ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. Classic three-component for
	// ─────────────────────────────────────────────
	fmt.Println("-- Classic for --")
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
	// i is NOT accessible here — scoped to the loop

	// ─────────────────────────────────────────────
	// 2. While-style (condition only)
	// ─────────────────────────────────────────────
	fmt.Println("\n-- While-style --")
	n := 1
	for n < 100 {
		n *= 2
	}
	fmt.Println("First power of 2 >= 100:", n)

	// ─────────────────────────────────────────────
	// 3. Infinite loop with break
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Infinite loop + break --")
	count := 0
	for {
		if count >= 5 {
			break
		}
		fmt.Printf("%d ", count)
		count++
	}
	fmt.Println()

	// ─────────────────────────────────────────────
	// 4. continue — skip iteration
	// ─────────────────────────────────────────────
	fmt.Println("\n-- continue (skip odds) --")
	for i := 0; i < 10; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// ─────────────────────────────────────────────
	// 5. Nested loops
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Nested loops --")
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("(%d,%d) ", i, j)
		}
		fmt.Println()
	}

	// ─────────────────────────────────────────────
	// 6. Multiple variables in for init
	// ─────────────────────────────────────────────
	fmt.Println("-- Multiple variables --")
	for i, j := 0, 10; i < j; i, j = i+1, j-1 {
		fmt.Printf("i=%d, j=%d\n", i, j)
	}

	// ─────────────────────────────────────────────
	// 7. GOTCHA: No while or do-while keywords
	// ─────────────────────────────────────────────
	// while (condition) { }     // ERROR: Go has no 'while'
	// do { } while (condition)  // ERROR: Go has no 'do'
	// Use `for condition { }` for while, `for { ... if !cond { break } }` for do-while

	// ─────────────────────────────────────────────
	// 8. GOTCHA: for without braces
	// ─────────────────────────────────────────────
	// for i := 0; i < 5; i++
	//     fmt.Println(i)  // ERROR: braces required

	// ─────────────────────────────────────────────
	// Comparison: Go vs Rust
	// ─────────────────────────────────────────────
	// Go:   for i := 0; i < 10; i++ { }
	// Rust: for i in 0..10 { }
	// Go:   for condition { }     (while)
	// Rust: while condition { }
	// Go:   for { }               (infinite)
	// Rust: loop { }
}
