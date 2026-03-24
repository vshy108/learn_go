//go:build ignore

// Section 11, Topic 88: io.Reader and io.Writer
//
// The two most important interfaces in Go:
//   type Reader interface { Read(p []byte) (n int, err error) }
//   type Writer interface { Write(p []byte) (n int, err error) }
//
// Everything connects through these: files, network, buffers, HTTP, compression.
//
// Run: go run examples/s11_io_reader_writer.go

package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fmt.Println("=== io.Reader and io.Writer ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. strings.Reader → io.Reader
	// ─────────────────────────────────────────────
	r := strings.NewReader("Hello, io.Reader!")
	buf := make([]byte, 8)
	for {
		n, err := r.Read(buf)
		if n > 0 {
			fmt.Printf("Read %d bytes: %q\n", n, buf[:n])
		}
		if err == io.EOF {
			break
		}
	}

	// ─────────────────────────────────────────────
	// 2. bytes.Buffer → Reader and Writer
	// ─────────────────────────────────────────────
	fmt.Println("\n-- bytes.Buffer --")
	var b bytes.Buffer
	b.WriteString("Hello, ")
	b.WriteString("Buffer!")
	fmt.Println("Buffer:", b.String())

	// Read from buffer:
	out := make([]byte, 5)
	n, _ := b.Read(out)
	fmt.Printf("Read: %q, remaining: %q\n", out[:n], b.String())

	// ─────────────────────────────────────────────
	// 3. io.Copy — connect reader to writer
	// ─────────────────────────────────────────────
	fmt.Println("\n-- io.Copy --")
	src := strings.NewReader("Copy me to stdout!\n")
	written, _ := io.Copy(os.Stdout, src)
	fmt.Printf("(wrote %d bytes)\n", written)

	// ─────────────────────────────────────────────
	// 4. Custom Reader
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Custom Reader --")
	zr := &ZeroReader{}
	p := make([]byte, 10)
	n, _ = zr.Read(p)
	fmt.Printf("ZeroReader: %v (read %d bytes of zeros)\n", p[:n], n)

	// ─────────────────────────────────────────────
	// 5. Custom Writer
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Custom Writer --")
	cw := &CountWriter{}
	fmt.Fprintf(cw, "Hello, %s!", "world")
	fmt.Printf("CountWriter: %d bytes written\n", cw.Total)

	// ─────────────────────────────────────────────
	// Why Reader/Writer are important:
	// ─────────────────────────────────────────────
	// - os.File implements both
	// - net.Conn implements both
	// - http.Response.Body is io.ReadCloser
	// - json.NewDecoder takes io.Reader
	// - gzip.NewWriter takes io.Writer
	// → Composable, testable, decoupled code
}

// Reads zeros forever
type ZeroReader struct{}

func (z *ZeroReader) Read(p []byte) (int, error) {
	for i := range p {
		p[i] = 0
	}
	return len(p), nil
}

// Counts bytes written
type CountWriter struct {
	Total int
}

func (cw *CountWriter) Write(p []byte) (int, error) {
	cw.Total += len(p)
	return len(p), nil
}
