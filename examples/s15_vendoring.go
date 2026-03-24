//go:build ignore

// Section 15, Topic 115: Vendoring
//
// Vendoring copies all dependencies into a vendor/ directory in your project.
// This ensures reproducible builds without network access.
//
// Commands:
//   go mod vendor     — create/update vendor/
//   go mod tidy       — clean up go.mod first
//   go build -mod=vendor  — force using vendor/
//
// Since Go 1.14, if vendor/ exists, it's used automatically.
//
// GOTCHA: vendor/ must be committed to version control for reproducibility.
// GOTCHA: go mod vendor overwrites vendor/ completely.
//
// Run: go run examples/s15_vendoring.go

package main

import "fmt"

func main() {
	fmt.Println("=== Vendoring ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. Creating vendor directory
	// ─────────────────────────────────────────────
	fmt.Println("-- Commands --")
	fmt.Println("  go mod tidy     # clean up go.mod/go.sum")
	fmt.Println("  go mod vendor   # copy deps to vendor/")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 2. Vendor directory structure
	// ─────────────────────────────────────────────
	fmt.Println("-- Structure --")
	fmt.Print(`
  vendor/
  ├── modules.txt          ← tracks vendored modules
  ├── github.com/
  │   └── gin-gonic/
  │       └── gin/
  │           ├── gin.go
  │           └── ...
  └── golang.org/
      └── x/
          └── sync/
              └── ...
`)

	// ─────────────────────────────────────────────
	// 3. When to vendor
	// ─────────────────────────────────────────────
	fmt.Println("-- When to vendor --")
	fmt.Println("  ✓ CI/CD without internet access")
	fmt.Println("  ✓ Enterprise environments with proxy restrictions")
	fmt.Println("  ✓ Ensuring build reproducibility")
	fmt.Println("  ✓ Auditing dependencies")
	fmt.Println()
	fmt.Println("  ✗ Small projects with access to module proxy")
	fmt.Println("  ✗ When vendor/ bloats the repository")

	// ─────────────────────────────────────────────
	// 4. Go module proxy
	// ─────────────────────────────────────────────
	// GOPROXY=https://proxy.golang.org,direct (default)
	// Module proxy caches modules — faster and more reliable than git.
	// GONOSUMCHECK and GONOSUMDB for private modules.
}
