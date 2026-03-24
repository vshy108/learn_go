//go:build ignore

// Section 19, Topic 142: Middleware Pattern
//
// Middleware wraps an HTTP handler to add cross-cutting concerns:
// logging, auth, rate limiting, CORS, etc.
//
// Pattern:
//   func Middleware(next http.Handler) http.Handler {
//       return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//           // before
//           next.ServeHTTP(w, r)
//           // after
//       })
//   }
//
// Run: go run examples/s19_middleware_pattern.go

package main

import (
	"fmt"
	"net/http"
	"time"
)

// ─────────────────────────────────────────────
// Middleware type
// ─────────────────────────────────────────────
type Middleware func(http.Handler) http.Handler

// ─────────────────────────────────────────────
// Logging middleware
// ─────────────────────────────────────────────
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		fmt.Printf("[LOG] %s %s %v\n", r.Method, r.URL.Path, time.Since(start))
	})
}

// ─────────────────────────────────────────────
// Auth middleware
// ─────────────────────────────────────────────
func RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// ─────────────────────────────────────────────
// CORS middleware
// ─────────────────────────────────────────────
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// ─────────────────────────────────────────────
// Chain multiple middleware
// ─────────────────────────────────────────────
func Chain(handler http.Handler, middlewares ...Middleware) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}
	return handler
}

func main() {
	fmt.Println("=== Middleware Pattern ===")
	fmt.Println()

	// Handler:
	hello := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	// Manual wrapping:
	// handler := Logging(RequireAuth(CORS(hello)))

	// Or with Chain helper:
	handler := Chain(hello, Logging, CORS)

	fmt.Println("Middleware chain created:")
	fmt.Println("  Request → Logging → CORS → Handler → Response")

	// Demonstrate (normally you'd use http.ListenAndServe):
	_ = handler

	fmt.Println(`
// Usage:
mux := http.NewServeMux()
mux.HandleFunc("GET /api/public", publicHandler)

// Apply middleware:
protectedHandler := Chain(
    http.HandlerFunc(secretHandler),
    Logging,
    RequireAuth,
    CORS,
)
mux.Handle("GET /api/secret", protectedHandler)

http.ListenAndServe(":8080", Logging(mux))
`)
}
