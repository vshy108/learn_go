//go:build ignore

// Section 17, Topic 126: Mocking and Interfaces for Testing
//
// Go uses interfaces for dependency injection and mocking.
// No special mocking framework needed (though testify/mock exists).
//
// Pattern:
//   1. Define interface for external dependency
//   2. Implement real and mock versions
//   3. Inject via function parameter or struct field
//
// Run: go run examples/s17_mocking_interfaces.go

package main

import "fmt"

// ─────────────────────────────────────────────
// 1. Production interface and struct
// ─────────────────────────────────────────────
type UserStore interface {
	GetUser(id int) (User, error)
	SaveUser(u User) error
}

type User struct {
	ID   int
	Name string
}

// Service that depends on UserStore (not concrete implementation):
type UserService struct {
	store UserStore
}

func (s *UserService) GetUserName(id int) (string, error) {
	user, err := s.store.GetUser(id)
	if err != nil {
		return "", fmt.Errorf("get user name: %w", err)
	}
	return user.Name, nil
}

// ─────────────────────────────────────────────
// 2. Real implementation
// ─────────────────────────────────────────────
type PostgresStore struct {
	// db *sql.DB
}

func (p *PostgresStore) GetUser(id int) (User, error) {
	// Real database query...
	return User{ID: id, Name: "Real User"}, nil
}

func (p *PostgresStore) SaveUser(u User) error {
	return nil
}

// ─────────────────────────────────────────────
// 3. Mock implementation for testing
// ─────────────────────────────────────────────
type MockStore struct {
	users     map[int]User
	getErr    error
	saveErr   error
	saveCalls []User
}

func (m *MockStore) GetUser(id int) (User, error) {
	if m.getErr != nil {
		return User{}, m.getErr
	}
	user, ok := m.users[id]
	if !ok {
		return User{}, fmt.Errorf("user %d not found", id)
	}
	return user, nil
}

func (m *MockStore) SaveUser(u User) error {
	m.saveCalls = append(m.saveCalls, u)
	return m.saveErr
}

func main() {
	fmt.Println("=== Mocking with Interfaces ===")
	fmt.Println()

	// ─────────────────────────────────────────────
	// Production usage
	// ─────────────────────────────────────────────
	fmt.Println("-- Production --")
	prodService := &UserService{store: &PostgresStore{}}
	name, _ := prodService.GetUserName(1)
	fmt.Println("Name:", name)

	// ─────────────────────────────────────────────
	// Test usage (with mock)
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Test (mock) --")
	mock := &MockStore{
		users: map[int]User{
			1: {ID: 1, Name: "Alice"},
			2: {ID: 2, Name: "Bob"},
		},
	}
	testService := &UserService{store: mock}

	name, err := testService.GetUserName(1)
	fmt.Printf("GetUserName(1): name=%q, err=%v\n", name, err)

	name, err = testService.GetUserName(99)
	fmt.Printf("GetUserName(99): name=%q, err=%v\n", name, err)

	// ─────────────────────────────────────────────
	// Test example
	// ─────────────────────────────────────────────
	fmt.Println("\n-- Example test --")
	fmt.Print(`
func TestGetUserName(t *testing.T) {
    mock := &MockStore{
        users: map[int]User{1: {ID: 1, Name: "Alice"}},
    }
    svc := &UserService{store: mock}

    name, err := svc.GetUserName(1)
    assert.NoError(t, err)
    assert.Equal(t, "Alice", name)
}

func TestGetUserName_NotFound(t *testing.T) {
    mock := &MockStore{users: map[int]User{}}
    svc := &UserService{store: mock}

    _, err := svc.GetUserName(99)
    assert.Error(t, err)
}
`)
}
