package services

import (
	"testing"

	"github.com/google/uuid"
	"github.com/sebsvt/prototype/repository"
	"github.com/stretchr/testify/assert"
)

func TestUserCreateCase(t *testing.T) {
	type testCase struct {
		name        string
		email       string
		password    string
		expectedErr error
	}
	testcases := []testCase{
		{
			name:        "invalid email should return err invalid email",
			email:       "thisis@co",
			password:    "Noway0asfiosajf1!",
			expectedErr: ErrInvalidEmail,
		},
		{
			name:        "invalid email should return err invalid email (case 2)",
			email:       "@asfail.co",
			password:    "Noway0a1fkwuso!",
			expectedErr: ErrInvalidEmail,
		},
		{
			name:        "invalid email should return err invalid email (case 3)",
			email:       "asfail.co",
			password:    "Noway0a1fkwuso!",
			expectedErr: ErrInvalidEmail,
		},
		{
			name:        "insecure password should return err insecure password",
			email:       "vithchata@valid.com",
			password:    "admin123",
			expectedErr: ErrInsecurePassword,
		},
		{
			name:        "secure password should return nil error",
			email:       "vithchata@valid.com",
			password:    "!Admin1231sdfqw",
			expectedErr: nil,
		},
		{
			name:        "can create user when email is not already in use",
			email:       "vithchataya.test@gmail.com",
			password:    "NowaywthfAi11oij20a1!",
			expectedErr: nil,
		},
		{
			name:        "can not create user when email already in use",
			email:       "vithchataya.test@gmail.com",
			password:    "Nowaysadfoi10a1!",
			expectedErr: ErrUserEmailAlreadyInUse,
		},
	}
	mockRepo := repository.NewUserRepositoryInMemory()
	userService := NewUserService(mockRepo)

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := userService.CreateNewUser(UserCreatedModel{
				Email:    tc.email,
				Password: tc.password,
			})
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}

func TestUserGetCase(t *testing.T) {
	mockRepo := repository.NewUserRepositoryInMemory()
	userService := NewUserService(mockRepo)

	user_id_1, _ := userService.CreateNewUser(UserCreatedModel{
		Email:    "testing@gmail.com",
		Password: "ayy!oqsAsfdsafsd",
	})
	user_id_2, _ := userService.CreateNewUser(UserCreatedModel{
		Email:    "testing2@gmail.com",
		Password: "Ayyotony1Af!q9",
	})

	type testCase struct {
		name        string
		userID      uuid.UUID
		email       string
		expectedErr error
	}
	testcases := []testCase{
		{
			name:        "can get user form id 1 and got same user",
			userID:      user_id_1,
			email:       "testing@gmail.com",
			expectedErr: nil,
		},
		{
			name:        "can get user form id 2 and got same user",
			userID:      user_id_2,
			email:       "testing2@gmail.com",
			expectedErr: nil,
		},
		{
			name:        "get user not found error if user from an id is not exists",
			userID:      uuid.New(),
			email:       "",
			expectedErr: ErrUserNotFound,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			user, err := userService.FromID(tc.userID)
			if tc.expectedErr != nil {
				assert.Equal(t, tc.expectedErr, err)
				assert.Nil(t, user)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.email, user.Email)
			}
		})
	}
}

func TestDeleteUserCases(t *testing.T) {
	type testCase struct {
		name        string
		userID      uuid.UUID
		expectedErr error
	}

	mockRepo := repository.NewUserRepositoryInMemory()
	userService := NewUserService(mockRepo)

	// Create some users to test deletion
	userID1, _ := userService.CreateNewUser(UserCreatedModel{
		Email:    "user1@example.com",
		Password: "SecurePassword1!",
	})
	userID2, _ := userService.CreateNewUser(UserCreatedModel{
		Email:    "user2@example.com",
		Password: "AnotherSecurePassword2!",
	})

	testcases := []testCase{
		{
			name:        "can delete existing user",
			userID:      userID1,
			expectedErr: nil,
		},
		{
			name:        "cannot delete non-existent user",
			userID:      uuid.New(), // Generate a new UUID that doesn't exist in the repo
			expectedErr: ErrUserNotFound,
		},
		{
			name:        "can delete existing user and ensure user is removed",
			userID:      userID2,
			expectedErr: nil,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			err := userService.DeleteFromUserID(tc.userID)

			if tc.expectedErr != nil {
				assert.Equal(t, tc.expectedErr, err)
			} else {
				assert.NoError(t, err)

				// Verify the user has been deleted
				user, getErr := userService.FromID(tc.userID)
				assert.Equal(t, ErrUserNotFound, getErr)
				assert.Nil(t, user)
			}
		})
	}
}

// func TestUpdateEmailCases(t *testing.T) {
// 	type testCase struct {
// 		name        string
// 		userID      uuid.UUID
// 		newEmail    string
// 		expectedErr error
// 	}

// 	mockRepo := repository.NewUserRepositoryInMemory()
// 	userService := NewUserService(mockRepo)

// 	// Create some users for testing
// 	userID1, _ := userService.CreateNewUser(UserCreatedModel{
// 		Email:    "user1@example.com",
// 		Password: "SecurePassword1!",
// 	})
// 	userID2, _ := userService.CreateNewUser(UserCreatedModel{
// 		Email:    "user2@example.com",
// 		Password: "AnotherSecurePassword2!",
// 	})

// 	testcases := []testCase{
// 		{
// 			name:        "can update email for existing user",
// 			userID:      userID1,
// 			newEmail:    "newuser1@example.com",
// 			expectedErr: nil,
// 		},
// 		{
// 			name:        "cannot update email to invalid email",
// 			userID:      userID1,
// 			newEmail:    "invalidemail",
// 			expectedErr: ErrInvalidEmail,
// 		},
// 		{
// 			name:        "cannot update email to same email",
// 			userID:      userID2,
// 			newEmail:    "user2@example.com",
// 			expectedErr: ErrCanNotChangeToSameEmail,
// 		},
// 		{
// 			name:        "cannot update email if email is already in use",
// 			userID:      userID1,
// 			newEmail:    "user2@example.com",
// 			expectedErr: ErrUserEmailAlreadyInUse,
// 		},
// 		{
// 			name:        "cannot update email for non-existent user",
// 			userID:      uuid.New(), // Non-existent UUID
// 			newEmail:    "newuser@example.com",
// 			expectedErr: ErrUserNotFound,
// 		},
// 	}

// 	for _, tc := range testcases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			err := userService.UpdateEmailFromUserID(tc.userID, tc.newEmail)

// 			if tc.expectedErr != nil {
// 				assert.Equal(t, tc.expectedErr, err)

// 				// Additional checks for some test cases
// 				if err == ErrUserNotFound || err == ErrCanNotChangeToSameEmail {
// 					return
// 				}

// 				// Verify the email was not updated if an error occurred
// 				user, getErr := userService.FromID(tc.userID)
// 				if assert.NoError(t, getErr) {
// 					assert.Equal(t, "user1@example.com", user.Email) // Check against old email
// 				}
// 			} else {
// 				assert.NoError(t, err)

// 				// Verify the email has been updated
// 				user, getErr := userService.FromID(tc.userID)
// 				if assert.NoError(t, getErr) {
// 					assert.Equal(t, tc.newEmail, user.Email)
// 				}
// 			}
// 		})
// 	}
// }
