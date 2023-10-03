package storage

import "github.com/godofprodev/simple-pass/internal/models"

type Storage interface {
	CreateUser(user *models.User) error
}
