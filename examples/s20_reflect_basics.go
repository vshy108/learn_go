//go:build ignore

// Section 20, Topic 145: Reflection (reflect package)
//
// Reflection lets you inspect and manipulate types/values at runtime.
//
// Core types:
//   reflect.Type  — describes a Go type
//   reflect.Value — holds a Go value
//
// GOTCHA: Reflection is SLOW — 10-100x slower than direct access.
// GOTCHA: Reflection can panic if you call wrong methods on wrong kinds.
// GOTCHA: Can only set exported fields.
// GOTCHA: reflect.DeepEqual handles nested structs, slices, maps.
//
// Rule of thumb: avoid reflection unless writing frameworks/serialization.
//
// Run: go run examples/s20_reflect_basics.go

package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"email"`
	age   int    // unexported
}

func (u User) Greet() string {
	return "Hello, " + u.Name
}

func main() {
	fmt.Println("=== Reflection ===")
	fmt.Println()

	u := User{Name: "Alice", Email: "alice@go.dev", age: 30}

	// ─────────────────────────────────────────────
	// 1. reflect.TypeOf — inspect type
	// ─────────────────────────────────────────────
	fmt.Println("-- Type inspection --")
	t := reflect.TypeOf(u)
	fmt.Println("Type:", t)            // main.User
	fmt.Println("Kind:", t.Kind())     // struct
	fmt.Println("Name:", t.Name())     // User
	fmt.Println("NumField:", t.NumField())

	// ─────────────────────────────────────────────
	// 2. Iterate fields and struct tags
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Fields and tags --")
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Printf("  %s: type=%s, exported=%t, json=%q, validate=%q\n",
			f.Name, f.Type, f.IsExported(),
			f.Tag.Get("json"), f.Tag.Get("validate"))
	}

	// ─────────────────────────────────────────────
	// 3. reflect.ValueOf — inspect values
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Value inspection --")
	v := reflect.ValueOf(u)
	fmt.Println("Value:", v)
	fmt.Println("Name field:", v.FieldByName("Name"))
	fmt.Println("Email field:", v.FieldByName("Email"))
	// v.FieldByName("age") works but v.FieldByName("age").Interface() would panic

	// ─────────────────────────────────────────────
	// 4. Setting values (requires pointer)
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Setting values --")
	vp := reflect.ValueOf(&u).Elem() // Elem() dereferences pointer
	nameField := vp.FieldByName("Name")

	if nameField.CanSet() {
		nameField.SetString("Bob")
		fmt.Println("After set:", u.Name) // Bob
	}

	// GOTCHA: Can't set unexported fields:
	ageField := vp.FieldByName("age")
	fmt.Println("Can set 'age':", ageField.CanSet()) // false

	// ─────────────────────────────────────────────
	// 5. Calling methods via reflection
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Calling methods --")
	method := reflect.ValueOf(u).MethodByName("Greet")
	results := method.Call(nil)
	fmt.Println("Greet():", results[0])

	// ─────────────────────────────────────────────
	// 6. reflect.DeepEqual
	// ─────────────────────────────────────────────
	fmt.Println("\n-- DeepEqual --")
	a := []int{1, 2, 3}
	b := []int{1, 2, 3}
	c := []int{1, 2, 4}
	fmt.Println("a == b:", reflect.DeepEqual(a, b)) // true
	fmt.Println("a == c:", reflect.DeepEqual(a, c)) // false

	m1 := map[string]int{"a": 1}
	m2 := map[string]int{"a": 1}
	fmt.Println("m1 == m2:", reflect.DeepEqual(m1, m2)) // true

	// ─────────────────────────────────────────────
	// 7. Dynamic type checking
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Dynamic type --")
	inspectType(42)
	inspectType("hello")
	inspectType([]int{1, 2})
	inspectType(map[string]int{"a": 1})
}

func inspectType(x interface{}) {
	v := reflect.ValueOf(x)
	fmt.Printf("  Value: %-12v  Kind: %-8s  Type: %s\n",
		v, v.Kind(), v.Type())
}
