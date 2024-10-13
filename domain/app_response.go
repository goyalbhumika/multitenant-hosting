package domain

type AppResponse struct {
	Name string `json:"name"`
	Port int    `json:"port"`
	DNS  string `json:"dns"`
}
