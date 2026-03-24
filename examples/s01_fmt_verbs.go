//go:build ignore

// Section 1, Topic 3: fmt Package — Println, Printf, Sprintf, Format Verbs
//
// The `fmt` package is Go's primary formatting package, similar to Python's
// str.format() or Rust's format!/println! macros.
//
// Key functions:
//   - Println:  prints with spaces between args, adds newline
//   - Printf:   formatted output (like C's printf), NO newline
//   - Sprintf:  like Printf but returns a string instead of printing
//   - Fprintf:  prints to an io.Writer (file, buffer, etc.)
//   - Errorf:   returns an error with formatted message
//
// GOTCHA: Printf does NOT add a newline — you must include \n yourself.
// GOTCHA: Println adds spaces between ALL arguments, which may not be what you want.
//
// Run: go run examples/s01_fmt_verbs.go

package main

import "fmt"

func main() {
	fmt.Println("=== fmt Package Deep Dive ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. General verbs (work for any type)
	// ─────────────────────────────────────────────
	name := "Go"
	age := 15
	pi := 3.14159

	fmt.Println("-- General Verbs --")
	fmt.Printf("%%v  (default format):    %v\n", name)                 // Go
	fmt.Printf("%%v  (struct default):    %v\n", struct{ X int }{42})  // {42}
	fmt.Printf("%%+v (struct with names): %+v\n", struct{ X int }{42}) // {X:42}
	fmt.Printf("%%#v (Go syntax repr):    %#v\n", name)                // "Go"
	fmt.Printf("%%T  (type of value):     %T\n", pi)                   // float64
	fmt.Printf("%%%%  (literal percent):   %%\n")

	// ─────────────────────────────────────────────
	// 2. Integer verbs
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Integer Verbs --")
	n := 42
	fmt.Printf("%%d  (decimal):        %d\n", n)   // 42
	fmt.Printf("%%b  (binary):         %b\n", n)   // 101010
	fmt.Printf("%%o  (octal):          %o\n", n)   // 52
	fmt.Printf("%%O  (octal with 0o):  %O\n", n)   // 0o52
	fmt.Printf("%%x  (hex lowercase):  %x\n", n)   // 2a
	fmt.Printf("%%X  (hex uppercase):  %X\n", n)   // 2A
	fmt.Printf("%%c  (character):      %c\n", 65)  // A (Unicode code point)
	fmt.Printf("%%U  (Unicode):        %U\n", 'A') // U+0041
	fmt.Printf("%%q  (quoted char):    %q\n", 'A') // 'A'

	// ─────────────────────────────────────────────
	// 3. Float verbs
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Float Verbs --")
	f := 3.14159265
	fmt.Printf("%%f  (decimal point):  %f\n", f)    // 3.141593
	fmt.Printf("%%e  (scientific):     %e\n", f)    // 3.141593e+00
	fmt.Printf("%%E  (scientific):     %E\n", f)    // 3.141593E+00
	fmt.Printf("%%g  (compact):        %g\n", f)    // 3.14159265
	fmt.Printf("%%.2f (2 decimals):    %.2f\n", f)  // 3.14
	fmt.Printf("%%9.2f (width 9, 2dp): %9.2f\n", f) // "     3.14"

	// ─────────────────────────────────────────────
	// 4. String verbs
	// ─────────────────────────────────────────────
	fmt.Println("\n-- String Verbs --")
	s := "Hello, 世界"
	fmt.Printf("%%s  (plain string):   %s\n", s) // Hello, 世界
	fmt.Printf("%%q  (quoted string):  %q\n", s) // "Hello, 世界"
	fmt.Printf("%%x  (hex of bytes):   %x\n", s) // 48656c6c6f2c20e4b896e7958c

	// ─────────────────────────────────────────────
	// 5. Width and padding
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Width and Padding --")
	fmt.Printf("|%10d| (right-aligned)\n", age)         // |        15|
	fmt.Printf("|%-10d| (left-aligned)\n", age)         // |15        |
	fmt.Printf("|%010d| (zero-padded)\n", age)          // |0000000015|
	fmt.Printf("|%+d|       (always show sign)\n", age) // |+15|
	fmt.Printf("|%+d|       (negative sign)\n", -age)   // |-15|
	fmt.Printf("|%10s| (right-aligned string)\n", "Go") // |        Go|
	fmt.Printf("|%-10s| (left-aligned string)\n", "Go") // |Go        |

	// ─────────────────────────────────────────────
	// 6. Boolean and pointer verbs
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Boolean and Pointer --")
	b := true
	fmt.Printf("%%t  (boolean):        %t\n", b) // true
	x := 42
	fmt.Printf("%%p  (pointer):        %p\n", &x) // 0xc0000b4008 (varies)

	// ─────────────────────────────────────────────
	// 7. Sprintf — returns a string
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Sprintf (returns string, does not print) --")
	result := fmt.Sprintf("Name: %s, Age: %d, Pi: %.2f", name, age, pi)
	fmt.Println(result)

	// ─────────────────────────────────────────────
	// 8. GOTCHA: Wrong verb for type
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Gotcha: Wrong verb --")
	// Using %d for a string won't panic at runtime, but prints a warning format.
	// go vet catches these as compile errors, so we show them commented out:
	//
	//   fmt.Printf("%d\n", "oops")  // output: %!d(string=oops)
	//   fmt.Printf("%s\n", 42)      // output: %!s(int=42)
	//
	// Go does NOT crash — it prints a diagnostic. Rust would refuse to compile.
	// Use %v as the universal verb if you don't know the type:
	fmt.Printf("%%v for string: %v\n", "oops") // oops
	fmt.Printf("%%v for int:    %v\n", 42)     // 42

	// ─────────────────────────────────────────────
	// 9. GOTCHA: Missing/extra arguments
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Gotcha: Argument count mismatch --")
	// go vet also catches argument count mismatches, so we show them commented:
	//
	//   fmt.Printf("Missing: %d %d\n", 1)      // output: 1 %!d(MISSING)
	//   fmt.Printf("Extra: %d\n", 1, 2, 3)      // output: 1%!(EXTRA int=2, int=3)
	//
	// go vet will catch these at static analysis time!
	fmt.Println("(See source comments for examples — go vet flags them as errors.)")
}
