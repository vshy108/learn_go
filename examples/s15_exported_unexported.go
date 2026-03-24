//go:build ignore

// Section 15, Topic 113: Exported vs Unexported
//
// Go's visibility rule is simple:
//   Uppercase first letter → exported (public)
//   Lowercase first letter → unexported (private to package)
//
// Applies to: functions, types, variables, constants, struct fields, methods.
//
// GOTCHA: Embedded struct with exported fields: fields are accessible
//         even if the type is unexported.
// GOTCHA: JSON marshaling only works with exported fields.
//
// Run: go run examples/s15_exported_unexported.go

package main

import (
	"encoding/json"
	"fmt"
)

// Exported (visible outside package):
type User struct {
	Name  string // exported field
	Email string // exported field
	age   int    // unexported field
}

// unexported type (only usable within this package):
type config struct {
	debug bool
	port  int
}

// Exported function:
func NewUser(name, email string, age int) User {
	return User{Name: name, Email: email, age: age}
}

// unexported function:
func validateEmail(email string) bool {
	return len(email) > 0
}

func main() {
	fmt.Println("=== Exported vs Unexported ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. Within the same package, everything is accessible
	// ─────────────────────────────────────────────
	u := User{Name: "Alice", Email: "alice@example.com", age: 30}
	fmt.Printf("User: %+v\n", u)
	fmt.Printf("Age (unexported): %d\n", u.age) // ok within same package

	c := config{debug: true, port: 8080}
	fmt.Printf("Config: %+v\n", c)

	fmt.Println("Valid email:", validateEmail("test@test.com"))

	// ─────────────────────────────────────────────
	// 2. From another package, only exported is visible
	// ─────────────────────────────────────────────
	// import "mypackage"
	// mypackage.User{Name: "Bob"}     // ok — User and Name are exported
	// mypackage.User{age: 25}         // ERROR — age is unexported
	// mypackage.config{}              // ERROR — config is unexported
	// mypackage.validateEmail("x")    // ERROR — validateEmail is unexported

	// ─────────────────────────────────────────────
	// 3. JSON and unexported fields
	// ─────────────────────────────────────────────
	fmt.Println("\n-- JSON --")
	u2 := User{Name: "Bob", Email: "bob@test.com", age: 25}

	data, _ := json.Marshal(u2)
	fmt.Println("JSON:", string(data))
	// {"Name":"Bob","Email":"bob@test.com"}
	// Note: age is NOT in JSON (unexported)

	// ─────────────────────────────────────────────
	// 4. Conventions
	// ─────────────────────────────────────────────
	// - Export only what's necessary (minimal API surface)
	// - Unexported types can have exported methods (interface satisfaction)
	// - Constructor functions: NewXxx() returns unexported type through interface

	// ─────────────────────────────────────────────
	// 5. internal/ directory
	// ─────────────────────────────────────────────
	// Packages in internal/ are only importable by code in the parent tree:
	//   myproject/internal/auth  — only myproject/* can import this
	//   outside/pkg              — CANNOT import myproject/internal/auth
}
