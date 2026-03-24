//go:build ignore

// Section 4, Topic 34: Recursion (No TCO in Go)
//
// Go supports recursion but has NO tail-call optimization (TCO).
// Deep recursion will overflow the stack (goroutine stack starts at ~8KB,
// grows dynamically but has limits).
//
// GOTCHA: No TCO means convert tail-recursive to iterative for large N.
// GOTCHA: Goroutine stack is segmented and growable but not infinite.
//
// Run: go run examples/s04_recursion.go

package main

import "fmt"

// 1. Classic factorial
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

// 2. Fibonacci (naive - exponential)
func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

// 3. Iterative factorial (preferred for large N)
func factorialIterative(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

// 4. Binary search (recursive)
func binarySearch(arr []int, target, lo, hi int) int {
	if lo > hi {
		return -1
	}
	mid := lo + (hi-lo)/2
	switch {
	case arr[mid] == target:
		return mid
	case arr[mid] < target:
		return binarySearch(arr, target, mid+1, hi)
	default:
		return binarySearch(arr, target, lo, mid-1)
	}
}

func main() {
	fmt.Println("=== Recursion ===")
	fmt.Println()

	// Factorial
	fmt.Println("-- Factorial --")
	for _, n := range []int{0, 1, 5, 10} {
		fmt.Printf("  %d! = %d\n", n, factorial(n))
	}

	// Fibonacci
	fmt.Println("\n-- Fibonacci --")
	for i := 0; i <= 10; i++ {
		fmt.Printf("  fib(%d) = %d\n", i, fib(i))
	}

	// Binary search
	fmt.Println("\n-- Binary search --")
	arr := []int{1, 3, 5, 7, 9, 11, 13}
	fmt.Println("  Array:", arr)
	fmt.Println("  Search 7:", binarySearch(arr, 7, 0, len(arr)-1))
	fmt.Println("  Search 4:", binarySearch(arr, 4, 0, len(arr)-1))

	// Iterative vs recursive
	fmt.Println("\n-- Iterative preferred for large N --")
	fmt.Println("  factorial(20) =", factorialIterative(20))
}
