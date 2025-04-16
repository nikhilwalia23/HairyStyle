package models

type User struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Role       string `json:"role"` // user or stylist
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	IsVerified bool   `json:"is_verified"`
	AadharCard string `json:"aadhar_card"` // 12-digit string
	Password   string `json:"password"`    // 8–20 characters
}
