package entity

import "time"

type Reception struct {
	ID        string     `json:"id" db:"id"`
	DateTime  time.Time  `json:"date_time" db:"date_time"`
	PvzID     string     `json:"pvz_id" db:"pvz_id"`
	Status    string     `json:"status" db:"status"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	ClosedAt  *time.Time `json:"closed_at" db:"closed_at"`
}

type ReceptionResponse struct {
	Reception Reception `json:"reception"`
	Products  []Product `json:"products"`
}
