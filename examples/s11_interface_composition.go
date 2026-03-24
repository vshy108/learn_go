//go:build ignore

// Section 11, Topic 85: Interface Embedding / Composition
//
// Interfaces can embed other interfaces, combining their method sets.
// This is Go's way of creating larger interfaces from smaller ones.
//
// GOTCHA: Go standard library uses this heavily:
//   io.ReadWriter = io.Reader + io.Writer
//   io.ReadCloser = io.Reader + io.Closer
//
// Run: go run examples/s11_interface_composition.go

package main

import "fmt"

// ─────────────────────────────────────────────
// Small interfaces
// ─────────────────────────────────────────────
type Reader interface {
	Read() string
}

type Writer interface {
	Write(data string)
}

type Closer interface {
	Close() error
}

// ─────────────────────────────────────────────
// Composed interfaces
// ─────────────────────────────────────────────
type ReadWriter interface {
	Reader
	Writer
}

type ReadWriteCloser interface {
	Reader
	Writer
	Closer
}

// ─────────────────────────────────────────────
// A type satisfying ReadWriteCloser
// ─────────────────────────────────────────────
type File struct {
	name   string
	buffer string
	closed bool
}

func (f *File) Read() string {
	if f.closed {
		return ""
	}
	return f.buffer
}

func (f *File) Write(data string) {
	if !f.closed {
		f.buffer += data
	}
}

func (f *File) Close() error {
	f.closed = true
	return nil
}

func main() {
	fmt.Println("=== Interface Composition ===")
	fmt.Println()

	f := &File{name: "test.txt"}

	// ─────────────────────────────────────────────
	// 1. Satisfies all composed interfaces
	// ─────────────────────────────────────────────
	var rw ReadWriter = f
	rw.Write("Hello, ")
	rw.Write("World!")
	fmt.Printf("Read: %q\n", rw.Read())

	// ─────────────────────────────────────────────
	// 2. Narrowing interface
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Interface narrowing --")
	var rwc ReadWriteCloser = f
	var r Reader = rwc // wider → narrower is fine
	fmt.Printf("Read only: %q\n", r.Read())
	// r.Write("test")  // ERROR: Reader has no Write method

	// ─────────────────────────────────────────────
	// 3. Function accepting composed interface
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Function example --")
	f2 := &File{name: "data.txt"}
	processFile(f2)

	// ─────────────────────────────────────────────
	// 4. Stdlib examples of interface composition
	// ─────────────────────────────────────────────
	// type ReadWriter interface {
	//     Reader
	//     Writer
	// }
	// type ReadCloser interface {
	//     Reader
	//     Closer
	// }
	// type ReadWriteCloser interface {
	//     Reader
	//     Writer
	//     Closer
	// }

	_ = rwc
}

func processFile(f ReadWriteCloser) {
	f.Write("processed data")
	fmt.Printf("Data: %q\n", f.Read())
	_ = f.Close()
	fmt.Println("File closed")
}
