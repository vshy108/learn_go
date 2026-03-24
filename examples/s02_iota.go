//go:build ignore

// Section 2, Topic 12: iota Enumerator
//
// iota is Go's auto-incrementing constant generator.
// It resets to 0 at each const block and increments per line.
//
// GOTCHA: iota increments per SPEC LINE, not per use.
// GOTCHA: Skipping a value with _ still increments iota.
// GOTCHA: iota resets in each new const block.
//
// Run: go run examples/s02_iota.go

package main

import "fmt"

// 1. Basic enum
type Weekday int

const (
	Sunday    Weekday = iota // 0
	Monday                   // 1
	Tuesday                  // 2
	Wednesday                // 3
	Thursday                 // 4
	Friday                   // 5
	Saturday                 // 6
)

// 2. Skip values with _
type Permission int

const (
	Read    Permission = 1 << iota // 1
	Write                          // 2
	Execute                        // 4
	_                              // 8 (skipped)
	Admin                          // 16
)

// 3. Custom starting value
type HTTPStatus int

const (
	StatusOK            HTTPStatus = iota + 200 // 200
	StatusCreated                                // 201
	StatusAccepted                               // 202
	StatusNoContent     HTTPStatus = iota + 200 // 203
)

// 4. iota resets in new block
const (
	A = iota // 0 (reset)
	B        // 1
	C        // 2
)

func main() {
	fmt.Println("=== iota ===")
	fmt.Println()

	// Weekdays
	fmt.Println("-- Weekdays --")
	fmt.Printf("Sunday=%d, Monday=%d, Saturday=%d\n", Sunday, Monday, Saturday)

	// Bit flags
	fmt.Println("\n-- Permissions (bit flags) --")
	fmt.Printf("Read=%d, Write=%d, Execute=%d, Admin=%d\n",
		Read, Write, Execute, Admin)
	perms := Read | Write
	fmt.Printf("Read|Write = %d, has Read: %t, has Execute: %t\n",
		perms, perms&Read != 0, perms&Execute != 0)

	// HTTP Status
	fmt.Println("\n-- HTTP Status --")
	fmt.Printf("OK=%d, Created=%d, Accepted=%d, NoContent=%d\n",
		StatusOK, StatusCreated, StatusAccepted, StatusNoContent)

	// Reset demonstration
	fmt.Println("\n-- Reset per block --")
	fmt.Printf("A=%d, B=%d, C=%d\n", A, B, C)
}
