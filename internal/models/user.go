package models

import "github.com/google/uuid"

type User struct {
	Id       uuid.UUID `json:"id" db:"id"`
	Username string    `json:"username" db:"username"`
	Password []byte    `json:"-" db:"password"`
}

type RegisterUserParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginUserParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
