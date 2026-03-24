//go:build ignore

// Section 4, Topic 32: defer - Deferred Function Calls
//
// defer schedules a function call to run when the surrounding function returns.
// Deferred calls execute in LIFO (last-in, first-out) order.
//
// GOTCHA: Arguments are evaluated at defer time, not execution time.
// GOTCHA: Deferred closures capture variables by reference.
// GOTCHA: defer runs even when the function panics.
// GOTCHA: Named return values can be modified in deferred functions.
//
// Run: go run examples/s04_defer.go

package main

import "fmt"

func main() {
	fmt.Println("=== defer ===")
	fmt.Println()

	// 1. Basic defer
	fmt.Println("-- Basic --")
	fmt.Println("First")
	defer fmt.Println("Deferred (runs last)")
	fmt.Println("Second")

	// 2. LIFO order
	fmt.Println("\n-- LIFO order --")
	for i := 1; i <= 3; i++ {
		defer fmt.Printf("  deferred %d\n", i)
	}
	fmt.Println("  After loop")

	// 3. Arguments evaluated at defer time
	fmt.Println("\n-- Arg evaluation --")
	x := 10
	defer fmt.Printf("  Deferred x=%d (captured at defer time)\n", x)
	x = 20
	fmt.Printf("  Current x=%d\n", x)

	// 4. Closure captures reference
	fmt.Println("\n-- Closure capture --")
	y := 10
	defer func() {
		fmt.Printf("  Closure y=%d (reference, sees final value)\n", y)
	}()
	y = 999

	// 5. Named return modification
	fmt.Println("\n-- Named return --")
	result := namedReturn()
	fmt.Println("  namedReturn() =", result)

	// 6. Resource cleanup pattern
	fmt.Println("\n-- Cleanup pattern --")
	processFile()
}

func namedReturn() (result int) {
	defer func() {
		result += 100 // modify the named return
	}()
	return 42 // sets result=42, then defer adds 100 -> 142
}

func processFile() {
	fmt.Println("  Open file")
	defer fmt.Println("  Close file (deferred)")
	fmt.Println("  Read file")
	fmt.Println("  Process data")
}
