package auth

import (
	"github.com/godofprodev/simple-pass/internal/models"
	"github.com/godofprodev/simple-pass/internal/response"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

var SecretKey = os.Getenv("JWT_SECRET")

func GenerateJWT(user *models.User) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Subject:   user.Username,
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
	})

	return claims.SignedString([]byte(SecretKey))
}

func AuthenticatedUser(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, response.ErrUnauthorized()
		}

		return []byte(SecretKey), nil
	})

	if err != nil {
		return response.ErrUnauthorized()
	}

	claims := token.Claims.(*jwt.RegisteredClaims)

	if token.Valid {
		if time.Now().Unix() > claims.ExpiresAt.Unix() {
			return response.ErrUnauthorized()
		}

		c.Locals("user", claims.Subject)

		return c.Next()
	}

	return response.ErrUnauthorized()
}
