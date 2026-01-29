package service

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"shellrean.id/back-end/domain"
	"shellrean.id/back-end/dto"
)

type bookingService struct {
	bookingRepository domain.BookingRepository
}

func NewBookingService(bookingRepository domain.BookingRepository) domain.BookingService {
	return &bookingService{
		bookingRepository: bookingRepository,
	}
}
func (b *bookingService) Index(ctx context.Context) ([]dto.BookingData, error) {
	bookings, err := b.bookingRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	var bookingData []dto.BookingData
	for _, v := range bookings {
		bookingData = append(bookingData, dto.BookingData{
			ID:          v.ID,
			UserID:      v.UserID,
			BookingDate: v.BookingDate,
			StartTime:   v.StartTime,
			EndTime:     v.EndTime,
			Status:      v.Status,
			Notes:       v.Notes,
		})
	}
	return bookingData, nil
}

// Create implements [domain.BookingService].
func (b *bookingService) Create(ctx context.Context, req dto.CreateBookingRequest) error {
	booking := domain.Booking{
		ID:          uuid.NewString(),
		UserID:      req.UserID,
		BookingDate: req.BookingDate,
		StartTime:   req.StartTime,
		EndTime:     req.EndTime,
		Status:      req.Status,
		Notes:       req.Notes,
	}
	return b.bookingRepository.Save(ctx, &booking)
}

func (b *bookingService) Update(ctx context.Context, req dto.UpdateBookingRequest) error {
	persisted, err := b.bookingRepository.FindById(ctx, req.ID)
	if err != nil {
		return err
	}
	if persisted.ID == "" {
		return errors.New("Data booking tidak ditemukan")
	}
	persisted.Notes = req.Notes
	persisted.Status = req.Status
	return b.bookingRepository.Update(ctx, &persisted)
}
func (b *bookingService) Delete(ctx context.Context, id string) error {
	exist, err := b.bookingRepository.FindById(ctx, id)
	if err != nil {
		return err
	}
	if exist.ID == "" {
		return errors.New("Data booking tidak ditemukan")
	}
	return b.bookingRepository.Delete(ctx, id)
}
func (b *bookingService) Show(ctx context.Context, id string) (dto.BookingData, error) {
	persisted, err := b.bookingRepository.FindById(ctx, id)
	if err != nil {
		return dto.BookingData{}, err
	}
	if persisted.ID == "" {
		return dto.BookingData{}, errors.New("data booking tidak ditemukan")
	}
	return dto.BookingData{
		ID:          persisted.ID,
		UserID:      persisted.UserID,
		BookingDate: persisted.BookingDate,
		StartTime:   persisted.StartTime,
		EndTime:     persisted.EndTime,
		Status:      persisted.Status,
		Notes:       persisted.Notes,
	}, nil
}
