//go:build ignore

// Section 12, Topic 94: panic, recover, and defer for Error Handling
//
// panic: Stops normal execution, unwinds the stack, runs deferred functions.
// recover: Catches a panic inside a deferred function. Returns the panic value.
// defer: Schedules a function to run when the enclosing function returns.
//
// GOTCHA: panic is NOT for normal error handling. Use return error.
// GOTCHA: recover only works inside a deferred function.
// GOTCHA: recover returns nil if no panic occurred.
//
// When to panic:
//   - Truly unrecoverable situations (corrupt state)
//   - Programming bugs (should never happen in correct code)
//   - Package initialization failures
//
// Run: go run examples/s12_panic_recover.go

package main

import "fmt"

func main() {
	fmt.Println("=== panic, recover, defer ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. Basic panic
	// ─────────────────────────────────────────────
	// panic("something terrible happened")
	// → prints stack trace and exits

	// ─────────────────────────────────────────────
	// 2. recover catches panic
	// ─────────────────────────────────────────────
	fmt.Println("Before safeDiv")
	result := safeDiv(10, 0)
	fmt.Printf("safeDiv(10, 0) = %d\n", result) // 0, recovered from panic
	fmt.Println("After safeDiv — program continues!")

	result = safeDiv(10, 2)
	fmt.Printf("safeDiv(10, 2) = %d\n", result)

	// ─────────────────────────────────────────────
	// 3. Defer runs on panic
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Defer during panic --")
	func() {
		defer fmt.Println("  deferred: cleanup runs even on panic")
		defer fmt.Println("  deferred: this runs first (LIFO)")
		fmt.Println("  About to trigger controlled panic...")
		// Use a nested recover to catch this:
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("  Recovered: %v\n", r)
			}
		}()
		panic("controlled panic")
	}()
	fmt.Println("Continued after recovered panic")

	// ─────────────────────────────────────────────
	// 4. Converting panic to error
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Panic to error --")
	val, err := safeParse("hello")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Value:", val)
	}

	// ─────────────────────────────────────────────
	// 5. Common stdlib panics
	// ─────────────────────────────────────────────
	// - Index out of bounds: a[10] on a 5-element array
	// - Nil pointer dereference: var p *int; *p
	// - Sending on closed channel
	// - Type assertion without comma-ok: i.(int) on a string

	// ─────────────────────────────────────────────
	// 6. When to use panic vs error
	// ─────────────────────────────────────────────
	// Use error:
	//   - File not found, network error, invalid input
	//   - Any expected failure condition
	// Use panic:
	//   - Impossible state (logic error, programmer mistake)
	//   - Failed to initialize critical resource
	//   - Index out of bounds (runtime does this)
}

// safeDiv recovers from division-by-zero panic
func safeDiv(a, b int) (result int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("  Recovered from panic: %v\n", r)
			result = 0
		}
	}()
	return a / b // panics if b == 0
}

// safeParse converts panics to errors
func safeParse(s string) (val int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("parse panic: %v", r)
		}
	}()
	// Simulate a function that panics on bad input
	if s == "" {
		panic("empty input")
	}
	if s[0] < '0' || s[0] > '9' {
		panic("not a number")
	}
	return int(s[0] - '0'), nil
}
