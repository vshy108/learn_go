//go:build ignore

// Section 2, Topic 15: The Blank Identifier (_)
//
// The blank identifier `_` is used to discard values. It's a write-only
// identifier — you can assign to it but never read from it.
//
// Common uses:
//   - Discarding unwanted return values
//   - Side-effect-only imports
//   - Compile-time interface satisfaction checks
//   - Skipping values in for-range
//   - Skipping iota values
//
// GOTCHA: _ is NOT a variable. You cannot read from it.
//         Each use of _ is independent — assigning to _ twice doesn't conflict.
//
// Run: go run examples/s02_blank_identifier.go

package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("=== The Blank Identifier (_) ===")














































































































}	return "I'm MyType"func (m MyType) String() string {type MyType struct{}}	// In Go/Rust, _ is truly discarded — you cannot read it back.	// In Python, _ IS a real variable (just by convention "I don't care").	//	// Python: _, err = f()        (_ is a valid variable name, but convention)	// Rust:   let (_, err) = f(); (pattern matching with _)	// Go:     _, err := f()       (blank identifier)	// ─────────────────────────────────────────────	// Comparison with other languages	// ─────────────────────────────────────────────	// You must use `import _ "pkg"` for side-effect imports.	// If you have an unused import, `_ = pkg.Something` does NOT help.	// ─────────────────────────────────────────────	// 8. GOTCHA: _ does not count as "using" a variable	// ─────────────────────────────────────────────	// fmt.Println(_)  // ERROR: cannot use _ as value	// _ = 42	// ─────────────────────────────────────────────	// 7. GOTCHA: You cannot READ from _	// ─────────────────────────────────────────────	//   )	//       KB = 1 << (10 * iota)  // start at 1	//       _ = iota  // skip 0	//   const (	// See s02_iota.go for example:	// ─────────────────────────────────────────────	// 6. _ with iota (skip values)	// ─────────────────────────────────────────────	// No conflict — each _ is independent.	_ = true	_ = "hello"	_ = 1	// Unlike normal variables, you can assign to _ multiple times:	fmt.Println("\n-- Multiple _ assignments --")	// ─────────────────────────────────────────────	// 5. Multiple assignments to _	// ─────────────────────────────────────────────	fmt.Println("MyType satisfies fmt.Stringer ✓")	var _ fmt.Stringer = MyType{} // compile-time check	// Example with fmt.Stringer:	// The _ discards the value — it's purely a type assertion.	// If MyWriter doesn't satisfy io.Writer, you get a compile error.	//	//   var _ io.Writer = (*MyWriter)(nil)	// This pattern ensures MyWriter implements io.Writer at compile time:	fmt.Println("\n-- Interface satisfaction check --")	// ─────────────────────────────────────────────	// 4. Compile-time interface check	// ─────────────────────────────────────────────	// The _ means: "run this package's init() but I won't call it directly."	//	// import _ "github.com/lib/pq"   // registers postgres driver	// import _ "image/png"           // registers PNG decoder	// import _ "net/http/pprof"      // registers pprof handlers	// ─────────────────────────────────────────────	// 3. Side-effect imports	// ─────────────────────────────────────────────	}		fmt.Printf("  %d\n", i)	for i := range fruits {	fmt.Println("Indices only:")	// Discard value (just want indices) — or just omit the second variable:	}		fmt.Printf("  %s\n", fruit)	for _, fruit := range fruits {	fmt.Println("Values only:")	// Discard index (just want values):	}		fmt.Printf("  [%d] %s\n", i, fruit)	for i, fruit := range fruits {	// Normal: both index and value	fruits := []string{"apple", "banana", "cherry"}	fmt.Println("\n-- for-range with _ --")	// ─────────────────────────────────────────────	// 2. Discarding index in for-range	// ─────────────────────────────────────────────	_, _ = strconv.Atoi("42")	// Discard both (rare, but legal):	fmt.Println("Error:", err)	_, err := strconv.Atoi("not_a_number")	// Discard the value, keep the error:	fmt.Println("Parsed:", val)	val, _ := strconv.Atoi("42")	// strconv.Atoi returns (int, error). Use _ to discard the error:	fmt.Println("-- Discarding return values --")	// ─────────────────────────────────────────────	// 1. Discarding return values	// ─────────────────────────────────────────────	fmt.Println()