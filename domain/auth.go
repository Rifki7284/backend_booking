package domain

import (
	"context"

	"shellrean.id/back-end/dto"
)

type AuthService interface {
	Login(ctx context.Context, req dto.AuthRequest) (dto.AuthResponse, error)
}
