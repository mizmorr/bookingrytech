package domain

import "github.com/google/uuid"

type Book struct {
	ID              uuid.UUID `json:"id"`
	Title           string    `json:"title,omitempty"`
	Author          string    `json:"author,omitempty"`
	PublicationYear int8      `json:"publication_year,omitempty"`
}
