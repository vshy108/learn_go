//go:build ignore

// Section 5, Topic 40: Labels, break, continue, goto
//
// Go supports labels for break, continue, and goto:
//   - break LABEL:    exits the labeled loop
//   - continue LABEL: skips to next iteration of labeled loop
//   - goto LABEL:     jumps to label (rarely used, discouraged)
//
// GOTCHA: `break` in a switch only exits the switch, not an enclosing loop.
//         Use a label to break the loop from inside a switch or select.
// GOTCHA: goto cannot jump over variable declarations.
// GOTCHA: goto is almost never needed — use sparingly for error cleanup.
//
// Run: go run examples/s05_labels_break_continue.go

package main

import "fmt"

func main() {
	fmt.Println("=== Labels, break, continue, goto ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. break with label (exit outer loop)
	// ─────────────────────────────────────────────
	fmt.Println("-- break with label --")
Outer:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 2 && j == 3 {
				fmt.Printf("Breaking at (%d,%d)\n", i, j)
				break Outer // exits BOTH loops
			}
			fmt.Printf("(%d,%d) ", i, j)
		}
		fmt.Println()
	}

	// ─────────────────────────────────────────────
	// 2. continue with label (skip to next outer iteration)
	// ─────────────────────────────────────────────
	fmt.Println("\n-- continue with label --")
OuterLoop:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if j == 1 {
				continue OuterLoop // skip rest of inner loop, go to next i
			}
			fmt.Printf("(%d,%d) ", i, j)
		}
	}
	fmt.Println()

	// ─────────────────────────────────────────────
	// 3. break with label in switch (common pattern)
	// ─────────────────────────────────────────────
	fmt.Println("\n-- break label in switch --")
Loop:
	for i := 0; i < 10; i++ {
		switch {
		case i == 5:
			fmt.Println("\nBreaking loop from switch at i=5")
			break Loop // without label, this would only exit the switch
		default:
			fmt.Printf("%d ", i)
		}
	}

	// ─────────────────────────────────────────────
	// 4. goto (use sparingly!)
	// ─────────────────────────────────────────────
	fmt.Println("\n-- goto --")
	n := 0
increment:
	n++
	if n < 5 {
		goto increment
	}
	fmt.Println("n after goto loop:", n)

	// ─────────────────────────────────────────────
	// 5. goto for error cleanup (acceptable use)
	// ─────────────────────────────────────────────
	fmt.Println("\n-- goto for cleanup --")
	if err := processSteps(); err != nil {
		fmt.Println("Error:", err)
	}

	// ─────────────────────────────────────────────
	// 6. GOTCHA: goto cannot skip variable declarations
	// ─────────────────────────────────────────────
	// goto skip
	// x := 42          // ERROR: goto skip jumps over declaration of x
	// skip:
	// fmt.Println(x)

	// ─────────────────────────────────────────────
	// 7. Labels must be used
	// ─────────────────────────────────────────────
	// Unused labels are compile errors (like unused variables):
	// unusedLabel:  // ERROR: label unusedLabel defined and not used
}

func processSteps() error {
	fmt.Println("  Step 1: allocate resources")
	// If step 2 fails, jump to cleanup
	if err := step2(); err != nil {
		goto cleanup
	}
	fmt.Println("  Step 3: finalize")
	return nil

cleanup:
	fmt.Println("  Cleaning up after failure")
	return fmt.Errorf("step 2 failed")
}

func step2() error {
	return fmt.Errorf("simulated error")
}
