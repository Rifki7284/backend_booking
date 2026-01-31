package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
	"shellrean.id/back-end/dto"
)

type Properties struct {
	ID          string `db:"id" gorm:"primaryKey"`
	OwnerID     string `db:"owner_id"`
	Name        string `db:"name"`
	Address     string `db:"address"`
	Description string `db:"description"`
	Rooms       []Room `gorm:"foreignKey:PropertyID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
type PropertiesRepository interface {
	FindAll(ctx context.Context) ([]Properties, error)
	FindByOwner(ctx context.Context, id string) ([]Properties, error)
	FindById(ctx context.Context, id string) (Properties, error)
	FindByIdAndOwner(ctx context.Context, id string, id_owner string) (Properties, error)
	Create(ctx context.Context, c *Properties) error
	Update(ctx context.Context, c *Properties) error
	Delete(ctx context.Context, id string) error
}

type PropertiesService interface {
	Index(ctx context.Context) ([]dto.PropertiesData, error)
	IndexByOwner(ctx context.Context, id string) ([]dto.PropertiesData, error)
	Create(ctx context.Context, req dto.CreatePropertiesRequest) error
	Update(ctx context.Context, req dto.UpdatePropertiesRequest, id string) error
	Delete(ctx context.Context, id string, id_owner string) error
	Show(ctx context.Context, id string) (dto.PropertiesData, error)
}
