package repository

import (
	"context"

	"gorm.io/gorm"
	"shellrean.id/back-end/domain"
)

type propertiesRepository struct {
	db *gorm.DB
}

func (p *propertiesRepository) FindByIdAndOwner(ctx context.Context, id string, id_owner string) (domain.Properties, error) {
	var result domain.Properties
	err := p.db.
		Where("id = ? AND owner_id = ? AND deleted_at IS NULL", id, id_owner).
		WithContext(ctx).
		Find(&result).
		Error
	return result, err
}

func (p *propertiesRepository) FindByOwner(ctx context.Context, id string) ([]domain.Properties, error) {
	var result []domain.Properties
	err := p.db.
		Preload("Rooms").
		Where("owner_id = ? AND deleted_at IS NULL", id).
		WithContext(ctx).
		Find(&result).
		Error
	return result, err
}

func NewProperties(db *gorm.DB) domain.PropertiesRepository {
	return &propertiesRepository{
		db: db,
	}
}

func (p *propertiesRepository) Create(ctx context.Context, c *domain.Properties) error {
	return p.db.
		WithContext(ctx).
		Create(c).
		Error
}

func (p *propertiesRepository) Delete(ctx context.Context, id string) error {
	return p.db.
		WithContext(ctx).
		Where("id = ?", id).
		Delete(&domain.Properties{}).
		Error
}

func (p *propertiesRepository) FindAll(ctx context.Context) ([]domain.Properties, error) {
	var result []domain.Properties
	err := p.db.
		Preload("Rooms").
		Where("deleted_at IS NULL").
		WithContext(ctx).
		Find(&result).
		Error
	return result, err
}

func (p *propertiesRepository) FindById(ctx context.Context, id string) (domain.Properties, error) {
	var result domain.Properties
	err := p.db.WithContext(ctx).
		Preload("Rooms").
		Where("id = ? AND deleted_at IS NULL", id).
		First(&result).
		Error
	return result, err
}

func (p *propertiesRepository) Update(ctx context.Context, c *domain.Properties) error {
	return p.db.
		Where("id = ?", c.ID).
		Model(&domain.Properties{}).
		WithContext(ctx).
		Updates(c).
		Error
}
