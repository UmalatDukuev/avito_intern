package repository

import (
	"github.com/jmoiron/sqlx"
)

type ReceptionPostgres struct {
	db *sqlx.DB
}

func NewReceptionPostgres(db *sqlx.DB) *ReceptionPostgres {
	return &ReceptionPostgres{db: db}
}

func (r *ReceptionPostgres) CreateReception(pvzID string) (string, error) {
	var id string
	query := `
		INSERT INTO receptions (pvz_id, status)
		VALUES ($1, 'in_progress') 
		RETURNING id
	`
	row := r.db.QueryRow(query, pvzID)
	if err := row.Scan(&id); err != nil {
		return "", err
	}

	return id, nil
}
