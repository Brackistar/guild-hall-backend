package models

type User struct {
	Id    uint64   `json:"id"`
	Role  UserRole `json:"role"`
	Email string   `json:"email"`
}
