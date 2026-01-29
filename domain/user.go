package domain

import "context"

type User struct {
	ID       string `db:"id"`
	Name     string `db:"name"`
	Email    string `db:"email"`
	Password string `db:"password"`
	Role     string `db:"role"`
	Phone    string `db:"phone"`
}
type UserRepository interface {
	FindByEmail(ctx context.Context, email string) (User, error)
}
