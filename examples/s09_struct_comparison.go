//go:build ignore

// Section 9, Topic 72: Struct Comparison and Equality
//
// Structs are comparable if ALL their fields are comparable.
// Comparison is field-by-field.
//
// GOTCHA: Structs with slice, map, or function fields are NOT comparable with ==.
//         Use reflect.DeepEqual or manual comparison for those.
// GOTCHA: Field ORDER in the definition doesn't affect comparison (same types needed).
//
// Run: go run examples/s09_struct_comparison.go

package main

import (
	"fmt"
	"reflect"
)

type Point struct {
	X, Y int
}

type Person struct {
	Name string
	Age  int
}

// NOT comparable with == (has slice field):
type Team struct {
	Name    string
	Members []string
}

func main() {
	fmt.Println("=== Struct Comparison ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. Comparable structs
	// ─────────────────────────────────────────────
	p1 := Point{1, 2}
	p2 := Point{1, 2}
	p3 := Point{3, 4}
	fmt.Printf("Point{1,2} == Point{1,2}: %t\n", p1 == p2) // true
	fmt.Printf("Point{1,2} == Point{3,4}: %t\n", p1 == p3) // false
	fmt.Printf("Point{1,2} != Point{3,4}: %t\n", p1 != p3) // true

	// ─────────────────────────────────────────────
	// 2. Different types are not comparable
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Different types --")
	a := Person{Name: "Alice", Age: 30}
	b := Person{Name: "Alice", Age: 30}
	fmt.Printf("Person == Person: %t\n", a == b)
	// Point{1,2} == Person{"A",1}  // compile error: different types

	// ─────────────────────────────────────────────
	// 3. Non-comparable structs
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Non-comparable (has slice) --")
	t1 := Team{Name: "A", Members: []string{"Alice"}}
	t2 := Team{Name: "A", Members: []string{"Alice"}}
	_ = t1
	_ = t2
	// t1 == t2  // compile error: struct containing []string cannot be compared
	fmt.Println("Team contains []string → cannot use ==")

	// Use reflect.DeepEqual instead:
	fmt.Printf("DeepEqual: %t\n", reflect.DeepEqual(t1, t2))

	// ─────────────────────────────────────────────
	// 4. Structs as map keys
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Map keys --")
	// Only comparable structs can be map keys:
	grid := map[Point]string{
		{0, 0}: "origin",
		{1, 0}: "right",
	}
	fmt.Println("Grid:", grid)
	// map[Team]string{}  // compile error: Team is not comparable

	// ─────────────────────────────────────────────
	// 5. Zero struct is valid comparison target
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Zero value comparison --")
	var zero Point
	pt := Point{0, 0}
	fmt.Printf("zero == Point{0,0}: %t\n", zero == pt)

	var zeroPerson Person
	fmt.Printf("Is zero person: %t\n", zeroPerson == Person{})
}
