package models

type Acceptance struct {
	ID        int    `json:"id" db:"id"`
	PvzID     int    `json:"pvz_id" db:"pvz_id"`
	CreatedAt string `json:"created_at" db:"created_at"`
	Status    string `json:"status" db:"status"`
}
