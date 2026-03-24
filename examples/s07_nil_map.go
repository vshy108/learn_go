//go:build ignore

// Section 7, Topic 55: nil Map Gotcha
//
// A nil map behaves like an empty map for reads, but PANICS on write.
//
// var m map[string]int  → m is nil
// m["key"]              → returns 0 (zero value), no panic
// len(m)                → returns 0, no panic
// for range m {}        → no iterations, no panic
// delete(m, "key")      → no-op, no panic
// m["key"] = 1          → PANIC: assignment to entry in nil map
//
// GOTCHA: Always initialize maps before writing to them!
//
// Run: go run examples/s07_nil_map.go

package main

import "fmt"

func main() {
	fmt.Println("=== nil Map Gotcha ===")
	fmt.Println()

	var m map[string]int

	// ─────────────────────────────────────────────
	// 1. Reads from nil map: work fine
	// ─────────────────────────────────────────────
	fmt.Println("-- Reads from nil map --")
	fmt.Printf("nil check: %t\n", m == nil)  // true
	fmt.Printf("len: %d\n", len(m))          // 0
	fmt.Printf("m[\"key\"]: %d\n", m["key"]) // 0 (zero value)
	_, ok := m["key"]
	fmt.Printf("comma-ok: %t\n", ok) // false

	// ─────────────────────────────────────────────
	// 2. Range over nil map: works (zero iterations)
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Range over nil map --")
	// for k, v := range m { ... } // runs zero times, no panic
	fmt.Println("Range over nil map: safe (zero iterations, no panic)")

	// ─────────────────────────────────────────────
	// 3. Delete from nil map: no-op
	// ─────────────────────────────────────────────
	delete(m, "key") // no panic
	fmt.Println("delete from nil: no panic")

	// ─────────────────────────────────────────────
	// 4. Write to nil map: PANIC!
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Write to nil map --")
	// Uncomment to see the panic:
	// m["key"] = 1  // panic: assignment to entry in nil map

	// This is a very common bug, especially in structs:
	type Config struct {
		Settings map[string]string
	}
	var cfg Config
	fmt.Printf("cfg.Settings nil: %t\n", cfg.Settings == nil)
	// cfg.Settings["theme"] = "dark"  // PANIC!

	// Fix: initialize the map
	cfg.Settings = make(map[string]string)
	cfg.Settings["theme"] = "dark"
	fmt.Printf("After init: %v\n", cfg.Settings)

	// ─────────────────────────────────────────────
	// 5. Constructor pattern to avoid nil maps
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Constructor pattern --")
	cfg2 := NewConfig()
	cfg2.Settings["theme"] = "light"
	fmt.Printf("Via constructor: %v\n", cfg2.Settings)
}

type Config struct {
	Settings map[string]string
}

func NewConfig() *Config {
	return &Config{
		Settings: make(map[string]string),
	}
}
