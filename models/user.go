package models

type User struct {
	ID    int    `json:"id" db:"id"`
	Email string `json:"email" db:"email"`
	Role  string `json:"role" db:"role"`
}
