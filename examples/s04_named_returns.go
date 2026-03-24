//go:build ignore

// Section 4, Topic 28: Named Return Values (Naked Return)
//
// Go allows naming return values in the function signature. Named returns
// are initialized to their zero values and can be set during the function.
// A bare `return` (naked return) returns the current values of named returns.
//
// GOTCHA: Naked returns are discouraged in long functions — they hurt readability.
//         Only use them in short functions (a few lines).
// GOTCHA: Named returns are still useful for documentation even without naked returns.
// GOTCHA: Named return values can be shadowed by local variables (subtle bug!).
//
// Run: go run examples/s04_named_returns.go

package main

import (
	"errors"
	"fmt"
)

// ─────────────────────────────────────────────
// 1. Named returns with naked return
// ─────────────────────────────────────────────
func divide(a, b float64) (result float64, err error) {
	if b == 0 {
		err = errors.New("division by zero")
		return // naked return: returns result=0.0, err=<error>
	}
	result = a / b
	return // naked return: returns result=<computed>, err=nil
}

// ─────────────────────────────────────────────
// 2. Named returns for documentation (without naked return)
// ─────────────────────────────────────────────
func swap(a, b string) (first, second string) {
	first = b
	second = a
	return first, second // explicit return — more readable
}

// ─────────────────────────────────────────────
// 3. Named returns with defer (common pattern)
// ─────────────────────────────────────────────
func doWork() (result int, err error) {
	// Named returns + defer is powerful for cleanup:
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("recovered: %v", r)
			result = -1
		}
	}()

	// Simulate work
	result = 42
	return
}

func main() {
	fmt.Println("=== Named Return Values ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// Named returns in action
	// ─────────────────────────────────────────────
	result, err := divide(10, 3)
	fmt.Printf("divide(10, 3): result=%.4f, err=%v\n", result, err)

	result, err = divide(10, 0)
	fmt.Printf("divide(10, 0): result=%.4f, err=%v\n", result, err)

	// ─────────────────────────────────────────────
	// Named returns for readability
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Named returns as documentation --")
	f, s := swap("hello", "world")
	fmt.Printf("swap: first=%s, second=%s\n", f, s)
	// The signature `(first, second string)` documents what each return means.

	// ─────────────────────────────────────────────
	// Named returns with defer
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Named returns + defer --")
	r, e := doWork()
	fmt.Printf("doWork: result=%d, err=%v\n", r, e)

	// ─────────────────────────────────────────────
	// GOTCHA: Shadowing named returns
	// ─────────────────────────────────────────────
	// func bad() (err error) {
	//     if true {
	//         err := fmt.Errorf("oops")  // := creates NEW local variable!
	//         fmt.Println(err)            // prints "oops"
	//     }
	//     return  // returns nil! The named `err` was never set.
	// }
	// Fix: use = instead of := to assign to the named return.

	// ─────────────────────────────────────────────
	// Style guidance
	// ─────────────────────────────────────────────
	// DO:  Use named returns for documentation in exported functions
	// DO:  Use named returns with defer for error handling
	// DO:  Use naked returns only in very short functions (< 5 lines)
	// DON'T: Use naked returns in long functions — hard to track what's returned
	// DON'T: Mix naked and explicit returns in the same function

	fmt.Println("\nNamed returns are best for short functions and defer patterns.")
}
