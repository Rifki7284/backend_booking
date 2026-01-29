package repository

import (
	"context"

	"gorm.io/gorm"
	"shellrean.id/back-end/domain"
)

type roomRepository struct {
	db *gorm.DB
}

func NewRoom(db *gorm.DB) *roomRepository {
	return &roomRepository{
		db: db,
	}
}

func (r *roomRepository) FindAll(ctx context.Context) ([]domain.Room, error) {
	var rooms []domain.Room
	err := r.db.WithContext(ctx).Where("deleted_at IS NULL").Find(&rooms).Error
	return rooms, err
}
func (r *roomRepository) Create(
	ctx context.Context,
	c *domain.Room,
) error {
	return r.db.
		WithContext(ctx).
		Create(c).
		Error
}

func (r *roomRepository) Delete(ctx context.Context, id string) error {
	return r.db.
		WithContext(ctx).
		Where("id = ?", id).
		Delete(&domain.Room{}).
		Error
}
