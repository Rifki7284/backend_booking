package repository

import (
	"context"

	"gorm.io/gorm"
	"shellrean.id/back-end/domain"
)

type bookingRepository struct {
	db *gorm.DB
}

func NewBooking(db *gorm.DB) domain.BookingRepository {
	return &bookingRepository{
		db: db,
	}
}

func (b *bookingRepository) Delete(ctx context.Context, id string) error {
	return b.db.
		WithContext(ctx).
		Where("id = ?", id).
		Delete(&domain.Booking{}).
		Error
}

func (b *bookingRepository) FindAll(
	ctx context.Context,
) ([]domain.Booking, error) {

	var result []domain.Booking

	err := b.db.
		WithContext(ctx).
		Where("deleted_at IS NULL").
		Find(&result).
		Error

	return result, err
}

func (b *bookingRepository) Create(
	ctx context.Context,
	c *domain.Booking,
) error {

	return b.db.
		WithContext(ctx).
		Create(c).
		Error
}

func (b *bookingRepository) Update(
	ctx context.Context,
	c *domain.Booking,
) error {

	return b.db.
		WithContext(ctx).
		Model(&domain.Booking{}).
		Where("id = ?", c.ID).
		Updates(c).
		Error
}

func (b *bookingRepository) FindById(
	ctx context.Context,
	id string,
) (domain.Booking, error) {

	var booking domain.Booking

	err := b.db.
		WithContext(ctx).
		Where("id = ? AND deleted_at IS NULL", id).
		First(&booking).
		Error

	return booking, err
}
func (b *bookingRepository) FindByUser(
	ctx context.Context,
	id string,
	id_user string,
) (domain.Booking, error) {

	var booking domain.Booking

	err := b.db.
		WithContext(ctx).
		Where("id = ? AND user_id = ? AND deleted_at IS NULL", id, id_user).
		First(&booking).
		Error

	return booking, err
}
