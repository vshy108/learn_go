//go:build ignore

// Section 4, Topic 33: init() Function — Special Initialization
//
// The `init()` function is automatically called before main():
//   1. Package-level variable declarations are evaluated
//   2. init() functions run (in source order within a file)
//   3. main() runs
//
// Special rules:
//   - init() takes no arguments and returns nothing
//   - A file can have MULTIPLE init() functions (all will run)
//   - A package can have init() in multiple files (run in file-name order)
//   - init() cannot be called or referenced
//
// GOTCHA: init() runs BEFORE main(). Use sparingly — it makes testing harder.
// GOTCHA: import cycle detection means init() can't create circular deps.
// GOTCHA: Side-effect imports (import _ "pkg") are ONLY for init() effects.
//
// Run: go run examples/s04_init_function.go

package main

import "fmt"

// Package-level variables are initialized FIRST (before init)
var config = "default"
































































}	//   4. main()	//   3. main's package-level vars → main's init()	//   2. A's package-level vars → A's init()	//   1. B's package-level vars → B's init()	// If main imports package A which imports package B:	// ─────────────────────────────────────────────	// Init order across packages:	// ─────────────────────────────────────────────	// _ = init  // ERROR: cannot use init as value	// init()  // ERROR: undefined: init	// ─────────────────────────────────────────────	// GOTCHA: Cannot call init() directly	// ─────────────────────────────────────────────	//   // Call SetupDB() from main() and test setup	//   func SetupDB() (*DB, error) { return connectToDatabase() }	// GOOD:	//	//   func init() { db = connectToDatabase() }	// BAD:	//	// functions that you call from main() or test setup.	// can't be controlled in tests. Prefer explicit initialization	// Because init() runs automatically, side effects in init()	// ─────────────────────────────────────────────	// GOTCHA: init() makes testing harder	// ─────────────────────────────────────────────	fmt.Println("5. Run sanity checks")	fmt.Println("4. Initialize package-level state")	fmt.Println("3. Validate configuration")	fmt.Println("2. Register image decoders:     import _ \"image/png\"")	fmt.Println("1. Register database drivers:   import _ \"github.com/lib/pq\"")	fmt.Println("-- Common init() use cases --")	// ─────────────────────────────────────────────	// Common use cases for init()	// ─────────────────────────────────────────────	// 4. main(): prints config = "fully_initialized"	// 3. init() #2: prints, sets config = "fully_initialized"	// 2. init() #1: prints, sets config = "initialized"	// 1. Package-level var: config = "default"	// ─────────────────────────────────────────────	// Execution order:	// ─────────────────────────────────────────────	fmt.Println()	fmt.Println("main() running — config is:", config)func main() {}	config = "fully_initialized"	fmt.Println("init() #2 running — config is:", config)func init() {// Second init function (yes, multiple init() in same file is legal!)}	config = "initialized"	fmt.Println("init() #1 running — config was:", config)func init() {// First init function