# Interview Prep: CRUD Service

> Focus: Building CRUD services in Go, database interactions, full implementation patterns

---

## Q1: Walk me through how you'd build a CRUD service in Go.

**Answer:**

I'd structure the service in layers:

```
cmd/
  server/
    main.go              ← entry point, wires everything together
internal/
  handler/
    vessel_handler.go    ← HTTP handlers (parse request, call service, write response)
  service/
    vessel_service.go    ← business logic
  repository/
    vessel_repo.go       ← database access (SQL queries)
  model/
    vessel.go            ← domain types (structs)
  middleware/
    auth.go              ← authentication, logging, CORS
```

**Why this layering?**
- **Handler** → deals with HTTP concerns (reading params, writing JSON)
- **Service** → contains business rules (validation, authorization logic)
- **Repository** → talks to the database only
- Each layer depends only on the one below it. You can test each in isolation.

**Go example — full CRUD flow for a "Vessel" resource:**

### Model:
```go
package model

import "time"

type Vessel struct {
    ID        string    `json:"id" db:"id"`
    Name      string    `json:"name" db:"name"`
    IMO       string    `json:"imo" db:"imo"`
    Status    string    `json:"status" db:"status"`
    CreatedAt time.Time `json:"createdAt" db:"created_at"`
    UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
}

type CreateVesselRequest struct {
    Name   string `json:"name"`
    IMO    string `json:"imo"`
    Status string `json:"status"`
}

type UpdateVesselRequest struct {
    Name   *string `json:"name,omitempty"`
    Status *string `json:"status,omitempty"`
}
```

Note: `*string` in the update request lets us distinguish between "field not sent" (nil) and "field explicitly set to empty string".

---

## Q2: Show me the repository layer — how do you interact with PostgreSQL in Go?

**Answer:**

```go
package repository

import (
    "context"
    "database/sql"
    "fmt"

    "github.com/google/uuid"
    "myapp/internal/model"
)

type VesselRepository struct {
    db *sql.DB
}

func NewVesselRepository(db *sql.DB) *VesselRepository {
    return &VesselRepository{db: db}
}

// Create inserts a new vessel and returns the created record
func (r *VesselRepository) Create(ctx context.Context, req model.CreateVesselRequest) (*model.Vessel, error) {
    id := uuid.New().String()
    query := `
        INSERT INTO vessels (id, name, imo, status, created_at, updated_at)
        VALUES ($1, $2, $3, $4, NOW(), NOW())
        RETURNING id, name, imo, status, created_at, updated_at
    `
    var v model.Vessel
    err := r.db.QueryRowContext(ctx, query, id, req.Name, req.IMO, req.Status).
        Scan(&v.ID, &v.Name, &v.IMO, &v.Status, &v.CreatedAt, &v.UpdatedAt)
    if err != nil {
        return nil, fmt.Errorf("create vessel: %w", err)
    }
    return &v, nil
}

// GetByID retrieves a single vessel
func (r *VesselRepository) GetByID(ctx context.Context, id string) (*model.Vessel, error) {
    query := `SELECT id, name, imo, status, created_at, updated_at FROM vessels WHERE id = $1`
    var v model.Vessel
    err := r.db.QueryRowContext(ctx, query, id).
        Scan(&v.ID, &v.Name, &v.IMO, &v.Status, &v.CreatedAt, &v.UpdatedAt)
    if err == sql.ErrNoRows {
        return nil, nil  // Not found → return nil, nil (caller decides 404)
    }
    if err != nil {
        return nil, fmt.Errorf("get vessel by id: %w", err)
    }
    return &v, nil
}

// List retrieves paginated vessels
func (r *VesselRepository) List(ctx context.Context, limit, offset int) ([]model.Vessel, int, error) {
    // Get total count
    var total int
    err := r.db.QueryRowContext(ctx, `SELECT COUNT(*) FROM vessels`).Scan(&total)
    if err != nil {
        return nil, 0, fmt.Errorf("count vessels: %w", err)
    }

    query := `
        SELECT id, name, imo, status, created_at, updated_at
        FROM vessels
        ORDER BY created_at DESC
        LIMIT $1 OFFSET $2
    `
    rows, err := r.db.QueryContext(ctx, query, limit, offset)
    if err != nil {
        return nil, 0, fmt.Errorf("list vessels: %w", err)
    }
    defer rows.Close()

    var vessels []model.Vessel
    for rows.Next() {
        var v model.Vessel
        if err := rows.Scan(&v.ID, &v.Name, &v.IMO, &v.Status, &v.CreatedAt, &v.UpdatedAt); err != nil {
            return nil, 0, fmt.Errorf("scan vessel: %w", err)
        }
        vessels = append(vessels, v)
    }
    return vessels, total, rows.Err()
}

// Update applies a partial update
func (r *VesselRepository) Update(ctx context.Context, id string, req model.UpdateVesselRequest) (*model.Vessel, error) {
    query := `
        UPDATE vessels
        SET name = COALESCE($2, name),
            status = COALESCE($3, status),
            updated_at = NOW()
        WHERE id = $1
        RETURNING id, name, imo, status, created_at, updated_at
    `
    var v model.Vessel
    err := r.db.QueryRowContext(ctx, query, id, req.Name, req.Status).
        Scan(&v.ID, &v.Name, &v.IMO, &v.Status, &v.CreatedAt, &v.UpdatedAt)
    if err == sql.ErrNoRows {
        return nil, nil
    }
    if err != nil {
        return nil, fmt.Errorf("update vessel: %w", err)
    }
    return &v, nil
}

// Delete removes a vessel
func (r *VesselRepository) Delete(ctx context.Context, id string) error {
    result, err := r.db.ExecContext(ctx, `DELETE FROM vessels WHERE id = $1`, id)
    if err != nil {
        return fmt.Errorf("delete vessel: %w", err)
    }
    rowsAffected, _ := result.RowsAffected()
    if rowsAffected == 0 {
        return fmt.Errorf("vessel not found")
    }
    return nil
}
```

**Key points to mention:**
- Always use **parameterized queries** (`$1`, `$2`) — never string concatenation (SQL injection prevention)
- Use `context.Context` for cancellation and timeouts
- Use `RETURNING` clause (PostgreSQL) to avoid a second query
- `COALESCE($2, name)` lets PATCH update only non-nil fields
- `sql.ErrNoRows` is the idiomatic way to check "not found" in Go
- Always `defer rows.Close()` and check `rows.Err()` after iteration

---

## Q3: Show me the HTTP handler layer.

**Answer:**

```go
package handler

import (
    "encoding/json"
    "net/http"
    "strconv"

    "myapp/internal/model"
    "myapp/internal/repository"
)

type VesselHandler struct {
    repo *repository.VesselRepository
}

func NewVesselHandler(repo *repository.VesselRepository) *VesselHandler {
    return &VesselHandler{repo: repo}
}

// POST /api/v1/vessels
func (h *VesselHandler) Create(w http.ResponseWriter, r *http.Request) {
    var req model.CreateVesselRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        writeError(w, http.StatusBadRequest, "invalid request body")
        return
    }

    // Validation
    if req.Name == "" || req.IMO == "" {
        writeError(w, http.StatusBadRequest, "name and imo are required")
        return
    }

    vessel, err := h.repo.Create(r.Context(), req)
    if err != nil {
        writeError(w, http.StatusInternalServerError, "failed to create vessel")
        return
    }

    writeJSON(w, http.StatusCreated, vessel)
}

// GET /api/v1/vessels/:id
func (h *VesselHandler) GetByID(w http.ResponseWriter, r *http.Request) {
    id := r.PathValue("id") // Go 1.22+ net/http routing

    vessel, err := h.repo.GetByID(r.Context(), id)
    if err != nil {
        writeError(w, http.StatusInternalServerError, "failed to fetch vessel")
        return
    }
    if vessel == nil {
        writeError(w, http.StatusNotFound, "vessel not found")
        return
    }

    writeJSON(w, http.StatusOK, vessel)
}

// GET /api/v1/vessels
func (h *VesselHandler) List(w http.ResponseWriter, r *http.Request) {
    page, _ := strconv.Atoi(r.URL.Query().Get("page"))
    limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
    if page < 1 { page = 1 }
    if limit < 1 || limit > 100 { limit = 20 }
    offset := (page - 1) * limit

    vessels, total, err := h.repo.List(r.Context(), limit, offset)
    if err != nil {
        writeError(w, http.StatusInternalServerError, "failed to list vessels")
        return
    }

    writeJSON(w, http.StatusOK, map[string]any{
        "data": vessels,
        "meta": map[string]any{
            "page":       page,
            "limit":      limit,
            "totalItems": total,
            "totalPages": (total + limit - 1) / limit,
        },
    })
}

// PATCH /api/v1/vessels/:id
func (h *VesselHandler) Update(w http.ResponseWriter, r *http.Request) {
    id := r.PathValue("id")

    var req model.UpdateVesselRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        writeError(w, http.StatusBadRequest, "invalid request body")
        return
    }

    vessel, err := h.repo.Update(r.Context(), id, req)
    if err != nil {
        writeError(w, http.StatusInternalServerError, "failed to update vessel")
        return
    }
    if vessel == nil {
        writeError(w, http.StatusNotFound, "vessel not found")
        return
    }

    writeJSON(w, http.StatusOK, vessel)
}

// DELETE /api/v1/vessels/:id
func (h *VesselHandler) Delete(w http.ResponseWriter, r *http.Request) {
    id := r.PathValue("id")

    if err := h.repo.Delete(r.Context(), id); err != nil {
        writeError(w, http.StatusNotFound, "vessel not found")
        return
    }

    w.WriteHeader(http.StatusNoContent)
}

// --- helpers ---

func writeJSON(w http.ResponseWriter, status int, data any) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    json.NewEncoder(w).Encode(data)
}

func writeError(w http.ResponseWriter, status int, message string) {
    writeJSON(w, status, map[string]any{
        "error": map[string]string{"message": message},
    })
}
```

---

## Q4: How do you wire everything together in `main.go`?

**Answer:**

```go
package main

import (
    "database/sql"
    "log"
    "net/http"
    "os"

    _ "github.com/lib/pq"
    "myapp/internal/handler"
    "myapp/internal/repository"
)

func main() {
    // Connect to PostgreSQL
    db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
    if err != nil {
        log.Fatal("failed to connect to database:", err)
    }
    defer db.Close()

    if err := db.Ping(); err != nil {
        log.Fatal("database unreachable:", err)
    }

    // Wire dependencies
    vesselRepo := repository.NewVesselRepository(db)
    vesselHandler := handler.NewVesselHandler(vesselRepo)

    // Routes (Go 1.22+ pattern matching)
    mux := http.NewServeMux()
    mux.HandleFunc("POST /api/v1/vessels", vesselHandler.Create)
    mux.HandleFunc("GET /api/v1/vessels", vesselHandler.List)
    mux.HandleFunc("GET /api/v1/vessels/{id}", vesselHandler.GetByID)
    mux.HandleFunc("PATCH /api/v1/vessels/{id}", vesselHandler.Update)
    mux.HandleFunc("DELETE /api/v1/vessels/{id}", vesselHandler.Delete)

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    log.Printf("Server starting on :%s", port)
    log.Fatal(http.ListenAndServe(":"+port, mux))
}
```

**Key talking points:**
- Use Go 1.22+ `net/http` for routing — no need for Gorilla Mux or Chi for simple services
- Dependency injection via constructor functions (no global state)
- Config from environment variables (12-factor app)
- `defer db.Close()` ensures cleanup

---

## Q5: What database schema would you use?

**Answer:**

```sql
CREATE TABLE vessels (
    id         UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name       VARCHAR(255) NOT NULL,
    imo        VARCHAR(7) UNIQUE NOT NULL,
    status     VARCHAR(50) NOT NULL DEFAULT 'active',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Index for listing queries (ORDER BY created_at DESC)
CREATE INDEX idx_vessels_created_at ON vessels (created_at DESC);

-- Index for filtering by status
CREATE INDEX idx_vessels_status ON vessels (status);
```

**Why these choices:**
- `UUID` for IDs — no sequential guessing, safe for distributed systems
- `TIMESTAMPTZ` (not `TIMESTAMP`) — always store with timezone for correctness
- `UNIQUE` on IMO — prevents duplicate vessel registrations at the DB level
- Indexes on columns used in `WHERE` and `ORDER BY` clauses

---

## Q6: How do you handle database migrations?

**Answer:**

I use a migration tool like `golang-migrate` or `goose`:

```
migrations/
  001_create_vessels.up.sql
  001_create_vessels.down.sql
  002_add_fleet_id.up.sql
  002_add_fleet_id.down.sql
```

Each migration has an `up` (apply) and `down` (rollback):

```sql
-- 002_add_fleet_id.up.sql
ALTER TABLE vessels ADD COLUMN fleet_id UUID REFERENCES fleets(id);

-- 002_add_fleet_id.down.sql
ALTER TABLE vessels DROP COLUMN fleet_id;
```

Run with: `migrate -path ./migrations -database $DATABASE_URL up`

**Best practices:**
- Never edit a migration that's been applied — create a new one
- Migrations run in CI/CD before deployment
- Keep migrations small and reversible
- Test both `up` and `down` in development

---

## Q7: How do you test a CRUD service?

**Answer:**

Three levels of testing:

### 1. Unit tests (handler layer, mock the repository)
```go
func TestCreateVessel_InvalidBody(t *testing.T) {
    req := httptest.NewRequest("POST", "/api/v1/vessels", strings.NewReader("invalid"))
    w := httptest.NewRecorder()

    handler := NewVesselHandler(nil) // repo not needed for this test
    handler.Create(w, req)

    if w.Code != http.StatusBadRequest {
        t.Errorf("expected 400, got %d", w.Code)
    }
}
```

### 2. Integration tests (real database, test full flow)
```go
func TestVesselCRUD(t *testing.T) {
    db := setupTestDB(t) // connect to test PostgreSQL
    repo := repository.NewVesselRepository(db)

    // Create
    vessel, err := repo.Create(context.Background(), model.CreateVesselRequest{
        Name: "Test Vessel", IMO: "1234567", Status: "active",
    })
    require.NoError(t, err)
    require.Equal(t, "Test Vessel", vessel.Name)

    // Read
    fetched, err := repo.GetByID(context.Background(), vessel.ID)
    require.NoError(t, err)
    require.Equal(t, vessel.ID, fetched.ID)

    // Update
    newName := "Updated Vessel"
    updated, err := repo.Update(context.Background(), vessel.ID, model.UpdateVesselRequest{
        Name: &newName,
    })
    require.NoError(t, err)
    require.Equal(t, "Updated Vessel", updated.Name)

    // Delete
    err = repo.Delete(context.Background(), vessel.ID)
    require.NoError(t, err)

    // Verify deleted
    gone, err := repo.GetByID(context.Background(), vessel.ID)
    require.NoError(t, err)
    require.Nil(t, gone)
}
```

### 3. API tests (HTTP-level, test the full stack)
```go
func TestListVessels_Paginated(t *testing.T) {
    // seed 25 vessels...
    resp, err := http.Get(server.URL + "/api/v1/vessels?page=2&limit=10")
    require.NoError(t, err)
    require.Equal(t, 200, resp.StatusCode)

    var body map[string]any
    json.NewDecoder(resp.Body).Decode(&body)

    data := body["data"].([]any)
    require.Len(t, data, 10)

    meta := body["meta"].(map[string]any)
    require.Equal(t, float64(25), meta["totalItems"])
}
```

---

## Q8: How do you handle concurrent updates / race conditions?

**Answer:**

Use **optimistic locking** with a version column:

```sql
ALTER TABLE vessels ADD COLUMN version INT NOT NULL DEFAULT 1;
```

```go
func (r *VesselRepository) Update(ctx context.Context, id string, req UpdateVesselRequest, expectedVersion int) (*Vessel, error) {
    query := `
        UPDATE vessels
        SET name = COALESCE($2, name),
            status = COALESCE($3, status),
            version = version + 1,
            updated_at = NOW()
        WHERE id = $1 AND version = $4
        RETURNING id, name, imo, status, version, created_at, updated_at
    `
    var v Vessel
    err := r.db.QueryRowContext(ctx, query, id, req.Name, req.Status, expectedVersion).
        Scan(&v.ID, &v.Name, &v.IMO, &v.Status, &v.Version, &v.CreatedAt, &v.UpdatedAt)
    if err == sql.ErrNoRows {
        return nil, ErrConflict  // 409 Conflict
    }
    return &v, err
}
```

The client sends `"version": 3` with the update. If another request already incremented it to 4, the `WHERE version = 3` matches zero rows → conflict detected → return 409.

---

## Q9: What about database connection pooling?

**Answer:**

Go's `database/sql` has a built-in connection pool. Configure it:

```go
db, _ := sql.Open("postgres", databaseURL)
db.SetMaxOpenConns(25)              // max simultaneous connections
db.SetMaxIdleConns(10)              // keep 10 idle connections ready
db.SetConnMaxLifetime(5 * time.Minute)  // recycle connections every 5 min
db.SetConnMaxIdleTime(1 * time.Minute)  // close idle connections after 1 min
```

**Why it matters:**
- PostgreSQL has a max connections limit (default 100)
- Without limits, under load you'd exhaust connections and get errors
- Idle connection recycling prevents stale connections after network changes

---

## Q10: How do you handle transactions in Go?

**Answer:**

Use `db.BeginTx` for operations that must be atomic:

```go
func (r *VesselRepository) TransferVessel(ctx context.Context, vesselID, fromFleetID, toFleetID string) error {
    tx, err := r.db.BeginTx(ctx, nil)
    if err != nil {
        return fmt.Errorf("begin tx: %w", err)
    }
    defer tx.Rollback() // no-op if already committed

    // Step 1: Remove from old fleet
    _, err = tx.ExecContext(ctx, `UPDATE vessels SET fleet_id = NULL WHERE id = $1 AND fleet_id = $2`, vesselID, fromFleetID)
    if err != nil {
        return fmt.Errorf("remove from fleet: %w", err)
    }

    // Step 2: Add to new fleet
    _, err = tx.ExecContext(ctx, `UPDATE vessels SET fleet_id = $2 WHERE id = $1`, vesselID, toFleetID)
    if err != nil {
        return fmt.Errorf("add to fleet: %w", err)
    }

    // Step 3: Log the transfer
    _, err = tx.ExecContext(ctx, `INSERT INTO transfer_log (vessel_id, from_fleet, to_fleet) VALUES ($1, $2, $3)`, vesselID, fromFleetID, toFleetID)
    if err != nil {
        return fmt.Errorf("log transfer: %w", err)
    }

    return tx.Commit()
}
```

**Key pattern:** `defer tx.Rollback()` is safe even if `tx.Commit()` succeeds — it becomes a no-op. This ensures rollback on any early return with error.
