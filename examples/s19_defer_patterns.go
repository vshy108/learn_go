//go:build ignore

// Section 19, Topic 138: Defer, Panic, Recover Patterns
//
// Advanced patterns using defer, panic, and recover.
//
// defer: Schedule cleanup that runs when function returns.
// panic: Unwind stack (use for unrecoverable errors only).
// recover: Catch panic in deferred function.
//
// GOTCHA: Deferred calls run in LIFO order (last deferred = first executed).
// GOTCHA: Deferred function arguments are evaluated at defer time, not execution time.
// GOTCHA: Named return values can be modified in deferred functions.
//
// Run: go run examples/s19_defer_patterns.go

package main

import "fmt"

func main() {
	fmt.Println("=== Defer Patterns ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. LIFO order
	// ─────────────────────────────────────────────
	fmt.Println("-- LIFO --")
	defer fmt.Println("main: deferred 1")
	defer fmt.Println("main: deferred 2")
	defer fmt.Println("main: deferred 3")
	// Output: 3, 2, 1 (after main returns)

	// ─────────────────────────────────────────────
	// 2. Arguments evaluated at defer time
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Argument capture --")
	x := 10
	defer fmt.Printf("Deferred x=%d (captured at defer time)\n", x)
	x = 20
	fmt.Printf("x=%d (current value)\n", x)

	// ─────────────────────────────────────────────
	// 3. Defer with closure (captures by reference)
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Closure capture --")
	y := 10
	defer func() {
		fmt.Printf("Deferred closure y=%d (captures reference)\n", y)
	}()
	y = 20

	// ─────────────────────────────────────────────
	// 4. Named return + defer (modify return value)
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Named return --")
	result := doubleAndAdd(5)
	fmt.Println("doubleAndAdd(5):", result)

	// ─────────────────────────────────────────────
	// 5. Resource cleanup with defer
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Resource cleanup --")
	processResource()

	// ─────────────────────────────────────────────
	// 6. Recover pattern: convert panic to error
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Panic to error --")
	err := safeOperation()
	fmt.Println("Error:", err)
	fmt.Println("Program continues after recovered panic")
}

func doubleAndAdd(n int) (result int) {
	defer func() {
		result += 10 // modify named return value
	}()
	return n * 2 // sets result=10, then defer adds 10 → 20
}

func processResource() {
	fmt.Println("  Acquiring resource")
	defer fmt.Println("  Releasing resource (always runs)")

	fmt.Println("  Processing...")
	// Even if this panics, the resource is released
}

func safeOperation() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("recovered: %v", r)
		}
	}()

	// Simulate something that panics:
	panic("something bad happened")
}
