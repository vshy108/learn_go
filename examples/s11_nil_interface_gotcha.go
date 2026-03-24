//go:build ignore

// Section 11, Topic 89: nil Interface vs Interface Holding nil
//
// This is one of Go's most confusing gotchas.
//
// An interface value has two components: (type, value).
//   - nil interface: (nil, nil) — both type and value are nil
//   - Interface holding nil: (*SomeType, nil) — type is set, value is nil
//
// GOTCHA: An interface holding a nil pointer is NOT nil!
//         if err != nil { ... } can be true even when the underlying value is nil.
//
// Run: go run examples/s11_nil_interface_gotcha.go

package main

import "fmt"

type MyError struct {
	Message string
}

func (e *MyError) Error() string {
	return e.Message
}

func main() {
	fmt.Println("=== nil Interface Gotcha ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. nil interface
	// ─────────────────────────────────────────────
	var err error                                                  // (nil, nil)
	fmt.Printf("nil interface: %v, is nil: %t\n", err, err == nil) // true

	// ─────────────────────────────────────────────
	// 2. Interface holding nil pointer
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Interface holding nil --")
	var myErr *MyError                                         // nil pointer of type *MyError
	fmt.Printf("myErr: %v, is nil: %t\n", myErr, myErr == nil) // true

	err = myErr // assign nil *MyError to error interface
	// err is now (*MyError, nil) — type is set!
	fmt.Printf("err: %v, is nil: %t\n", err, err == nil) // FALSE!

	// ─────────────────────────────────────────────
	// 3. The bug in practice
	// ─────────────────────────────────────────────
	fmt.Println("\n-- The bug --")
	err = getError(false)
	if err != nil {
		fmt.Println("Bug! err is not nil even though no error occurred")
		fmt.Printf("  Type: %T, Value: %v\n", err, err)
	}

	// Fix:
	err = getErrorFixed(false)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Fixed: err is properly nil")
	}

	// ─────────────────────────────────────────────
	// 4. How to check properly
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Proper checks --")
	// Use reflect or type assertion to check if underlying is nil:
	err = getError(false)
	if err != nil {
		if me, ok := err.(*MyError); ok && me == nil {
			fmt.Println("Interface holds nil *MyError")
		}
	}

	// ─────────────────────────────────────────────
	// Rule: ALWAYS return nil explicitly for error interface
	// ─────────────────────────────────────────────
	// BAD:
	//   var err *MyError
	//   return err        // returns non-nil interface!
	//
	// GOOD:
	//   return nil         // returns nil interface
}

// BAD: returns non-nil error interface even when no error
func getError(fail bool) error {
	var err *MyError // nil pointer
	if fail {
		err = &MyError{Message: "failed"}
	}
	return err // BUG: returns (*MyError, nil) — not nil interface!
}

// GOOD: returns nil explicitly
func getErrorFixed(fail bool) error {
	if fail {
		return &MyError{Message: "failed"}
	}
	return nil // returns (nil, nil) — properly nil interface
}
