package service

import (
	"context"
	"fmt"
	"github.com/netlify/netlify-go"
	"multitenant-hosting/domain"
)

type cloudDeploySvc struct {
	netlifyClient *netlify.Client
}

func NewcloudDeploySvc() DeployInstance {
	return &cloudDeploySvc{
		netlifyClient: netlify.NewClient(&netlify.Config{AccessToken: "//access token here"}),
	}
}

func (svc *cloudDeploySvc) DeployAppInstance(ctx context.Context, appID string) (*domain.DeployResponse, error) {
	// Create a new site
	dns := fmt.Sprintf("%s.gravityfalls42.hitchhiker", appID)
	site, _, err := svc.netlifyClient.Sites.Create(&netlify.SiteAttributes{
		Name:         appID,
		CustomDomain: dns,
	})
	if err != nil {
		return nil, err
	}
	// Deploy the site on netlify
	site.Deploys.Create("/Users/bhumikagoyal/Documents/misc/netlify-index.zip")
	return &domain.DeployResponse{
		DNS: dns,
	}, nil
}
