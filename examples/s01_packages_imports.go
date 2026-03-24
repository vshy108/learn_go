//go:build ignore

// Section 1, Topic 2: Packages and Imports
//
// Go code is organized into packages. Every .go file must declare its package
// as the first statement. The package name matches the directory name by convention.
//
// The `import` keyword brings other packages into scope.
// You can import one at a time or use grouped (factored) imports.
//
// GOTCHA: Go WILL NOT COMPILE if you import a package and don't use it.
//         This is enforced by the compiler, not just a linter warning.
//         Use the blank identifier `_` for side-effect-only imports (e.g., database drivers).
//
// GOTCHA: Import paths are strings, not identifiers. The last element of the
//         path becomes the package name: "math/rand" → rand.Intn(...)
//
// Run: go run examples/s01_packages_imports.go

package main

import (
	// --- Grouped (factored) import — the Go-standard way ---
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

// You CAN also write individual imports, but grouped is idiomatic:
//   import "fmt"
//   import "math"
// The `goimports` tool will auto-group them for you.

func main() {
	fmt.Println("=== Packages and Imports ===")

	// --- Using imported packages ---
	// The package name is the last segment of the import path.
	fmt.Println("Pi:", math.Pi)             // math package → math.Pi
	fmt.Println("Sqrt(16):", math.Sqrt(16)) // math.Sqrt

	// math/rand: imported as just "rand"
	// GOTCHA: In Go < 1.20, rand.Seed() was needed. Since Go 1.20,
	// the global functions are automatically seeded.
	fmt.Println("Random int:", rand.Intn(100))

	// strings package
	fmt.Println("Upper:", strings.ToUpper("hello"))
	fmt.Println("Contains:", strings.Contains("hello world", "world"))

	// time package
	fmt.Println("Now:", time.Now().Format("2006-01-02 15:04:05"))
	// GOTCHA: Go uses a reference time "Mon Jan 2 15:04:05 MST 2006"
	// (01/02 03:04:05PM '06 -0700) for formatting — NOT strftime codes!

	// --- Exported names ---
	// In Go, a name is exported if it begins with a capital letter.
	// math.Pi is exported (capital P). math.pi would be unexported.
	// You CANNOT access unexported names from outside their package.
	// fmt.Println(math.pi) // ERROR: cannot refer to unexported name math.pi

	// --- Aliased imports ---
	// You can rename an import to avoid conflicts or for convenience.
	// Example (not used here to avoid unused import error):
	//   import (
	//       mrand "math/rand"
	//       crand "crypto/rand"
	//   )

	// --- Blank identifier import (side-effect import) ---
	// Some packages need to be imported just for their init() side effects:
	//   import _ "net/http/pprof"   // registers pprof HTTP handlers
	//   import _ "github.com/lib/pq"  // registers PostgreSQL driver
	// The _ means "import but I won't reference it directly."

	// --- Dot import (discouraged) ---
	// import . "fmt" — lets you call Println() without the fmt prefix.
	// This is almost NEVER used — it pollutes the namespace and confuses readers.

	// --- Internal packages ---
	// Any package under a directory named "internal/" can only be imported
	// by code in the parent tree. This is enforced by the Go toolchain.
	// Example: foo/internal/bar can only be imported by foo/ and its children.

	fmt.Println("\nAll imports demonstrated successfully!")
}
