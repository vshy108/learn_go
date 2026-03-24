//go:build ignore

// Section 20, Topic 146: unsafe package
//
// The unsafe package bypasses Go's type safety. It should be avoided
// unless you need low-level memory manipulation (e.g., syscalls, CGo).
//
// Key functions:
//   unsafe.Sizeof(x)      — size of x in bytes (compile-time)
//   unsafe.Alignof(x)     — alignment of x
//   unsafe.Offsetof(s.f)  — byte offset of field f in struct s
//   unsafe.Pointer(p)     — generic pointer (can convert between types)
//
// GOTCHA: unsafe operations may break across Go versions.
// GOTCHA: No garbage collector tracking for unsafe.Pointer arithmetic.
// GOTCHA: Incorrect use causes undefined behavior.
//
// Run: go run examples/s20_unsafe_pointer.go

package main

import (
	"fmt"
	"unsafe"
)

type Example struct {
	A bool   // 1 byte
	B int32  // 4 bytes
	C int64  // 8 bytes
	D bool   // 1 byte
}

type Compact struct {
	C int64  // 8 bytes
	B int32  // 4 bytes
	A bool   // 1 byte
	D bool   // 1 byte
	// Total: 16 bytes (better alignment)
}

func main() {
	fmt.Println("=== unsafe Package ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. Sizeof — type sizes
	// ─────────────────────────────────────────────
	fmt.Println("-- Sizes --")
	fmt.Println("bool:", unsafe.Sizeof(true))
	fmt.Println("int8:", unsafe.Sizeof(int8(0)))
	fmt.Println("int16:", unsafe.Sizeof(int16(0)))
	fmt.Println("int32:", unsafe.Sizeof(int32(0)))
	fmt.Println("int64:", unsafe.Sizeof(int64(0)))
	fmt.Println("int:", unsafe.Sizeof(int(0))) // platform-dependent (8 on 64-bit)
	fmt.Println("string:", unsafe.Sizeof(""))  // 16 (pointer + length)
	fmt.Println("slice:", unsafe.Sizeof([]int{})) // 24 (pointer + length + cap)

	// ─────────────────────────────────────────────
	// 2. Struct padding and alignment
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Struct layout (padding matters!) --")
	var e Example
	fmt.Printf("Example: size=%d, align=%d\n", unsafe.Sizeof(e), unsafe.Alignof(e))
	fmt.Printf("  A (bool):  offset=%d, size=%d\n", unsafe.Offsetof(e.A), unsafe.Sizeof(e.A))
	fmt.Printf("  B (int32): offset=%d, size=%d\n", unsafe.Offsetof(e.B), unsafe.Sizeof(e.B))
	fmt.Printf("  C (int64): offset=%d, size=%d\n", unsafe.Offsetof(e.C), unsafe.Sizeof(e.C))
	fmt.Printf("  D (bool):  offset=%d, size=%d\n", unsafe.Offsetof(e.D), unsafe.Sizeof(e.D))
	// Example is likely 24 bytes due to padding

	var c Compact
	fmt.Printf("\nCompact: size=%d, align=%d\n", unsafe.Sizeof(c), unsafe.Alignof(c))
	fmt.Printf("  C (int64): offset=%d\n", unsafe.Offsetof(c.C))
	fmt.Printf("  B (int32): offset=%d\n", unsafe.Offsetof(c.B))
	fmt.Printf("  A (bool):  offset=%d\n", unsafe.Offsetof(c.A))
	fmt.Printf("  D (bool):  offset=%d\n", unsafe.Offsetof(c.D))
	// Compact is likely 16 bytes (better field ordering)

	// GOTCHA: Struct field order affects memory usage!
	// Order fields from largest to smallest to minimize padding.

	// ─────────────────────────────────────────────
	// 3. unsafe.Pointer conversions
	// ─────────────────────────────────────────────
	fmt.Println("\n-- unsafe.Pointer --")
	x := int64(42)
	// *int64 → unsafe.Pointer → *float64 (reinterpret bits)
	p := unsafe.Pointer(&x)
	fp := (*float64)(p)
	fmt.Printf("int64 %d bits reinterpreted as float64: %g\n", x, *fp)

	// ─────────────────────────────────────────────
	// 4. String internals via unsafe
	// ─────────────────────────────────────────────
	fmt.Println("\n-- String header --")
	s := "Hello, Go!"
	// A string is a struct { ptr *byte; len int }
	type stringHeader struct {
		Data uintptr
		Len  int
	}
	sh := (*stringHeader)(unsafe.Pointer(&s))
	fmt.Printf("String: %q → Data=%x, Len=%d\n", s, sh.Data, sh.Len)

	// ─────────────────────────────────────────────
	// Safety rules (from Go spec):
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Rules for valid unsafe.Pointer use --")
	fmt.Println("1. Convert *T to unsafe.Pointer (and back)")
	fmt.Println("2. Convert unsafe.Pointer to uintptr (for printing)")
	fmt.Println("3. Convert uintptr back to unsafe.Pointer ONLY in same expression")
	fmt.Println("4. Use with reflect.Value.Pointer() / UnsafeAddr()")
	fmt.Println("5. Use with syscall.Syscall()")
	fmt.Println()
	fmt.Println("AVOID: storing uintptr in variables (GC may move objects)")
}
