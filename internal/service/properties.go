package service

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"shellrean.id/back-end/domain"
	"shellrean.id/back-end/dto"
)

type propertiesService struct {
	propertiesRepository domain.PropertiesRepository
}

func NewPropertiesService(pr domain.PropertiesRepository) domain.PropertiesService {
	return &propertiesService{
		propertiesRepository: pr,
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
	return p.propertiesRepository.Create(ctx, &properties)

}
func (p *propertiesService) Update(ctx context.Context, req dto.UpdatePropertiesRequest, id string) error {
	persisted, err := p.propertiesRepository.FindByIdAndOwner(ctx, req.ID, id)
	if err != nil {
		return err
	}
	if persisted.ID == "" {
		return errors.New("Property not found")
	}
	persisted.Name = req.Name
	persisted.Address = req.Address
	persisted.Description = req.Description
	return p.propertiesRepository.Update(ctx, &persisted)
}
func (p *propertiesService) Delete(ctx context.Context, id string, id_owner string) error {
	exist, err := p.propertiesRepository.FindByIdAndOwner(ctx, id, id_owner)
	if err != nil {
		return err
	}
	if exist.ID == "" {
		return errors.New("Property not found")
	}
	return p.propertiesRepository.Delete(ctx, id)
}

func (p *propertiesService) Index(ctx context.Context) ([]dto.PropertiesData, error) {
	poperties, err := p.propertiesRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	var propertiesData []dto.PropertiesData

	for _, v := range poperties {
		var roomsDto []dto.RoomData
		for _, r := range v.Rooms {
			roomsDto = append(roomsDto, dto.RoomData{
				ID:            r.ID,
				Name:          r.Name,
				PropertyID:    r.PropertyID,
				Capacity:      r.Capacity,
				PricePerNight: r.PricePerNight,
				Description:   r.Description,
			})
		}
		propertiesData = append(propertiesData, dto.PropertiesData{
			ID:          v.ID,
			OwnerID:     v.OwnerID,
			Name:        v.Name,
			Address:     v.Address,
			Description: v.Description,
			Rooms:       roomsDto,
		})
	}
	return propertiesData, nil
}
func (p *propertiesService) IndexByOwner(ctx context.Context, id string) ([]dto.PropertiesData, error) {
	poperties, err := p.propertiesRepository.FindByOwner(ctx, id)
	if err != nil {
		return nil, err
	}
	var propertiesData []dto.PropertiesData

	for _, v := range poperties {
		var roomsDto []dto.RoomData
		for _, r := range v.Rooms {
			roomsDto = append(roomsDto, dto.RoomData{
				ID:            r.ID,
				Name:          r.Name,
				PropertyID:    r.PropertyID,
				Capacity:      r.Capacity,
				PricePerNight: r.PricePerNight,
				Description:   r.Description,
			})
		}
		propertiesData = append(propertiesData, dto.PropertiesData{
			ID:          v.ID,
			OwnerID:     v.OwnerID,
			Name:        v.Name,
			Address:     v.Address,
			Description: v.Description,
			Rooms:       roomsDto,
		})
	}
	return propertiesData, nil
}
func (p *propertiesService) Show(ctx context.Context, id string) (dto.PropertiesData, error) {
	properties, err := p.propertiesRepository.FindById(ctx, id)
	if err != nil {
		return dto.PropertiesData{}, err
	}
	var roomsDto []dto.RoomData
	for _, r := range properties.Rooms {
		roomsDto = append(roomsDto, dto.RoomData{
			ID:            r.ID,
			Name:          r.Name,
			PropertyID:    r.PropertyID,
			Capacity:      r.Capacity,
			PricePerNight: r.PricePerNight,
			Description:   r.Description,
		})
	}

	return dto.PropertiesData{
		ID:          properties.ID,
		OwnerID:     properties.OwnerID,
		Name:        properties.Name,
		Address:     properties.Address,
		Description: properties.Description,
		Rooms:       roomsDto,
	}, nil
}
