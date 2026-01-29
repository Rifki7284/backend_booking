package dto

type BookingData struct {
	ID           string `db:"id"`
	UserID       string `db:"user_id"`
	RoomID       string `db:"room_id"`
	CheckInDate  string `db:"check_in_date"`
	CheckOutDate string `db:"check_out_date"`
	Nights       int    `db:"nights"`
	Status       string `db:"status"`
	Notes        string `db:"notes"`
}
type CreateBookingRequest struct {
	UserID       string `json:"user_id" validate:"required,uuid4"`
	RoomID       string `json:"room_id" validate:"required,uuid4"`
	CheckInDate  string `json:"check_in_date" validate:"required,datetime=2006-01-02"`
	CheckOutDate string `json:"check_out_date" validate:"required,datetime=2006-01-02,gtfield=CheckInDate"`
	Notes        string `db:"notes"`
}
type UpdateBookingRequest struct {
	ID     string `json:"-"`
	Status string `db:"status" validate:"required,oneof=scheduled completed cancelled cancelled"`
	Notes  string `db:"notes"`
}
