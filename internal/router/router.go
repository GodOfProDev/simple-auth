package router

import (
	"fmt"
	"github.com/godofprodev/simple-pass/internal"
	"github.com/godofprodev/simple-pass/internal/config"
	"github.com/godofprodev/simple-pass/internal/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Router struct {
	app *fiber.App
}

func New() *Router {
	app := fiber.New(fiber.Config{
		ErrorHandler: customErrorHandler,
	})

	return &Router{app: app}
}

func (r *Router) Listen(s *config.ServerConfig) error {
	err := r.app.Listen(fmt.Sprintf("%v:%v", s.Host, s.Port))
	if err != nil {
		return err
	}

	return nil
}

func (r *Router) RegisterMiddlewares() {
	r.app.Use(logger.New(logger.Config{
		Format:   "${cyan}[${time}] ${white}| ${status} | ${latency} | ${white}${ip} | ${method} | ${white}${path}\n",
		TimeZone: "UTC",
	}))
	r.app.Use(cors.New())
}

func (r *Router) RegisterHandlers() {
	h := handlers.New()

	v1 := r.app.Group("/v1")

	v1.Get("/ping", h.HandlePing)
	v1.Post("/register", h.HandleRegister)
	v1.Post("/login", h.HandleRegister)
}

func customErrorHandler(c *fiber.Ctx, err error) error {
	switch e := err.(type) {
	case internal.APIError:
		return c.Status(e.Status).JSON(e)
	case internal.APISuccessData:
		return c.Status(e.Status).JSON(e.Data)
	case internal.APISuccessResponse:
		return c.Status(e.Status).JSON(e)
	default:
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{"message": "internal server error"})
	}
}
