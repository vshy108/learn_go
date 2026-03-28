# Interview Prep: General Full Stack & Behavioral

> Focus: TypeScript/React, Git, CI/CD, Docker, Agile, and behavioral questions

---

# Part A: TypeScript & React

---

## Q1: What are the key differences between TypeScript and JavaScript?

**Answer:**

| Feature | JavaScript | TypeScript |
|---------|-----------|-----------|
| Type system | Dynamic (runtime) | Static (compile-time) |
| Errors caught | At runtime | At compile time |
| Interfaces | No | Yes |
| Enums | No | Yes |
| Generics | No | Yes |

**Why TypeScript matters for a team:**
- Catches bugs before code runs (typos, wrong argument types, missing properties)
- Self-documenting — types serve as documentation
- Better IDE support (autocomplete, refactoring)
- Safer refactoring — rename a field and the compiler shows every place to update

```typescript
// TypeScript catches this at compile time
interface Vessel {
  id: string;
  name: string;
  imo: string;
  status: 'active' | 'inactive'; // union type = only these values allowed
}

function getVesselLabel(vessel: Vessel): string {
  return `${vessel.name} (${vessel.imo})`;
}

getVesselLabel({ id: '1', name: 'MV Pacific' }); // TS Error: missing 'imo' and 'status'
```

---

## Q2: Explain React component lifecycle and hooks.

**Answer:**

Modern React uses **function components + hooks** (not class components).

### Core hooks:
```tsx
function VesselDashboard() {
  // State — triggers re-render when updated
  const [vessels, setVessels] = useState<Vessel[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  // Effect — runs side effects (API calls, subscriptions)
  useEffect(() => {
    async function fetchVessels() {
      try {
        const data = await api.getVessels();
        setVessels(data);
      } catch (err) {
        setError('Failed to load vessels');
      } finally {
        setLoading(false);
      }
    }
    fetchVessels();
  }, []); // empty dependency array = run once on mount

  // Memo — expensive computation, recalculate only when deps change
  const activeVessels = useMemo(
    () => vessels.filter(v => v.status === 'active'),
    [vessels]
  );

  // Callback — stable function reference for child components
  const handleSelect = useCallback((id: string) => {
    // ...
  }, []);

  if (loading) return <Spinner />;
  if (error) return <Alert>{error}</Alert>;

  return <VesselList vessels={activeVessels} onSelect={handleSelect} />;
}
```

**Key points:**
- `useState` — local component state
- `useEffect` — side effects (fetch data, event listeners). Cleanup with return function
- `useMemo` — cache expensive computations
- `useCallback` — cache function references (prevent unnecessary child re-renders)
- `useRef` — mutable reference that doesn't trigger re-render

---

## Q3: How do you manage state in a React application?

**Answer:**

It depends on the scope and complexity:

| Scope | Solution | When to use |
|-------|---------|-------------|
| **Local** | `useState` | Form inputs, toggles, single component |
| **Shared (small)** | `useContext` + `useReducer` | Theme, auth, 2-3 components sharing state |
| **Server state** | React Query / TanStack Query | API data (caching, refetch, loading states) |
| **Complex global** | Zustand or Redux Toolkit | Large app, many components, complex interactions |

**For a data-heavy dashboard, I'd recommend React Query:**
```tsx
function useVessels() {
  return useQuery({
    queryKey: ['vessels'],
    queryFn: () => api.getVessels(),
    staleTime: 30_000, // data is fresh for 30s
    refetchInterval: 60_000, // auto-refresh every 60s (live dashboard)
  });
}

function VesselList() {
  const { data, isLoading, error } = useVessels();
  // React Query handles caching, deduplication, background refetching
}
```

**Why React Query for a cybersecurity dashboard:**
- Auto-refetch keeps data fresh (critical for monitoring)
- Built-in loading/error states
- Cache deduplication — multiple components using the same query share one request
- Optimistic updates for quick UI feedback

---

## Q4: What is a design system / component library and how do you contribute to one?

**Answer:**

A **design system** is a collection of reusable UI components with consistent styling and behavior. (The JD mentions "contributing to shared design systems.")

**Structure:**
```
components/
  Button/
    Button.tsx        ← component
    Button.styles.ts  ← styled-components or CSS modules
    Button.test.tsx   ← tests
    Button.stories.tsx ← Storybook stories
    index.ts          ← barrel export
```

**Building a reusable component:**
```tsx
interface ButtonProps {
  variant: 'primary' | 'secondary' | 'danger';
  size?: 'sm' | 'md' | 'lg';
  loading?: boolean;
  disabled?: boolean;
  children: React.ReactNode;
  onClick?: () => void;
}

function Button({ variant, size = 'md', loading, disabled, children, onClick }: ButtonProps) {
  return (
    <button
      className={cn(styles.base, styles[variant], styles[size])}
      disabled={disabled || loading}
      onClick={onClick}
    >
      {loading ? <Spinner size="sm" /> : children}
    </button>
  );
}
```

**What makes a good design system contribution:**
- Components are **composable** (small, single responsibility)
- **Typed props** with clear defaults
- **Accessible** (keyboard navigation, ARIA attributes, focus management)
- **Documented** with Storybook for visual testing
- Consistent with the design tokens (colors, spacing, typography)

---

# Part B: Git & Version Control

---

## Q5: What Git workflow do you follow?

**Answer:**

I follow a **feature branch workflow** (common in Agile teams):

```
main ─────────────────────────────────── (production-ready)
  └── feature/vessel-alerts ──────────── (my work)
       ├── commit: add alert model
       ├── commit: implement alert API
       └── commit: add alert list UI
       → PR → code review → merge to main
```

**Daily workflow:**
```bash
git checkout main
git pull origin main
git checkout -b feature/vessel-alerts    # new branch from latest main

# ... work ...
git add -p                               # stage changes interactively (review diffs)
git commit -m "feat: add alert list endpoint"

git push origin feature/vessel-alerts    # push to remote
# → Open PR for code review
```

**Branch naming convention:**
- `feature/description` — new functionality
- `fix/description` — bug fix
- `chore/description` — refactoring, deps, config

**Commit message convention (Conventional Commits):**
- `feat: add vessel alert notifications`
- `fix: handle null IMO in vessel list`
- `refactor: extract validation logic`
- `test: add integration tests for alert API`

---

## Q6: How do you handle merge conflicts?

**Answer:**

1. **Pull the latest main** into your branch:
   ```bash
   git fetch origin
   git rebase origin/main  # or git merge origin/main
   ```

2. **Resolve conflicts** — Git marks them:
   ```
   <<<<<<< HEAD (your changes)
   const MAX_ALERTS = 100;
   =======
   const MAX_ALERTS = 50;
   >>>>>>> origin/main (their changes)
   ```

3. **Choose the correct resolution** — don't just blindly pick one side. Understand both changes.

4. **Test after resolving** — run tests to ensure the merge is correct.

**Rebase vs Merge:**
| | Rebase | Merge |
|-|--------|-------|
| History | Linear, clean | Preserves branch history |
| When | Interactive rebase before PR to clean up | When merging PR to main |
| Risk | Rewrites history (never on shared branches) | Safe on shared branches |

**My preference:** Rebase locally to clean up my commits, then merge (via PR) to main.

---

## Q7: What is `git rebase -i` and when do you use it?

**Answer:**

Interactive rebase lets you clean up your commit history before opening a PR:

```bash
git rebase -i HEAD~4  # edit last 4 commits
```

Opens an editor:
```
pick abc1234 add vessel model
pick def5678 WIP: trying alert query        ← squash this
pick ghi9012 fix typo                        ← squash this
pick jkl3456 implement vessel alert API
```

Change to:
```
pick abc1234 add vessel model
squash def5678 WIP: trying alert query
squash ghi9012 fix typo
pick jkl3456 implement vessel alert API
```

Result: 2 clean commits instead of 4 messy ones.

**Use cases:**
- **squash** — combine "WIP" and "fix typo" commits
- **reword** — fix a commit message
- **reorder** — move commits around
- **drop** — remove a commit entirely

**Rule:** Only rebase commits that haven't been pushed to a shared branch.

---

# Part C: CI/CD, Docker & Cloud

---

## Q8: Explain a typical CI/CD pipeline.

**Answer:**

```
Push to feature branch
    │
    ▼
┌─────────── CI (Continuous Integration) ───────────┐
│  1. Lint       → eslint, golangci-lint             │
│  2. Type check → tsc --noEmit                      │
│  3. Unit tests → go test, jest                     │
│  4. Build      → go build, npm run build           │
│  5. Integration tests → test against real DB       │
│  6. Security scan → dependency audit, SAST         │
└───────────────────────────────────────────────────┘
    │
    ▼ (all green → merge PR)
    │
┌─────────── CD (Continuous Deployment) ────────────┐
│  7. Build Docker image                             │
│  8. Push to container registry (ECR/ACR/GCR)      │
│  9. Deploy to staging                              │
│  10. Run smoke tests                               │
│  11. Deploy to production (manual approval or auto)│
│  12. Health check                                  │
└───────────────────────────────────────────────────┘
```

**GitHub Actions example:**
```yaml
name: CI
on:
  pull_request:
    branches: [main]

jobs:
  test:
    runs-on: ubuntu-latest
    services:
      postgres:
        image: postgres:16
        env:
          POSTGRES_DB: testdb
          POSTGRES_PASSWORD: test
        ports: ['5432:5432']
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v5
        with:
          go-version: '1.22'
      - run: go test ./...
      - run: golangci-lint run
```

---

## Q9: How do you containerize a Go service with Docker?

**Answer:**

**Multi-stage Dockerfile** (small, secure image):

```dockerfile
# Stage 1: Build
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server

# Stage 2: Run
FROM alpine:3.19
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /app/server .
EXPOSE 8080
USER nobody
ENTRYPOINT ["./server"]
```

**Why multi-stage?**
- Build stage has Go toolchain (~800MB) — not needed at runtime
- Final image is ~15MB (just the binary + ca-certificates)
- `USER nobody` — don't run as root (security best practice)
- `CGO_ENABLED=0` — static binary, no C dependencies needed

**Docker Compose for local development:**
```yaml
services:
  api:
    build: .
    ports: ['8080:8080']
    environment:
      DATABASE_URL: postgres://user:pass@db:5432/myapp?sslmode=disable
    depends_on:
      db:
        condition: service_healthy

  db:
    image: postgres:16
    environment:
      POSTGRES_DB: myapp
      POSTGRES_USER: user
      POSTGRES_PASSWORD: pass
    ports: ['5432:5432']
    healthcheck:
      test: ['CMD-SHELL', 'pg_isready -U user -d myapp']
      interval: 5s
      retries: 5
```

---

## Q10: What is Kubernetes at a high level? (JD mentions it)

**Answer:**

Kubernetes (K8s) orchestrates containers across multiple machines.

**Key concepts:**
| Concept | What it does |
|---------|-------------|
| **Pod** | Smallest unit — one or more containers running together |
| **Deployment** | Manages pods — handles scaling, rolling updates |
| **Service** | Stable network endpoint for a set of pods (load balancing) |
| **Ingress** | Routes external HTTP traffic to services |
| **ConfigMap / Secret** | Inject config and credentials into pods |

**Simple deployment:**
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: vessel-api
spec:
  replicas: 3
  selector:
    matchLabels:
      app: vessel-api
  template:
    spec:
      containers:
        - name: api
          image: myorg/vessel-api:v1.2.0
          ports:
            - containerPort: 8080
          env:
            - name: DATABASE_URL
              valueFrom:
                secretKeyRef:
                  name: db-credentials
                  key: url
          livenessProbe:
            httpGet:
              path: /healthz
              port: 8080
```

**Key talking point:** I understand the basics of container orchestration. At L3, I'd deploy to existing K8s infrastructure with guidance from senior engineers, and I'm eager to learn more about it.

---

# Part D: Agile & Collaboration

---

## Q11: Describe your experience with Agile/Scrum.

**Answer:**

**Sprint workflow I'm familiar with:**

| Ceremony | What I do |
|----------|-----------|
| **Sprint Planning** | Help estimate stories, pick work from the backlog, ask clarifying questions |
| **Daily Standup** | Share what I did yesterday, what I'm doing today, any blockers |
| **Backlog Refinement** | Help break down large stories, flag technical risks early |
| **Sprint Review/Demo** | Demo my completed features to stakeholders |
| **Retrospective** | Share what went well, suggest improvements, commit to action items |

**How I manage my work:**
- Break stories into small tasks (< 1 day each)
- Update the board (Jira/Linear) as I progress
- Flag blockers early — don't wait until standup
- "Done" means deployed and verified, not just "code merged"

---

## Q12: How do you approach code reviews?

**Answer:**

### As a reviewer:
- **Read the PR description first** — understand the context and intent
- **Check the happy path** — does the logic make sense?
- **Check error paths** — what happens when things fail?
- **Look for security issues** — SQL injection, XSS, exposed secrets
- **Check tests** — are edge cases covered?
- Comment constructively: "Have you considered...?" not "This is wrong."
- Approve with minor suggestions rather than blocking on non-critical items

### As the author:
- Write a clear PR description (what, why, how to test)
- Keep PRs small (< 400 lines if possible)
- Self-review before requesting reviews
- Respond to all comments — explain reasoning or accept and fix
- View feedback as a learning opportunity, not criticism

---

# Part E: Behavioral Questions

---

## Q13: "Tell me about a time you took ownership of a feature end-to-end."

**Answer (STAR format):**

**Situation:** At my previous role, I was assigned to build a notification feature for our monitoring dashboard.

**Task:** Own it from design through deployment — not just write code, but ensure it works in production.

**Action:**
- Collaborated with the designer to clarify edge cases in the UX
- Built the React frontend components and the Go API endpoint
- Wrote unit and integration tests
- Set up database migration for the notification preferences table
- Deployed to staging, tested manually, then promoted to production
- Monitored logs and metrics after deployment to catch any issues

**Result:** The feature launched without incidents. I proactively found a performance issue during monitoring (N+1 query) and fixed it before users noticed.

**Key message:** "Done" means production-ready and monitored, not just code-merged.

---

## Q14: "How do you handle receiving critical feedback on your code?"

**Answer:**

I welcome it. Code reviews are one of the best ways to learn.

**Example:** A senior engineer pointed out that my error handling was inconsistent — some handlers returned 500 for validation errors. I didn't take it personally. I:
1. Thanked them for catching it
2. Asked for their recommended pattern
3. Refactored the whole handler layer to use a consistent `AppError` type
4. Shared the pattern with the team so others could use it too

**Mindset:** The code isn't "mine" — it belongs to the team. A reviewer finding a bug means fewer bugs in production. That's a win.

---

## Q15: "How do you approach learning a new technology you haven't used before?"

**Answer:**

My approach:
1. **Read the official docs first** — understand the mental model, not just the API
2. **Build something small** — a toy project to get hands dirty (like this learn_go repo!)
3. **Read real code** — open-source projects show production patterns
4. **Pair with someone experienced** — ask questions, understand the "why" behind patterns
5. **Teach/document what I learn** — writing it down solidifies understanding

**Example:** When learning Go, I built example files for every concept (types, functions, error handling, concurrency) and wrote a cheatsheet. The act of writing forces me to understand deeply, not just superficially.

---

## Q16: "Describe how you ensure quality in your code."

**Answer:**

Quality at multiple levels:

1. **Before writing code:** Understand requirements. Ask clarifying questions. Consider edge cases.
2. **While writing code:**
   - Strong types (TypeScript strict mode, Go's type system)
   - Small functions with single responsibility
   - Meaningful names — code should read like prose
3. **Automated checks:**
   - Linters: `eslint`, `golangci-lint`
   - Type checking: `tsc --noEmit`
   - Tests: unit, integration, and API-level
4. **Manual checks:**
   - Self-review my PR diff before requesting review
   - Test manually in the browser / with curl
5. **After deployment:**
   - Monitor logs and error rates
   - Check performance metrics

**Philosophy:** Quality isn't a phase — it's baked into every step.

---

# Part F: System Design (Lightweight)

---

## Q17: "How would you design a real-time alert dashboard?"

**Answer:**

**Architecture:**
```
┌──────────┐     ┌──────────┐     ┌─────────────┐
│  React   │◄───►│  Go API  │◄───►│ PostgreSQL  │
│ Frontend │     │ Server   │     │ (alerts,    │
└──────────┘     └──────────┘     │  vessels)   │
     ▲                ▲           └─────────────┘
     │                │
     │           ┌────┴────┐
     └───────────│ WebSocket│ (real-time push)
                 └─────────┘
```

**Components:**
- **REST API** for CRUD (list alerts, acknowledge, filter by vessel)
- **WebSocket** for real-time updates (new alert → push to connected dashboards)
- **PostgreSQL** for persistent storage
- **Redis** (optional) for pub/sub if multiple API server instances
- **React Query** for data fetching + WebSocket integration for live updates

**Frontend approach:**
```tsx
function AlertDashboard() {
  const { data: alerts } = useQuery({ queryKey: ['alerts'], queryFn: fetchAlerts });

  // WebSocket for real-time updates
  useEffect(() => {
    const ws = new WebSocket('wss://api.example.com/ws/alerts');
    ws.onmessage = (event) => {
      const newAlert = JSON.parse(event.data);
      queryClient.setQueryData(['alerts'], (old) => [newAlert, ...old]);
    };
    return () => ws.close();
  }, []);

  return <AlertTable alerts={alerts} />;
}
```

**Key decisions to discuss:**
- Polling vs WebSocket → WebSocket for low-latency alerts
- Pagination for historical alerts, real-time push for new ones
- Rate-limit the WebSocket to prevent flooding the UI
- Optimistic UI for "acknowledge" action (update locally, sync to server)

---

## Quick Reference: Key Technologies from the JD

```
┌─────────────────────────────────────────────────────┐
│ Required                                            │
│  ✓ TypeScript + React (frontend)                    │
│  ✓ Go or Rust (backend) ← you're learning Go       │
│  ✓ PostgreSQL / MySQL (relational DB)               │
│  ✓ Git                                              │
│                                                     │
│ Preferred (mention familiarity, willingness to learn)│
│  ○ Design systems / component libraries             │
│  ○ CI/CD (GitHub Actions, Jenkins)                  │
│  ○ Docker + Kubernetes                              │
│  ○ Elastic Search + Kibana                          │
│  ○ NoSQL / time-series DBs (InfluxDB)               │
│  ○ Agile / Scrum                                    │
│  ○ AWS / Azure / GCP                                │
└─────────────────────────────────────────────────────┘
```
