package store

import (
	"fmt"
	"multitenant-hosting/domain"
)

type Store interface {
	CreateApp(app *domain.App) error
	GetApp(name string) *domain.App
	UpdateAppState(status string, name string) error
	UpdateAppPort(port int, name string) error
	UpdateAppDNS(dns string, name string) error
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

func (s *store) UpdateAppState(status string, name string) error {
	s.apps[name].Status = status
	return nil
}

func (s *store) UpdateAppDNS(dns string, name string) error {
	s.apps[name].DNS = &domain.DNS{ARecord: dns}
	return nil
}

func (s *store) UpdateAppPort(port int, name string) error {
	s.apps[name].Port = port
	return nil
}
