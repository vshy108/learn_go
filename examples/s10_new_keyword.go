//go:build ignore

// Section 10, Topic 75: new() Function
//
// `new(T)` allocates zeroed memory for type T and returns *T.
// It's rarely used in Go — composite literals (&T{}) are preferred.
//
// GOTCHA: new() returns a pointer to the ZERO VALUE. Fields are not initialized.
// GOTCHA: new() is useful for basic types: new(int) gives you a *int pointing to 0.
//
// Run: go run examples/s10_new_keyword.go

package main

import "fmt"

type Config struct {
	Host string
	Port int
}

func main() {
	fmt.Println("=== new() Function ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. new with basic types
	// ─────────────────────────────────────────────
	ip := new(int) // *int pointing to 0
	fmt.Printf("new(int): %d (type: %T)\n", *ip, ip)

	sp := new(string) // *string pointing to ""
	fmt.Printf("new(string): %q\n", *sp)

	bp := new(bool) // *bool pointing to false
	fmt.Printf("new(bool): %t\n", *bp)

	// ─────────────────────────────────────────────
	// 2. new with structs
	// ─────────────────────────────────────────────
	fmt.Println("\n-- new with structs --")
	cfg := new(Config) // *Config, all fields zero
	fmt.Printf("new(Config): %+v\n", *cfg)

	cfg.Host = "localhost"
	cfg.Port = 8080
	fmt.Printf("After setting: %+v\n", *cfg)

	// ─────────────────────────────────────────────
	// 3. new vs composite literal (preferred)
	// ─────────────────────────────────────────────
	fmt.Println("\n-- new vs &T{} --")

	// These are equivalent:
	p1 := new(Config)
	p2 := &Config{}
	fmt.Printf("new():  %+v (type: %T)\n", *p1, p1)
	fmt.Printf("&T{}:   %+v (type: %T)\n", *p2, p2)

	// But &T{} lets you initialize fields:
	p3 := &Config{Host: "example.com", Port: 443}
	fmt.Printf("&T{..}: %+v\n", *p3)

	// new(Config) cannot set fields inline — this is why &T{} is preferred.

	// ─────────────────────────────────────────────
	// 4. GOTCHA: new doesn't initialize maps/slices
	// ─────────────────────────────────────────────
	fmt.Println("\n-- new doesn't initialize maps --")
	type App struct {
		Settings map[string]string
	}
	app := new(App)
	fmt.Printf("app.Settings: %v (nil=%t)\n", app.Settings, app.Settings == nil)
	// app.Settings["key"] = "val"  // PANIC: nil map!
	// You still need to: app.Settings = make(map[string]string)

	// ─────────────────────────────────────────────
	// 5. new for getting pointer to basic type inline
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Pointer to literal --")
	// You can't do: p := &42  // ERROR
	// But you can use a helper or new:
	p := newInt(42)
	fmt.Printf("Pointer to 42: %d\n", *p)
}

// Helper to get pointer to a value (common pattern pre-generics)
func newInt(v int) *int {
	return &v
}
