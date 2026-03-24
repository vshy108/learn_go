//go:build ignore

// Section 15, Topic 112: Go Modules (go.mod, go.sum)
//
// Go modules manage dependencies since Go 1.11 (default since 1.16).
//
// Key commands:
//   go mod init <module-path>  — create new module
//   go mod tidy                — add missing / remove unused dependencies
//   go get <package>@<version> — add or update dependency
//   go mod download            — download all dependencies
//   go mod verify              — verify dependency integrity
//   go mod vendor              — copy deps to vendor/
//   go mod graph               — print dependency graph
//
// Files:
//   go.mod — module path, Go version, dependencies
//   go.sum — cryptographic hashes for dependency integrity
//
// GOTCHA: Always commit go.sum to version control.
// GOTCHA: go mod tidy should be run before committing.
//
// Run: go run examples/s15_go_modules.go

package main

import "fmt"

func main() {
	fmt.Println("=== Go Modules ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. go.mod file structure
	// ─────────────────────────────────────────────
	fmt.Println("-- go.mod --")
	fmt.Println(`
module github.com/user/myproject

go 1.22

require (
    github.com/gin-gonic/gin v1.9.1
    golang.org/x/sync v0.6.0
)

require (
    // indirect dependencies (auto-managed)
    github.com/some/dep v1.0.0 // indirect
)
`)

	// ─────────────────────────────────────────────
	// 2. Semantic versioning
	// ─────────────────────────────────────────────
	fmt.Println("-- Semantic versioning --")
	// v1.2.3 → Major.Minor.Patch
	// Major: breaking changes
	// Minor: new features (backwards compatible)
	// Patch: bug fixes
	//
	// go get example.com/pkg@v1.2.3   — exact version
	// go get example.com/pkg@latest   — latest version
	// go get example.com/pkg@v1       — latest v1.x.x

	// ─────────────────────────────────────────────
	// 3. Major version suffixes
	// ─────────────────────────────────────────────
	// v0.x.x and v1.x.x: import "example.com/pkg"
	// v2.x.x: import "example.com/pkg/v2"
	// This allows v1 and v2 to coexist!

	// ─────────────────────────────────────────────
	// 4. Common workflows
	// ─────────────────────────────────────────────
	fmt.Println("-- Common workflows --")
	fmt.Println("  New project:  go mod init github.com/user/project")
	fmt.Println("  Add dep:      go get github.com/gin-gonic/gin")
	fmt.Println("  Update dep:   go get -u github.com/gin-gonic/gin")
	fmt.Println("  Update all:   go get -u ./...")
	fmt.Println("  Clean up:     go mod tidy")
	fmt.Println("  List deps:    go list -m all")

	// ─────────────────────────────────────────────
	// 5. Replace and exclude
	// ─────────────────────────────────────────────
	// In go.mod:
	// replace github.com/broken/pkg => github.com/fixed/pkg v1.0.0
	// replace github.com/pkg => ../local/pkg  // local development
	// exclude github.com/bad/pkg v1.0.0
}
