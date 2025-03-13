package domain

import "github.com/google/uuid"

type Book struct {
	ID              uuid.UUID `json:"id" validate:"required"`
	Title           string    `json:"title"`
	Author          string    `json:"author"`
	PublicationYear uint16    `json:"publication_year"`
}
