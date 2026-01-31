package domain

import (
	"context"
	"time"

	"shellrean.id/back-end/dto"
)

type Reviews struct {
	ID        string  `db:"id" gorm:"primaryKey"`
	BookingID string  `db:"booking_id"`
	Booking   Booking `gorm:"foreignKey:BookingID"`
	Rating    float32 `db:"rating"`
	Comment   string  `db:"comment"`
	CreatedAt time.Time
}

func (Reviews) TableName() string {
	return "room_reviews"
}

type ReviewsRepository interface {
	Create(ctx context.Context, c *Reviews) error
}

type ReviewsService interface {
	Create(ctx context.Context, req dto.CreateReviewRequest) error
}
