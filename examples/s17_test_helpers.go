//go:build ignore

// Section 17, Topic 125: Test Helpers and testify
//
// testing.T helpers and popular assertion libraries.
//
// t.Helper() marks a function as a test helper, so error messages
// report the caller's line number instead of the helper's.
//
// Popular libraries:
//   github.com/stretchr/testify — assert, require, mock, suite
//
// GOTCHA: Go's stdlib has no assert. You write if-checks manually.
// GOTCHA: testify is third-party — some teams prefer stdlib-only tests.
//
// Run: go run examples/s17_test_helpers.go

package main

import "fmt"

func main() {
	fmt.Println("=== Test Helpers & testify ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. Custom test helper
	// ─────────────────────────────────────────────
	fmt.Print(`
// helpers_test.go
func assertEqual(t *testing.T, got, want int) {
    t.Helper()  // marks this as helper
    if got != want {
        t.Errorf("got %d, want %d", got, want)
        // Error will report CALLER's line number, not this line
    }
}

func TestAdd(t *testing.T) {
    assertEqual(t, add(2, 3), 5)   // line 15 — reported on error
    assertEqual(t, add(-1, 1), 0)  // line 16
}
`)

	// ─────────────────────────────────────────────
	// 2. testify assert
	// ─────────────────────────────────────────────
	fmt.Println("-- testify assert --")
	fmt.Print(`
import "github.com/stretchr/testify/assert"

func TestWithAssert(t *testing.T) {
    assert.Equal(t, 5, add(2, 3))
    assert.NotEqual(t, 0, add(2, 3))
    assert.Nil(t, err)
    assert.NotNil(t, result)
    assert.True(t, isPrime(7))
    assert.Contains(t, "hello world", "world")
    assert.Len(t, slice, 3)
    assert.Error(t, err)
    assert.NoError(t, err)
    assert.ErrorIs(t, err, ErrNotFound)
}
`)

	// ─────────────────────────────────────────────
	// 3. testify require (stops on failure)
	// ─────────────────────────────────────────────
	fmt.Println("-- testify require --")
	fmt.Print(`
import "github.com/stretchr/testify/require"

func TestWithRequire(t *testing.T) {
    result, err := fetchData()
    require.NoError(t, err)       // STOPS test if error
    require.NotNil(t, result)     // STOPS test if nil
    assert.Equal(t, "data", result.Name)  // continues on failure
}
`)

	// ─────────────────────────────────────────────
	// 4. Setup and teardown
	// ─────────────────────────────────────────────
	fmt.Println("-- Setup/Teardown --")
	fmt.Print(`
func TestMain(m *testing.M) {
    // Global setup
    setup()

    code := m.Run()  // run all tests

    // Global teardown
    teardown()

    os.Exit(code)
}

// Per-test setup:
func TestSomething(t *testing.T) {
    // setup
    db := setupDB(t)
    t.Cleanup(func() {
        db.Close()  // runs after test, even on failure
    })
    // ... test code
}
`)

	// ─────────────────────────────────────────────
	// 5. t.Parallel()
	// ─────────────────────────────────────────────
	fmt.Println("-- Parallel tests --")
	fmt.Print(`
func TestA(t *testing.T) {
    t.Parallel()  // run in parallel with other parallel tests
    // ...
}

func TestB(t *testing.T) {
    t.Parallel()
    // ...
}
`)
}
