//go:build ignore

// Section 16, Topic 119: Generic Types (Structs)
//
// Types can also be generic — structs, interfaces, etc.
//
//   type Stack[T any] struct { ... }
//   type Pair[K comparable, V any] struct { ... }
//
// Methods on generic types must repeat the type parameter:
//   func (s *Stack[T]) Push(val T) { ... }
//
// GOTCHA: Methods CANNOT introduce additional type parameters.
//         Only the type-level parameters are available.
//
// Run: go run examples/s16_generic_types.go

package main

import "fmt"

// ─────────────────────────────────────────────
// 1. Generic Stack
// ─────────────────────────────────────────────
type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(val T) {
	s.items = append(s.items, val)
}

func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	if len(s.items) == 0 {
		return zero, false
	}
	val := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return val, true
}

func (s *Stack[T]) Peek() (T, bool) {
	var zero T
	if len(s.items) == 0 {
		return zero, false
	}
	return s.items[len(s.items)-1], true
}

func (s *Stack[T]) Len() int {
	return len(s.items)
}

// ─────────────────────────────────────────────
// 2. Generic Pair
// ─────────────────────────────────────────────
type Pair[A, B any] struct {
	First  A
	Second B
}

func NewPair[A, B any](a A, b B) Pair[A, B] {
	return Pair[A, B]{First: a, Second: b}
}

// ─────────────────────────────────────────────
// 3. Generic Result (like Rust's Result<T, E>)
// ─────────────────────────────────────────────
type Result[T any] struct {
	Value T
	Err   error
}

func Ok[T any](val T) Result[T] {
	return Result[T]{Value: val}
}

func Fail[T any](err error) Result[T] {
	return Result[T]{Err: err}
}

func (r Result[T]) Unwrap() T {
	if r.Err != nil {
		panic(r.Err)
	}
	return r.Value
}

func main() {
	fmt.Println("=== Generic Types ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// Stack of int
	// ─────────────────────────────────────────────
	intStack := &Stack[int]{}
	intStack.Push(1)
	intStack.Push(2)
	intStack.Push(3)
	fmt.Printf("Stack len: %d\n", intStack.Len())

	if val, ok := intStack.Pop(); ok {
		fmt.Println("Popped:", val) // 3
	}

	// ─────────────────────────────────────────────
	// Stack of string
	// ─────────────────────────────────────────────
	fmt.Println("\n-- String stack --")
	strStack := &Stack[string]{}
	strStack.Push("hello")
	strStack.Push("world")
	if val, ok := strStack.Peek(); ok {
		fmt.Println("Peek:", val) // "world"
	}

	// ─────────────────────────────────────────────
	// Pair
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Pair --")
	p := NewPair("name", 42)
	fmt.Printf("Pair: (%v, %v)\n", p.First, p.Second)

	p2 := NewPair(3.14, true)
	fmt.Printf("Pair: (%v, %v)\n", p2.First, p2.Second)

	// ─────────────────────────────────────────────
	// Result
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Result --")
	r1 := Ok(42)
	fmt.Println("Ok:", r1.Unwrap())

	r2 := Fail[int](fmt.Errorf("something failed"))
	fmt.Printf("Fail: err=%v\n", r2.Err)
}
