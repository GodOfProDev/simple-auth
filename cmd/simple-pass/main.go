package main

import (
	"github.com/godofprodev/simple-pass/internal/router"
	"log/slog"
)

func main() {
	r := router.New()

	r.RegisterMiddlewares()
	r.RegisterHandlers()

	err := r.Listen()
	if err != nil {
		slog.Error("there was an issue listening to port 8080: ", err)
	}
}
