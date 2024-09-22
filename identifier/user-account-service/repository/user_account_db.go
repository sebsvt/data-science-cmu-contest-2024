package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type userAccountRepositoryDB struct {
	db *sqlx.DB
}

func NewUserAccountRepositoryDB(db *sqlx.DB) UserAccountRepository {
	return userAccountRepositoryDB{db: db}
}

// DeleteByUserID implements UserAccountRepository.
func (repo userAccountRepositoryDB) DeleteByUserID(user_id string) error {
	query := "DELETE FROM user_accounts WHERE user_id = $1"

	result, err := repo.db.Exec(query, user_id)
	if err != nil {
		return fmt.Errorf("error deleting user account with ID %s: %v", user_id, err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no account found with ID %s", user_id)
	}
	return nil
}

// FromEmail implements UserAccountRepository.
func (repo userAccountRepositoryDB) FromEmail(email string) (*UserAccount, error) {
	var user_account UserAccount
	query := "SELECT user_id, firstname, lastname, email, date_of_birth, created_at FROM user_accounts WHERE email = $1"
	if err := repo.db.Get(&user_account, query, email); err != nil {
		return nil, err
	}
	return &user_account, nil
}

// FromID implements UserAccountRepository.
func (repo userAccountRepositoryDB) FromID(user_id string) (*UserAccount, error) {
	var user_account UserAccount
	query := "SELECT user_id, firstname, lastname, email, date_of_birth, created_at FROM user_accounts WHERE user_id = $1"
	if err := repo.db.Get(&user_account, query, user_id); err != nil {
		return nil, err
	}
	return &user_account, nil
}

// Save implements UserAccountRepository.
func (repo userAccountRepositoryDB) Save(entity UserAccount) error {
	var existing UserAccount
	query := "SELECT user_id FROM user_accounts WHERE user_id = $1"
	err := repo.db.Get(&existing, query, entity.UserID)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	if existing.UserID != "" {
		updateQuery := `
			UPDATE user_accounts
			SET firstname = $1, lastname = $2, email = $3, date_of_birth = $4
			WHERE user_id = $5
		`
		_, err := repo.db.Exec(updateQuery, entity.FirstName, entity.LastName, entity.Email, entity.DateOfBirth, entity.UserID)
		return err
	} else {
		insertQuery := `
			INSERT INTO user_accounts (user_id, firstname, lastname, email, date_of_birth, created_at)
			VALUES ($1, $2, $3, $4, $5, $6)
		`
		_, err := repo.db.Exec(insertQuery, entity.UserID, entity.FirstName, entity.LastName, entity.Email, entity.DateOfBirth, time.Now())
		return err
	}
}
