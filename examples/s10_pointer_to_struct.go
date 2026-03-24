//go:build ignore

// Section 10, Topic 79: Pointer to Struct — Automatic Dereferencing
//
// Go automatically dereferences struct pointers with dot notation:
//   p.Field  is equivalent to  (*p).Field
//
// This makes working with struct pointers feel natural.
//
// Run: go run examples/s10_pointer_to_struct.go

package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	fmt.Println("=== Pointer to Struct ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. Auto-dereference with dot notation
	// ─────────────────────────────────────────────
	p := &Person{Name: "Alice", Age: 30}

	// These are equivalent:
	fmt.Println(p.Name)    // auto-deref
	fmt.Println((*p).Name) // explicit deref (verbose, rarely used)

	// Modify through pointer:
	p.Age = 31
	fmt.Printf("After modify: %+v\n", *p)

	// ─────────────────────────────────────────────
	// 2. Common pattern: return pointer from constructor
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Constructor pattern --")
	user := NewPerson("Bob", 25)
	fmt.Printf("user.Name: %s\n", user.Name) // auto-deref
	user.Age++
	fmt.Printf("After birthday: %+v\n", *user)

	// ─────────────────────────────────────────────
	// 3. Struct pointer in slice
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Slice of pointers --")
	people := []*Person{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 25},
		{Name: "Charlie", Age: 35},
	}
	for _, p := range people {
		p.Age++ // modifies the original (p is already a pointer)
	}
	for _, p := range people {
		fmt.Printf("  %s: %d\n", p.Name, p.Age)
	}

	// ─────────────────────────────────────────────
	// 4. GOTCHA: range with value structs
	// ─────────────────────────────────────────────
	fmt.Println("\n-- range gotcha --")
	values := []Person{
		{Name: "X", Age: 1},
		{Name: "Y", Age: 2},
	}
	for _, v := range values {
		v.Age = 99 // modifies the copy, not the original!
		_ = v      // acknowledge the copy is discarded
	}
	fmt.Printf("After range: %+v (unchanged!)\n", values)

	// Fix: use index:
	for i := range values {
		values[i].Age = 99
	}
	fmt.Printf("After index: %+v (modified)\n", values)

	// Or use slice of pointers (as above)
}

func NewPerson(name string, age int) *Person {
	return &Person{Name: name, Age: age}
}
