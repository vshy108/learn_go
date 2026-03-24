//go:build ignore

// Section 10, Topic 76: nil Pointers
//
// The zero value of any pointer type is nil.
// Dereferencing a nil pointer causes a runtime panic.
//
// GOTCHA: Always check for nil before dereferencing pointers.
// GOTCHA: Methods on nil pointers CAN work if the method doesn't dereference.
//
// Run: go run examples/s10_nil_pointers.go

package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

// Method that works on nil receiver:
func (n *Node) IsEmpty() bool {
	return n == nil
}

// Method that panics on nil receiver:
func (n *Node) GetValue() int {
	return n.Value // panics if n is nil
}

func main() {
	fmt.Println("=== nil Pointers ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. Zero value is nil
	// ─────────────────────────────────────────────
	var p *int
	fmt.Printf("p = %v, is nil: %t\n", p, p == nil)

	var sp *string
	fmt.Printf("sp = %v, is nil: %t\n", sp, sp == nil)

	// ─────────────────────────────────────────────
	// 2. Dereferencing nil panics
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Dereference nil = panic --")
	// fmt.Println(*p)  // PANIC: runtime error: invalid memory address
	fmt.Println("(skipping dereference to avoid panic)")

	// ─────────────────────────────────────────────
	// 3. Safe nil check pattern
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Safe nil check --")
	if p != nil {
		fmt.Println("Value:", *p)
	} else {
		fmt.Println("Pointer is nil, skipping dereference")
	}

	// ─────────────────────────────────────────────
	// 4. Methods on nil receiver
	// ─────────────────────────────────────────────
	fmt.Println("\n-- nil receiver --")
	var node *Node
	fmt.Printf("IsEmpty: %t\n", node.IsEmpty()) // works! returns true
	// node.GetValue()  // PANIC: nil pointer dereference

	// ─────────────────────────────────────────────
	// 5. Common nil pointer scenarios
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Common scenarios --")

	// Uninitialized struct field:
	type App struct {
		Config *Config
	}
	var app App
	fmt.Printf("app.Config is nil: %t\n", app.Config == nil)
	// app.Config.Name  // PANIC

	// Function returning nil:
	result := findUser("nonexistent")
	if result != nil {
		fmt.Printf("Found: %+v\n", *result)
	} else {
		fmt.Println("User not found (nil returned)")
	}

	// ─────────────────────────────────────────────
	// 6. nil in linked list
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Linked list --")
	head := &Node{Value: 1, Next: &Node{Value: 2, Next: &Node{Value: 3}}}
	for n := head; n != nil; n = n.Next {
		fmt.Printf("%d → ", n.Value)
	}
	fmt.Println("nil")
}

type Config struct {
	Name string
}

func findUser(name string) *Config {
	if name == "admin" {
		return &Config{Name: "admin"}
	}
	return nil
}
