//go:build ignore

// Section 9, Topic 67: Methods with Value Receiver
//
// Methods are functions with a receiver argument.
// Value receivers operate on a COPY of the struct.
//
// func (r ReceiverType) MethodName() ReturnType { ... }
//
// GOTCHA: Value receivers cannot modify the original struct.
// GOTCHA: Value receiver methods can be called on both values AND pointers.
//
// Run: go run examples/s09_methods_value_receiver.go

package main

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float64
}

// Value receiver — operates on a copy
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Cannot modify the original:
func (c Circle) TryDouble() {
	c.Radius *= 2 // modifies the copy only!
}

func (c Circle) String() string {
	return fmt.Sprintf("Circle(r=%.2f)", c.Radius)
}

func main() {
	fmt.Println("=== Methods — Value Receiver ===")
	fmt.Println()

	c := Circle{Radius: 5.0}
	fmt.Printf("Circle: %v\n", c)
	fmt.Printf("Area: %.2f\n", c.Area())
	fmt.Printf("Perimeter: %.2f\n", c.Perimeter())

	// ─────────────────────────────────────────────
	// Value receiver can't modify original
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Value receiver can't modify --")
	c.TryDouble()
	fmt.Printf("After TryDouble: %v (unchanged!)\n", c)

	// ─────────────────────────────────────────────
	// Can call value receiver methods on pointers too
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Called on pointer --")
	cp := &Circle{Radius: 3.0}
	fmt.Printf("Pointer.Area: %.2f\n", cp.Area()) // Go auto-dereferences

	// ─────────────────────────────────────────────
	// Methods on non-struct types
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Methods on custom types --")
	var m MyInt = 42
	fmt.Printf("MyInt %d doubled: %d\n", m, m.Double())
}

// You can define methods on any named type (defined in the same package)
type MyInt int

func (m MyInt) Double() MyInt {
	return m * 2
}
