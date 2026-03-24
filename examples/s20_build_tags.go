//go:build ignore

// Section 20, Topic 147: Build Tags & Conditional Compilation
//
// Build tags control which files are included in a build.
// Go 1.17+ uses //go:build syntax (replaces // +build).
//
// Common uses:
//   - Platform-specific code (linux, windows, darwin)
//   - Feature flags
//   - Test/debug builds
//   - Excluding files from normal builds (//go:build ignore)
//
// GOTCHA: //go:build must be FIRST line in the file (before package).
// GOTCHA: Old syntax (// +build) is still supported but deprecated.
//
// Run: go run examples/s20_build_tags.go

package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("=== Build Tags & Conditional Compilation ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. Runtime platform detection
	// ─────────────────────────────────────────────
	fmt.Println("-- Runtime info --")
	fmt.Println("OS:      ", runtime.GOOS)      // linux, darwin, windows
	fmt.Println("Arch:    ", runtime.GOARCH)     // amd64, arm64
	fmt.Println("Compiler:", runtime.Compiler)   // gc
	fmt.Println("Version: ", runtime.Version())  // go1.22.x

	// ─────────────────────────────────────────────
	// 2. Build tag syntax
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Build tag syntax (//go:build) --")
	fmt.Println(`
// Single constraint:
//go:build linux

// OR:
//go:build linux || darwin

// AND:
//go:build linux && amd64

// NOT:
//go:build !windows

// Complex:
//go:build (linux || darwin) && amd64

// Ignore file:
//go:build ignore
`)

	// ─────────────────────────────────────────────
	// 3. File naming convention (alternative to tags)
	// ─────────────────────────────────────────────
	fmt.Println("-- File naming convention --")
	fmt.Println(`
Go auto-includes files based on name suffixes:
  file_linux.go      → only on Linux
  file_windows.go    → only on Windows
  file_darwin.go     → only on macOS
  file_amd64.go      → only on amd64
  file_linux_amd64.go → Linux AND amd64

This is equivalent to putting a build tag at the top.
`)

	// ─────────────────────────────────────────────
	// 4. Custom build tags
	// ─────────────────────────────────────────────
	fmt.Println("-- Custom build tags --")
	fmt.Println(`
// In file: feature_flag.go
//go:build feature_x

// Build with:
// go build -tags feature_x

// Multiple tags:
// go build -tags "feature_x,debug"
`)

	// ─────────────────────────────────────────────
	// 5. Practical example: debug logging
	// ─────────────────────────────────────────────
	fmt.Println("-- Debug build pattern --")
	fmt.Println(`
// debug.go:
//go:build debug

package myapp
func debugLog(msg string) { log.Println("[DEBUG]", msg) }

// release.go:
//go:build !debug

package myapp
func debugLog(msg string) {} // no-op

// Build:
// go build -tags debug    → includes debug logging
// go build                → no debug logging
`)

	// ─────────────────────────────────────────────
	// 6. Why this file uses //go:build ignore
	// ─────────────────────────────────────────────
	fmt.Println("-- Why //go:build ignore --")
	fmt.Println(`
All example files in this project use:
  //go:build ignore

This prevents 'go build' from including them
(each file has its own main function, which would conflict).

To run an individual example:
  go run examples/s20_build_tags.go

'go run' compiles the specified file directly,
ignoring the build tag.
`)
}
