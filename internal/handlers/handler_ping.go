package handlers

import (
	"github.com/godofprodev/simple-pass/internal/response"
	"github.com/gofiber/fiber/v2"
)

func (h *Handlers) HandlePing(c *fiber.Ctx) error {
	return response.SuccessMessage("pong")
}
