package services

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/sebsvt/prototype/logs"
	"github.com/sebsvt/prototype/repository"
)

type userProfileService struct {
	user_profile_repo repository.UserProfileRepository
}

func NewUserProfileService(user_profile_repo repository.UserProfileRepository) UserProfileService {
	return userProfileService{user_profile_repo: user_profile_repo}
}

// CreateNewUserProfile implements UserProfileService.
func (srv userProfileService) CreateNewUserProfile(user_id uuid.UUID, entity CreateUserProfileModel) error {
	user_profile, err := srv.user_profile_repo.FromUserID(user_id)
	if err != nil && err != sql.ErrNoRows {
		logs.Error(err)
		return ErrUnexpectedError
	}
	if user_profile != nil {
		return ErrUserAlreadyHasProfile
	}
	err = srv.user_profile_repo.Create(repository.UserProfile{
		UserID:      user_id,
		FirstName:   entity.FirstName,
		LastName:    entity.LastName,
		Address:     entity.Address,
		PhoneNumber: entity.PhoneNumber,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	})
	return err
}

// GetUserProfile implements UserProfileService.
func (srv userProfileService) GetUserProfile(user_id uuid.UUID) (*UserProfileModel, error) {
	user_profile, err := srv.user_profile_repo.FromUserID(user_id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrUserProflieNotFound
		}
		logs.Error(err)
		return nil, ErrUnexpectedError
	}
	return &UserProfileModel{
		UserID:      user_profile.UserID,
		FirstName:   user_profile.FirstName,
		LastName:    user_profile.LastName,
		PhoneNumber: user_profile.PhoneNumber,
		Address:     user_profile.Address,
		CreatedAt:   user_profile.CreatedAt,
		UpdatedAt:   user_profile.UpdatedAt,
	}, nil
}

// UpdateUserProfile implements UserProfileService.
func (srv userProfileService) UpdateUserProfile(user_id uuid.UUID, entity UpdateUserProfileModel) error {
	err := srv.user_profile_repo.Update(repository.UserProfile{
		UserID:      user_id,
		FirstName:   entity.FirstName,
		LastName:    entity.LastName,
		PhoneNumber: entity.PhoneNumber,
		Address:     entity.Address,
		UpdatedAt:   time.Now(),
	})
	return err
}
