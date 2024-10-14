package mocks

import (
	"context"
	"github.com/stretchr/testify/mock"
	"multitenant-hosting/domain"
)

type DeployAppServiceMock struct {
	mock.Mock
}

func (m *DeployAppServiceMock) DeployApp(ctx context.Context, name, deployType string) (*domain.DeployResponse, error) {
	args := m.Called(ctx, name, deployType)
	if deployResp, ok := args.Get(0).(*domain.DeployResponse); ok {
		return deployResp, args.Error(1)
	}
	return nil, args.Error(1)
}
