package dto

type BookingData struct {
	ID          string `db:"id"`
	UserID      string `db:"user_id"`
	BookingDate string `db:"booking_date"`
	StartTime   string `db:"start_time"`
	EndTime     string `db:"end_time"`
	Status      string `db:"status"`
	Notes       string `db:"notes"`
}

type CreateBookingRequest struct {
	UserID      string `db:"user_id" validate:"required,gt=0"`
	BookingDate string `db:"booking_date" validate:"required,datetime=2006-01-02"`
	StartTime   string `validate:"required,datetime=15:04"`
	EndTime     string `validate:"required,datetime=15:04"`
	Status      string `db:"status" validate:"required,oneof=scheduled completed cancelled"`
	Notes       string `db:"notes"`
}
type UpdateBookingRequest struct {
	ID     string `json:"-"`
	Status string `db:"status" validate:"required,oneof=scheduled completed cancelled cancelled"`
	Notes  string `db:"notes"`
}
