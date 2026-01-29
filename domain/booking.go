package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
	"shellrean.id/back-end/dto"
)

type Booking struct {
	ID           string `db:"id"`
	UserID       string `db:"user_id"`
	RoomID       string `db:"room_id"`
	CheckInDate  string `db:"check_in_date"`
	CheckOutDate string `db:"check_out_date"`
	Nights       int    `db:"nights"`
	Status       string `db:"status"`
	Notes        string `db:"notes"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
type BookingRepository interface {
	FindAll(ctx context.Context) ([]Booking, error)
	FindById(ctx context.Context, id string) (Booking, error)
	Create(ctx context.Context, c *Booking) error
	Update(ctx context.Context, c *Booking) error
	Delete(ctx context.Context, id string) error
}
type BookingService interface {
	Index(ctx context.Context) ([]dto.BookingData, error)
	Create(ctx context.Context, req dto.CreateBookingRequest) error
	Update(ctx context.Context, req dto.UpdateBookingRequest) error
	Delete(ctx context.Context, id string) error
	Show(ctx context.Context, id string) (dto.BookingData, error)
}
