package entity

type Reception struct {
	ID       int    `json:"id" db:"id"`
	DateTime string `json:"date_time" db:"date_time"`
	PvzID    int    `json:"pvz_id" db:"pvz_id"`
	Status   string `json:"status" db:"status"`
}
