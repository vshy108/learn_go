//go:build ignore

// Section 9, Topic 65: Defining and Creating Structs
//
// Structs are Go's primary way to group related data together.
// They are VALUE TYPES — assignment and passing copies all fields.
//
// Go has NO classes, NO inheritance. Structs + interfaces = Go's OOP.
//
// GOTCHA: Zero value of a struct has all fields set to their zero values.
// GOTCHA: Structs are comparable if all fields are comparable.
// GOTCHA: Field order matters for memory layout (padding).
//
// Run: go run examples/s09_defining_structs.go

package main

import "fmt"

// ─────────────────────────────────────────────
// 1. Basic struct definition
// ─────────────────────────────────────────────
type Person struct {
	Name string
	Age  int
}

// ─────────────────────────────────────────────
// 2. Struct with mixed types
// ─────────────────────────────────────────────
type Address struct {
	Street string
	City   string
	Zip    string
}

type Employee struct {
	Person  // embedded struct
	Address Address
	Salary  float64
	Active  bool
}

func main() {
	fmt.Println("=== Defining and Creating Structs ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 3. Creating structs (multiple ways)
	// ─────────────────────────────────────────────

	// Named fields (preferred — order doesn't matter):
	p1 := Person{Name: "Alice", Age: 30}
	fmt.Println("Named:", p1)

	// Positional (ALL fields required, order matters):
	p2 := Person{"Bob", 25}
	fmt.Println("Positional:", p2)

	// Zero value:
	var p3 Person
	fmt.Printf("Zero: %+v (Name=%q, Age=%d)\n", p3, p3.Name, p3.Age)

	// With new (returns pointer):
	p4 := new(Person)
	p4.Name = "Charlie"
	p4.Age = 35
	fmt.Printf("new: %+v\n", *p4)

	// Pointer literal:
	p5 := &Person{Name: "Diana", Age: 28}
	fmt.Printf("Pointer: %+v (type: %T)\n", *p5, p5)

	// ─────────────────────────────────────────────
	// 4. Accessing and modifying fields
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Field access --")
	p1.Age = 31
	fmt.Println("After birthday:", p1)

	// Pointer auto-dereference (no need for (*p5).Name):
	p5.Name = "Diana Updated"
	fmt.Println("Pointer field:", p5.Name)

	// ─────────────────────────────────────────────
	// 5. Struct comparison
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Comparison --")
	a := Person{Name: "Alice", Age: 30}
	b := Person{Name: "Alice", Age: 30}
	c := Person{Name: "Bob", Age: 25}
	fmt.Println("a == b:", a == b) // true
	fmt.Println("a == c:", a == c) // false

	// ─────────────────────────────────────────────
	// 6. Nested struct
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Nested --")
	emp := Employee{
		Person:  Person{Name: "Eve", Age: 40},
		Address: Address{Street: "123 Main", City: "NYC", Zip: "10001"},
		Salary:  85000,
		Active:  true,
	}
	fmt.Printf("Employee: %+v\n", emp)
	fmt.Println("Name (promoted):", emp.Name) // from embedded Person
}
