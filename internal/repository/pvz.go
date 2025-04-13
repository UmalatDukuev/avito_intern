package repository

import (
	"avito_intern/internal/handler/entity"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type PVZPostgres struct {
	db *sqlx.DB
}

func NewPVZPostgres(db *sqlx.DB) *PVZPostgres {
	return &PVZPostgres{db: db}
}

func (r *PVZPostgres) CreatePVZ(pvz entity.PVZ) (string, error) {
	var id string
	query := fmt.Sprintf("INSERT INTO %s (city, registration_date) values ($1, $2) RETURNING id", pvzTable)

	row := r.db.QueryRow(query, pvz.City, pvz.RegistrationDate)
	if err := row.Scan(&id); err != nil {
		return "0", err
	}
	return id, nil
}
