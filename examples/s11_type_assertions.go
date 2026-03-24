//go:build ignore

// Section 11, Topic 83: Type Assertions
//
// Type assertion extracts the concrete type from an interface value.
//   val := i.(Type)          — panics if wrong type
//   val, ok := i.(Type)      — returns false if wrong type
//
// GOTCHA: Asserting wrong type without comma-ok causes PANIC.
// GOTCHA: Can only type-assert on interface values, not concrete types.
//
// Run: go run examples/s11_type_assertions.go

package main

import "fmt"

func main() {
	fmt.Println("=== Type Assertions ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. Basic type assertion
	// ─────────────────────────────────────────────
	var i any = "hello"

	s := i.(string)
	fmt.Printf("Asserted string: %q\n", s)

	// ─────────────────────────────────────────────
	// 2. GOTCHA: Wrong type = panic!
	// ─────────────────────────────────────────────
	// n := i.(int)  // PANIC: interface conversion: interface {} is string, not int

	// ─────────────────────────────────────────────
	// 3. Safe assertion with comma-ok
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Comma-ok --")
	if s, ok := i.(string); ok {
		fmt.Printf("Is string: %q\n", s)
	}
	if n, ok := i.(int); ok {
		fmt.Printf("Is int: %d\n", n)
	} else {
		fmt.Println("Not an int")
	}

	// ─────────────────────────────────────────────
	// 4. Asserting from specific interface
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Specific interface --")
	type Stringer interface {
		String() string
	}

	type MyType struct{ Name string }
	mt := MyType{Name: "test"}
	_ = mt

	var val any = 42
	if _, ok := val.(Stringer); ok {
		fmt.Println("Implements Stringer")
	} else {
		fmt.Println("Does NOT implement Stringer")
	}

	// ─────────────────────────────────────────────
	// 5. Multiple assertions
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Process different types --")
	values := []any{42, "hello", 3.14, true, nil}
	for _, v := range values {
		process(v)
	}
}

func process(v any) {
	if s, ok := v.(string); ok {
		fmt.Printf("  string: %q (len=%d)\n", s, len(s))
		return
	}
	if n, ok := v.(int); ok {
		fmt.Printf("  int: %d (doubled=%d)\n", n, n*2)
		return
	}
	fmt.Printf("  other: %v (%T)\n", v, v)
}
