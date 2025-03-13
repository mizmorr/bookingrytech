package inmemory

import (
	"context"

	"github.com/google/uuid"
	"github.com/mizmorr/ingrytech/internal/store"
	"github.com/mizmorr/ingrytech/internal/store/model"
)

func (repo *Repo) GetAll(ctx context.Context) ([]*model.Book, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	if len(repo.storage) == 0 {
		return []*model.Book{}, store.ErrBooksNotFound
	}

	return repo.getAll(ctx)
}

func (repo *Repo) getAll(ctx context.Context) ([]*model.Book, error) {
	values := make([]*model.Book, 0, len(repo.storage))

	for _, value := range repo.storage {
		values = append(values, value)
	}

	return values, nil
}

func (repo *Repo) Get(ctx context.Context, id uuid.UUID) (*model.Book, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	val, ok := repo.storage[id]
	if !ok {
		return nil, store.ErrBookNotFound
	}
	return val, nil
}

func (repo *Repo) Delete(ctx context.Context, id uuid.UUID) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	if _, ok := repo.storage[id]; ok {
		delete(repo.storage, id)
		return nil
	}
	return store.ErrBookNotFound
}

func (repo *Repo) Update(ctx context.Context, newBookData *model.Book) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	if _, ok := repo.storage[newBookData.ID]; ok {
		return repo.update(ctx, newBookData)
	}

	return store.ErrBookNotFound
}

func (repo *Repo) update(ctx context.Context, newBookData *model.Book) error {
	val := repo.storage[newBookData.ID]

	if newBookData.Title != undefinedString {
		val.Title = newBookData.Title
	}

	if newBookData.Author != undefinedString {
		val.Author = newBookData.Author
	}

	if newBookData.PublicationYear != undefinedInt {
		val.PublicationYear = newBookData.PublicationYear
	}
	return nil
}

func (repo *Repo) Create(ctx context.Context, book *model.Book) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	repo.storage[book.ID] = book
	return nil
}
