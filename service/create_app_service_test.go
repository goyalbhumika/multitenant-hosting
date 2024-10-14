package service_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"multitenant-hosting/constants"
	"multitenant-hosting/domain"
	errors2 "multitenant-hosting/errors"
	"multitenant-hosting/service"
	"multitenant-hosting/service/mocks"
	mock2 "multitenant-hosting/store/mocks"
)

func TestCreateAppService_CreateApp(t *testing.T) {
	// Create mocks for store and deploy app service
	storeMock := new(mock2.StoreMock)
	deployAppSvcMock := new(mocks.DeployAppServiceMock)

	// Create instance of the service using the mocks
	createAppSvc := service.NewCreateAppService(storeMock, deployAppSvcMock)

	t.Run("successful app creation and deployment", func(t *testing.T) {
		deployResponse := &domain.DeployResponse{
			DNS:  "testapp.example.com",
			Port: 8080,
		}

		storeMock.On("GetApp", "testapp").Return(nil)
		storeMock.On("CreateApp", mock.AnythingOfType("*domain.App")).Return(nil)
		deployAppSvcMock.On("DeployApp", mock.Anything, "testapp", "cloud").Return(deployResponse, nil)
		storeMock.On("UpdateAppState", constants.StatusDeployed, "testapp").Return(nil)
		storeMock.On("UpdateAppDNS", "testapp.example.com", "testapp").Return(nil)
		storeMock.On("UpdateAppPort", 8080, "testapp").Return(nil)

		ctx := context.Background()
		resp, err := createAppSvc.CreateApp(ctx, "testapp", "cloud")

		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, "testapp", resp.Name)
		assert.Equal(t, "testapp.example.com", resp.DNS)
		assert.Equal(t, 8080, resp.Port)
		storeMock.AssertExpectations(t)
		deployAppSvcMock.AssertExpectations(t)
	})

	t.Run("app already exists", func(t *testing.T) {
		existingApp := &domain.App{
			Name:      "ExistingApp",
			ID:        "ExistingApp",
			CreatedAt: time.Now(),
			Status:    constants.StatusDeployed,
		}
		storeMock.On("GetApp", "ExistingApp").Return(existingApp)

		ctx := context.Background()
		resp, err := createAppSvc.CreateApp(ctx, "ExistingApp", "cloud")

		assert.Error(t, err)
		assert.Nil(t, resp)
		assert.Equal(t, errors2.ErrAppAlreadyExists, err)
		storeMock.AssertExpectations(t)
	})

	t.Run("deployment failure", func(t *testing.T) {
		storeMock.On("GetApp", "FailApp").Return(nil)
		storeMock.On("CreateApp", mock.AnythingOfType("*domain.App")).Return(nil)
		deployAppSvcMock.On("DeployApp", mock.Anything, "FailApp", "cloud").Return(nil, errors.New("deployment error"))
		storeMock.On("UpdateAppState", constants.StatusDeploymentFailed, "FailApp").Return(nil)

		ctx := context.Background()
		resp, err := createAppSvc.CreateApp(ctx, "FailApp", "cloud")

		assert.Error(t, err)
		assert.Nil(t, resp)
		assert.Equal(t, "deployment error", err.Error())
		storeMock.AssertExpectations(t)
		deployAppSvcMock.AssertExpectations(t)
	})
}
