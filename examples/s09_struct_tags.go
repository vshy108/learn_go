//go:build ignore

// Section 9, Topic 70: Struct Tags
//
// Struct tags are metadata strings attached to struct fields.
// They're used by encoding/json, database ORMs, validation libraries, etc.
//
// Syntax: `key:"value" key2:"value2"`
//
// GOTCHA: Tags are raw strings — typos produce no compile error!
// GOTCHA: Tags are invisible at runtime without reflection.
// GOTCHA: Only exported fields are visible to encoding/json.
//
// Run: go run examples/s09_struct_tags.go

package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type User struct {
	ID        int    `json:"id" db:"user_id"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
	Email     string `json:"email" db:"email"`
	Password  string `json:"-"` // "-" means omit from JSON
	Age       int    `json:"age,omitempty"`
	Internal  string // no tag — uses field name "Internal" in JSON
}

func main() {
	fmt.Println("=== Struct Tags ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. JSON marshaling with tags
	// ─────────────────────────────────────────────
	u := User{
		ID:        1,
		FirstName: "Alice",
		LastName:  "Smith",
		Email:     "alice@example.com",
		Password:  "secret123",
		Age:       0, // omitempty: omitted because zero value
		Internal:  "data",
	}

	data, _ := json.MarshalIndent(u, "", "  ")
	fmt.Println("JSON output:")
	fmt.Println(string(data))
	// Note: Password is omitted (tag: "-"), Age is omitted (omitempty + zero)

	// ─────────────────────────────────────────────
	// 2. JSON unmarshaling
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Unmarshal --")
	jsonStr := `{"id": 2, "first_name": "Bob", "last_name": "Jones", "email": "bob@example.com", "age": 25}`
	var u2 User
	_ = json.Unmarshal([]byte(jsonStr), &u2)
	fmt.Printf("Unmarshaled: %+v\n", u2)

	// ─────────────────────────────────────────────
	// 3. Reading tags via reflection
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Reading tags --")
	t := reflect.TypeOf(User{})
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		jsonTag := field.Tag.Get("json")
		dbTag := field.Tag.Get("db")
		fmt.Printf("  %-10s json:%-15s db:%s\n", field.Name, jsonTag, dbTag)
	}

	// ─────────────────────────────────────────────
	// 4. Common tag options
	// ─────────────────────────────────────────────
	// json:"name"            — custom JSON key
	// json:"-"               — skip this field
	// json:"name,omitempty"  — omit if zero value
	// json:",string"         — encode int/bool as JSON string
	// db:"column_name"       — database column mapping
	// validate:"required"    — validation rules (third-party)
	// yaml:"key"             — YAML key name

	// ─────────────────────────────────────────────
	// 5. GOTCHA: Tag typos are silent
	// ─────────────────────────────────────────────
	// type Bad struct {
	//     Name string `josn:"name"` // typo! "josn" instead of "json"
	// }
	// This compiles fine but json package ignores the misspelled tag.
	// Use `go vet` to catch some tag errors.
}
