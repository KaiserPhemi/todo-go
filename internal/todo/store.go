package todo

import (
	"fmt"
	"sync"
)

// interfaces
type Store interface {
	List() []Todo
	Get(id string) (Todo, error)
	Create(t Todo) Todo
	Update(id string, t Todo) (Todo, error)
	Delete(id string) error
}

type MemoryStore struct {
	mu     sync.RWMutex
	todos  map[string]Todo
	nextID int
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		todos:  make(map[string]Todo),
		nextID: 1,
	}
}

// fetch alla todos
func (s *MemoryStore) List() []Todo {
	s.mu.RLock()
	defer s.mu.RUnlock()
	result := make([]Todo, 0, len(s.todos))
	for _, t := range s.todos {
		result = append(result, t)
	}
	return result
}

// get todo by id
func (s *MemoryStore) Get(id string) (Todo, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	t, ok := s.todos[id]
	if !ok {
		return Todo{}, fmt.Errorf("todo %q not found", id)
	}
	return t, nil
}

// create a todo
func (s *MemoryStore) Create(t Todo) Todo {
	s.mu.Lock()
	defer s.mu.Unlock()

	id := fmt.Sprintf("%d", s.nextID)
	s.nextID++
	t.ID = id
	s.todos[id] = t
	return t
}

// update todo
func (s *MemoryStore) Update(id string, t Todo) (Todo, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	existing, ok := s.todos(id)
	if !ok {
		return Todo{}, fmt.Errorf("Todo %q not found", id)
	}

	if t.Title != "" {
		existing.Title = t.Title
	}
	existing.Completed = t.Completed

	s.todos[id] = existing
	return existing, nil
}

func (s *MemoryStore) Delete(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.todos[id]; !ok {
		return fmt.Errorf("todo %q not found", id)
	}
	delete(s.todos, id)
	return nil
}
