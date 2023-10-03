package handlers

import (
	"github.com/godofprodev/simple-pass/internal/models"
	"github.com/godofprodev/simple-pass/internal/response"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func (h *Handlers) HandleLogin(c *fiber.Ctx) error {
	params := new(models.LoginUserParams)

	if err := c.BodyParser(&params); err != nil {
		return response.ErrParsingParams()
	}

	if params.Username == "" {
		return response.ErrRequired("user")
	}
	if params.Password == "" {
		return response.ErrRequired("password")
	}

	user, err := h.store.GetUser(params.Username)
	if err != nil {
		return response.ErrNotFound(params.Username)
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(params.Password)); err != nil {
		return response.ErrIncorrectPassword()
	}

	return response.SuccessGotten(user)
}
