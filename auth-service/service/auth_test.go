// auth-service/service/auth_test.go
package service

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
    // Import repository Anda di sini
)

// MockRepository untuk unit test
type MockAuthRepo struct { mock.Mock }

func (m *MockAuthRepo) GetByEmail(email string) (*models.User, error) {
    args := m.Called(email)
    return args.Get(0).(*models.User), args.Error(1)
}

func TestLogin_Success(t *testing.T) {
    mockRepo := new(MockAuthRepo)
    // Skenario: User ditemukan di database
    mockRepo.On("GetByEmail", "user@test.com").Return(&models.User{Email: "user@test.com", Password: "hashed_password"}, nil)

    authService := NewAuthService(mockRepo) // Anda harus buat constructor ini
    user, err := authService.Login("user@test.com", "password") // Asumsi password benar

    assert.NoError(t, err)
    assert.NotNil(t, user)
}