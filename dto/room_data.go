package dto

type RoomData struct {
	ID            string  `db:"id"`
	Name          string  `db:"name"`
	PropertyID    string  `db:"property_id"`
	Capacity      int     `db:"capacity"`
	PricePerNight float64 `db:"price_per_night"`
	Description   string  `db:"description"`
}
type CreateRoomRequest struct {
	Name          string  `json:"name" validate:"required"`
	PropertyID    string  `json:"property_id" validate:"required"`
	Capacity      int     `json:"capacity" validate:"required,gt=0"`
	PricePerNight float64 `json:"price_per_night" validate:"required,gt=0"`
	Description   string  `json:"description" validate:"required"`
}
type UpdateRoomRequest struct {
	ID            string  `json:"-"`
	Name          string  `json:"name" validate:"required"`
	PropertyID    string  `json:"property_id" validate:"required"`
	Capacity      int     `json:"capacity" validate:"required,gt=0"`
	PricePerNight float64 `json:"price_per_night" validate:"required,gt=0"`
	Description   string  `json:"description" validate:"required"`
}
