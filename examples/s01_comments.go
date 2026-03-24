//go:build ignore

// Section 1, Topic 4: Comments
//
// Go supports two styles of comments — same as C/C++/Rust.
//
// Doc comments (godoc) have special conventions:
//   - Start with the name of the thing being documented
//   - Use complete sentences
//   - First sentence becomes the synopsis in generated docs
//
// GOTCHA: Go has NO doc-comment attributes like Rust's /// or //!.
//         Doc comments are just regular // comments placed directly
//         before the declaration, with no blank line in between.
//
// Run: go run examples/s01_comments.go

package main

import "fmt"

// ─────────────────────────────────────────────
// 1. Single-line comments
// ─────────────────────────────────────────────

// Greeting returns a personalized greeting string.
// This is a doc comment — it starts with the function name "Greeting".
// Godoc will pick this up and display it in documentation.
func Greeting(name string) string {
	return "Hello, " + name + "!" // inline comment after code
}

// ─────────────────────────────────────────────
// 2. Multi-line comments (block comments)
// ─────────────────────────────────────────────

/*
Block comments can span multiple lines.
They are often used for:
  - Package-level documentation
  - Temporarily disabling code
  - Long explanations

Unlike Rust's doc comments (///), Go doesn't have a special syntax
for doc comments — just place // comments before the declaration.
*/

// ─────────────────────────────────────────────
// 3. Package doc comment
// ─────────────────────────────────────────────
//
// By convention, package documentation goes in a comment before the
// `package` declaration, OR in a separate file called `doc.go`.
//
// For package main (executables), the doc comment describes the program:
//
//   // Command myapp does something useful.
//   // It supports the following flags: ...
//   package main
//
// For library packages:
//
//   // Package strings implements simple functions to manipulate
//   // UTF-8 encoded strings.
//   package strings

// ─────────────────────────────────────────────
// 4. Doc comment conventions (godoc)
// ─────────────────────────────────────────────

// Add returns the sum of two integers.
//
// Doc comments should:
//   - Start with the name of the thing (Add, not "This function adds...")
//   - Use complete sentences with proper punctuation.
//   - Have a blank comment line to separate paragraphs.
//
// Code blocks in doc comments are indented:
//
//	result := Add(2, 3)
//	fmt.Println(result) // 5
func Add(a, b int) int {
	return a + b
}

// ─────────────────────────────────────────────
// 5. Deprecated marker
// ─────────────────────────────────────────────

// OldFunction does something the old way.
//
// Deprecated: Use NewFunction instead.
// The "Deprecated:" prefix is recognized by godoc and IDEs.
func OldFunction() string {
	return "old"
}

func main() {
	fmt.Println("=== Go Comments ===")
	fmt.Println()

	fmt.Println(Greeting("Gopher"))
	fmt.Println("2 + 3 =", Add(2, 3))

	// --- Build tags (special comments) ---
	// Build tags are placed at the top of a file to control compilation:
	//
	//   //go:build linux && amd64
	//
	// Before Go 1.17, the syntax was:
	//   // +build linux,amd64
	//
	// Common build tags:
	//   //go:build ignore       — file is never compiled (useful for docs/scripts)
	//   //go:build integration  — only compile with: go test -tags=integration

	// --- Compiler directives (//go: comments) ---
	// These are NOT regular comments — they're instructions to the compiler:
	//   //go:generate  — run code generation commands
	//   //go:embed     — embed files at compile time (Go 1.16+)
	//   //go:noinline  — prevent function inlining
	//   //go:nosplit   — prevent stack-split check
	//   //go:linkname  — access unexported symbols (dangerous!)

	// --- Special case: comments and semicolons ---
	// Go uses implicit semicolons. A comment at end of line counts as
	// a line ending, so the semicolon is inserted before the comment.
	// This is why you can't put opening braces on the next line:
	//
	//   if true    // semicolon inserted here!
	//   {          // ERROR: unexpected semicolon or newline before {
	//       ...
	//   }

	fmt.Println("\nComment types demonstrated. Run `go doc` to see generated docs.")
}
