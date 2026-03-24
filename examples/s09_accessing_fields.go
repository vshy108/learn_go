//go:build ignore

// Section 9, Topic 66: Accessing and Modifying Struct Fields
//
// Fields are accessed with dot notation. Go auto-dereferences pointers to structs.
//
// GOTCHA: Structs are value types — passing to a function copies them.
//         Use a pointer if you want to modify the original.
//
// Run: go run examples/s09_accessing_fields.go

package main

import "fmt"

type Point struct {
	X, Y float64
}

type Rectangle struct {
	TopLeft     Point
	BottomRight Point
}

func main() {
	fmt.Println("=== Accessing Fields ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. Direct access
	// ─────────────────────────────────────────────
	p := Point{X: 3.0, Y: 4.0}
	fmt.Printf("X=%.1f, Y=%.1f\n", p.X, p.Y)

	// Modify:
	p.X = 10.0
	fmt.Printf("After modify: %+v\n", p)

	// ─────────────────────────────────────────────
	// 2. Pointer auto-dereference
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Pointer auto-deref --")
	pp := &Point{X: 1, Y: 2}
	pp.X = 100 // same as (*pp).X = 100
	fmt.Printf("Via pointer: %+v\n", *pp)

	// ─────────────────────────────────────────────
	// 3. Nested struct access
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Nested --")
	rect := Rectangle{
		TopLeft:     Point{0, 10},
		BottomRight: Point{20, 0},
	}
	fmt.Printf("TopLeft.X = %.0f\n", rect.TopLeft.X)
	rect.BottomRight.Y = -5
	fmt.Printf("Rect: %+v\n", rect)

	// ─────────────────────────────────────────────
	// 4. Passing struct to functions
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Passing to functions --")
	pt := Point{5, 5}
	moveByValue(pt)
	fmt.Printf("After moveByValue: %+v (unchanged)\n", pt)
	moveByPointer(&pt)
	fmt.Printf("After moveByPointer: %+v (modified)\n", pt)
}

func moveByValue(p Point) {
	p.X += 10 // modifies copy
}

func moveByPointer(p *Point) {
	p.X += 10 // modifies original (auto-deref)
}
