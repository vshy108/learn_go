//go:build ignore

// Section 11, Topic 81: Implicit Interface Satisfaction
//
// Go interfaces are satisfied IMPLICITLY — no "implements" keyword.
// If a type has all the methods an interface requires, it satisfies it.
//
// This decouples packages: the type doesn't need to know about the interface.
// Interfaces can be defined by the consumer, not the provider.
//
// GOTCHA: Pointer receiver method: *T satisfies, T does NOT.
// GOTCHA: Value receiver method: both T and *T satisfy.
//
// Run: go run examples/s11_implicit_implementation.go

package main

import "fmt"

type Speaker interface {
	Speak() string
}

type Dog struct{ Name string }
type Cat struct{ Name string }

// Both satisfy Speaker without declaring it:
func (d Dog) Speak() string { return d.Name + " says Woof!" }
func (c Cat) Speak() string { return c.Name + " says Meow!" }

// ─────────────────────────────────────────────
// Pointer vs value receiver matters!
// ─────────────────────────────────────────────
type Mutable interface {
	Mutate()
}

type Foo struct{ Val int }

func (f *Foo) Mutate() { f.Val++ } // pointer receiver

func main() {
	fmt.Println("=== Implicit Implementation ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. Both satisfy Speaker
	// ─────────────────────────────────────────────
	animals := []Speaker{
		Dog{Name: "Rex"},
		Cat{Name: "Whiskers"},
	}
	for _, a := range animals {
		fmt.Println(a.Speak())
	}

	// ─────────────────────────────────────────────
	// 2. Consumer-defined interfaces
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Consumer-defined interface --")
	// I can define an interface that existing types already satisfy:
	type Namer interface {
		Speak() string // Dog already has this!
	}
	var n Namer = Dog{Name: "Buddy"}
	fmt.Println(n.Speak())

	// ─────────────────────────────────────────────
	// 3. Pointer receiver gotcha
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Pointer receiver --")
	// Foo has Mutate() with pointer receiver

	var m Mutable
	// m = Foo{Val: 1}    // ERROR: Foo does not implement Mutable
	m = &Foo{Val: 1} // OK: *Foo implements Mutable
	m.Mutate()
	fmt.Printf("After Mutate: %+v\n", m)

	// ─────────────────────────────────────────────
	// 4. Compile-time interface check
	// ─────────────────────────────────────────────
	// Verify a type satisfies an interface at compile time:
	var _ Speaker = Dog{}  // compile-time assertion
	var _ Speaker = Cat{}  // compile-time assertion
	var _ Mutable = &Foo{} // must be pointer
}
