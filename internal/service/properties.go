package service

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"shellrean.id/back-end/domain"
	"shellrean.id/back-end/dto"
)

type propertiesService struct {
	propertiesService domain.PropertiesRepository
}

func NewPropertiesService(pr domain.PropertiesRepository) domain.PropertiesService {
	return &propertiesService{
		propertiesService: pr,
	}
}

func (p *propertiesService) Create(ctx context.Context, req dto.CreatePropertiesRequest) error {
	properties := domain.Properties{
		ID:          uuid.NewString(),
		OwnerID:     req.OwnerID,
		Name:        req.Name,
		Address:     req.Address,
		Description: req.Description,
	}
	return p.propertiesService.Create(ctx, &properties)

}
func (p *propertiesService) Delete(ctx context.Context, id string) error {
	exist, err := p.propertiesService.FindById(ctx, id)
	if err != nil {
		return err
	}
	if exist.ID == "" {
		return errors.New("Property not found")
	}
	return p.propertiesService.Delete(ctx, id)
}

// Index implements [domain.PropertiesService].
func (p *propertiesService) Index(ctx context.Context) ([]dto.PropertiesData, error) {
	poperties, err := p.propertiesService.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	var propertiesData []dto.PropertiesData
	for _, v := range poperties {
		propertiesData = append(propertiesData, dto.PropertiesData{
			ID:          v.ID,
			OwnerID:     v.OwnerID,
			Name:        v.Name,
			Address:     v.Address,
			Description: v.Description,
		})
	}
	return propertiesData, nil
}
func (p *propertiesService) Show(ctx context.Context, id string) (dto.PropertiesData, error) {
	properties, err := p.propertiesService.FindById(ctx, id)
	if err != nil {
		return dto.PropertiesData{}, err
	}
	return dto.PropertiesData{
		ID:          properties.ID,
		OwnerID:     properties.OwnerID,
		Name:        properties.Name,
		Address:     properties.Address,
		Description: properties.Description,
	}, nil
}

// Update implements [domain.PropertiesService].
func (p *propertiesService) Update(ctx context.Context, req dto.UpdatePropertiesRequest) error {
	persisted, err := p.propertiesService.FindById(ctx, req.ID)
	if err != nil {
		return err
	}
	if persisted.ID == "" {
		return errors.New("Property not found")
	}
	persisted.Name = req.Name
	persisted.Address = req.Address
	persisted.Description = req.Description
	persisted.OwnerID = req.OwnerID
	return p.propertiesService.Update(ctx, &persisted)
}
