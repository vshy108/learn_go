//go:build ignore

// Section 18, Topic 129: fmt (Formatted I/O)
//
// The fmt package implements formatted I/O.
// Most important functions:
//   Print, Println, Printf      — write to stdout
//   Sprint, Sprintf, Sprintln   — return a string
//   Fprint, Fprintf, Fprintln   — write to io.Writer
//   Scan, Scanf, Scanln         — read from stdin
//   Errorf                      — return formatted error
//
// Run: go run examples/s18_fmt_deep_dive.go

package main

import (
	"fmt"
	"os"
)

type Point struct {
	X, Y int
}

func (p Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

func main() {
	fmt.Println("=== fmt Deep Dive ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. Format verbs
	// ─────────────────────────────────────────────
	n := 42
	f := 3.14159
	s := "hello"
	p := Point{10, 20}

	fmt.Printf("%%v:  %v\n", p)    // default: (10, 20) — uses String()
	fmt.Printf("%%+v: %+v\n", p)   // with field names: {X:10 Y:20}
	fmt.Printf("%%#v: %#v\n", p)   // Go syntax: main.Point{X:10, Y:20}
	fmt.Printf("%%T:  %T\n", p)    // type: main.Point
	fmt.Printf("%%d:  %d\n", n)    // decimal integer
	fmt.Printf("%%b:  %b\n", n)    // binary: 101010
	fmt.Printf("%%o:  %o\n", n)    // octal: 52
	fmt.Printf("%%x:  %x\n", n)    // hex: 2a
	fmt.Printf("%%X:  %X\n", n)    // hex uppercase: 2A
	fmt.Printf("%%f:  %f\n", f)    // float: 3.141590
	fmt.Printf("%%.2f: %.2f\n", f) // 2 decimal places: 3.14
	fmt.Printf("%%e:  %e\n", f)    // scientific: 3.141590e+00
	fmt.Printf("%%s:  %s\n", s)    // string
	fmt.Printf("%%q:  %q\n", s)    // quoted: "hello"
	fmt.Printf("%%p:  %p\n", &n)   // pointer

	// ─────────────────────────────────────────────
	// 2. Width and precision
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Width --")
	fmt.Printf("|%10d|\n", 42)     // right-aligned, width 10
	fmt.Printf("|%-10d|\n", 42)    // left-aligned
	fmt.Printf("|%010d|\n", 42)    // zero-padded
	fmt.Printf("|%10.2f|\n", 3.14) // float width + precision

	// ─────────────────────────────────────────────
	// 3. Sprint — return string
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Sprint --")
	msg := fmt.Sprintf("User %s has %d points", "Alice", 100)
	fmt.Println(msg)

	// ─────────────────────────────────────────────
	// 4. Fprint — write to any io.Writer
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Fprint --")
	fmt.Fprintf(os.Stderr, "This goes to stderr\n")
	fmt.Fprintln(os.Stdout, "This goes to stdout")

	// ─────────────────────────────────────────────
	// 5. Errorf
	// ─────────────────────────────────────────────
	err := fmt.Errorf("failed to load %s: %w", "config", fmt.Errorf("not found"))
	fmt.Println("\nError:", err)
}
