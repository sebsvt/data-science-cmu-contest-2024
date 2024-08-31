package services

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrUserProflieNotFound   = errors.New("user's profile not found")
	ErrUserAlreadyHasProfile = errors.New("user already has a profile")
)

type CreateUserProfileModel struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
}

type UpdateUserProfileModel struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
}

type UserProfileModel struct {
	UserID      uuid.UUID `json:"user_id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	PhoneNumber string    `json:"phone_number"`
	Address     string    `json:"address"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UserProfileService interface {
	CreateNewUserProfile(user_id uuid.UUID, entity CreateUserProfileModel) error
	GetUserProfile(user_id uuid.UUID) (*UserProfileModel, error)
	UpdateUserProfile(user_id uuid.UUID, entity UpdateUserProfileModel) error
}
