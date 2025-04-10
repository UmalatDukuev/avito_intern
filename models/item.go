package models

type Item struct {
	ID           int    `json:"id" db:"id"`
	AcceptanceID int    `json:"acceptance_id" db:"acceptance_id"`
	Type         string `json:"type" db:"type"`
	ReceivedAt   string `json:"received_at" db:"received_at"`
}
