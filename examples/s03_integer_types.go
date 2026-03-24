//go:build ignore

// Section 3, Topic 17: Integer Types — int8/16/32/64, uint8/16/32/64, int, uint, uintptr
//
// Go provides explicit-width integers and platform-sized integers.
//
// Signed:   int8 (-128..127), int16, int32, int64
// Unsigned: uint8 (0..255), uint16, uint32, uint64
// Platform: int (32 or 64 bit), uint (32 or 64 bit), uintptr (pointer-sized)
//
// GOTCHA: int is NOT an alias for int32 or int64. It's a distinct type.
//         On a 64-bit machine, int is 64 bits but int != int64.
// GOTCHA: uintptr is for low-level pointer arithmetic. Don't use it for general math.
//
// Run: go run examples/s03_integer_types.go

package main

import (
	"fmt"
	"math"
	"unsafe"
)

func main() {
	fmt.Println("=== Integer Types ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. Signed integer ranges
	// ─────────────────────────────────────────────
	fmt.Println("-- Signed integer ranges --")
	fmt.Printf("int8:  %d to %d  (%d bytes)\n", math.MinInt8, math.MaxInt8, unsafe.Sizeof(int8(0)))
	fmt.Printf("int16: %d to %d  (%d bytes)\n", math.MinInt16, math.MaxInt16, unsafe.Sizeof(int16(0)))
	fmt.Printf("int32: %d to %d  (%d bytes)\n", math.MinInt32, math.MaxInt32, unsafe.Sizeof(int32(0)))
	fmt.Printf("int64: %d to %d  (%d bytes)\n", math.MinInt64, math.MaxInt64, unsafe.Sizeof(int64(0)))
	fmt.Printf("int:   %d to %d  (%d bytes)\n", math.MinInt, math.MaxInt, unsafe.Sizeof(int(0)))

	// ─────────────────────────────────────────────
	// 2. Unsigned integer ranges
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Unsigned integer ranges --")
	fmt.Printf("uint8:  0 to %d  (%d bytes)\n", math.MaxUint8, unsafe.Sizeof(uint8(0)))
	fmt.Printf("uint16: 0 to %d  (%d bytes)\n", math.MaxUint16, unsafe.Sizeof(uint16(0)))
	fmt.Printf("uint32: 0 to %d  (%d bytes)\n", math.MaxUint32, unsafe.Sizeof(uint32(0)))
	// MaxUint64 needs special handling — it overflows int
	fmt.Printf("uint64: 0 to %d  (%d bytes)\n", uint64(math.MaxUint64), unsafe.Sizeof(uint64(0)))
	fmt.Printf("uint:   0 to %d  (%d bytes)\n", uint(math.MaxUint), unsafe.Sizeof(uint(0)))

	// ─────────────────────────────────────────────
	// 3. Platform-sized types
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Platform-sized types --")
	fmt.Printf("int size:     %d bytes\n", unsafe.Sizeof(int(0)))
	fmt.Printf("uint size:    %d bytes\n", unsafe.Sizeof(uint(0)))
	fmt.Printf("uintptr size: %d bytes\n", unsafe.Sizeof(uintptr(0)))
	// On 64-bit: all are 8 bytes. On 32-bit: all are 4 bytes.

	// ─────────────────────────────────────────────
	// 4. Aliases: byte and rune
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Aliases --")
	var b byte = 255 // byte is alias for uint8
	var r rune = '世' // rune is alias for int32
	fmt.Printf("byte: %d (alias for uint8), size=%d\n", b, unsafe.Sizeof(b))
	fmt.Printf("rune: %d '%c' (alias for int32), size=%d\n", r, r, unsafe.Sizeof(r))

	// ─────────────────────────────────────────────
	// 5. Integer literals in different bases
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Integer literal bases --")
	decimal := 42
	binary := 0b101010 // Go 1.13+
	octal := 0o52      // Go 1.13+ (also 052 works, but 0o is clearer)
	hex := 0x2A
	fmt.Printf("Decimal: %d, Binary: 0b%b, Octal: 0o%o, Hex: 0x%X\n",
		decimal, binary, octal, hex)

	// Visual separator with underscores:
	billion := 1_000_000_000
	hexColor := 0xFF_AA_CC
	fmt.Printf("Billion: %d, Color: #%06X\n", billion, hexColor)

	// ─────────────────────────────────────────────
	// 6. GOTCHA: int != int64 (distinct types)
	// ─────────────────────────────────────────────
	fmt.Println("\n-- int != int64 --")
	var a int = 42
	var b64 int64 = 42
	// if a == b64 { }  // ERROR: mismatched types int and int64
	if int64(a) == b64 {
		fmt.Println("Equal after explicit conversion")
	}

	// ─────────────────────────────────────────────
	// 7. When to use which integer type
	// ─────────────────────────────────────────────
	// - int:    general purpose (loop counters, indices, etc.)
	// - int64:  when you need guaranteed 64-bit (timestamps, file sizes)
	// - int32:  when interfacing with 32-bit APIs
	// - uint8:  byte manipulation, network protocols
	// - uint:   rare — mainly for bitwise operations
	// - uintptr: only for unsafe pointer arithmetic
}
