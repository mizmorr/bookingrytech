package delivery

import "errors"

var (
	ErrInvalidUUID    = errors.New("invalid UUID format")
	ErrBookNotFound   = errors.New("book not found")
	ErrBooksNotFound  = errors.New("books not found")
	ErrInternal       = errors.New("internal server error")
	ErrBadRequestBody = errors.New("request body incorrect")
)
