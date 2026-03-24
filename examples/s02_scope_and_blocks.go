//go:build ignore

// Section 2, Topic 14: Variable Scope and Block Scoping
//
// Go uses block scoping. Inner blocks can shadow outer variables.
//
// GOTCHA: := in if/for creates a NEW variable scoped to that block.
// GOTCHA: Package-level vars are accessible from all functions in the package.
// GOTCHA: Shadowing can cause subtle bugs (the outer var is unchanged).
//
// Run: go run examples/s02_scope_and_blocks.go

package main

import "fmt"

var global = "I am global"

func main() {
	fmt.Println("=== Scope and Blocks ===")
	fmt.Println()

	// 1. Function scope
	x := 10
	fmt.Println("Function scope x:", x)

	// 2. Block scope
	{
		y := 20
		fmt.Println("Block scope y:", y)
		fmt.Println("Can see x:", x)
	}
	// fmt.Println(y) // ERROR: y not defined

	// 3. Shadowing
	fmt.Println("\n-- Shadowing --")
	val := "outer"
	fmt.Println("Before if:", val)
	if true {
		val := "inner" // NEW variable, shadows outer
		fmt.Println("Inside if:", val)
	}
	fmt.Println("After if:", val) // still "outer"

	// 4. if initializer scope
	fmt.Println("\n-- if initializer --")
	if n := computeValue(); n > 5 {
		fmt.Println("n > 5:", n)
	}
	// fmt.Println(n) // ERROR: n not in scope

	// 5. for loop scope
	fmt.Println("\n-- for scope --")
	for i := 0; i < 3; i++ {
		fmt.Printf("i=%d ", i)
	}
	fmt.Println()
	// fmt.Println(i) // ERROR: i not in scope

	// 6. Package scope
	fmt.Println("\n-- Package scope --")
	fmt.Println(global)
	helper()
}

func computeValue() int {
	return 10
}

func helper() {
	fmt.Println("helper() can see global:", global)
}
