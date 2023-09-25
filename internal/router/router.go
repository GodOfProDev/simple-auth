package router

import (
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
	v1 := r.app.Group("/v1")

	v1.Get("/signup", func(ctx *fiber.Ctx) error {
		return ctx.SendString("HELLO WORLD")
	})
}
