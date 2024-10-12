package service

import (
	service "multitenant-hosting/service/deploy"
	"multitenant-hosting/store"
)

type Registry struct {
	CreateAppSvc CreateAppService
	DeployAppSvc service.DeployAppService
}

func NewRegistry(store store.Store) *Registry {
	deployAppSvc := service.NewDeployAppService()
	return &Registry{
		CreateAppSvc: NewCreateAppService(store, deployAppSvc),
		DeployAppSvc: deployAppSvc,
	}
}
