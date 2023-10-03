package config

import (
	"os"
	"strconv"
)

type ServerConfig struct {
	Host string
	Port int
}

func NewServerConfig() (*ServerConfig, error) {
	portStr := os.Getenv("PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, err
	}

	return &ServerConfig{Host: os.Getenv("HOST"), Port: port}, nil
}
