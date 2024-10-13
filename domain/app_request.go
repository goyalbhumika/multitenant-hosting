package domain

type AppRequest struct {
	Name       string `json:"name"`
	DeployType string `json:"deploy_type"`
}
