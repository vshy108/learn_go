//go:build ignore

// Section 4, Topic 33: init() Function Special Behavior
//
// init() runs automatically before main(). Each file can have multiple init().
// Used for package initialization, registering drivers, etc.
//
// Order: package-level var init -> init() functions -> main()
// Multiple init() in same file: executed in order of appearance.
// Multiple files: alphabetical order (but don't depend on this).
//
// GOTCHA: init() takes no arguments and returns nothing.
// GOTCHA: init() cannot be called explicitly.
// GOTCHA: Overuse of init() makes code harder to test.
//
// Run: go run examples/s04_init_function.go

package main

import "fmt"

// 1. Package-level vars initialize first
var greeting = initGreeting()

func initGreeting() string {
	fmt.Println("1. Package var initialized")
	return "Hello"
}

// 2. First init
func init() {
	fmt.Println("2. First init() called")
}

// 3. Second init (yes, multiple init() allowed!)
func init() {
	fmt.Println("3. Second init() called")
}

func main() {
	fmt.Println("4. main() called")
	fmt.Println("   greeting =", greeting)

	fmt.Println("\n-- Order summary --")
	fmt.Println("1. Package-level var initializers")
	fmt.Println("2. init() functions (in source order)")
	fmt.Println("3. main()")
}
