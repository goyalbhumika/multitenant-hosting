package service

import (
	"context"
	"log"
	"multitenant-hosting/constants"
	"multitenant-hosting/domain"
)

type DeployAppService interface {
	DeployApp(ctx context.Context, appId, deployType string) (*domain.DeployResponse, error)
}

type DeployInstance interface {
	DeployAppInstance(ctx context.Context, appID string) (*domain.DeployResponse, error)
}

type deployAppService struct {
	localDeploySvc DeployInstance
	cloudDeploySvc DeployInstance
}

func NewDeployAppService(localDeploySvc, cloudDeploySvc DeployInstance) DeployAppService {
	return &deployAppService{localDeploySvc: localDeploySvc, cloudDeploySvc: cloudDeploySvc}
}

func (svc *deployAppService) DeployApp(ctx context.Context, appId, deployType string) (*domain.DeployResponse, error) {
	var deployResp *domain.DeployResponse
	var err error

	if deployType == constants.DeployLocal {
		deployResp, err = svc.localDeploySvc.DeployAppInstance(ctx, appId)
	} else {
		deployResp, err = svc.cloudDeploySvc.DeployAppInstance(ctx, appId)
	}
	if err != nil {
		return nil, err
	}
	log.Printf("App deployment for %s successful", appId)
	return deployResp, nil
}
