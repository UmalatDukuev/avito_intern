package entity

type Reception struct {
	ID       string `json:"id" db:"id"`
	DateTime string `json:"date_time" db:"date_time"`
	PvzID    string `json:"pvz_id" db:"pvz_id"`
	Status   string `json:"status" db:"status"`
}
