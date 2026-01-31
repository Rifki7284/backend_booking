package repository

import (
	"context"

	"gorm.io/gorm"
	"shellrean.id/back-end/domain"
)

type roomRepository struct {
	db *gorm.DB
}

func (r *roomRepository) FindByIdAndOwner(ctx context.Context, id string, id_owner string) (domain.Room, error) {
	var result domain.Room
	err := r.db.
		Preload("Property").
		Joins("JOIN properties ON properties.id = rooms.property_id").
		Where("rooms.id = ? AND properties.owner_id = ?", id, id_owner).
		WithContext(ctx).
		Find(&result).
		Error
	return result, err
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
func (r *roomRepository) FindById(
	ctx context.Context,
	id string,
) (domain.Room, error) {
	var room domain.Room
	err := r.db.
		WithContext(ctx).
		Where("id = ? AND deleted_at IS NULL", id).
		First(&room).
		Error

	return room, err
}

func (r *roomRepository) Delete(ctx context.Context, id string) error {
	return r.db.
		WithContext(ctx).
		Where("id = ?", id).
		Delete(&domain.Room{}).
		Error
}
func (r *roomRepository) Update(
	ctx context.Context,
	c *domain.Room,
) error {
	return r.db.
		WithContext(ctx).
		Model(&domain.Room{}).
		Where("id = ?", c.ID).
		Updates(c).
		Error
}
