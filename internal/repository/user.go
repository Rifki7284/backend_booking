package repository

import (
	"context"

	"gorm.io/gorm"
	"shellrean.id/back-end/domain"
)

type UserRepository struct {
	db *gorm.DB
}

func (u *UserRepository) Save(
	ctx context.Context,
	user *domain.User,
) (domain.User, error) {

	if err := u.db.WithContext(ctx).Create(user).Error; err != nil {
		return domain.User{}, err
	}

	return *user, nil
}

func NewUser(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) FindByEmail(ctx context.Context, email string) (usr domain.User, err error) {
	var user domain.User
	err = u.db.WithContext(ctx).Where("email = ? AND deleted_at IS NULL", email).First(&user).Error
	return user, err
}
