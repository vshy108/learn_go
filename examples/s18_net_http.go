//go:build ignore

// Section 18, Topic 132: net/http (HTTP Client & Server)
//
// Go's stdlib includes a production-quality HTTP server and client.
// No framework needed for many use cases.
//
// Run: go run examples/s18_net_http.go

package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

func main() {
	fmt.Println("=== net/http ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// 1. HTTP Client — GET request
	// ─────────────────────────────────────────────
	fmt.Println("-- HTTP Client --")

	// Simple GET:
	resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		fmt.Println("GET error:", err)
	} else {
		defer resp.Body.Close()
		fmt.Printf("Status: %s\n", resp.Status)

		// Read first 200 bytes:
		body := make([]byte, 200)
		n, _ := resp.Body.Read(body)
		fmt.Printf("Body (first %d bytes): %s...\n", n, string(body[:n]))
	}

	// ─────────────────────────────────────────────
	// 2. HTTP Client with timeout
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Custom client --")
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	resp2, err := client.Get("https://httpbin.org/ip")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		defer resp2.Body.Close()
		body, _ := io.ReadAll(resp2.Body)
		fmt.Println("IP response:", strings.TrimSpace(string(body)))
	}

	// ─────────────────────────────────────────────
	// 3. HTTP Server (example code — not started here)
	// ─────────────────────────────────────────────
	fmt.Println("\n-- HTTP Server (code example) --")
	fmt.Println(`
// Basic HTTP server:
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
})

http.HandleFunc("/api/users", func(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        w.Header().Set("Content-Type", "application/json")
        fmt.Fprint(w, ` + "`" + `[{"name":"Alice"},{"name":"Bob"}]` + "`" + `)
    case http.MethodPost:
        // Read body: body, _ := io.ReadAll(r.Body)
        w.WriteHeader(http.StatusCreated)
    default:
        http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
    }
})

log.Fatal(http.ListenAndServe(":8080", nil))
`)

	// ─────────────────────────────────────────────
	// 4. http.ServeMux (router)
	// ─────────────────────────────────────────────
	fmt.Println("-- ServeMux (Go 1.22 enhanced) --")
	fmt.Println(`
mux := http.NewServeMux()
mux.HandleFunc("GET /users", listUsers)
mux.HandleFunc("POST /users", createUser)
mux.HandleFunc("GET /users/{id}", getUser)  // Go 1.22 path params!
http.ListenAndServe(":8080", mux)
`)

	// ─────────────────────────────────────────────
	// GOTCHA: Always set timeouts in production
	// ─────────────────────────────────────────────
	fmt.Println("-- Production server --")
	fmt.Println(`
srv := &http.Server{
    Addr:         ":8080",
    Handler:      mux,
    ReadTimeout:  5 * time.Second,
    WriteTimeout: 10 * time.Second,
    IdleTimeout:  120 * time.Second,
}
log.Fatal(srv.ListenAndServe())
`)
}
