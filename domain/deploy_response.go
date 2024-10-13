package domain

type DeployResponse struct {
	Port int    `json:"port"`
	DNS  string `json:"dns"`
}
