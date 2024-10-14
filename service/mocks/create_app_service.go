package mocks

import (
	"context"

	"multitenant-hosting/domain"

	"github.com/stretchr/testify/mock"
)

type CreateAppServiceMock struct {
	mock.Mock
}

func (m *CreateAppServiceMock) CreateApp(ctx context.Context, name, deployType string) (*domain.AppResponse, error) {
	args := m.Called(ctx, name, deployType)

	if args.Get(0) != nil {
		return args.Get(0).(*domain.AppResponse), args.Error(1)
	}
	return nil, args.Error(1)
}
