//go:build ignore

// Section 1, Topic 5: go run, go build, go install
//
// Go provides a toolchain for building and running programs:
//   - go run:     compile + execute in one step (temp binary, for development)
//   - go build:   compile to a binary in the current directory
//   - go install:  compile and install to $GOPATH/bin (or $GOBIN)
//
// GOTCHA: `go run` creates a temp binary and deletes it after execution.
//         It's NOT suitable for production — always use `go build` for deployments.
//
// GOTCHA: `go run` only works with `package main` files.
//         `go run *.go` on a directory with multiple package main files
//         will fail if they have conflicting main() functions.
//
// Run: go run examples/s01_go_run_build.go

package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	fmt.Println("=== go run, go build, go install ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. go run — compile and run in one step
	// ─────────────────────────────────────────────
	// Usage:
	//   go run main.go           # single file
	//   go run main.go utils.go  # multiple files in same package
	//   go run .                 # all .go files in current directory
	//
	// Internally, go run:
	//   1. Compiles the file(s) to a temp directory
	//   2. Executes the resulting binary
	//   3. Cleans up the temp binary
	//
	// EDGE CASE: If your program creates files relative to the binary path,
	// `go run` will create them in a temp directory, not your working directory.
	// Use os.Getwd() for the working directory.

	// ─────────────────────────────────────────────
	// 2. go build — compile to binary
	// ─────────────────────────────────────────────
	// Usage:
	//   go build                 # builds package in current dir → binary named after dir
	//   go build -o myapp        # custom output name
	//   go build ./cmd/server    # build a specific sub-package
	//
	// Cross-compilation (one of Go's superpowers!):
	//   GOOS=linux GOARCH=amd64 go build -o myapp-linux
	//   GOOS=windows GOARCH=amd64 go build -o myapp.exe
	//   GOOS=darwin GOARCH=arm64 go build -o myapp-mac
	//
	// The resulting binary is STATICALLY LINKED by default — no dependencies needed!
	// This makes Go binaries extremely portable (just copy and run).

	// ─────────────────────────────────────────────
	// 3. go install — build and install
	// ─────────────────────────────────────────────
	// Usage:
	//   go install                              # install current module
	//   go install golang.org/x/tools/cmd/goimports@latest  # install a tool
	//
	// Installs to:
	//   $GOBIN (if set)
	//   $GOPATH/bin (default: ~/go/bin)
	//
	// Make sure $GOPATH/bin is in your $PATH!

	// ─────────────────────────────────────────────
	// 4. Other useful go commands
	// ─────────────────────────────────────────────
	// go vet ./...        — static analysis (catches common bugs)
	// go fmt ./...        — format all Go files (canonical style)
	// go test ./...       — run all tests
	// go mod tidy         — clean up go.mod and go.sum
	// go mod download     — download dependencies
	// go doc fmt.Println  — view documentation
	// go env              — show Go environment variables
	// go generate ./...   — run //go:generate directives
	// go clean            — remove build artifacts

	// ─────────────────────────────────────────────
	// 5. Runtime information
	// ─────────────────────────────────────────────
	fmt.Println("Go version:", runtime.Version())            // e.g., go1.22.0
	fmt.Println("OS/Arch:", runtime.GOOS+"/"+runtime.GOARCH) // e.g., darwin/arm64
	fmt.Println("Num CPUs:", runtime.NumCPU())

	// Working directory (always reliable, unlike binary path with `go run`)
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting working dir:", err)
	} else {
		fmt.Println("Working directory:", wd)
	}

	// Executable path (temp path with `go run`, real path with `go build`)
	exe, err := os.Executable()
	if err != nil {
		fmt.Println("Error getting executable:", err)
	} else {
		fmt.Println("Executable path:", exe)
	}

	fmt.Println("\n--- Build flags cheat sheet ---")
	fmt.Println("go build -v          # verbose (show packages being compiled)")
	fmt.Println("go build -x          # print commands being executed")
	fmt.Println("go build -race       # enable race detector")
	fmt.Println("go build -ldflags '-s -w'  # strip debug info (smaller binary)")
	fmt.Println("go build -gcflags '-m'     # show compiler optimizations")
	fmt.Println("go build -tags integration # include files with build tag")
}
