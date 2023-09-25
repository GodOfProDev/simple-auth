package main

import (
	"github.com/godofprodev/simple-pass/internal/config"
	"github.com/godofprodev/simple-pass/internal/router"
	"log/slog"
)

func main() {
	v, err := initViper()
	if err != nil {
		slog.Error("failed to initialize viper: ", err)
	}

	serverCfg := config.NewServerConfig(v)

	r := router.New()

	r.RegisterMiddlewares()
	r.RegisterHandlers()

	err = r.Listen(serverCfg)
	if err != nil {
		slog.Error("there was an issue listening to port 8080: ", err)
	}
}
