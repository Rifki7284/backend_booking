package service

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"shellrean.id/back-end/domain"
	"shellrean.id/back-end/dto"
)

type roomService struct {
	roomRepository domain.RoomRepository
}

func NewRoomService(roomRepository domain.RoomRepository) domain.RoomService {
	return &roomService{
		roomRepository: roomRepository,
	}
}

func (r roomService) Index(ctx context.Context) ([]dto.RoomData, error) {
	rooms, err := r.roomRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	var roomData []dto.RoomData
	for _, v := range rooms {
		roomData = append(roomData, dto.RoomData{
			ID:            v.ID,
			Name:          v.Name,
			Capacity:      v.Capacity,
			PricePerNight: v.PricePerNight,
			Description:   v.Description,
		})
	}
	return roomData, nil
}
func (r roomService) Create(ctx context.Context, req dto.CreateRoomRequest) error {
	//userID, ok := util.GetUserID(ctx)
	room := domain.Room{
		ID: uuid.NewString(),
		//OwnerID:       userID,
		Name:          req.Name,
		Capacity:      req.Capacity,
		PricePerNight: req.PricePerNight,
		Description:   req.Description,
		PropertyID:    req.PropertyID,
	}
	return r.roomRepository.Create(ctx, &room)
}
func (r *roomService) Update(ctx context.Context, req dto.UpdateRoomRequest) error {
	persisted, err := r.roomRepository.FindById(ctx, req.ID)
	if err != nil {
		return err
	}
	if persisted.ID == "" {
		return errors.New("Room not found")
	}
	persisted.Capacity = req.Capacity
	persisted.Name = req.Name
	persisted.PropertyID = req.PropertyID
	persisted.PricePerNight = req.PricePerNight
	persisted.Description = req.Description
	return r.roomRepository.Update(ctx, &persisted)
}
