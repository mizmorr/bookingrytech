package model

import "github.com/google/uuid"

type Book struct {
	ID              uuid.UUID
	Title           string
	Author          string
	PublicationYear int8
}
