package service

import (
	"context"
	"multitenant-hosting/constants"
	"multitenant-hosting/domain"
	"multitenant-hosting/errors"
	service "multitenant-hosting/service/deploy"
	"multitenant-hosting/store"
	"time"
)

type CreateAppService interface {
	CreateApp(ctx context.Context, name, deployType string) (*domain.AppResponse, error)
}

type createAppService struct {
	deployAppSvc service.DeployAppService
	store        store.Store
}

func NewCreateAppService(store store.Store, deployAppSvc service.DeployAppService) CreateAppService {
	return &createAppService{store: store, deployAppSvc: deployAppSvc}
}

func (svc *createAppService) CreateApp(ctx context.Context, name string, deployType string) (*domain.AppResponse, error) {
	if app := svc.store.GetApp(name); app != nil {
		return nil, errors.ErrAppAlreadyExists
	}
	app := &domain.App{
		Name:      name,
		ID:        name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Status:    constants.StatusCreated,
	}

	//Create the entry in the database
	svc.store.CreateApp(app)

	// Deploy the app and create DNS endpoint
	deployResp, err := svc.deployAppSvc.DeployApp(ctx, name, deployType)
	if err != nil {
		svc.store.UpdateAppState(constants.StatusDeploymentFailed, name)
		return nil, err
	}
	svc.store.UpdateAppState(constants.StatusDeployed, name)
	svc.store.UpdateAppDNS(deployResp.DNS, name)
	svc.store.UpdateAppPort(deployResp.Port, name)
	resp := &domain.AppResponse{
		Name: name,
		Port: deployResp.Port,
		DNS:  deployResp.DNS,
	}
	return resp, err
}
