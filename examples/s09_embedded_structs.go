//go:build ignore

// Section 9, Topic 69: Embedded Structs — Composition Over Inheritance
//
// Go has NO inheritance. Instead, use struct embedding (composition).
// Embedded struct's fields and methods are "promoted" — accessible directly.
//
// GOTCHA: Embedding is NOT inheritance. The outer type doesn't "become" the inner type.
// GOTCHA: If two embeds have the same field name, you must disambiguate.
// GOTCHA: Promoted methods use the inner type's receiver, not the outer type.
//
// Run: go run examples/s09_embedded_structs.go

package main

import "fmt"

// ─────────────────────────────────────────────
// Base types
// ─────────────────────────────────────────────
type Address struct {
	Street string
	City   string
	State  string
}

func (a Address) FullAddress() string {
	return fmt.Sprintf("%s, %s, %s", a.Street, a.City, a.State)
}

// ─────────────────────────────────────────────
// Embedding
// ─────────────────────────────────────────────
type Employee struct {
	Name    string
	Address // embedded (not "Address Address")
	Salary  float64
}

// ─────────────────────────────────────────────
// Ambiguous embedding
// ─────────────────────────────────────────────
type Logger struct {
	Name string // same field name as in Employee area
}

func (l Logger) Log(msg string) {
	fmt.Printf("[%s] %s\n", l.Name, msg)
}

type Manager struct {
	Employee
	Logger
	Level int
}

func main() {
	fmt.Println("=== Embedded Structs ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. Promoted fields
	// ─────────────────────────────────────────────
	emp := Employee{
		Name:    "Alice",
		Address: Address{Street: "123 Main", City: "NYC", State: "NY"},
		Salary:  90000,
	}

	// Promoted field access:
	fmt.Println("City:", emp.City) // same as emp.Address.City
	fmt.Println("Full:", emp.FullAddress())

	// ─────────────────────────────────────────────
	// 2. Disambiguating same-named fields
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Ambiguity --")
	mgr := Manager{
		Employee: Employee{
			Name:    "Bob",
			Address: Address{Street: "456 Oak", City: "LA", State: "CA"},
			Salary:  120000,
		},
		Logger: Logger{Name: "manager-logger"},
		Level:  3,
	}

	// mgr.Name is ambiguous — both Employee and Logger have Name
	// Must qualify:
	fmt.Println("Employee Name:", mgr.Employee.Name)
	fmt.Println("Logger Name:", mgr.Logger.Name)

	// Non-ambiguous promoted method:
	mgr.Log("starting work")       // from Logger
	fmt.Println("City:", mgr.City) // from Employee → Address
}
