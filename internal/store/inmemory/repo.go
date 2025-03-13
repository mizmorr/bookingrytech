package inmemory

import (
	"sync"

	"github.com/google/uuid"
	"github.com/mizmorr/ingrytech/internal/store/model"
)

type Repo struct {
	storage map[uuid.UUID]*model.Book
	mu      sync.RWMutex
}

func NewInMemoryRepo() *Repo {
	return &Repo{
		storage: make(map[uuid.UUID]*model.Book),
		mu:      sync.RWMutex{},
	}
}
