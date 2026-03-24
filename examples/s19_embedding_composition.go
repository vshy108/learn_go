//go:build ignore

// Section 19, Topic 139: Embedding (Composition over Inheritance)
//
// Go has no inheritance. Instead, it uses EMBEDDING for code reuse.
// Embedding promotes fields and methods of the embedded type.
//
// GOTCHA: Embedding is NOT inheritance — it's delegation/forwarding.
// GOTCHA: The embedded type's methods receive the embedded type as receiver, not the outer type.
// GOTCHA: If outer and embedded have same method, outer wins (shadowing).
//
// Run: go run examples/s19_embedding_composition.go

package main

import "fmt"

// ─────────────────────────────────────────────
// Base type
// ─────────────────────────────────────────────
type Animal struct {
	Name string
	Age  int
}

func (a Animal) Describe() string {
	return fmt.Sprintf("%s (age %d)", a.Name, a.Age)
}

func (a Animal) Breathe() {
	fmt.Printf("  %s is breathing\n", a.Name)
}

// ─────────────────────────────────────────────
// Embedding Animal into Dog
// ─────────────────────────────────────────────
type Dog struct {
	Animal // embedded — Dog "has an" Animal
	Breed  string
}

func (d Dog) Bark() {
	fmt.Printf("  %s says Woof!\n", d.Name)
}

// Override (shadow) a method:
func (d Dog) Describe() string {
	return fmt.Sprintf("%s the %s (age %d)", d.Name, d.Breed, d.Age)
}

// ─────────────────────────────────────────────
// Multiple embedding
// ─────────────────────────────────────────────
type Logger struct{}

func (l Logger) Log(msg string) {
	fmt.Printf("  [LOG] %s\n", msg)
}

type Server struct {
	Logger // embed logger
	Host   string
	Port   int
}

func main() {
	fmt.Println("=== Embedding / Composition ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. Promoted fields and methods
	// ─────────────────────────────────────────────
	d := Dog{
		Animal: Animal{Name: "Rex", Age: 5},
		Breed:  "Labrador",
	}

	// Promoted fields (from Animal):
	fmt.Println("Name:", d.Name) // same as d.Animal.Name
	fmt.Println("Age:", d.Age)   // same as d.Animal.Age
	fmt.Println("Breed:", d.Breed)

	// Promoted methods:
	d.Breathe() // from Animal
	d.Bark()    // from Dog

	// Shadowed method (Dog's Describe wins):
	fmt.Println("Describe:", d.Describe())

	// Original still accessible:
	fmt.Println("Animal.Describe:", d.Animal.Describe())

	// ─────────────────────────────────────────────
	// 2. Interface satisfaction through embedding
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Interface satisfaction --")
	type Describer interface {
		Describe() string
	}

	var desc Describer = d // Dog satisfies Describer (via its own Describe)
	fmt.Println("Via interface:", desc.Describe())

	// ─────────────────────────────────────────────
	// 3. Multiple embedding
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Multiple embedding --")
	s := Server{Host: "localhost", Port: 8080}
	s.Log("Server starting") // from embedded Logger

	// ─────────────────────────────────────────────
	// 4. Embedding interfaces
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Embedding interfaces --")
	type ReadWriter interface {
		Read([]byte) (int, error)
		Write([]byte) (int, error)
	}
	// This is interface composition, covered in section 11
}
