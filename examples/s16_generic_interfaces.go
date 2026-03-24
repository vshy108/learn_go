//go:build ignore

// Section 16, Topic 120: Generic Interfaces
//
// Interfaces can be generic too, though this is less common.
// More importantly, existing interfaces serve as constraints for generics.
//
// GOTCHA: Interface type elements (type unions) can only be used as constraints,
//         not as regular interface types.
//
// Run: go run examples/s16_generic_interfaces.go

package main

import "fmt"

// ─────────────────────────────────────────────
// 1. Generic interface
// ─────────────────────────────────────────────
type Container[T any] interface {
	Add(T)
	Get(index int) T
	Len() int
}

// Implementation:
type SliceContainer[T any] struct {
	items []T
}

func (sc *SliceContainer[T]) Add(val T) {
	sc.items = append(sc.items, val)
}

func (sc *SliceContainer[T]) Get(index int) T {
	return sc.items[index]
}

func (sc *SliceContainer[T]) Len() int {
	return len(sc.items)
}

// ─────────────────────────────────────────────
// 2. Constraint interfaces (type sets)
// ─────────────────────────────────────────────
type Numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

type Addable interface {
	~int | ~float64 | ~string
}

func Add[T Addable](a, b T) T {
	return a + b
}

// ─────────────────────────────────────────────
// 3. Interface with both methods and type elements
// ─────────────────────────────────────────────
type StringLike interface {
	~string
	Len() int
}

func main() {
	fmt.Println("=== Generic Interfaces ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// Using Container interface
	// ─────────────────────────────────────────────
	var c Container[string] = &SliceContainer[string]{}
	c.Add("hello")
	c.Add("world")
	fmt.Printf("Container: len=%d, [0]=%s, [1]=%s\n", c.Len(), c.Get(0), c.Get(1))

	var intC Container[int] = &SliceContainer[int]{}
	intC.Add(10)
	intC.Add(20)
	fmt.Printf("IntContainer: len=%d, [0]=%d\n", intC.Len(), intC.Get(0))

	// ─────────────────────────────────────────────
	// Using Addable constraint
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Addable --")
	fmt.Println("Add ints:", Add(3, 5))
	fmt.Println("Add floats:", Add(1.5, 2.5))
	fmt.Println("Add strings:", Add("hello, ", "world"))

	// ─────────────────────────────────────────────
	// GOTCHA: Type set interfaces can't be used as types
	// ─────────────────────────────────────────────
	// var x Numeric = 42  // ERROR: cannot use Numeric as type
	// Numeric can only be used as a constraint: func F[T Numeric](x T) T
}
