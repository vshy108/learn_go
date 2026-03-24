//go:build ignore

// Section 1, Topic 1: Hello World — package main, func main
//
// Every Go program starts with a package declaration. The `main` package is
// special — it defines an executable program (not a library).
//
// The `main()` function in the `main` package is the entry point.
// Unlike Rust/C, Go's main() takes no arguments and returns no value.
// (Use os.Args for CLI arguments, os.Exit() for exit codes.)
//
// GOTCHA: If you declare `package main` but have no `func main()`, the
// compiler will error: "runtime.main_main·f: function main is undeclared in the main package"
//
// GOTCHA: You cannot have multiple `main()` functions in the same package
// (same directory). Each example file here works because we run them
// individually with `go run`, not as a combined package.
//
// Run: go run examples/s01_hello_world.go

package main

import "fmt"

func main() {
	// fmt.Println adds a newline at the end (like Rust's println! or Python's print()).
	// It returns (n int, err error) — the number of bytes written and any error.
	// Most people ignore the return values for stdout printing.
	fmt.Println("Hello, World!")

	// --- Edge case: Println with multiple arguments ---
	// Println inserts spaces between arguments and adds a newline.
	// This is different from Printf which requires explicit formatting.
	fmt.Println("Hello,", "Go", "World!") // Output: Hello, Go World!

	// --- Edge case: Println with no arguments ---
	// Just prints a newline (useful for spacing output).
	fmt.Println()

	// --- Special case: Print (no "ln") ---
	// Print does NOT add a newline. Spaces are added between operands
	// only when NEITHER is a string.
	fmt.Print("No newline here. ")
	fmt.Print("This continues on the same line.\n")

	// --- Special case: main() must be in package main ---
	// If you change `package main` to `package foo`, you get:
	//   "go run: cannot run non-main package"
	// A package named "main" but without func main() also won't compile.

	// --- Comparison with other languages ---
	// Python: no main() required, just top-level code (or if __name__ == "__main__")
	// Rust:   fn main() in any file, but binary crate must have exactly one
	// Go:     func main() in package main — exactly one entry point per binary
}
