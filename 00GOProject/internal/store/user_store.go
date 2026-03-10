package store

import (
	"sync"

	"gobackend/internal/models"
)

type UserStore struct {
	mu      sync.RWMutex
	users   map[int]models.User
	counter int
}

func NewUserStore() *UserStore {
	return &UserStore{
		users: make(map[int]models.User),
	}
}

func (s *UserStore) GetAll() []models.User {
	s.mu.RLock()
	defer s.mu.RUnlock()

	list := make([]models.User, 0, len(s.users))
	for _, u := range s.users {
		list = append(list, u)
	}
	return list
}

func (s *UserStore) GetByID(id int) (models.User, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	u, ok := s.users[id]
	return u, ok
}

func (s *UserStore) Create(u models.User) models.User {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.counter++
	u.ID = s.counter
	s.users[u.ID] = u
	return u
}

func (s *UserStore) Update(id int, u models.User) (models.User, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.users[id]; !ok {
		return models.User{}, false
	}
	u.ID = id
	s.users[id] = u
	return u, true
}

func (s *UserStore) Delete(id int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.users[id]; !ok {
		return false
	}
	delete(s.users, id)
	return true
}
