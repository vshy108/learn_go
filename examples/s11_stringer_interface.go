//go:build ignore

// Section 11, Topic 81: Stringer Interface (fmt.Stringer)
//
// The fmt.Stringer interface controls how a type is printed:
//   type Stringer interface {
//       String() string
//   }
//
// Any type implementing String() will have that output used by
// fmt.Println, fmt.Printf with %v / %s, etc.
//
// GOTCHA: Don't call fmt.Sprintf("%s", t) inside String() — infinite recursion!
// GOTCHA: Pointer receiver vs value receiver matters for interface satisfaction.
//
// Run: go run examples/s11_stringer_interface.go

package main

import "fmt"

// ─────────────────────────────────────────────
// 1. Basic Stringer
// ─────────────────────────────────────────────
type Point struct {
	X, Y int
}

func (p Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

// ─────────────────────────────────────────────
// 2. IP Address example
// ─────────────────────────────────────────────
type IPAddress [4]byte

func (ip IPAddress) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

// ─────────────────────────────────────────────
// 3. Enum with String
// ─────────────────────────────────────────────
type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func (d Weekday) String() string {
	names := [...]string{"Sunday", "Monday", "Tuesday", "Wednesday",
		"Thursday", "Friday", "Saturday"}
	if d < Sunday || d > Saturday {
		return "Unknown"
	}
	return names[d]
}

func main() {
	fmt.Println("=== Stringer Interface ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// Automatic use by fmt
	// ─────────────────────────────────────────────
	p := Point{3, 4}
	fmt.Println("Point:", p)             // uses String()
	fmt.Printf("Printf %%v: %v\n", p)   // uses String()
	fmt.Printf("Printf %%s: %s\n", p)   // uses String()

	ip := IPAddress{192, 168, 1, 1}
	fmt.Println("IP:", ip)

	day := Wednesday
	fmt.Println("Day:", day)
	fmt.Println("Tomorrow:", day+1)

	// ─────────────────────────────────────────────
	// Use with interface
	// ─────────────────────────────────────────────
	fmt.Println("\n-- As interface --")
	stringers := []fmt.Stringer{p, ip, Monday}
	for _, s := range stringers {
		fmt.Printf("  %s\n", s)
	}
}
