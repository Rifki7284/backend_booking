package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
	"shellrean.id/back-end/dto"
)

type Room struct {
	ID            string  `db:"id"`
	Name          string  `db:"name"`
	PropertyID    string  `db:"property_id"`
	Capacity      int     `db:"capacity"`
	PricePerNight float64 `db:"price_per_night"`
	Description   string  `db:"description"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

type RoomRepository interface {
	FindAll(ctx context.Context) ([]Room, error)
	FindById(ctx context.Context, id string) (Room, error)
	Create(ctx context.Context, c *Room) error
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, c *Room) error
}
type RoomService interface {
	Index(ctx context.Context) ([]dto.RoomData, error)
	Create(ctx context.Context, req dto.CreateRoomRequest) error
	Update(ctx context.Context, req dto.UpdateRoomRequest) error
}
