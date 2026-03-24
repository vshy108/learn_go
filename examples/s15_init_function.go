//go:build ignore

// Section 15, Topic 114: init() Function
//
// init() is a special function called automatically before main().
// Each file can have multiple init() functions.
// init() cannot be called explicitly.
//
// Execution order:
//   1. Package-level variables initialized
//   2. init() functions run (in source file order, then file order)
//   3. main() runs
//
// For imported packages: dependency init() runs first (bottom-up).
//
// GOTCHA: Overusing init() makes code harder to test and reason about.
// GOTCHA: init() has no parameters and no return values.
// GOTCHA: Side-effect imports (_ "pkg") run that package's init().
//
// Run: go run examples/s15_init_function.go

package main

import "fmt"

// Package-level variables (evaluated first):
var (
	appName    = "MyApp"
	appVersion = computeVersion()
)

func computeVersion() string {
	fmt.Println("[computeVersion] Computing version...")
	return "1.0.0"
}

// First init (runs after package vars):
func init() {
	fmt.Println("[init 1] First init function")
	fmt.Printf("[init 1] appName=%s, appVersion=%s\n", appName, appVersion)
}

// Second init in same file (runs after first):
func init() {
	fmt.Println("[init 2] Second init function")
}

func main() {
	fmt.Println("\n=== init() Function ===")
	fmt.Println("main() is running now")
	fmt.Printf("App: %s v%s\n", appName, appVersion)

	// ─────────────────────────────────────────────
	// Output order:
	// ─────────────────────────────────────────────
	// [computeVersion] Computing version...
	// [init 1] First init function
	// [init 1] appName=MyApp, appVersion=1.0.0
	// [init 2] Second init function
	// === init() Function ===
	// main() is running now
	// App: MyApp v1.0.0

	// ─────────────────────────────────────────────
	// Common uses of init():
	// ─────────────────────────────────────────────
	// 1. Register database drivers: _ "github.com/lib/pq"
	// 2. Set up package-level configuration
	// 3. Validate build environment
	// 4. Register codecs or formats: _ "image/png"
	//
	// Best practice:
	// - Keep init() simple and fast
	// - Avoid init() for complex logic (use explicit initialization)
	// - Don't rely on init() ordering across packages
}
