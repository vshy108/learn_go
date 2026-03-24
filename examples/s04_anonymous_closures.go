//go:build ignore

// Section 4, Topic 31: Anonymous Functions and Closures
//
// Anonymous functions have no name and can be declared inline.
// Closures are anonymous functions that capture variables from their
// enclosing scope.
//
// GOTCHA: Closures capture variables BY REFERENCE (the variable itself,
//         not a copy of its value). This is a common source of bugs in loops.
// GOTCHA: The classic loop closure bug: goroutines capturing loop variable.
//
// Run: go run examples/s04_anonymous_closures.go

package main

import "fmt"

func main() {
	fmt.Println("=== Anonymous Functions and Closures ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. Basic anonymous function
	// ─────────────────────────────────────────────
	fmt.Println("-- Anonymous function --")
	square := func(x int) int {
		return x * x
	}
	fmt.Println("square(5) =", square(5))

	// ─────────────────────────────────────────────
	// 2. Immediately invoked function expression (IIFE)
	// ─────────────────────────────────────────────
	fmt.Println("\n-- IIFE --")
	result := func(a, b int) int {
		return a + b
	}(3, 5) // immediately called with (3, 5)
	fmt.Println("IIFE result:", result)

	// ─────────────────────────────────────────────
	// 3. Closures — capturing variables
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Closures --")
	counter := makeCounter()
	fmt.Println("count:", counter()) // 1
	fmt.Println("count:", counter()) // 2
	fmt.Println("count:", counter()) // 3

	// Each call to makeCounter creates a new independent counter:
	counter2 := makeCounter()
	fmt.Println("counter2:", counter2()) // 1 (independent)
	fmt.Println("counter1:", counter())  // 4 (continues)

	// ─────────────────────────────────────────────
	// 4. Closures capture by reference
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Capture by reference --")
	x := 10
	increment := func() {
		x++ // modifies the outer x
	}
	increment()
	increment()
	fmt.Println("x after 2 increments:", x) // 12

	// ─────────────────────────────────────────────
	// 5. GOTCHA: Loop closure bug
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Loop closure bug --")

	// BUG: All functions capture the SAME variable `i`
	funcs := make([]func(), 5)
	for i := 0; i < 5; i++ {
		funcs[i] = func() {
			fmt.Printf("%d ", i) // captures `i` by reference
		}
	}
	fmt.Print("Buggy:  ")
	for _, f := range funcs {
		f() // prints "5 5 5 5 5" — all see final value of i
	}
	fmt.Println()

	// FIX 1: Shadow the variable with a local copy
	for i := 0; i < 5; i++ {
		i := i // shadow with new variable
		funcs[i] = func() {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Print("Fixed1: ")
	for _, f := range funcs {
		f() // prints "0 1 2 3 4"
	}
	fmt.Println()

	// FIX 2: Pass as parameter
	for i := 0; i < 5; i++ {
		funcs[i] = func(n int) func() {
			return func() { fmt.Printf("%d ", n) }
		}(i)
	}
	fmt.Print("Fixed2: ")
	for _, f := range funcs {
		f() // prints "0 1 2 3 4"
	}
	fmt.Println()

	// NOTE: Since Go 1.22, the loop variable is per-iteration by default!
	// So the bug above is fixed in Go 1.22+. But the pattern is still
	// important to understand for older code and goroutines.

	// ─────────────────────────────────────────────
	// 6. Closure as function argument
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Closure as argument --")
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evens := filter(numbers, func(n int) bool {
		return n%2 == 0
	})
	fmt.Println("Evens:", evens)

	greaterThan5 := filter(numbers, func(n int) bool {
		return n > 5
	})
	fmt.Println("Greater than 5:", greaterThan5)
}

func makeCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func filter(nums []int, predicate func(int) bool) []int {
	var result []int
	for _, n := range nums {
		if predicate(n) {
			result = append(result, n)
		}
	}
	return result
}
