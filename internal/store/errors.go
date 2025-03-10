package store

import "errors"

var (
	ErrBookNotFound  = errors.New("book not found")
	ErrBooksNotFound = errors.New("books not found")
)
