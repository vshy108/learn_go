//go:build ignore

// Section 15, Topic 111: Package Declaration and Imports
//
// Every Go file belongs to a package (declared on line 1 after build tags).
// package main → executable program (has func main())
// package xyz  → library package (importable by other code)
//
// Import syntax:
//   import "fmt"
//   import (
//       "fmt"
//       "os"
//   )
//
// GOTCHA: Unused imports are compile errors (use _ for side-effects).
// GOTCHA: Package name = directory name (by convention, not enforced).
// GOTCHA: Only Uppercase identifiers are exported (visible outside package).
//
// Run: go run examples/s15_packages_imports.go

package main

import (
	"fmt"
	"math"
	"strings"

	// Aliased import:
	str "strconv"
	// Side-effect import (init function only):
	// _ "image/png"
	// Dot import (puts all exports in current scope — avoid):
	// . "fmt"
)

func main() {
	fmt.Println("=== Packages & Imports ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. Using standard library packages
	// ─────────────────────────────────────────────
	fmt.Println("Pi:", math.Pi)
	fmt.Println("Upper:", strings.ToUpper("hello"))

	// ─────────────────────────────────────────────
	// 2. Aliased import
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Aliased import --")
	n, _ := str.Atoi("42") // str is alias for strconv
	fmt.Println("Parsed:", n)

	// ─────────────────────────────────────────────
	// 3. Export rules
	// ─────────────────────────────────────────────
	// Uppercase → exported (public):
	//   fmt.Println, math.Pi, strings.Builder
	// Lowercase → unexported (private to package):
	//   Only accessible within the same package

	// ─────────────────────────────────────────────
	// 4. Package structure
	// ─────────────────────────────────────────────
	//   myproject/
	//   ├── go.mod
	//   ├── main.go         → package main
	//   ├── handler/
	//   │   └── handler.go  → package handler
	//   └── model/
	//       └── user.go     → package model
	//
	// Import: "myproject/handler"

	// ─────────────────────────────────────────────
	// 5. init() function
	// ─────────────────────────────────────────────
	// Each package can have init() functions.
	// They run automatically before main().
	// Used for initialization (DB connections, config loading).
	// A file can have multiple init() functions.
	fmt.Println("\n-- init() --")
	fmt.Println("init() already ran (see output above)")
}

func init() {
	fmt.Println("[init] Package initialized!")
}
