package repository

import (
	"avito_intern/internal/handler/entity"
	"fmt"

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

func (r *ReceptionPostgres) GetActiveReception(pvzID string) (*entity.Reception, error) {
	var reception entity.Reception
	query := `
		SELECT id, pvz_id, status
		FROM receptions
		WHERE pvz_id = $1 AND status = 'in_progress'
		LIMIT 1
	`
	// Выполняем запрос
	err := r.db.Get(&reception, query, pvzID)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			// Если не нашли запись, возвращаем nil
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get active reception: %v", err)
	}

	return &reception, nil
}
