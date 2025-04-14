package entity

import "time"

type PVZ struct {
	ID               string    `json:"id" db:"id"`
	RegistrationDate time.Time `json:"registration_date" db:"registration_date"`
	City             string    `json:"city" db:"city"`
}

type PVZResponse struct {
	PVZ        PVZ                 `json:"pvz"`
	Receptions []ReceptionResponse `json:"receptions"`
}
