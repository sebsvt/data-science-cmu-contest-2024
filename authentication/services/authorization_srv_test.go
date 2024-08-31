package services

import (
	"testing"
	"time"

	"github.com/sebsvt/prototype/repository"
	"github.com/stretchr/testify/assert"
)

func TestAuth_SignIn(t *testing.T) {
	// Setup
	type testCase struct {
		name         string
		email        string
		password     string
		expectedErr  error
		expectTokens bool
	}

	userRepo := repository.NewUserRepositoryInMemory()
	authService := NewAuth(userRepo, []byte("your-secret-key"), time.Minute*15, time.Hour*24)
	userSrv := NewUserService(userRepo)

	// Add a test user
	_, err := userSrv.CreateNewUser(UserCreatedModel{
		Email:    "test@example.com",
		Password: "!AiyaAiyayay",
	})
	assert.NoError(t, err)

	testcases := []testCase{
		{
			name:         "does not exist email should return user not found",
			email:        "lolipop@gmail.com",
			password:     "!SeC4urePew",
			expectedErr:  ErrAuthenticationFailed,
			expectTokens: false,
		},
		{
			name:         "exist email but wrong password should return authentication failed",
			email:        "test@example.com",
			password:     "!SeC4urePew",
			expectedErr:  ErrAuthenticationFailed,
			expectTokens: false,
		},
		{
			name:         "exist email and correct password should return tokens and has no error",
			email:        "test@example.com",
			password:     "!AiyaAiyayay",
			expectedErr:  nil,
			expectTokens: true,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			accessToken, refreshToken, err := authService.SignIn(tc.email, tc.password)
			assert.Equal(t, tc.expectedErr, err)
			if tc.expectTokens {
				assert.NotEmpty(t, accessToken)
				assert.NotEmpty(t, refreshToken)
			} else {
				assert.Empty(t, accessToken)
				assert.Empty(t, refreshToken)
			}
		})
	}
}

func TestAuth_RefreshToken(t *testing.T) {
	// Setup
	type testCase struct {
		name            string
		existingRFTK    string
		expectedErr     error
		expectNewTokens bool
	}

	userRepo := repository.NewUserRepositoryInMemory()
	authService := NewAuth(userRepo, []byte("your-secret-key"), time.Minute*15, time.Hour*24)
	userSrv := NewUserService(userRepo)

	// Add a test user
	_, err := userSrv.CreateNewUser(UserCreatedModel{
		Email:    "test@example.com",
		Password: "!AiyaAiyayay",
	})
	assert.NoError(t, err)

	// Generate initial tokens
	_, refreshToken, err := authService.SignIn("test@example.com", "!AiyaAiyayay")
	assert.NoError(t, err)

	testcases := []testCase{
		{
			name:            "invalid refresh token should return error",
			existingRFTK:    "invalidToken",
			expectedErr:     ErrBadClaim,
			expectNewTokens: false,
		},
		{
			name:            "valid refresh token should return new tokens",
			existingRFTK:    refreshToken,
			expectedErr:     nil,
			expectNewTokens: true,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			newAccessToken, newRefreshToken, err := authService.Refresh(tc.existingRFTK)
			assert.Equal(t, tc.expectedErr, err)
			if tc.expectNewTokens {
				assert.NotEmpty(t, newAccessToken)
				assert.NotEmpty(t, newRefreshToken)
			} else {
				assert.Empty(t, newAccessToken)
				assert.Empty(t, newRefreshToken)
			}
		})
	}
}
