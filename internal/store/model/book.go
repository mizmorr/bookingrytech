package model

import "github.com/google/uuid"

type Book struct {
	ID              uuid.UUID
	Title           string
	Author          string
	PublicationYear uint16
}

func (Book) TableName() string {
	return "books"
}
