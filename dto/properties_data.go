package dto

// dto/properties.go
type PropertiesData struct {
	ID          string     `json:"id"`
	OwnerID     string     `json:"owner_id"`
	Name        string     `json:"name"`
	Address     string     `json:"address"`
	Description string     `json:"description"`
	Rooms       []RoomData `json:"rooms"`
}
type CreatePropertiesRequest struct {
	OwnerID     string `json:"owner_id" validate:"required,uuid4"`
	Name        string `json:"name" validate:"required"`
	Address     string `json:"address" validate:"required"`
	Description string `json:"description" validate:"required"`
}
type UpdatePropertiesRequest struct {
	ID          string `json:"-"`
	OwnerID     string `json:"owner_id" validate:"required,uuid4"`
	Name        string `json:"name" validate:"required"`
	Address     string `json:"address" validate:"required"`
	Description string `json:"description" validate:"required"`
}
