package config

import (
	"os"
)

type Config struct {
	NetlifyToken  string
	IndexFilePath string
}

var Configuration *Config

func SetConfig() {
	//TODO: read this from yaml config
	netlifyToken := os.Getenv("NETLIFY_TOKEN")
	path := os.Getenv("INDEX_FILE_PATH")
	Configuration = &Config{NetlifyToken: netlifyToken, IndexFilePath: path}
}

func (c *Config) GetNetlifyToken() string {
	return c.NetlifyToken
}

func (c *Config) GetIndexFilePath() string {
	return c.IndexFilePath
}
