package service

import (
	"context"
	"log"
)

type DeployAppService interface {
	DeployApp(ctx context.Context, appId string) error
}

type deployAppService struct {
	LocalAppInstance
}

func NewDeployAppService() DeployAppService {
	return &deployAppService{LocalAppInstance{port: 9000}}
}

func (svc *deployAppService) DeployApp(ctx context.Context, appId string) error {
	svc.mu.Lock()
	svc.port++
	svc.mu.Unlock()
	err := startAppInstance(ctx, appId, svc.port)
	if err != nil {
		return err
	}
	log.Printf("App deployment for %s successful", appId)
	return nil
}
