package repository

import (
	"context"

	"gorm.io/gorm"
	"shellrean.id/back-end/domain"
)

type reviewRepository struct {
	db *gorm.DB
}

func NewReview(db *gorm.DB) domain.ReviewsRepository {
	return &reviewRepository{
		db: db,
	}
}
func (r *reviewRepository) Create(ctx context.Context, c *domain.Reviews) error {
	return r.db.
		WithContext(ctx).
		Create(c).
		Error
}
