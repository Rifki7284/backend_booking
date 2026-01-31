package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
	"shellrean.id/back-end/dto"
)

type Room struct {
	ID            string     `db:"id" gorm:"primaryKey"`
	Name          string     `db:"name"`
	PropertyID    string     `db:"property_id"`
	Property      Properties `gorm:"foreignKey:PropertyID"`
	Capacity      int        `db:"capacity"`
	PricePerNight float64    `db:"price_per_night"`
	Description   string     `db:"description"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

type RoomRepository interface {
	FindAll(ctx context.Context) ([]Room, error)
	FindById(ctx context.Context, id string) (Room, error)
	FindByIdAndOwner(ctx context.Context, id string, id_owner string) (Room, error)
	Create(ctx context.Context, c *Room) error
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, c *Room) error
}
type RoomService interface {
	Index(ctx context.Context) ([]dto.RoomData, error)
	Create(ctx context.Context, req dto.CreateRoomRequest) error
	Update(ctx context.Context, req dto.UpdateRoomRequest, id string) error
	Delete(ctx context.Context, id string, id_owner string) error
	// Show(ctx context.Context, id string) (dto.PropertiesData, error)
}
