package router

import (
	"github.com/godofprodev/simple-pass/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	app *fiber.App
}

func New() *Router {
	app := fiber.New()

	return &Router{app: app}
}

func (r *Router) Listen() error {
	err := r.app.Listen(":8080")
	if err != nil {
		return err
	}

	return nil
}

func (r *Router) RegisterMiddlewares() {

}

func (r *Router) RegisterHandlers() {
	h := handlers.New()

	v1 := r.app.Group("/v1")

	v1.Get("/ping", h.HandlePing)
}
