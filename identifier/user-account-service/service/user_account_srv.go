package service

import (
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/standardise-software/user-account-service/repository"
)

type userAccountService struct {
	user_account_repo repository.UserAccountRepository
}

func NewUserAccountService(user_account_repo repository.UserAccountRepository) UserAccountService {
	return userAccountService{user_account_repo: user_account_repo}
}

// CreateNewUserAccount implements UserAccountService.
func (srv userAccountService) CreateNewUserAccount(entity CreatedAccount) (string, error) {
	id := uuid.New().String()
	user, err := srv.user_account_repo.FromEmail(entity.Email)
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}
	if user != nil {
		return "", errors.New("user email already exist")
	}
	new_user := repository.UserAccount{
		UserID:      id,
		FirstName:   entity.FirstName,
		LastName:    entity.LastName,
		Email:       entity.Email,
		DateOfBirth: entity.DateOfBirth,
		CreatedAt:   time.Now(),
	}
	if err := srv.user_account_repo.Save(new_user); err != nil {
		return "", nil
	}
	return id, nil
}

// DeleteUserAccount implements UserAccountService.
func (srv userAccountService) DeleteUserAccount(user_id string) error {
	_, err := srv.user_account_repo.FromID(user_id)
	if err != nil {
		return err
	}
	err = srv.DeleteUserAccount(user_id)
	return err
}

// GetAccountFromEmail implements UserAccountService.
func (srv userAccountService) GetAccountFromEmail(email string) (*UserAccountResponse, error) {
	user_account, err := srv.user_account_repo.FromEmail(email)
	if err != nil {
		return nil, err
	}
	return &UserAccountResponse{
		UserID:      user_account.UserID,
		FirstName:   user_account.FirstName,
		LastName:    user_account.LastName,
		Email:       user_account.Email,
		DateOfBirth: user_account.DateOfBirth,
		CreatedAt:   user_account.CreatedAt,
	}, nil
}

// GetAccountFromID implements UserAccountService.
func (srv userAccountService) GetAccountFromID(user_id string) (*UserAccountResponse, error) {
	user_account, err := srv.user_account_repo.FromID(user_id)
	if err != nil {
		return nil, err
	}
	return &UserAccountResponse{
		UserID:      user_account.UserID,
		FirstName:   user_account.FirstName,
		LastName:    user_account.LastName,
		Email:       user_account.Email,
		DateOfBirth: user_account.DateOfBirth,
		CreatedAt:   user_account.CreatedAt,
	}, nil
}
