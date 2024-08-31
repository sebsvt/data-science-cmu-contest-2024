package repository

import (
	"errors"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type userProfileRepository struct {
	db *sqlx.DB
}

func NewUserProfileRepository(db *sqlx.DB) UserProfileRepository {
	return userProfileRepository{db: db}
}

// Create implements UserProfileRepository.
func (repo userProfileRepository) Create(entity UserProfile) error {
	query := `INSERT INTO user_profiles (user_id, first_name, last_name, phone_number, address, created_at, updated_at)
	          VALUES (:user_id, :first_name, :last_name, :phone_number, :address, :created_at, :updated_at)`
	_, err := repo.db.NamedExec(query, entity)
	return err
}

// Delete implements UserProfileRepository.
func (repo userProfileRepository) Delete(user_id uuid.UUID) error {
	query := `DELETE FROM user_profiles WHERE user_id = $1`
	result, err := repo.db.Exec(query, user_id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("no rows were deleted")
	}
	return nil
}

// FromUserID implements UserProfileRepository.
func (repo userProfileRepository) FromUserID(user_id uuid.UUID) (*UserProfile, error) {
	var profile UserProfile
	query := `SELECT * FROM user_profiles WHERE user_id = $1`
	err := repo.db.Get(&profile, query, user_id)
	if err != nil {
		return nil, err
	}
	return &profile, nil
}

// Update implements UserProfileRepository.
func (repo userProfileRepository) Update(entity UserProfile) error {
	query := `UPDATE user_profiles SET first_name = :first_name, last_name = :last_name, phone_number = :phone_number, address = :address, updated_at = :updated_at WHERE user_id = :user_id`
	_, err := repo.db.NamedExec(query, entity)
	return err
}
