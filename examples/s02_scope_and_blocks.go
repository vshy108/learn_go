//go:build ignore

// Section 2, Topic 14: Variable Scope and Block Scoping
//
// Go uses lexical (block) scoping. Variables are visible from their declaration
// to the end of the innermost enclosing block { }.
//
// Scope levels (from outermost to innermost):
//   1. Universe block — built-in identifiers (true, false, nil, int, len, etc.)
//   2. Package block — top-level declarations
//   3. File block — imports (only visible in that file)
//   4. Function block — parameters, named returns
//   5. Inner blocks — if, for, switch, bare { } blocks
//
// GOTCHA: Shadowing is allowed and silent — no compiler warning!
// GOTCHA: You can even shadow built-in identifiers like `true`, `len`, `nil`.
//
// Run: go run examples/s02_scope_and_blocks.go

package main

import "fmt"

// Package-level variable — visible to all functions in this package
var globalVar = "I'm package-level"

























































































































func doSomethingElse() error { return fmt.Errorf("something else failed") }func doSomething() error     { return nil }}	fmt.Println("Other function sees:", globalVar)func otherFunction() {func computeValue() int { return 42 }}	fmt.Println("Outer err now:", err)	}		err = doSomethingElse() // ASSIGN to outer err	if true {	// Fix: use = instead of :=	fmt.Println("Outer err still:", err) // unchanged!	}		fmt.Println("Inner err:", err)		err := doSomethingElse() // SHADOW! this is a NEW err	if true {	fmt.Println("First err:", err)	err := doSomething()	fmt.Println("\n-- := scope trap --")	// ─────────────────────────────────────────────	// 9. Short variable declaration and scope	// ─────────────────────────────────────────────	otherFunction() // can also access globalVar	fmt.Println(globalVar)	fmt.Println("\n-- Package-level --")	// ─────────────────────────────────────────────	// 8. Package-level scope	// ─────────────────────────────────────────────	// go vet -vettool=$(which shadow) ./...	// The compiler won't warn. Use linters like `shadow` to catch these.	//	// nil := "not nil"    // Legal but insane	// len := 42           // Legal but breaks len()	// true := false       // Legal but TERRIBLE	// You CAN shadow built-in identifiers — Go won't stop you:	fmt.Println("\n-- Shadowing built-ins (DON'T DO THIS!) --")	// ─────────────────────────────────────────────	// 7. GOTCHA: Shadowing built-in identifiers	// ─────────────────────────────────────────────	fmt.Println("After block:", val) // still "outer"	}		fmt.Println("Inside block:", val)		val := "inner" // NEW variable, shadows outer	{	fmt.Println("Before block:", val)	val := "outer"	fmt.Println("\n-- Shadowing --")	// ─────────────────────────────────────────────	// 6. GOTCHA: Variable shadowing	// ─────────────────────────────────────────────	// fmt.Println(v)  // ERROR: v not declared	}		fmt.Println("Non-positive:", v)	default:		fmt.Println("Positive:", v)	case v > 0:	switch v := 42; {	fmt.Println("\n-- switch initializer scope --")	// ─────────────────────────────────────────────	// 5. switch initializer scope	// ─────────────────────────────────────────────	// fmt.Println(n)  // ERROR: n not declared	}		fmt.Println("Non-positive:", n) // n is visible here too!	} else {		fmt.Println("Positive:", n)	if n := computeValue(); n > 0 {	// Variables declared in if initializer are scoped to the entire if/else chain:	fmt.Println("\n-- if initializer scope --")	// ─────────────────────────────────────────────	// 4. if initializer scope	// ─────────────────────────────────────────────	// fmt.Println(i)  // ERROR: i not declared	fmt.Println()	}		fmt.Printf("i=%d ", i) // i is scoped to the for block	for i := 0; i < 3; i++ {	fmt.Println("\n-- for loop scope --")	// ─────────────────────────────────────────────	// 3. for loop scope	// ─────────────────────────────────────────────	// fmt.Println(z)  // ERROR: z not declared	}		fmt.Println("Bare block, z:", z)		z := 30	{	// Bare block:	// fmt.Println(y)  // ERROR: y not declared in this scope	}		fmt.Println("Inside if, x:", x) // outer x is visible		fmt.Println("Inside if, y:", y)		y := 20 // only visible inside this if block	if true {	fmt.Println("\n-- Block scope --")	// ─────────────────────────────────────────────	// 2. Block scope (if, for, bare blocks)	// ─────────────────────────────────────────────	fmt.Println("Function scope x:", x)	x := 10	// ─────────────────────────────────────────────	// 1. Function scope	// ─────────────────────────────────────────────	fmt.Println()	fmt.Println("=== Scope and Blocks ===")func main() {