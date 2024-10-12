package service

import (
	"context"
	"multitenant-hosting/domain"
	"multitenant-hosting/errors"
	service "multitenant-hosting/service/deploy"
	"multitenant-hosting/store"
	"time"
)

type CreateAppService interface {
	CreateApp(ctx context.Context, name string) error
}

type createAppService struct {
	deployAppSvc service.DeployAppService
	store        store.Store
}

func NewCreateAppService(store store.Store, deployAppSvc service.DeployAppService) CreateAppService {
	return &createAppService{store: store, deployAppSvc: deployAppSvc}
}

func (svc *createAppService) CreateApp(ctx context.Context, name string) error {
	app := &domain.App{
		Name:      name,
		ID:        name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := svc.store.CreateApp(app)
	if err != nil {
		return errors.ErrAppAlreadyExists
	}
	err = svc.deployAppSvc.DeployApp(ctx, name)
	return err
}
