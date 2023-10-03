package handlers

import (
	"github.com/godofprodev/simple-pass/internal/response"
	"github.com/gofiber/fiber/v2"
)

func (h *Handlers) HandlePing(c *fiber.Ctx) error {
	username := c.Locals("user")

	user, err := h.store.GetUser(username.(string))
	if err != nil {
		return response.ErrNotFound(username.(string))
	}
	return response.SuccessGotten(user)
}
