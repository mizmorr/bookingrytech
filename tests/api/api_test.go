package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/mizmorr/ingrytech/internal/domain"
	"github.com/mizmorr/ingrytech/internal/store/model"
	"github.com/stretchr/testify/assert"
)

const apiBaseURL = "http://localhost:8080/api/v1/books"

var curID = uuid.New()

func TestCreateBook(t *testing.T) {
	newBook := &domain.Book{
		ID:              curID,
		Title:           "New Book",
		Author:          "Author Name",
		PublicationYear: 2024,
	}

	body, err := json.Marshal(newBook)
	if err != nil {
		t.Fatalf("Error marshaling book: %v", err)
	}

	req, err := http.NewRequest("POST", apiBaseURL+"/create", bytes.NewReader(body))
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Error when calling endpoint: %v", err)
	}

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestGetBooks(t *testing.T) {
	req, err := http.NewRequest("GET", apiBaseURL+"/", nil)
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Error when calling endpoint: %v", err)
	}

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	var books []*domain.Book
	err = json.NewDecoder(resp.Body).Decode(&books)
	if err != nil {
		t.Fatalf("Error decoding response: %v", err)
	}

	assert.NotNil(t, books)
}

func TestGetBookByID(t *testing.T) {
	bookID := curID
	req, err := http.NewRequest("GET", apiBaseURL+"/"+bookID.String(), nil)
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Error when calling endpoint: %v", err)
	}

	if bookID.String() != uuid.Nil.String() {
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var book domain.Book
		err = json.NewDecoder(resp.Body).Decode(&book)
		if err != nil {
			t.Fatalf("Error decoding response: %v", err)
		}

		assert.Equal(t, bookID.String(), book.ID.String())
	} else {
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	}
}

func TestUpdateBook(t *testing.T) {
	bookID := curID
	updatedBook := model.Book{
		ID:              bookID,
		Title:           "Updated Book Title",
		Author:          "Updated Author",
		PublicationYear: 2026,
	}

	body, err := json.Marshal(updatedBook)
	if err != nil {
		t.Fatalf("Error marshaling book: %v", err)
	}

	req, err := http.NewRequest("POST", apiBaseURL+"/update", bytes.NewReader(body))
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Error when calling endpoint: %v", err)
	}

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	req, err = http.NewRequest("GET", apiBaseURL+"/"+bookID.String(), nil)
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}

	resp, err = client.Do(req)
	if err != nil {
		t.Fatalf("Error when calling endpoint: %v", err)
	}

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var updated model.Book

	err = json.NewDecoder(resp.Body).Decode(&updated)
	if err != nil {
		t.Fatalf("Error decoding response: %v", err)
	}

	assert.Equal(t, updatedBook.Title, updated.Title)
	assert.Equal(t, updatedBook.Author, updated.Author)
}

func TestDeleteBook(t *testing.T) {
	bookID := curID
	req, err := http.NewRequest("DELETE", apiBaseURL+"/"+bookID.String(), nil)
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Error when calling endpoint: %v", err)
	}

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestGetBookInvalidUUID(t *testing.T) {
	invalidID := "invalid-uuid"
	req, err := http.NewRequest("GET", apiBaseURL+"/"+invalidID, nil)
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Error when calling endpoint: %v", err)
	}

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestCreateBookError(t *testing.T) {
	req, err := http.NewRequest("POST", apiBaseURL+"/create", nil)
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Error when calling endpoint: %v", err)
	}

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestUpdateBookError(t *testing.T) {
	req, err := http.NewRequest("POST", apiBaseURL+"/update", nil)
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Error when calling endpoint: %v", err)
	}

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}
