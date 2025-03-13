package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/mizmorr/ingrytech/internal/store"
	"github.com/mizmorr/ingrytech/internal/store/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (r *PostgresRepo) GetAll(ctx context.Context) ([]*model.Book, error) {
	var books []*model.Book
	if err := r.db.WithContext(ctx).Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func (r *PostgresRepo) Get(ctx context.Context, id uuid.UUID) (*model.Book, error) {
	var book model.Book
	if err := r.db.WithContext(ctx).First(&book, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, store.ErrBookNotFound
		}
		return nil, err
	}
	return &book, nil
}

func (r *PostgresRepo) Create(ctx context.Context, book *model.Book) error {
	return r.db.WithContext(ctx).Create(book).Error
}

func (r *PostgresRepo) Update(ctx context.Context, newBookData *model.Book) error {
	result := r.db.WithContext(ctx).Model(&model.Book{}).Where("id = ?", newBookData.ID).Updates(newBookData)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *PostgresRepo) Delete(ctx context.Context, id uuid.UUID) error {
	result := r.db.WithContext(ctx).Delete(&model.Book{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
