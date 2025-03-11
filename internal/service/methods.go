package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/mizmorr/ingrytech/internal/domain"
	"github.com/mizmorr/ingrytech/internal/mappers"
)

func (svc *BookService) Get(ctx context.Context, id uuid.UUID) (*domain.Book, error) {
	bookDB, err := svc.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return mappers.BookToWeb(bookDB), nil
}

func (svc *BookService) Create(ctx context.Context, book *domain.Book) {
	bookDB := mappers.BookToDB(book)

	svc.repo.Create(ctx, bookDB)
}

func (svc *BookService) Update(ctx context.Context, book *domain.Book) error {
	bookDB := mappers.BookToDB(book)

	return svc.repo.Update(ctx, bookDB)
}

func (svc *BookService) GetAll(ctx context.Context) ([]*domain.Book, error) {
	books, err := svc.repo.GetAll(ctx)
	if err != nil {
		return []*domain.Book{}, err
	}
	booksToWeb := mappers.BooksToWeb(books)

	return booksToWeb, err
}

func (svc *BookService) Delete(ctx context.Context, id uuid.UUID) error {
	return svc.repo.Delete(ctx, id)
}
