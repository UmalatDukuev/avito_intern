package models

type Pvz struct {
	ID        int    `json:"id" db:"id"`
	City      string `json:"city" db:"city"`
	Creator   string `json:"creator" db:"creator"`
	CreatedAt string `json:"created_at" db:"created_at"`
}
