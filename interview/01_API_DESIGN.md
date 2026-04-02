# Interview Prep: Basic API Design

> Focus: RESTful API design, HTTP fundamentals, Go backend

---

## Q1: What is REST and what makes an API "RESTful"?

**Answer:**

REST (Representational State Transfer) is an architectural style for building web services. A RESTful API follows these principles:

1. **Client-Server separation** — frontend and backend are independent; they communicate over HTTP.
2. **Statelessness** — each request contains all information needed to process it; the server stores no session state between requests.
3. **Uniform Interface** — resources are identified by URIs, manipulated through representations (JSON), and use standard HTTP methods.
4. **Layered System** — the client doesn't need to know if it's talking to the actual server or a proxy/load balancer.
5. **Cacheability** — responses should indicate whether they can be cached.

In practice, a RESTful API means:
- Resources are nouns: `/api/v1/ships`, not `/api/v1/getShips`
- HTTP methods convey the action: `GET`, `POST`, `PUT`, `PATCH`, `DELETE`
- Status codes communicate the result: `200 OK`, `201 Created`, `404 Not Found`, etc.

---

## Q2: How do you design URL structures for a RESTful API?

**Answer:**

I follow a resource-oriented approach:

```
GET    /api/v1/vessels          → List all vessels
POST   /api/v1/vessels          → Create a new vessel
GET    /api/v1/vessels/:id      → Get a specific vessel
PUT    /api/v1/vessels/:id      → Full update of a vessel
PATCH  /api/v1/vessels/:id      → Partial update of a vessel
DELETE /api/v1/vessels/:id      → Delete a vessel
```

For nested resources (sub-resources):
```
GET    /api/v1/vessels/:id/alerts       → List alerts for a vessel
POST   /api/v1/vessels/:id/alerts       → Create an alert for a vessel
```

Key rules I follow:
- **Use plural nouns** (`/vessels`, not `/vessel`)
- **Use kebab-case** for multi-word resources (`/alert-rules`)
- **Version the API** (`/api/v1/...`) to allow non-breaking evolution
- **Use query parameters** for filtering, sorting, and pagination: `/vessels?status=active&sort=name&page=2&limit=20`
- **Avoid verbs in URLs** — the HTTP method is the verb

---

## Q3: Explain the difference between PUT and PATCH.

**Answer:**

| Aspect | PUT | PATCH |
|--------|-----|-------|
| Semantics | Replace the **entire** resource | Apply a **partial** update |
| Idempotent? | Yes | Can be, but not required |
| Request body | Full resource representation | Only the fields being changed |

**Example:**

Existing resource:
```json
{ "id": 1, "name": "Vessel A", "status": "active", "imo": "1234567" }
```

`PUT /vessels/1` with `{ "name": "Vessel B" }` → **replaces** the resource, other fields may be lost or zeroed:
```json
{ "id": 1, "name": "Vessel B", "status": "", "imo": "" }
```

`PATCH /vessels/1` with `{ "name": "Vessel B" }` → **merges**, other fields remain:
```json
{ "id": 1, "name": "Vessel B", "status": "active", "imo": "1234567" }
```

In practice, **PATCH is more commonly used** because clients usually want to update one or two fields, not send the entire object.

---

## Q4: What HTTP status codes should a well-designed API use?

**Answer:**

### Success (2xx)
| Code | Meaning | When to use |
|------|---------|-------------|
| `200 OK` | Request succeeded | GET, PUT, PATCH responses with body |
| `201 Created` | Resource created | POST that creates a resource |
| `204 No Content` | Success, no body | DELETE, or PUT/PATCH if no body returned |

### Client Errors (4xx)
| Code | Meaning | When to use |
|------|---------|-------------|
| `400 Bad Request` | Invalid input | Malformed JSON, validation errors |
| `401 Unauthorized` | Not authenticated | Missing or invalid auth token |
| `403 Forbidden` | Not authorized | Authenticated but lacks permission |
| `404 Not Found` | Resource doesn't exist | ID not found in database |
| `409 Conflict` | State conflict | Duplicate key, version conflict |
| `422 Unprocessable Entity` | Semantic error | Valid JSON but business rule violation |
| `429 Too Many Requests` | Rate limited | Client exceeded rate limit |

### Server Errors (5xx)
| Code | Meaning | When to use |
|------|---------|-------------|
| `500 Internal Server Error` | Unexpected failure | Unhandled exception, bug |
| `502 Bad Gateway` | Upstream failure | Dependent service is down |
| `503 Service Unavailable` | Overloaded / maintenance | Server temporarily can't handle requests |

**Key principle:** Never return `200` with an error in the body. Use proper status codes so clients (and monitoring) can react correctly.

---

## Q5: How would you design pagination for a list endpoint?

**Answer:**

I use **offset-based pagination** for simple cases and **cursor-based pagination** for large or real-time datasets.

### Offset-based (simpler, good for most CRUD):
```
GET /api/v1/vessels?page=2&limit=20
```

Response:
```json
{
  "data": [...],
  "meta": {
    "page": 2,
    "limit": 20,
    "totalItems": 143,
    "totalPages": 8
  }
}
```

### Cursor-based (better for large/live datasets):
```
GET /api/v1/alerts?limit=20&cursor=eyJpZCI6MTAwfQ==
```

Response:
```json
{
  "data": [...],
  "meta": {
    "nextCursor": "eyJpZCI6MTIwfQ==",
    "hasMore": true
  }
}
```

**Why cursor-based can be better:**
- Offset pagination breaks when items are inserted/deleted between pages
- `OFFSET 10000` in SQL is slow — it scans and discards 10000 rows
- Cursor uses `WHERE id > last_seen_id LIMIT 20`, which is O(1) with an index

---

## Q6: How do you handle API versioning?

**Answer:**

Three common strategies:

| Strategy | Example | Pros | Cons |
|----------|---------|------|------|
| **URL path** | `/api/v1/vessels` | Simple, explicit, cacheable | URL changes on version bump |
| **Header** | `Accept: application/vnd.myapp.v1+json` | Clean URLs | Hidden, harder to test in browser |
| **Query param** | `/vessels?version=1` | Easy to implement | Pollutes query string |

**I prefer URL path versioning** because:
- It's the most common and widely understood
- Easy to route at the infrastructure level (e.g., different deployments per version)
- Clear in documentation and logs

When bumping versions, I keep the old version running for a deprecation period and communicate timelines to consumers.

---

## Q7: How do you structure a JSON API response?

**Answer:**

I use a consistent **envelope pattern**:

### Success response:
```json
{
  "data": {
    "id": "abc-123",
    "name": "MV Pacific",
    "status": "active"
  }
}
```

### List response:
```json
{
  "data": [
    { "id": "abc-123", "name": "MV Pacific" },
    { "id": "def-456", "name": "MV Atlantic" }
  ],
  "meta": {
    "page": 1,
    "limit": 20,
    "totalItems": 2
  }
}
```

### Error response:
```json
{
  "error": {
    "code": "VALIDATION_ERROR",
    "message": "Invalid request body",
    "details": [
      { "field": "email", "message": "must be a valid email address" },
      { "field": "name", "message": "is required" }
    ]
  }
}
```

**Why this structure?**
- Consistent: clients always know where to find data vs errors
- Extensible: add `meta`, `links`, `warnings` without breaking
- Machine-readable: error `code` can be used for i18n or client-side logic

---

## Q8: What is idempotency and why does it matter in API design?

**Answer:**

**Idempotent** means making the same request multiple times produces the same result as making it once.

| Method | Idempotent? | Why |
|--------|-------------|-----|
| GET | Yes | Reading doesn't change state |
| PUT | Yes | Replacing with the same data yields the same result |
| DELETE | Yes | Deleting an already-deleted resource is a no-op (return 404 or 204) |
| PATCH | Usually yes | Depends on implementation |
| POST | **No** | Each call typically creates a new resource |

**Why it matters:**
- Network failures cause retries — if POST isn't idempotent, you can get duplicate records
- Solution: use an **idempotency key** (a client-generated UUID sent in a header)

```
POST /api/v1/orders
Idempotency-Key: 550e8400-e29b-41d4-a716-446655440000
```

The server stores the key and returns the cached response on retries instead of creating a duplicate.

---

## Q9: How would you design an API for a monitoring dashboard?

**Answer:**

For a monitoring platform, I'd design resources around the domain:

```
# Vessels
GET    /api/v1/vessels                        → List vessels with status summary
GET    /api/v1/vessels/:id                    → Vessel details + risk score

# Alerts
GET    /api/v1/vessels/:id/alerts             → Alerts for a vessel
POST   /api/v1/vessels/:id/alerts/:alertId/acknowledge → Acknowledge an alert

# Network Events
GET    /api/v1/vessels/:id/events             → Network events (cursor-paginated)
GET    /api/v1/vessels/:id/events/stats       → Aggregated stats over a time range

# Dashboard
GET    /api/v1/dashboard/summary              → Fleet-wide summary (counts, risk scores)
GET    /api/v1/dashboard/timeline             → Time-series data for charts
```

Key design decisions:
- **Alerts are sub-resources of vessels** because they always belong to a vessel
- **Cursor pagination for events** because there could be millions
- **Separate `/stats` endpoint** instead of computing in the client — push heavy aggregation to the server/database
- Use **query params for time ranges**: `?from=2026-01-01T00:00:00Z&to=2026-03-28T00:00:00Z`
- **Rate limit** the API (429) to prevent abuse, especially on heavy aggregation endpoints

---

## Q10: How would you handle authentication and authorization in an API?

**Answer:**

**Authentication** (who are you?):
- Use **JWT (JSON Web Tokens)** or **OAuth 2.0** with short-lived access tokens and longer refresh tokens
- Token sent in `Authorization: Bearer <token>` header
- Never in query parameters (logged in server access logs, browser history)

**Authorization** (what can you do?):
- **RBAC (Role-Based Access Control)** — assign roles like `admin`, `operator`, `viewer`
- Check permissions in middleware before the handler runs
- Return `401` if not authenticated, `403` if authenticated but not permitted

**Go middleware example:**
```go
func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("Authorization")
        if token == "" {
            http.Error(w, "unauthorized", http.StatusUnauthorized)
            return
        }
        claims, err := validateToken(strings.TrimPrefix(token, "Bearer "))
        if err != nil {
            http.Error(w, "unauthorized", http.StatusUnauthorized)
            return
        }
        ctx := context.WithValue(r.Context(), "user", claims)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}
```

**Security best practices:**
- Always validate tokens server-side
- Use HTTPS everywhere
- Set token expiry short (15 min access, 7 day refresh)
- Implement CORS properly for frontend clients
- Sanitize all inputs to prevent injection

---

## Quick Reference: HTTP Methods at a Glance

```
┌────────┬────────────┬────────────┬─────────────┐
│ Method │ CRUD       │ Idempotent │ Request Body│
├────────┼────────────┼────────────┼─────────────┤
│ GET    │ Read       │ Yes        │ No          │
│ POST   │ Create     │ No         │ Yes         │
│ PUT    │ Replace    │ Yes        │ Yes         │
│ PATCH  │ Update     │ Usually    │ Yes         │
│ DELETE │ Delete     │ Yes        │ Rarely      │
└────────┴────────────┴────────────┴─────────────┘
```
