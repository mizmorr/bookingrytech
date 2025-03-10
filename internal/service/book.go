package service

import (
	"github.com/google/uuid"
	"github.com/mizmorr/ingrytech/internal/store/model"
)

type Repository interface {
	GetAll() ([]*model.Book, error)
	getAll() ([]*model.Book, error)
	Get(id uuid.UUID) (*model.Book, error)
	Delete(id uuid.UUID) error
	Update(newBookData *model.Book) error
	update(newBookData *model.Book) error
	Create(book *model.Book)
}
