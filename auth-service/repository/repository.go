// auth-service/repository/repository.go
package repository

import (
	"context"
	"github.com/shakilaaulia/Dealan/auth-service/models"
)

// AuthRepository adalah interface atau "kontrak" yang mendefinisikan
// semua operasi database yang dibutuhkan oleh Auth Service.
type AuthRepository interface {
	// --- Kredensial Pengguna (auth_credentials) ---
	
	// CreateUser menyimpan data user/driver baru ke database saat registrasi
	CreateUser(ctx context.Context, user *models.User) error
	
	// GetUserByEmail mencari user berdasarkan email (untuk login via email)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	
	// GetUserByPhone mencari user berdasarkan nomor HP (untuk login via HP)
	GetUserByPhone(ctx context.Context, phone string) (*models.User, error)


	// --- Pengelolaan OTP (otp_codes) ---
	
	// SaveOTP menyimpan kode OTP untuk nomor HP tertentu
	SaveOTP(ctx context.Context, phone string, otpCode string) error
	
	// GetOTP mengambil kode OTP yang tersimpan untuk divalidasi
	GetOTP(ctx context.Context, phone string) (string, error)


	// --- Pengelolaan Token (refresh_tokens) ---
	
	// SaveRefreshToken menyimpan refresh token yang valid ke database
	SaveRefreshToken(ctx context.Context, userID string, token string) error
	
	// GetRefreshToken mengambil refresh token berdasarkan User ID
	GetRefreshToken(ctx context.Context, userID string) (string, error)
	
	// DeleteRefreshToken menghapus token saat user melakukan Logout
	DeleteRefreshToken(ctx context.Context, userID string) error
}