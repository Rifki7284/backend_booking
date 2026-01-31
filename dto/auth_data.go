package dto

type AuthRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
	Name     string `json:"name" validate:"required,min=3,max=100"`
	Phone    string `json:"phone" validate:"required,e164"`
	Role     string `json:"role" validate:"required"`
}

type AuthResponse struct {
	Token string `json:"token"`
}
