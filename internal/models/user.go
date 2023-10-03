package models

import "github.com/google/uuid"

type User struct {
	Id       uuid.UUID
	Username string
	Password []byte
}

type CreateUserParams struct {
	User     string
	Password string
}
