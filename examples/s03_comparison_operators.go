//go:build ignore

// Section 3, Topic 25: Comparison and Equality Operators
//
// Go's comparison operators: ==, !=, <, >, <=, >=
//
// Key rules:
//   - Both operands must be the SAME type (no implicit conversion)
//   - Structs are comparable if all fields are comparable
//   - Slices, maps, and functions are NOT comparable with == (except to nil)
//   - Interfaces are compared by dynamic type and value
//
// GOTCHA: Comparing incomparable types is a compile error (slices, maps, functions).
// GOTCHA: Interface comparison can panic if the dynamic type is not comparable.
// GOTCHA: Floating-point NaN is not equal to itself.
//
// Run: go run examples/s03_comparison_operators.go

package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y int
}

type WithSlice struct {
	Name  string
	Items []int // makes this struct non-comparable!
}

func main() {
	fmt.Println("=== Comparison Operators ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. Basic type comparisons
	// ─────────────────────────────────────────────
	fmt.Println("-- Basic comparisons --")
	num1, num1b, num2 := 1, 1, 2
	fmt.Printf("1 == 1:     %t\n", num1 == num1b)
	fmt.Printf("1 != 2:     %t\n", num1 != num2)
	fmt.Printf("1 < 2:      %t\n", num1 < num2)
	num2b := 2
	fmt.Printf("2 >= 2:     %t\n", num2 >= num2b)
	fmt.Printf("\"a\" < \"b\": %t (lexicographic)\n", "a" < "b")
	fmt.Printf("\"abc\" < \"abd\": %t\n", "abc" < "abd")

	// ─────────────────────────────────────────────
	// 2. MUST be same type
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Type matching --")
	// var a int32 = 42
	// var b int64 = 42
	// fmt.Println(a == b)  // ERROR: mismatched types int32 and int64
	// Fix: fmt.Println(int64(a) == b)
	var i32 int32 = 42
	var i64 int64 = 42
	fmt.Printf("int32(42) == int64(42): %t (after conversion)\n", int64(i32) == i64)

	// ─────────────────────────────────────────────
	// 3. Struct comparison
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Struct comparison --")
	p1 := Point{1, 2}
	p2 := Point{1, 2}
	p3 := Point{3, 4}
	fmt.Printf("Point{1,2} == Point{1,2}: %t\n", p1 == p2) // true
	fmt.Printf("Point{1,2} == Point{3,4}: %t\n", p1 == p3) // false

	// Structs with non-comparable fields CANNOT use ==:
	// w1 := WithSlice{"a", []int{1}}
	// w2 := WithSlice{"a", []int{1}}
	// fmt.Println(w1 == w2)  // ERROR: cannot compare (contains []int)

	// ─────────────────────────────────────────────
	// 4. Array comparison (fixed-size only)
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Array comparison --")
	arr1 := [3]int{1, 2, 3}
	arr2 := [3]int{1, 2, 3}
	arr3 := [3]int{1, 2, 4}
	fmt.Printf("[1,2,3] == [1,2,3]: %t\n", arr1 == arr2)
	fmt.Printf("[1,2,3] == [1,2,4]: %t\n", arr1 == arr3)
	// Arrays of different sizes are different types:
	// [3]int and [4]int — can't compare

	// ─────────────────────────────────────────────
	// 5. NOT comparable: slices, maps, functions
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Non-comparable types --")
	// s1 := []int{1, 2, 3}
	// s2 := []int{1, 2, 3}
	// fmt.Println(s1 == s2)  // ERROR: slice can only be compared to nil

	// Can only compare to nil:
	var s1 []int
	fmt.Printf("nil slice == nil: %t\n", s1 == nil)

	var m1 map[string]int
	fmt.Printf("nil map == nil:   %t\n", m1 == nil)

	var fn func()
	fmt.Printf("nil func == nil:  %t\n", fn == nil)

	// ─────────────────────────────────────────────
	// 6. Interface comparison
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Interface comparison --")
	var i1 interface{} = 42
	var i2 interface{} = 42
	var i3 interface{} = "42"
	fmt.Printf("interface{}(42) == interface{}(42):   %t\n", i1 == i2)
	fmt.Printf("interface{}(42) == interface{}(\"42\"): %t\n", i1 == i3)

	// GOTCHA: Comparing interfaces with non-comparable dynamic types PANICS:
	// var i4 interface{} = []int{1}
	// var i5 interface{} = []int{1}
	// fmt.Println(i4 == i5)  // PANIC at runtime!

	// ─────────────────────────────────────────────
	// 7. GOTCHA: Float comparison
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Float comparison gotchas --")
	fmt.Printf("0.1+0.2 == 0.3: %t (floating point surprise!)\n", 0.1+0.2 == 0.3)
	nan := math.NaN()
	fmt.Printf("NaN == NaN: %t (always false per IEEE 754)\n", nan == nan)

	// ─────────────────────────────────────────────
	// 8. Pointer comparison
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Pointer comparison --")
	x, y := 42, 42
	px, py := &x, &y
	fmt.Printf("&x == &y: %t (different addresses)\n", px == py)
	fmt.Printf("*px == *py: %t (same values)\n", *px == *py)
	pz := &x
	fmt.Printf("&x == &x: %t (same address)\n", px == pz)
}
