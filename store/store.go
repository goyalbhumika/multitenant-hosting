package store

import (
	"fmt"
	"multitenant-hosting/domain"
)

type Store interface {
	CreateApp(app *domain.App) error
	GetApp(name string) *domain.App
}

type store struct {
	apps map[string]*domain.App
}

func NewStore() Store {
	return &store{apps: make(map[string]*domain.App)}
}

func (s *store) CreateApp(app *domain.App) error {
	if _, ok := s.apps[app.Name]; ok {
		return fmt.Errorf("app %s already exists", app.Name)
	}
	s.apps[app.ID] = app
	return nil
}

func (s *store) GetApp(name string) *domain.App {
	return s.apps[name]
}
