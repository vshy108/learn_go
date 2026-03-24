//go:build ignore

// Section 11, Topic 80: Defining Interfaces
//
// An interface defines a set of method signatures.
// Any type that implements all methods satisfies the interface — IMPLICITLY.
// No "implements" keyword needed.
//
// GOTCHA: Interfaces are satisfied implicitly — nothing links a type to an interface.
// GOTCHA: Keep interfaces small. Go idiom: 1-2 methods per interface.
// GOTCHA: Accept interfaces, return concrete types.
//
// Run: go run examples/s11_defining_interfaces.go

package main

import (
	"fmt"
	"math"
)

// ─────────────────────────────────────────────
// Interface definition
// ─────────────────────────────────────────────
type Shape interface {
	Area() float64
	Perimeter() float64
}

// ─────────────────────────────────────────────
// Types that satisfy Shape (no explicit declaration)
// ─────────────────────────────────────────────
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64      { return math.Pi * c.Radius * c.Radius }
func (c Circle) Perimeter() float64 { return 2 * math.Pi * c.Radius }

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64      { return r.Width * r.Height }
func (r Rectangle) Perimeter() float64 { return 2 * (r.Width + r.Height) }

// ─────────────────────────────────────────────
// Function accepting the interface
// ─────────────────────────────────────────────
func printShape(s Shape) {
	fmt.Printf("  Type: %T\n", s)
	fmt.Printf("  Area: %.2f\n", s.Area())
	fmt.Printf("  Perimeter: %.2f\n", s.Perimeter())
}

func main() {
	fmt.Println("=== Defining Interfaces ===")
	fmt.Println()

	c := Circle{Radius: 5}
	r := Rectangle{Width: 4, Height: 6}

	fmt.Println("Circle:")
	printShape(c) // Circle satisfies Shape

	fmt.Println("\nRectangle:")
	printShape(r) // Rectangle satisfies Shape

	// ─────────────────────────────────────────────
	// Slice of interfaces (polymorphism)
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Polymorphism --")
	shapes := []Shape{
		Circle{Radius: 3},
		Rectangle{Width: 10, Height: 5},
		Circle{Radius: 1},
	}
	totalArea := 0.0
	for _, s := range shapes {
		totalArea += s.Area()
	}
	fmt.Printf("Total area: %.2f\n", totalArea)

	// ─────────────────────────────────────────────
	// Idiom: Small interfaces
	// ─────────────────────────────────────────────
	// io.Reader:  Read(p []byte) (n int, err error)
	// io.Writer:  Write(p []byte) (n int, err error)
	// fmt.Stringer: String() string
	// error:      Error() string
	// These are Go's most powerful interfaces — each has just 1 method.
}
