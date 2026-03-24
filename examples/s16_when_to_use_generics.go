//go:build ignore

// Section 16, Topic 121: When to Use Generics
//
// Generics are powerful but not always the right tool.
// Go community guidelines on when to use them.
//
// Run: go run examples/s16_when_to_use_generics.go

package main

import "fmt"

func main() {
	fmt.Println("=== When to Use Generics ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// USE generics when:
	// ─────────────────────────────────────────────
	fmt.Println("✓ USE generics for:")
	fmt.Println("  1. Container types (Stack, Queue, LinkedList)")
	fmt.Println("  2. Functions operating on slices/maps of any element type")
	fmt.Println("  3. When the behavior is the same regardless of type")
	fmt.Println("  4. When you'd otherwise use interface{} and type assertions")

	// Example: generic Set
	s := NewSet[string]()
	s.Add("a")
	s.Add("b")
	s.Add("a") // duplicate
	fmt.Printf("\nSet: %v (len: %d)\n", s.Items(), s.Len())
	fmt.Println("Contains 'a':", s.Contains("a"))
	fmt.Println("Contains 'c':", s.Contains("c"))

	// ─────────────────────────────────────────────
	// DON'T use generics when:
	// ─────────────────────────────────────────────
	fmt.Println("\n✗ DON'T use generics for:")
	fmt.Println("  1. Simply calling a method on type (use interface)")
	fmt.Println("  2. When the implementation differs per type (use interface)")
	fmt.Println("  3. When only one type is ever used")
	fmt.Println("  4. When it makes the code harder to read")

	// ─────────────────────────────────────────────
	// Example: Keys and Values helpers
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Map helpers --")
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println("Keys:", Keys(m))
	fmt.Println("Values:", Values(m))

	// ─────────────────────────────────────────────
	// Comparison: Go vs Rust generics
	// ─────────────────────────────────────────────
	// Go:  func F[T constraint](x T) T { ... }
	// Rust: fn f<T: Trait>(x: T) -> T { ... }
	// Both use constraints/bounds. Rust also has where clauses.
	// Go generics are simpler but less powerful (no specialization, no const generics).
}

// ─────────────────────────────────────────────
// Generic Set implementation
// ─────────────────────────────────────────────
type Set[T comparable] struct {
	m map[T]struct{}
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{m: make(map[T]struct{})}
}

func (s *Set[T]) Add(val T) {
	s.m[val] = struct{}{}
}

func (s *Set[T]) Remove(val T) {
	delete(s.m, val)
}

func (s *Set[T]) Contains(val T) bool {
	_, ok := s.m[val]
	return ok
}

func (s *Set[T]) Len() int {
	return len(s.m)
}

func (s *Set[T]) Items() []T {
	items := make([]T, 0, len(s.m))
	for k := range s.m {
		items = append(items, k)
	}
	return items
}

// ─────────────────────────────────────────────
// Generic map helpers
// ─────────────────────────────────────────────
func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func Values[K comparable, V any](m map[K]V) []V {
	vals := make([]V, 0, len(m))
	for _, v := range m {
		vals = append(vals, v)
	}
	return vals
}
