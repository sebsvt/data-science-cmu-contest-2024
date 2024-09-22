package service

import "time"

type CreatedAccount struct {
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	DateOfBirth string `json:"date_of_birth"`
}

type UserAccountResponse struct {
	UserID      string    `json:"user_id"`
	FirstName   string    `json:"firstname"`
	LastName    string    `json:"lastname"`
	Email       string    `json:"email"`
	DateOfBirth string    `json:"date_of_birth"`
	CreatedAt   time.Time `json:"created_at"`
}

type UserAccountService interface {
	GetAccountFromEmail(email string) (*UserAccountResponse, error)
	GetAccountFromID(user_id string) (*UserAccountResponse, error)
	CreateNewUserAccount(entity CreatedAccount) (string, error)
	DeleteUserAccount(user_id string) error
}
