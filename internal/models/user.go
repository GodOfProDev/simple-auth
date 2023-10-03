package models

import "github.com/google/uuid"

type User struct {
	Id       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Password []byte    `json:"-"`
}

type RegisterUserParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginUserParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
