package service

import (
	"context"

	"github.com/google/uuid"
	"shellrean.id/back-end/domain"
	"shellrean.id/back-end/dto"
)

type reviewService struct {
	reviewRepository domain.ReviewsRepository
}

func NewReviewService(rr domain.ReviewsRepository) domain.ReviewsService {
	return &reviewService{
		reviewRepository: rr,
	}
}
func (r *reviewService) Create(ctx context.Context, req dto.CreateReviewRequest) error {
	review := domain.Reviews{
		ID:        uuid.NewString(),
		BookingID: req.BookingID,
		Rating:    req.Rating,
		Comment:   req.Comment,
	}
	return r.reviewRepository.Create(ctx, &review)
}
