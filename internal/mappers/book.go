package mappers

import (
	"github.com/mizmorr/ingrytech/internal/domain"
	"github.com/mizmorr/ingrytech/internal/store/model"
)

func BookToWeb(bookDB *model.Book) *domain.Book {
	return &domain.Book{
		Author:          bookDB.Author,
		PublicationYear: bookDB.PublicationYear,
		ID:              bookDB.ID,
		Title:           bookDB.Title,
	}
}

func BookToDB(bookWeb *domain.Book) *model.Book {
	return &model.Book{
		Author:          bookWeb.Author,
		PublicationYear: bookWeb.PublicationYear,
		ID:              bookWeb.ID,
		Title:           bookWeb.Title,
	}
}

func BooksToWeb(books []*model.Book) []*domain.Book {
	var (
		booksWeb = make([]*domain.Book, 0, len(books))
		bookWeb  *domain.Book
	)

	for _, book := range books {
		bookWeb = &domain.Book{
			Author:          book.Author,
			PublicationYear: book.PublicationYear,
			ID:              book.ID,
			Title:           book.Title,
		}
		booksWeb = append(booksWeb, bookWeb)
	}

	return booksWeb
}
