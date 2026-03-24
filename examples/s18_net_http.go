//go:build ignore

// Section 18, Topic 132: net/http (HTTP Client and Server)
//
// Go's stdlib includes a production-quality HTTP server and client.
// No framework needed for many use cases.
//
// This example uses httptest.NewServer to create a LOCAL mock server,
// so it runs offline with no external dependencies.
//
// GOTCHA: Always close resp.Body (use defer resp.Body.Close()).
// GOTCHA: http.Get uses the default client with NO timeout -- set one in production.
// GOTCHA: Always set read/write timeouts on production servers.
//
// Run: go run examples/s18_net_http.go

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"time"
)

func main() {
	fmt.Println("=== net/http ===")
	fmt.Println()

	// Create a local mock server (replaces external APIs like httpbin.org)
	mux := http.NewServeMux()

	mux.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"url": "%s", "method": "%s", "host": "%s"}`,
			r.URL.String(), r.Method, r.Host)
	})

	mux.HandleFunc("/ip", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"origin": "127.0.0.1"}`)
	})

	mux.HandleFunc("/api/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode([]map[string]string{
				{"name": "Alice"},
				{"name": "Bob"},
			})
		case http.MethodPost:
			body, _ := io.ReadAll(r.Body)
			w.WriteHeader(http.StatusCreated)
			fmt.Fprintf(w, "Created: %s", body)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/delay", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(200 * time.Millisecond)
		fmt.Fprint(w, "delayed response")
	})

	// httptest.NewServer starts a REAL HTTP server on localhost
	srv := httptest.NewServer(mux)
	defer srv.Close()
	baseURL := srv.URL
	fmt.Println("Mock server at:", baseURL)

	// ─────────────────────────────────────────────
	// 1. HTTP GET
	// ─────────────────────────────────────────────
	fmt.Println("\n-- HTTP GET --")
	resp, err := http.Get(baseURL + "/get")
	if err != nil {
		fmt.Println("GET error:", err)
	} else {
		defer resp.Body.Close()
		fmt.Printf("Status: %s\n", resp.Status)
		body, _ := io.ReadAll(resp.Body)
		fmt.Printf("Body: %s\n", body)
	}

	// ─────────────────────────────────────────────
	// 2. Custom client with timeout
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Custom client --")
	client := &http.Client{Timeout: 5 * time.Second}

	resp2, err := client.Get(baseURL + "/ip")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		defer resp2.Body.Close()
		body, _ := io.ReadAll(resp2.Body)
		fmt.Println("Response:", strings.TrimSpace(string(body)))
	}

	// ─────────────────────────────────────────────
	// 3. POST with body
	// ─────────────────────────────────────────────
	fmt.Println("\n-- HTTP POST --")
	postBody := strings.NewReader(`{"name":"Charlie"}`)
	resp3, err := http.Post(baseURL+"/api/users", "application/json", postBody)
	if err != nil {
		fmt.Println("POST error:", err)
	} else {
		defer resp3.Body.Close()
		fmt.Printf("Status: %d %s\n", resp3.StatusCode, resp3.Status)
		body, _ := io.ReadAll(resp3.Body)
		fmt.Printf("Body: %s\n", body)
	}

	// ─────────────────────────────────────────────
	// 4. GET + JSON decode
	// ─────────────────────────────────────────────
	fmt.Println("\n-- GET + JSON decode --")
	resp4, err := client.Get(baseURL + "/api/users")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		defer resp4.Body.Close()
		var users []map[string]string
		json.NewDecoder(resp4.Body).Decode(&users)
		for _, u := range users {
			fmt.Printf("  User: %s\n", u["name"])
		}
	}

	// ─────────────────────────────────────────────
	// 5. Custom request with headers
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Custom request --")
	req, _ := http.NewRequest("GET", baseURL+"/get", nil)
	req.Header.Set("Authorization", "Bearer my-token")
	req.Header.Set("Accept", "application/json")
	resp5, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		defer resp5.Body.Close()
		body, _ := io.ReadAll(resp5.Body)
		fmt.Printf("Response: %s\n", body)
	}

	// ─────────────────────────────────────────────
	// 6. Timeout demo
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Timeout demo --")
	shortClient := &http.Client{Timeout: 100 * time.Millisecond}
	_, err = shortClient.Get(baseURL + "/delay")
	if err != nil {
		fmt.Println("Expected timeout error:", err)
	}

	// ─────────────────────────────────────────────
	// 7. httptest.NewRecorder (unit test without network)
	// ─────────────────────────────────────────────
	fmt.Println("\n-- httptest.NewRecorder --")
	recorder := httptest.NewRecorder()
	testReq := httptest.NewRequest("GET", "/ip", nil)
	mux.ServeHTTP(recorder, testReq)
	fmt.Printf("Recorder status: %d, body: %s\n",
		recorder.Code, recorder.Body.String())

	// ─────────────────────────────────────────────
	// 8. Summary
	// ─────────────────────────────────────────────
	fmt.Println("\n-- httptest summary --")
	fmt.Println("httptest.NewServer(handler)    -> real localhost server")
	fmt.Println("httptest.NewTLSServer(handler) -> same but with TLS")
	fmt.Println("httptest.NewRecorder()         -> capture response, no network")
	fmt.Println("httptest.NewRequest()          -> create test *http.Request")
	fmt.Println("Always: defer srv.Close()")

	fmt.Println("\n-- Production patterns --")
	fmt.Println("  mux := http.NewServeMux()")
	fmt.Println("  mux.HandleFunc(\"GET /users\", listUsers)       // Go 1.22")
	fmt.Println("  mux.HandleFunc(\"GET /users/{id}\", getUser)    // path params!")
	fmt.Println("  srv := &http.Server{Addr: \":8080\", Handler: mux,")
	fmt.Println("    ReadTimeout: 5*time.Second, WriteTimeout: 10*time.Second}")
	fmt.Println("  log.Fatal(srv.ListenAndServe())")
}
