package dto

import "time"

type BookingData struct {
	ID           string    `db:"id"`
	UserID       string    `db:"user_id"`
	RoomID       string    `db:"room_id"`
	CheckInDate  time.Time `db:"check_in_date"`
	CheckOutDate time.Time `db:"check_out_date"`
	Nights       int       `db:"nights"`
	Status       string    `db:"status"`
	Notes        string    `db:"notes"`
}
type CreateBookingRequest struct {
	UserID       string    `json:"user_id" validate:"required,uuid"`
	RoomID       string    `json:"room_id" validate:"required,uuid"`
	CheckInDate  time.Time `json:"check_in_date" validate:"required"`
	CheckOutDate time.Time `json:"check_out_date" validate:"required,gtfield=CheckInDate"`
	Notes        string    `json:"notes"`
}

type UpdateBookingRequest struct {
	ID     string `json:"-"`
	Status string `db:"status" validate:"required,oneof=scheduled completed cancelled cancelled"`
	Notes  string `db:"notes"`
}
