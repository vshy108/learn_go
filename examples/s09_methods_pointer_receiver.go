//go:build ignore

// Section 9, Topic 68: Methods with Pointer Receiver
//
// Pointer receivers can modify the struct they're called on.
// Use pointer receivers when:
//   1. You need to modify the receiver
//   2. The struct is large (avoid copying)
//   3. Consistency — if any method needs pointer receiver, use it for all
//
// GOTCHA: You can call pointer receiver methods on values (Go takes address automatically).
// GOTCHA: If a type has ANY pointer receiver methods, ALL methods should use pointer receivers.
//
// Run: go run examples/s09_methods_pointer_receiver.go

package main

import "fmt"

type Counter struct {
	count int
}

// Pointer receiver — can modify the struct
func (c *Counter) Increment() {
	c.count++
}

func (c *Counter) Add(n int) {
	c.count += n
}

func (c *Counter) Reset() {
	c.count = 0
}

// Value receiver for read-only operations is fine,
// but for consistency, often pointer receiver is used for all
func (c *Counter) Value() int {
	return c.count
}

func (c *Counter) String() string {
	return fmt.Sprintf("Counter(%d)", c.count)
}

func main() {
	fmt.Println("=== Methods — Pointer Receiver ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. Pointer receiver modifies original
	// ─────────────────────────────────────────────
	c := Counter{}
	c.Increment()
	c.Increment()
	c.Add(10)
	fmt.Printf("Count: %d\n", c.Value()) // 12

	c.Reset()
	fmt.Printf("After reset: %d\n", c.Value()) // 0

	// ─────────────────────────────────────────────
	// 2. Auto-addressing: value → pointer
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Auto-addressing --")
	// c is a value, but Go takes &c automatically:
	c.Increment() // Go translates to (&c).Increment()
	fmt.Printf("Called on value: %d\n", c.Value())

	// Explicit pointer:
	cp := &Counter{count: 100}
	cp.Increment()
	fmt.Printf("Called on pointer: %d\n", cp.Value())

	// ─────────────────────────────────────────────
	// 3. GOTCHA: Can't call pointer methods on non-addressable values
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Non-addressable values --")
	// Counter{}.Increment()  // ERROR: cannot call pointer method on Counter{}
	// The literal Counter{} is not addressable.
	// Must assign to variable first:
	temp := Counter{}
	temp.Increment()
	fmt.Printf("Temp: %d\n", temp.Value())

	// ─────────────────────────────────────────────
	// 4. Large struct — always use pointer receiver
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Large struct --")
	big := BigData{data: [1024]byte{1, 2, 3}}
	big.Process() // only copies the pointer, not 1024 bytes
	fmt.Printf("BigData first byte: %d\n", big.data[0])
}

type BigData struct {
	data [1024]byte
}

func (b *BigData) Process() {
	b.data[0] = 255
}
