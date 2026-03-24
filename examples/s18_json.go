//go:build ignore

// Section 18, Topic 131: encoding/json
//
// JSON encoding and decoding (marshal/unmarshal).
//
// Key functions:
//   json.Marshal(v)              — Go → JSON bytes
//   json.Unmarshal(data, &v)     — JSON bytes → Go
//   json.NewEncoder(w).Encode(v) — streaming encode
//   json.NewDecoder(r).Decode(&v) — streaming decode
//
// GOTCHA: Only exported fields are marshaled/unmarshaled.
// GOTCHA: Use struct tags to control JSON field names.
// GOTCHA: json.Number for precise number handling.
//
// Run: go run examples/s18_json.go

package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Email   string   `json:"email,omitempty"` // omit if empty
	Address *Address `json:"address,omitempty"`
	secret  string   // unexported — NOT in JSON
}

type Address struct {
	City    string `json:"city"`
	Country string `json:"country"`
}

func main() {
	fmt.Println("=== encoding/json ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. Marshal (Go → JSON)
	// ─────────────────────────────────────────────
	p := Person{
		Name:    "Alice",
		Age:     30,
		Email:   "alice@example.com",
		Address: &Address{City: "NYC", Country: "US"},
		secret:  "hidden",
	}

	data, _ := json.Marshal(p)
	fmt.Println("Marshal:", string(data))

	// Pretty print:
	pretty, _ := json.MarshalIndent(p, "", "  ")
	fmt.Println("\nPretty:\n" + string(pretty))

	// ─────────────────────────────────────────────
	// 2. Unmarshal (JSON → Go)
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Unmarshal --")
	jsonStr := `{"name":"Bob","age":25,"address":{"city":"London","country":"UK"}}`

	var p2 Person
	if err := json.Unmarshal([]byte(jsonStr), &p2); err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("Unmarshaled: %+v\n", p2)
	fmt.Printf("City: %s\n", p2.Address.City)

	// ─────────────────────────────────────────────
	// 3. omitempty
	// ─────────────────────────────────────────────
	fmt.Println("\n-- omitempty --")
	empty := Person{Name: "Charlie", Age: 0}
	data, _ = json.Marshal(empty)
	fmt.Println("With omitempty:", string(data))
	// Email and Address omitted because they're zero values

	// ─────────────────────────────────────────────
	// 4. map[string]any for dynamic JSON
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Dynamic JSON --")
	var dynamic map[string]any
	jsonDyn := `{"name":"Eve","scores":[95,87,92],"active":true}`
	json.Unmarshal([]byte(jsonDyn), &dynamic)
	fmt.Printf("Dynamic: %+v\n", dynamic)
	fmt.Printf("Name: %v (type: %T)\n", dynamic["name"], dynamic["name"])

	// GOTCHA: Numbers decoded as float64 by default
	scores := dynamic["scores"].([]any)
	fmt.Printf("First score: %v (type: %T)\n", scores[0], scores[0]) // float64!

	// ─────────────────────────────────────────────
	// 5. Struct tags
	// ─────────────────────────────────────────────
	// `json:"name"`          — use "name" as JSON key
	// `json:"name,omitempty"` — omit if zero value
	// `json:"-"`              — always omit
	// `json:",string"`        — encode number as string

	// ─────────────────────────────────────────────
	// 6. json.RawMessage (defer parsing)
	// ─────────────────────────────────────────────
	fmt.Println("\n-- RawMessage --")
	type Event struct {
		Type    string          `json:"type"`
		Payload json.RawMessage `json:"payload"` // parse later
	}
	evt := `{"type":"click","payload":{"x":100,"y":200}}`
	var e Event
	json.Unmarshal([]byte(evt), &e)
	fmt.Printf("Type: %s, Raw payload: %s\n", e.Type, string(e.Payload))
}
