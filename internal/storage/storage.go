package storage

import (
	"sync"
)

type Storage struct {
	data struct {
		Projects [][]byte
	}
	mu sync.Mutex
}

func New() *Storage {
	return &Storage{}
}

func (s *Storage) AppendProject(project []byte) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data.Projects = append(s.data.Projects, project)
}

func (s *Storage) GetProjects() [][]byte {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.data.Projects
}
