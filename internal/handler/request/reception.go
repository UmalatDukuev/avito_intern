package request

type CreateReception struct {
	PvzID string `json:"pvz_id" db:"pvz_id"`
}
