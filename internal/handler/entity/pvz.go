package entity

type PVZ struct {
	ID              int    `json:"id" db:"id"`
	RegistraionDate string `json:"registraion_date" db:"registraion_date"`
	City            string `json:"city" db:"city"`
}
