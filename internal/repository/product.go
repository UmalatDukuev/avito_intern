package repository

import (
	"avito_intern/internal/handler/entity"
	"avito_intern/internal/handler/response"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type ProductPostgres struct {
	db *sqlx.DB
}

func NewProductPostgres(db *sqlx.DB) *ProductPostgres {
	return &ProductPostgres{db: db}
}

func (r *ProductPostgres) AddProductToReception(product entity.Product) (string, error) {
	var id string
	query := `
		INSERT INTO products (type, reception_id, date_time)
		VALUES ($1, $2, $3) 
		RETURNING id
	`

	row := r.db.QueryRow(query, product.Type, product.ID, product.DateTime)
	if err := row.Scan(&id); err != nil {
		return "", err
	}

	return id, nil
}

func (r *ProductPostgres) GetActiveReception(pvzID string) (*response.ReceptionResponse, error) {
	var reception response.ReceptionResponse
	query := `
		SELECT id, pvz_id, status
		FROM receptions
		WHERE pvz_id = $1 AND status = 'in_progress'
		LIMIT 1
	`
	err := r.db.Get(&reception, query, pvzID)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get active reception: %v", err)
	}

	return &reception, nil
}

func (r *ProductPostgres) GetLastProduct(receptionID string) (*entity.Product, error) {
	var product entity.Product
	query := `
		SELECT id, type, reception_id
		FROM products
		WHERE reception_id = $1
		ORDER BY date_time DESC
		LIMIT 1
	`

	err := r.db.Get(&product, query, receptionID)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get last product: %v", err)
	}

	return &product, nil
}

func (r *ProductPostgres) DeleteProduct(productID string) error {
	query := `
		DELETE FROM products
		WHERE id = $1
	`
	_, err := r.db.Exec(query, productID)
	return err
}
