package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        string `db:"id"`
	Name      string `db:"name"`
	Email     string `db:"email"`
	Password  string `db:"password"`
	Role      string `db:"role"`
	Phone     string `db:"phone"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
type UserRepository interface {
	FindByEmail(ctx context.Context, email string) (User, error)
	Save(ctx context.Context, user *User) (User, error)
}
