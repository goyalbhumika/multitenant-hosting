package mocks

import (
	"github.com/stretchr/testify/mock"
	"multitenant-hosting/domain"
)

type StoreMock struct {
	mock.Mock
}

func (m *StoreMock) CreateApp(app *domain.App) error {
	args := m.Called(app)
	return args.Error(0)
}

func (m *StoreMock) GetApp(name string) *domain.App {
	args := m.Called(name)
	if app, ok := args.Get(0).(*domain.App); ok {
		return app
	}
	return nil
}

func (m *StoreMock) UpdateAppState(status string, name string) error {
	args := m.Called(status, name)
	return args.Error(0)
}

func (m *StoreMock) UpdateAppPort(port int, name string) error {
	args := m.Called(port, name)
	return args.Error(0)
}

func (m *StoreMock) UpdateAppDNS(dns string, name string) error {
	args := m.Called(dns, name)
	return args.Error(0)
}
