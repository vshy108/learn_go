# Interview Prep: Error Handling

> Focus: Error handling in Go, API error responses

---

## Q1: How does error handling work in Go? How is it different from exceptions?

**Answer:**

Go has **no exceptions**. Errors are **values** — functions return an `error` as the last return value, and the caller must explicitly check it.

```go
file, err := os.Open("config.yaml")
if err != nil {
    return fmt.Errorf("open config: %w", err)
}
defer file.Close()
```

**Comparison with exceptions (Java/Python/TypeScript):**

| Aspect | Go errors | Exceptions |
|--------|-----------|------------|
| Flow | Explicit — checked at call site | Implicit — bubbles up the stack |
| Visibility | Every error path is visible in code | Hidden — you don't know what throws |
| Performance | Cheap (just a value) | Expensive (stack unwinding) |
| Can be ignored? | Yes (assign to `_`) but linters catch it | Yes (no catch block) |
| Control flow | Linear — `if err != nil` | Non-linear — jumps to nearest catch |

**Why Go chose this:**
- Forces you to think about every failure point
- No hidden control flow — you read top to bottom
- Errors are first-class data you can inspect, wrap, and pass around

---

## Q2: What is error wrapping and why should you use it?

**Answer:**

Error wrapping adds context as an error propagates up the call stack, using `fmt.Errorf` with the `%w` verb:

```go
func GetVessel(ctx context.Context, id string) (*Vessel, error) {
    vessel, err := repo.FindByID(ctx, id)
    if err != nil {
        return nil, fmt.Errorf("get vessel %s: %w", id, err)
    }
    return vessel, nil
}
```

This produces an error chain: `get vessel abc-123: query row: connection refused`

**Why wrap?**
- **Debugging** — you see the full call path, not just "connection refused"
- **Unwrapping** — callers can still check the original error:

```go
if errors.Is(err, sql.ErrNoRows) {
    // handle not found
}
```

**`%w` vs `%v`:**
- `%w` — wraps the error (preserves the chain, can be unwrapped with `errors.Is`/`errors.As`)
- `%v` — formats it as a string (breaks the chain, original error lost)

**Rule:** Always use `%w` unless you intentionally want to hide the underlying error from callers.

---

## Q3: Explain `errors.Is()` and `errors.As()`.

**Answer:**

### `errors.Is(err, target)` — checks if any error in the chain matches a **value**

```go
// Sentinel error
var ErrNotFound = errors.New("not found")

func GetVessel(id string) (*Vessel, error) {
    v, err := repo.Find(id)
    if err != nil {
        return nil, fmt.Errorf("get vessel: %w", err) // wraps ErrNotFound
    }
    return v, nil
}

// Caller:
vessel, err := GetVessel("abc")
if errors.Is(err, ErrNotFound) {
    // true! errors.Is walks the full chain
    http.Error(w, "not found", 404)
}
```

### `errors.As(err, &target)` — checks if any error in the chain matches a **type**

```go
type ValidationError struct {
    Field   string
    Message string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("validation: %s %s", e.Field, e.Message)
}

// Caller:
var valErr *ValidationError
if errors.As(err, &valErr) {
    // extract structured info
    fmt.Println(valErr.Field, valErr.Message)
}
```

**When to use which:**
| Use | When |
|-----|------|
| `errors.Is` | Comparing against a known sentinel value (`sql.ErrNoRows`, `io.EOF`, custom `ErrNotFound`) |
| `errors.As` | Extracting a typed error to read its fields |

**Never do:** `if err.Error() == "not found"` — fragile string comparison that breaks with wrapping.

---

## Q4: How do you define custom error types in Go?

**Answer:**

### Sentinel errors (simple, no extra data):
```go
var (
    ErrNotFound     = errors.New("not found")
    ErrUnauthorized = errors.New("unauthorized")
    ErrConflict     = errors.New("conflict")
)
```

### Typed errors (carry structured data):
```go
type AppError struct {
    Code    int    // HTTP status code
    Message string // user-facing message
    Err     error  // underlying error (for logging)
}

func (e *AppError) Error() string {
    return e.Message
}

func (e *AppError) Unwrap() error {
    return e.Err
}

// Constructor helpers
func NewNotFound(resource string, err error) *AppError {
    return &AppError{
        Code:    http.StatusNotFound,
        Message: fmt.Sprintf("%s not found", resource),
        Err:     err,
    }
}

func NewBadRequest(message string) *AppError {
    return &AppError{
        Code:    http.StatusBadRequest,
        Message: message,
    }
}

func NewInternal(err error) *AppError {
    return &AppError{
        Code:    http.StatusInternalServerError,
        Message: "internal server error", // never expose internals to users
        Err:     err,
    }
}
```

**Usage in handler:**
```go
func (h *VesselHandler) GetByID(w http.ResponseWriter, r *http.Request) {
    vessel, err := h.service.GetVessel(r.Context(), r.PathValue("id"))
    if err != nil {
        var appErr *AppError
        if errors.As(err, &appErr) {
            writeError(w, appErr.Code, appErr.Message)
        } else {
            writeError(w, 500, "internal server error")
        }
        return
    }
    writeJSON(w, 200, vessel)
}
```

This pattern gives you:
- **Consistent API error responses** — the handler maps `AppError.Code` to HTTP status
- **Separation of concerns** — the service layer decides the error type, the handler maps it to HTTP
- **Safe logging** — log `appErr.Err` server-side, return `appErr.Message` to the client

---

## Q5: How do you handle errors in HTTP middleware?

**Answer:**

A common pattern is an **error-handling middleware** that catches `AppError` from handlers:

```go
// Define a handler type that returns an error
type AppHandler func(w http.ResponseWriter, r *http.Request) error

// Middleware that converts errors to HTTP responses
func ErrorHandler(h AppHandler) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        err := h(w, r)
        if err == nil {
            return
        }

        // Log the full error server-side
        log.Printf("ERROR: %s %s: %v", r.Method, r.URL.Path, err)

        var appErr *AppError
        if errors.As(err, &appErr) {
            writeError(w, appErr.Code, appErr.Message)
            return
        }

        // Unknown error → 500
        writeError(w, http.StatusInternalServerError, "internal server error")
    }
}

// Usage:
mux.HandleFunc("GET /api/v1/vessels/{id}", ErrorHandler(vesselHandler.GetByID))
```

Now the handler is cleaner — it just returns errors:
```go
func (h *VesselHandler) GetByID(w http.ResponseWriter, r *http.Request) error {
    id := r.PathValue("id")
    vessel, err := h.service.GetVessel(r.Context(), id)
    if err != nil {
        return err // middleware handles the mapping
    }
    writeJSON(w, 200, vessel)
    return nil
}
```

**Benefits:**
- All error-to-HTTP mapping lives in one place
- Handlers stay focused on happy path
- Consistent logging and response format

---

## Q6: What is `panic` and `recover` in Go? When should you use them?

**Answer:**

### `panic` — causes the program to crash (unwind the stack and exit)
```go
panic("something went terribly wrong")
```

### `recover` — catches a panic, must be called inside a `defer`
```go
func safeHandler(w http.ResponseWriter, r *http.Request) {
    defer func() {
        if r := recover(); r != nil {
            log.Printf("PANIC recovered: %v\n%s", r, debug.Stack())
            http.Error(w, "internal server error", 500)
        }
    }()

    // ... handler code that might panic
}
```

### When to use:

| Do | Don't |
|----|-------|
| **Recover** in HTTP middleware to prevent one bad request from crashing the server | **Panic** for expected errors (file not found, invalid input, DB error) |
| **Panic** for truly unrecoverable programmer bugs (e.g., nil map that should have been initialized) | **Panic** as flow control (like using exceptions) |
| **Panic** in `init()` if the app can't start (missing required config) | **Panic** in library code (return errors instead) |

**In practice:** 99% of the time, return errors. Panic only for things that should never happen if the code is correct. Always have a recovery middleware in HTTP servers.

---

## Q7: How do you handle validation errors and return them to the client?

**Answer:**

I create a structured validation error that collects multiple field errors:

```go
type FieldError struct {
    Field   string `json:"field"`
    Message string `json:"message"`
}

type ValidationError struct {
    Fields []FieldError
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("validation failed: %d errors", len(e.Fields))
}

func ValidateCreateVessel(req CreateVesselRequest) error {
    var errs []FieldError

    if req.Name == "" {
        errs = append(errs, FieldError{Field: "name", Message: "is required"})
    }
    if req.IMO == "" {
        errs = append(errs, FieldError{Field: "imo", Message: "is required"})
    } else if len(req.IMO) != 7 {
        errs = append(errs, FieldError{Field: "imo", Message: "must be exactly 7 characters"})
    }
    if req.Status != "" && req.Status != "active" && req.Status != "inactive" {
        errs = append(errs, FieldError{Field: "status", Message: "must be 'active' or 'inactive'"})
    }

    if len(errs) > 0 {
        return &ValidationError{Fields: errs}
    }
    return nil
}
```

**API response (400 Bad Request):**
```json
{
  "error": {
    "code": "VALIDATION_ERROR",
    "message": "validation failed: 2 errors",
    "details": [
      { "field": "name", "message": "is required" },
      { "field": "imo", "message": "must be exactly 7 characters" }
    ]
  }
}
```

**Why return all errors at once?** Instead of returning after the first error, collect them all so the caller can fix everything in one go. This is much better UX for any frontend consumer.

---

## Q8: How do you handle errors with database operations gracefully?

**Answer:**

Map database errors to domain errors in the repository layer:

```go
import (
    "database/sql"
    "errors"
    "github.com/lib/pq"
)

func (r *VesselRepository) Create(ctx context.Context, req CreateVesselRequest) (*Vessel, error) {
    var v Vessel
    err := r.db.QueryRowContext(ctx, query, ...).Scan(...)
    if err != nil {
        // Check for specific PostgreSQL errors
        var pgErr *pq.Error
        if errors.As(err, &pgErr) {
            switch pgErr.Code {
            case "23505": // unique_violation
                return nil, &AppError{
                    Code:    409,
                    Message: fmt.Sprintf("vessel with IMO %s already exists", req.IMO),
                    Err:     err,
                }
            case "23503": // foreign_key_violation
                return nil, &AppError{
                    Code:    400,
                    Message: "referenced fleet does not exist",
                    Err:     err,
                }
            }
        }
        // Unknown DB error → internal server error
        return nil, fmt.Errorf("create vessel: %w", err)
    }
    return &v, nil
}
```

**PostgreSQL error codes to know:**
| Code | Name | Meaning |
|------|------|---------|
| `23505` | `unique_violation` | Duplicate key (e.g., duplicate IMO) |
| `23503` | `foreign_key_violation` | Referenced row doesn't exist |
| `23502` | `not_null_violation` | Required field is NULL |
| `23514` | `check_violation` | CHECK constraint failed |
| `40001` | `serialization_failure` | Transaction conflict (retry) |

**Principle:** Never leak database errors to the client. `pq: duplicate key value violates unique constraint "vessels_imo_key"` is an implementation detail. Return `"vessel with IMO 1234567 already exists"` instead.

---

## Q9: How do you handle timeouts and context cancellation?

**Answer:**

Go's `context.Context` propagates deadlines and cancellations through the call chain:

```go
// Set a timeout on the API request
func (h *VesselHandler) GetByID(w http.ResponseWriter, r *http.Request) error {
    ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
    defer cancel()

    vessel, err := h.service.GetVessel(ctx, r.PathValue("id"))
    if err != nil {
        if errors.Is(err, context.DeadlineExceeded) {
            return &AppError{Code: 504, Message: "request timed out"}
        }
        if errors.Is(err, context.Canceled) {
            return nil // client disconnected, nothing to return
        }
        return err
    }
    writeJSON(w, 200, vessel)
    return nil
}
```

**How it works:**
1. HTTP server creates a context per request (already done by `net/http`)
2. Handler can add a timeout with `context.WithTimeout`
3. Context is passed to service → repository → database driver
4. If the database query takes too long, the context expires and the query is cancelled
5. If the client disconnects, `r.Context()` is cancelled automatically

**Why it matters:**
- Prevents slow queries from holding connections forever
- Frees resources when the client has already given up
- Cascades through the full call chain — one timeout protects everything

---

## Q10: How do you log errors properly?

**Answer:**

Use **structured logging** (not `fmt.Println`):

```go
import "log/slog"

// Set up structured logger
logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
    Level: slog.LevelInfo,
}))

// In the error handler middleware:
func ErrorHandler(h AppHandler) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        err := h(w, r)
        if err == nil {
            return
        }

        var appErr *AppError
        if errors.As(err, &appErr) {
            if appErr.Code >= 500 {
                // Server errors → log at ERROR with full details
                slog.Error("request failed",
                    "method", r.Method,
                    "path", r.URL.Path,
                    "status", appErr.Code,
                    "error", appErr.Err, // internal error for debugging
                    "request_id", r.Header.Get("X-Request-ID"),
                )
            } else {
                // Client errors → log at WARN (less noise)
                slog.Warn("client error",
                    "method", r.Method,
                    "path", r.URL.Path,
                    "status", appErr.Code,
                    "message", appErr.Message,
                )
            }
            writeError(w, appErr.Code, appErr.Message)
            return
        }

        // Unknown error
        slog.Error("unhandled error",
            "method", r.Method,
            "path", r.URL.Path,
            "error", err,
        )
        writeError(w, 500, "internal server error")
    }
}
```

**Output (JSON, ready for log aggregation):**
```json
{
  "time": "2026-03-28T10:30:00Z",
  "level": "ERROR",
  "msg": "request failed",
  "method": "GET",
  "path": "/api/v1/vessels/abc-123",
  "status": 500,
  "error": "get vessel: query row: connection refused",
  "request_id": "req-550e8400"
}
```

**Best practices:**
- Use `slog` (Go 1.21+ stdlib) — structured, leveled, JSON output
- **5xx errors → ERROR level** (alerts, needs attention)
- **4xx errors → WARN level** (client's fault, but track for patterns)
- Include `request_id` for tracing across services
- **Never log sensitive data** (tokens, passwords, PII)
- JSON logs integrate directly with log aggregation tools (Elastic Search/Kibana, Loki, etc.)

---

## Q11: What is the difference between recoverable and unrecoverable errors?

**Answer:**

| Type | Examples | How to handle |
|------|----------|---------------|
| **Recoverable** | Network timeout, DB connection lost, invalid user input, rate limit hit | Return error, retry, or tell the user |
| **Unrecoverable** | Out of memory, corrupted data, programmer bug (nil pointer), missing required config at startup | Panic/crash, fix the code, alert ops |

**In a web service:**
- Recoverable: return appropriate HTTP status (400, 404, 429, 503) and let the client retry
- Unrecoverable: the recovery middleware catches the panic, returns 500, and logs the stack trace

**Retry pattern for transient failures:**
```go
func withRetry(ctx context.Context, maxAttempts int, fn func() error) error {
    var err error
    for attempt := 1; attempt <= maxAttempts; attempt++ {
        err = fn()
        if err == nil {
            return nil
        }
        if !isTransient(err) {
            return err // don't retry non-transient errors
        }
        // Exponential backoff: 100ms, 200ms, 400ms...
        backoff := time.Duration(attempt) * 100 * time.Millisecond
        select {
        case <-time.After(backoff):
        case <-ctx.Done():
            return ctx.Err()
        }
    }
    return fmt.Errorf("after %d attempts: %w", maxAttempts, err)
}

func isTransient(err error) bool {
    return errors.Is(err, context.DeadlineExceeded) ||
        errors.Is(err, syscall.ECONNREFUSED) ||
        errors.Is(err, syscall.ECONNRESET)
}
```

---

## Quick Reference: Go Error Handling Patterns

```
┌─────────────────────────────────────────────────────────────┐
│                    Go Error Handling Cheatsheet              │
├─────────────────────────────────────────────────────────────┤
│                                                             │
│  1. Always check errors:                                    │
│     if err != nil { return fmt.Errorf("ctx: %w", err) }    │
│                                                             │
│  2. Wrap with context:                                      │
│     fmt.Errorf("fetch user %s: %w", id, err)               │
│                                                             │
│  3. Check sentinel:                                         │
│     errors.Is(err, sql.ErrNoRows)                           │
│                                                             │
│  4. Check type:                                             │
│     var appErr *AppError                                    │
│     errors.As(err, &appErr)                                 │
│                                                             │
│  5. Never:                                                  │
│     - Compare err.Error() strings                           │
│     - Panic for expected errors                             │
│     - Ignore errors (assign to _)                           │
│     - Expose internal errors to API clients                 │
│                                                             │
│  6. Always:                                                 │
│     - Use %w for wrapping (not %v)                          │
│     - Log server errors with request context                │
│     - Return all validation errors at once                  │
│     - Map DB errors to domain errors                        │
└─────────────────────────────────────────────────────────────┘
```
