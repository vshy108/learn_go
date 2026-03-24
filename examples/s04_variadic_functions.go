//go:build ignore

// Section 4, Topic 29: Variadic Functions (...)
//
// Variadic functions accept a variable number of arguments of the same type.
// The syntax uses `...` before the type: func name(args ...int)
//
// Inside the function, the variadic parameter is a slice of the given type.
//
// GOTCHA: Only the LAST parameter can be variadic.
// GOTCHA: To pass a slice to a variadic function, use the ... spread operator.
// GOTCHA: fmt.Println is variadic: func Println(a ...any) (int, error)
//
// Run: go run examples/s04_variadic_functions.go

package main

import "fmt"

// ─────────────────────────────────────────────
// 1. Basic variadic function
// ─────────────────────────────────────────────
func sum(numbers ...int) int {
	total := 0
	for _, n := range numbers {
		total += n
	}
	return total
}

// ─────────────────────────────────────────────
// 2. Variadic with other parameters
// ─────────────────────────────────────────────
func greetAll(greeting string, names ...string) {
	for _, name := range names {
		fmt.Printf("%s, %s!\n", greeting, name)
	}
}

// ─────────────────────────────────────────────
// 3. Variadic with any (interface{})
// ─────────────────────────────────────────────
func printAll(args ...any) {
	for i, arg := range args {
		fmt.Printf("  [%d] %v (%T)\n", i, arg, arg)
	}
}

func main() {
	fmt.Println("=== Variadic Functions ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// Basic usage — pass individual values
	// ─────────────────────────────────────────────
	fmt.Println("-- sum --")
	fmt.Println("sum() =", sum())                                   // 0 (empty slice)
	fmt.Println("sum(1) =", sum(1))                                 // 1
	fmt.Println("sum(1,2,3) =", sum(1, 2, 3))                       // 6
	fmt.Println("sum(1..10) =", sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)) // 55

	// ─────────────────────────────────────────────
	// Passing a slice with ... spread operator
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Passing slice with ... --")
	nums := []int{10, 20, 30}
	fmt.Println("sum(slice...) =", sum(nums...)) // spread the slice

	// GOTCHA: You can't mix individual args and spread:
	// sum(1, nums...)  // ERROR: too many arguments
	// sum(nums..., 1)  // ERROR: syntax error

	// ─────────────────────────────────────────────
	// With other parameters
	// ─────────────────────────────────────────────
	fmt.Println("\n-- greetAll --")
	greetAll("Hello", "Alice", "Bob", "Charlie")
	greetAll("Hi") // zero variadic args is OK

	names := []string{"Dave", "Eve"}
	greetAll("Hey", names...)

	// ─────────────────────────────────────────────
	// any variadic (like fmt.Println)
	// ─────────────────────────────────────────────
	fmt.Println("\n-- printAll (any) --")
	printAll(42, "hello", true, 3.14) // mixed types

	// ─────────────────────────────────────────────
	// Inside the function, variadic param IS a slice
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Variadic param is a slice --")
	showType(1, 2, 3)

	// ─────────────────────────────────────────────
	// GOTCHA: Spread shares the backing array
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Spread shares backing array --")
	original := []int{1, 2, 3}
	modifyFirst(original...)
	fmt.Println("After modifyFirst:", original) // [999 2 3] — modified!

	// ─────────────────────────────────────────────
	// Comparison: Go vs other languages
	// ─────────────────────────────────────────────
	// Go:     func sum(nums ...int)
	// Rust:   no variadic functions (use slices or macros)
	// Python: def sum(*nums)
	// JS:     function sum(...nums)
}

func showType(nums ...int) {
	fmt.Printf("  Type: %T, Value: %v\n", nums, nums) // []int
}

func modifyFirst(nums ...int) {
	if len(nums) > 0 {
		nums[0] = 999
	}
}
