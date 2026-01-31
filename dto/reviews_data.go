package dto

type CreateReviewRequest struct {
	BookingID string  `json:"booking_id" validate:"required"`
	Rating    float32 `json:"rating" validate:"required"`
	Comment   string  `json:"comment" validate:"required"`
}
