//go:build ignore

// Section 17, Topic: TDD (Test-Driven Development) in Go
//
// TDD cycle: RED → GREEN → REFACTOR
//   1. RED:    Write a failing test first
//   2. GREEN:  Write the minimum code to make it pass
//   3. REFACTOR: Clean up while keeping tests green
//
// Go makes TDD natural:
//   - go test is built-in (no extra frameworks)
//   - Table-driven tests pair well with incremental TDD
//   - Fast compile = tight feedback loop
//
// This file walks through a TDD example: building a Stack data structure
// step by step, showing each RED→GREEN→REFACTOR cycle.
//
// GOTCHA: In real TDD, tests live in *_test.go files. Here we simulate
//         the workflow in a single file with a mini test runner.
// GOTCHA: TDD is about design — tests drive the API, not just verification.
// GOTCHA: Don't write more production code than needed to pass the test.
//
// Run: go run examples/s17_tdd.go

package main

import "fmt"

// ═══════════════════════════════════════════════════
// Mini test runner (simulates testing.T)
// ═══════════════════════════════════════════════════
type T struct {
	name   string
	failed bool
}

func (t *T) Errorf(format string, args ...any) {
	t.failed = true
	fmt.Printf("    FAIL: %s\n", fmt.Sprintf(format, args...))
}

func run(name string, fn func(t *T)) {
	t := &T{name: name}
	fn(t)
	if t.failed {
		fmt.Printf("  FAIL %s\n", name)
	} else {
		fmt.Printf("  PASS %s\n", name)
	}
}

// ═══════════════════════════════════════════════════
// CYCLE 1 — RED: "A new stack should be empty"
//           GREEN: Create Stack with IsEmpty()
// ═══════════════════════════════════════════════════

// Production code (minimum to pass):
type Stack struct {
	items []int
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// ═══════════════════════════════════════════════════
// CYCLE 2 — RED: "Push adds an item, stack is not empty"
//           GREEN: Add Push()
// ═══════════════════════════════════════════════════

func (s *Stack) Push(val int) {
	s.items = append(s.items, val)
}

// ═══════════════════════════════════════════════════
// CYCLE 3 — RED: "Pop returns last pushed item"
//           GREEN: Add Pop()
// ═══════════════════════════════════════════════════

func (s *Stack) Pop() (int, bool) {
	if len(s.items) == 0 {
		return 0, false // CYCLE 4 drove this: "Pop on empty returns false"
	}
	val := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return val, true
}

// ═══════════════════════════════════════════════════
// CYCLE 5 — RED: "Peek returns top without removing"
//           GREEN: Add Peek()
// ═══════════════════════════════════════════════════

func (s *Stack) Peek() (int, bool) {
	if len(s.items) == 0 {
		return 0, false
	}
	return s.items[len(s.items)-1], true
}

// ═══════════════════════════════════════════════════
// CYCLE 6 — REFACTOR: Add Size() (noticed repeated len checks)
// ═══════════════════════════════════════════════════

func (s *Stack) Size() int {
	return len(s.items)
}

func main() {
	fmt.Println("=== TDD: Building a Stack ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// Cycle 1: New stack is empty
	// ─────────────────────────────────────────────
	fmt.Println("Cycle 1: New stack is empty")
	run("new stack should be empty", func(t *T) {
		s := NewStack()
		if !s.IsEmpty() {
			t.Errorf("expected empty stack")
		}
	})

	// ─────────────────────────────────────────────
	// Cycle 2: Push makes it non-empty
	// ─────────────────────────────────────────────
	fmt.Println("\nCycle 2: Push makes it non-empty")
	run("push one item, not empty", func(t *T) {
		s := NewStack()
		s.Push(42)
		if s.IsEmpty() {
			t.Errorf("expected non-empty after push")
		}
	})

	// ─────────────────────────────────────────────
	// Cycle 3: Pop returns last pushed
	// ─────────────────────────────────────────────
	fmt.Println("\nCycle 3: Pop returns last pushed (LIFO)")
	run("pop returns last pushed value", func(t *T) {
		s := NewStack()
		s.Push(1)
		s.Push(2)
		s.Push(3)
		val, ok := s.Pop()
		if !ok || val != 3 {
			t.Errorf("expected 3, got %d (ok=%v)", val, ok)
		}
	})

	run("pop removes item", func(t *T) {
		s := NewStack()
		s.Push(1)
		s.Push(2)
		s.Pop()
		val, ok := s.Pop()
		if !ok || val != 1 {
			t.Errorf("expected 1, got %d (ok=%v)", val, ok)
		}
	})

	// ─────────────────────────────────────────────
	// Cycle 4: Pop on empty stack
	// ─────────────────────────────────────────────
	fmt.Println("\nCycle 4: Pop on empty returns false")
	run("pop on empty returns false", func(t *T) {
		s := NewStack()
		_, ok := s.Pop()
		if ok {
			t.Errorf("expected ok=false for empty pop")
		}
	})

	// ─────────────────────────────────────────────
	// Cycle 5: Peek
	// ─────────────────────────────────────────────
	fmt.Println("\nCycle 5: Peek returns top without removing")
	run("peek returns top", func(t *T) {
		s := NewStack()
		s.Push(10)
		s.Push(20)
		val, ok := s.Peek()
		if !ok || val != 20 {
			t.Errorf("expected peek=20, got %d", val)
		}
	})

	run("peek does not remove", func(t *T) {
		s := NewStack()
		s.Push(10)
		s.Peek()
		if s.IsEmpty() {
			t.Errorf("peek should not remove item")
		}
	})

	run("peek on empty returns false", func(t *T) {
		s := NewStack()
		_, ok := s.Peek()
		if ok {
			t.Errorf("expected ok=false for empty peek")
		}
	})

	// ─────────────────────────────────────────────
	// Cycle 6: Size (refactor-driven)
	// ─────────────────────────────────────────────
	fmt.Println("\nCycle 6: Size")
	run("size tracks push/pop", func(t *T) {
		s := NewStack()
		if s.Size() != 0 {
			t.Errorf("expected size 0")
		}
		s.Push(1)
		s.Push(2)
		if s.Size() != 2 {
			t.Errorf("expected size 2, got %d", s.Size())
		}
		s.Pop()
		if s.Size() != 1 {
			t.Errorf("expected size 1, got %d", s.Size())
		}
	})

	// ─────────────────────────────────────────────
	// Summary
	// ─────────────────────────────────────────────
	fmt.Println("\n-- TDD Summary --")
	fmt.Println("Each cycle followed RED -> GREEN -> REFACTOR:")
	fmt.Println("  Cycle 1: Test IsEmpty -> implement struct + IsEmpty")
	fmt.Println("  Cycle 2: Test Push    -> implement Push")
	fmt.Println("  Cycle 3: Test Pop     -> implement Pop (basic)")
	fmt.Println("  Cycle 4: Test Pop empty -> add empty check (ok bool)")
	fmt.Println("  Cycle 5: Test Peek    -> implement Peek")
	fmt.Println("  Cycle 6: Refactor     -> extract Size (noticed pattern)")
	fmt.Println()
	fmt.Println("Key TDD principles:")
	fmt.Println("  - Write the test BEFORE the code")
	fmt.Println("  - Write MINIMUM code to pass (no speculative features)")
	fmt.Println("  - Refactor only when tests are green")
	fmt.Println("  - Each test should fail for a clear reason first")
	fmt.Println("  - Tests document expected behavior")
	fmt.Println()
	fmt.Println("In real Go:")
	fmt.Println("  // stack.go       -> production code")
	fmt.Println("  // stack_test.go  -> tests (go test ./...)")
	fmt.Println("  // Run specific:  go test -run TestPop ./...")
	fmt.Println("  // With coverage: go test -cover ./...")
	fmt.Println("  // Verbose:       go test -v ./...")
}
