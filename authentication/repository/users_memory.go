package repository

import (
	"database/sql"
	"errors"
	"sync"
	"time"

	"github.com/google/uuid"
)

// userRepositoryInMemory is an in-memory implementation of UserRepository.
type userRepositoryInMemory struct {
	users []User
	sync.Mutex
}

// NewUserRepositoryInMemory creates a new instance of userRepositoryInMemory.
func NewUserRepositoryInMemory() UserRepository {
	return &userRepositoryInMemory{}
}

// Create adds a new user to the in-memory store.
func (u *userRepositoryInMemory) Create(entity User) error {
	u.Lock()
	defer u.Unlock()

	// Check if the user already exists by email
	for _, user := range u.users {
		if user.Email == entity.Email {
			return errors.New("user with this email already exists")
		}
	}

	u.users = append(u.users, entity)
	return nil
}

// DeleteByUserID removes a user by their ID from the in-memory store.
func (u *userRepositoryInMemory) DeleteByUserID(userID uuid.UUID) error {
	u.Lock()
	defer u.Unlock()

	for i, user := range u.users {
		if user.UserID == userID {
			u.users = append(u.users[:i], u.users[i+1:]...)
			return nil
		}
	}

	return errors.New("user not found")
}

// FromEmail retrieves a user by their email address.
func (u *userRepositoryInMemory) FromEmail(email string) (*User, error) {
	u.Lock()
	defer u.Unlock()

	for _, user := range u.users {
		if user.Email == email {
			return &user, nil
		}
	}

	return nil, sql.ErrNoRows // Email not found
}

// FromUserID retrieves a user by their ID.
func (u *userRepositoryInMemory) FromUserID(userID uuid.UUID) (*User, error) {
	u.Lock()
	defer u.Unlock()

	for _, user := range u.users {
		if user.UserID == userID {
			return &user, nil
		}
	}

	return nil, sql.ErrNoRows // User ID not found
}

// Update updates an existing user in the in-memory store.
func (u *userRepositoryInMemory) Update(entity User) error {
	u.Lock()
	defer u.Unlock()

	for i, user := range u.users {
		if user.UserID == entity.UserID {
			u.users[i] = entity
			return nil
		}
	}

	return sql.ErrNoRows
}

func (u *userRepositoryInMemory) UpdateRefreshToken(userID uuid.UUID, refreshToken string, expiry time.Time) error {
	u.Lock()
	defer u.Unlock()

	for i, user := range u.users {
		if user.UserID == userID {
			u.users[i].RefreshToken = refreshToken
			u.users[i].RefreshTokenExpiry = expiry
			return nil
		}
	}

	return sql.ErrNoRows // User ID not found
}
