package handlers

import (
	"github.com/godofprodev/simple-pass/internal/models"
	"github.com/godofprodev/simple-pass/internal/response"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (h *Handlers) HandleRegister(c *fiber.Ctx) error {
	params := new(models.CreateUserParams)

	if err := c.BodyParser(&params); err != nil {
		return response.ErrParsingParams()
	}

	if params.User == "" {
		return response.ErrRequired("user")
	}
	if params.Password == "" {
		return response.ErrRequired("password")
	}

	encryptedPass, err := bcrypt.GenerateFromPassword([]byte(params.Password), 14)
	if err != nil {
		return response.ErrEncryptingPassword()
	}

	user := models.User{
		Id:       uuid.New(),
		Username: params.User,
		Password: encryptedPass,
	}

	if err := h.store.CreateUser(&user); err != nil {
		return response.ErrCreating("user")
	}

	return response.SuccessCreated(user)
}
