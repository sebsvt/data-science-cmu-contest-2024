package repository

import (
	"time"

	"github.com/google/uuid"
)

type UserProfile struct {
	UserID      uuid.UUID `db:"user_id"`
	FirstName   string    `db:"first_name"`
	LastName    string    `db:"last_name"`
	PhoneNumber string    `db:"phone_number"`
	Address     string    `db:"address"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

type UserProfileRepository interface {
	FromUserID(user_id uuid.UUID) (*UserProfile, error)
	Create(entitiy UserProfile) error
	Update(entity UserProfile) error
	Delete(user_id uuid.UUID) error
}
