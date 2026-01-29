package service

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"shellrean.id/back-end/domain"
	"shellrean.id/back-end/dto"
	"shellrean.id/back-end/internal/config"
)

type authService struct {
	conf           *config.Config
	userRepository domain.UserRepository
}

func NewAuthService(cnf *config.Config, userRepository domain.UserRepository) domain.AuthService {
	return authService{
		conf:           cnf,
		userRepository: userRepository,
	}
}
func (a authService) Login(ctx context.Context, req dto.AuthRequest) (dto.AuthResponse, error) {
	user, err := a.userRepository.FindByEmail(ctx, req.Email)
	if err != nil {
		return dto.AuthResponse{}, err
	}
	if user.ID == "" {
		return dto.AuthResponse{}, errors.New("Authentication failed: user not found")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return dto.AuthResponse{}, errors.New("Authentication failed: invalid credentials")
	}
	claim := jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Duration(a.conf.JWT.Exp) * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenStr, err := token.SignedString([]byte(a.conf.JWT.Key))
	if err != nil {
		return dto.AuthResponse{}, errors.New("Authentication failed: invalid credentials")
	}
	return dto.AuthResponse{
		Token: tokenStr,
	}, nil
}
