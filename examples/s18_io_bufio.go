//go:build ignore

// Section 18, Topic 134: io and bufio
//
// io: Core I/O interfaces (Reader, Writer, Closer).
// bufio: Buffered I/O for efficient reading/writing.
//
// Run: go run examples/s18_io_bufio.go

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

func main() {
	fmt.Println("=== io and bufio ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. io.ReadAll — read everything
	// ─────────────────────────────────────────────
	r := strings.NewReader("Hello, io.ReadAll!")
	data, _ := io.ReadAll(r)
	fmt.Println("ReadAll:", string(data))

	// ─────────────────────────────────────────────
	// 2. io.Copy — stream from reader to writer
	// ─────────────────────────────────────────────
	fmt.Println("\n-- io.Copy --")
	src := strings.NewReader("Copy this data")
	var dst bytes.Buffer
	n, _ := io.Copy(&dst, src)
	fmt.Printf("Copied %d bytes: %q\n", n, dst.String())

	// ─────────────────────────────────────────────
	// 3. io.TeeReader — read and copy simultaneously
	// ─────────────────────────────────────────────
	fmt.Println("\n-- TeeReader --")
	original := strings.NewReader("tee data")
	var copy1 bytes.Buffer
	tee := io.TeeReader(original, &copy1)
	data, _ = io.ReadAll(tee) // read through tee
	fmt.Printf("Read: %q, Copy: %q\n", string(data), copy1.String())

	// ─────────────────────────────────────────────
	// 4. io.MultiReader / io.MultiWriter
	// ─────────────────────────────────────────────
	fmt.Println("\n-- MultiReader --")
	r1 := strings.NewReader("Hello, ")
	r2 := strings.NewReader("World!")
	multi := io.MultiReader(r1, r2)
	combined, _ := io.ReadAll(multi)
	fmt.Println("Combined:", string(combined))

	// ─────────────────────────────────────────────
	// 5. bufio.Scanner — read lines
	// ─────────────────────────────────────────────
	fmt.Println("\n-- bufio.Scanner --")
	input := "line one\nline two\nline three\n"
	scanner := bufio.NewScanner(strings.NewReader(input))
	lineNum := 1
	for scanner.Scan() {
		fmt.Printf("  %d: %s\n", lineNum, scanner.Text())
		lineNum++
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Scanner error:", err)
	}

	// ─────────────────────────────────────────────
	// 6. bufio.Scanner — read words
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Scan words --")
	wordScanner := bufio.NewScanner(strings.NewReader("hello world foo bar"))
	wordScanner.Split(bufio.ScanWords)
	for wordScanner.Scan() {
		fmt.Printf("  word: %q\n", wordScanner.Text())
	}

	// ─────────────────────────────────────────────
	// 7. bufio.Writer — buffered writing
	// ─────────────────────────────────────────────
	fmt.Println("\n-- bufio.Writer --")
	var buf bytes.Buffer
	w := bufio.NewWriter(&buf)
	w.WriteString("buffered ")
	w.WriteString("output")
	w.Flush() // MUST flush to write to underlying writer
	fmt.Println("Buffered:", buf.String())

	// ─────────────────────────────────────────────
	// 8. io.LimitReader — limit bytes read
	// ─────────────────────────────────────────────
	fmt.Println("\n-- LimitReader --")
	unlimited := strings.NewReader("This is a long string that should be limited")
	limited := io.LimitReader(unlimited, 10)
	data, _ = io.ReadAll(limited)
	fmt.Printf("Limited: %q\n", string(data)) // first 10 bytes
}
