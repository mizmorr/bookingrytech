package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/mizmorr/ingrytech/internal/store/model"
)

type Repository interface {
	GetAll(ctx context.Context) ([]*model.Book, error)
	Get(ctx context.Context, id uuid.UUID) (*model.Book, error)
	Delete(ctx context.Context, id uuid.UUID) error
	Update(ctx context.Context, newBookData *model.Book) error
	Create(ctx context.Context, book *model.Book)
}

type BookService struct {
	repo Repository
}

func NewBookService(repo Repository) *BookService {
	return &BookService{
		repo: repo,
	}
}
