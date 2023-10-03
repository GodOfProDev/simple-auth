package models

import "github.com/google/uuid"

type User struct {
	Id       uuid.UUID
	Username string
	Password []byte
}

type RegisterUserParams struct {
	Username string
	Password string
}

type LoginUserParams struct {
	Username string
	Password string
}
