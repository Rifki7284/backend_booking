package dto

type PropertiesData struct {
	ID          string `db:"id"`
	OwnerID     string `db:"owner_id"`
	Name        string `db:"name"`
	Address     string `db:"address"`
	Description string `db:"description"`
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
