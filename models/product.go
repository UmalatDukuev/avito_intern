package models

import "time"

type Product struct {
	ID          string    `json:"id" db:"id"`
	DateTime    time.Time `json:"date_time" db:"date_time"`
	Type        string    `json:"type" db:"type"`
	ReceptionID string    `json:"reception_id" db:"reception_id"`
}
