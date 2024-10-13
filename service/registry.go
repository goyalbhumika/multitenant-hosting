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
	localDeploySvc := service.NewlocalDeploySvc(9000)
	cloudDeploySvc := service.NewcloudDeploySvc()
	deployAppSvc := service.NewDeployAppService(localDeploySvc, cloudDeploySvc)
	return &Registry{
		CreateAppSvc: NewCreateAppService(store, deployAppSvc),
		DeployAppSvc: deployAppSvc,
	}
}
