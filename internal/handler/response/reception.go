package response

import "time"

type ReceptionResponse struct {
	ID    string `json:"id" db:"id"`
	PvzID string `json:"pvz_id" db:"pvz_id"`

	Status   string    `json:"status" db:"status"`
	ClosedAt time.Time `json:"closed_at"`
}
