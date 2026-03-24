//go:build ignore

// Section 11, Topic 82: Empty Interface (any / interface{})
//
// The empty interface has zero methods, so ALL types satisfy it.
//   interface{}  — pre-Go 1.18
//   any          — Go 1.18+ alias for interface{}
//
// Used for generic containers, variadic functions, and JSON parsing.
//
// GOTCHA: You lose type safety with empty interface.
// GOTCHA: You must use type assertion to access the underlying value.
// GOTCHA: Prefer generics (Go 1.18+) over empty interface when possible.
//
// Run: go run examples/s11_empty_interface.go

package main

import "fmt"

func main() {
	fmt.Println("=== Empty Interface (any) ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. any can hold anything
	// ─────────────────────────────────────────────
	var a any
	a = 42
	fmt.Printf("int: %v (type: %T)\n", a, a)
	a = "hello"
	fmt.Printf("string: %v (type: %T)\n", a, a)
	a = []int{1, 2, 3}
	fmt.Printf("slice: %v (type: %T)\n", a, a)
	a = struct{ Name string }{"Alice"}
	fmt.Printf("struct: %v (type: %T)\n", a, a)

	// ─────────────────────────────────────────────
	// 2. Slice of any (heterogeneous)
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Heterogeneous slice --")
	items := []any{42, "hello", 3.14, true, nil}
	for i, item := range items {
		fmt.Printf("  [%d] %v (type: %T)\n", i, item, item)
	}

	// ─────────────────────────────────────────────
	// 3. Function accepting any
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Function with any --")
	printAnything(42)
	printAnything("world")
	printAnything([]int{1, 2})

	// fmt.Println already accepts any:
	// func Println(a ...any) (n int, err error)

	// ─────────────────────────────────────────────
	// 4. Map with any values
	// ─────────────────────────────────────────────
	fmt.Println("\n-- map[string]any --")
	data := map[string]any{
		"name":   "Alice",
		"age":    30,
		"scores": []int{95, 87, 92},
	}
	for k, v := range data {
		fmt.Printf("  %s: %v (%T)\n", k, v, v)
	}

	// ─────────────────────────────────────────────
	// 5. any == interface{} (they are identical)
	// ─────────────────────────────────────────────
	var old interface{} = 42
	var new_ any = 42
	fmt.Printf("\ninterface{}: %v, any: %v (same thing)\n", old, new_)

	// ─────────────────────────────────────────────
	// 6. GOTCHA: Must type-assert to use
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Must type-assert --")
	var val any = 42
	// result := val + 1  // ERROR: mismatched types any and int
	result := val.(int) + 1 // type assertion
	fmt.Printf("42 + 1 = %d\n", result)
}

func printAnything(v any) {
	fmt.Printf("  printAnything: %v (%T)\n", v, v)
}
