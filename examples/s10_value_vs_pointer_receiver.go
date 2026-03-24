//go:build ignore

// Section 10, Topic 78: Value vs Pointer Receiver — When to Use Which
//
// Value receiver (func (t T) method()):
//   - Cannot modify the receiver
//   - Works on both values and pointers
//   - Safe for concurrent use (each call gets a copy)
//
// Pointer receiver (func (t *T) method()):
//   - CAN modify the receiver
//   - Avoids copying large structs
//   - Required for mutation
//
// Rule of thumb: If ANY method needs pointer receiver, use pointer for ALL.
//
// Run: go run examples/s10_value_vs_pointer_receiver.go

package main

import (
	"fmt"
	"math"
)

// ─────────────────────────────────────────────
// Small, immutable → value receivers
// ─────────────────────────────────────────────
type Point struct {
	X, Y float64
}

func (p Point) Distance(other Point) float64 {
	dx := p.X - other.X
	dy := p.Y - other.Y
	return math.Sqrt(dx*dx + dy*dy)
}

func (p Point) String() string {
	return fmt.Sprintf("(%.1f, %.1f)", p.X, p.Y)
}

// ─────────────────────────────────────────────
// Needs mutation → pointer receivers
// ─────────────────────────────────────────────
type Stack struct {
	items []int
}

func (s *Stack) Push(val int) {
	s.items = append(s.items, val)
}

func (s *Stack) Pop() (int, bool) {
	if len(s.items) == 0 {
		return 0, false
	}
	top := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return top, true
}

func (s *Stack) Len() int {
	return len(s.items)
}

func main() {
	fmt.Println("=== Value vs Pointer Receiver ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. Value receiver (immutable)
	// ─────────────────────────────────────────────
	a := Point{3, 4}
	b := Point{0, 0}
	fmt.Printf("Distance %s to %s: %.2f\n", a, b, a.Distance(b))

	// ─────────────────────────────────────────────
	// 2. Pointer receiver (mutation)
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Stack (pointer receiver) --")
	s := &Stack{}
	s.Push(10)
	s.Push(20)
	s.Push(30)
	fmt.Printf("Len: %d\n", s.Len())

	if val, ok := s.Pop(); ok {
		fmt.Printf("Popped: %d\n", val)
	}
	fmt.Printf("Len after pop: %d\n", s.Len())

	// ─────────────────────────────────────────────
	// 3. Decision guidelines
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Guidelines --")
	fmt.Println("Use POINTER receiver when:")
	fmt.Println("  1. Method needs to modify the receiver")
	fmt.Println("  2. Struct is large (avoid copy overhead)")
	fmt.Println("  3. Consistency (if one method needs pointer, use for all)")
	fmt.Println()
	fmt.Println("Use VALUE receiver when:")
	fmt.Println("  1. Struct is small and immutable (Point, Time, Color)")
	fmt.Println("  2. Method is read-only")
	fmt.Println("  3. You want the type to be safe for concurrent use")

	// ─────────────────────────────────────────────
	// 4. GOTCHA: Interface satisfaction
	// ─────────────────────────────────────────────
	// If an interface requires method M, and M has a pointer receiver:
	//   - *T satisfies the interface
	//   - T does NOT satisfy the interface
	// If M has a value receiver:
	//   - Both T and *T satisfy the interface

	// ─────────────────────────────────────────────
	// 5. GOTCHA: Map values are not addressable
	// ─────────────────────────────────────────────
	// m := map[string]Stack{"a": {}}
	// m["a"].Push(1)  // ERROR: cannot call pointer method on map value
	// Fix: use map[string]*Stack
}
