// auth-service/models/user.go
package models

// User merepresentasikan data akun yang digunakan untuk login
type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Phone    string `json:"nomor_hp"`
	Password string `json:"-"` // Disembunyikan saat dikirim sebagai JSON, wajib di-hash menggunakan Bcrypt
	Role     string `json:"role"` // "user" atau "driver"
}