package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
	"shellrean.id/back-end/dto"
)

type Properties struct {
	ID          string `db:"id"`
	OwnerID     string `db:"owner_id"`
	Name        string `db:"name"`
	Address     string `db:"address"`
	Description string `db:"description"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
type PropertiesRepository interface {
	FindAll(ctx context.Context) ([]Properties, error)
	FindById(ctx context.Context, id string) (Properties, error)
	Create(ctx context.Context, c *Properties) error
	Update(ctx context.Context, c *Properties) error
	Delete(ctx context.Context, id string) error
}

type PropertiesService interface {
	Index(ctx context.Context) ([]dto.PropertiesData, error)
	Create(ctx context.Context, req dto.CreatePropertiesRequest) error
	Update(ctx context.Context, req dto.UpdatePropertiesRequest) error
	Delete(ctx context.Context, id string) error
	Show(ctx context.Context, id string) (dto.PropertiesData, error)
}
