package handlers

import (
	"github.com/godofprodev/simple-pass/internal/storage"
)

type Handlers struct {
	store storage.Storage
}

func New(store storage.Storage) *Handlers {
	return &Handlers{
		store: store,
	}
}
