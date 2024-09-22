package repository

import "time"

type UserAccount struct {
	UserID      string    `db:"user_id"`
	FirstName   string    `db:"firstname"`
	LastName    string    `db:"lastname"`
	Email       string    `db:"email"`
	DateOfBirth string    `db:"date_of_birth"`
	CreatedAt   time.Time `db:"created_at"`
}

type UserAccountRepository interface {
	Save(entity UserAccount) error
	FromID(user_id string) (*UserAccount, error)
	FromEmail(email string) (*UserAccount, error)
	DeleteByUserID(user_id string) error
}
