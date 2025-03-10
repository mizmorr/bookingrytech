package store

import (
	"sync"

	"github.com/google/uuid"
	"github.com/mizmorr/ingrytech/internal/store/model"
)

type Repo struct {
	storage map[uuid.UUID]*model.Book
	mu      sync.RWMutex
}
