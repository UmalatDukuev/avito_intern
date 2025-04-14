package repository

import (
	"avito_intern/internal/handler/entity"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"

	"github.com/jmoiron/sqlx"
)

type PVZPostgres struct {
	db *sqlx.DB
}

func NewPVZPostgres(db *sqlx.DB) *PVZPostgres {
	return &PVZPostgres{db: db}
}

func (r *PVZPostgres) CreatePVZ(pvz entity.PVZ) (string, error) {
	query, args, err := sq.Insert(pvzTable).
		Columns("city", "registration_date").
		Values(pvz.City, pvz.RegistrationDate).
		Suffix("RETURNING id").
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return "", err
	}

	var id string
	if err := r.db.QueryRow(query, args...).Scan(&id); err != nil {
		return "0", err
	}
	return id, nil
}

type pvzFlatRow struct {
	ID               string    `db:"id"`
	RegistrationDate time.Time `db:"registration_date"`
	City             string    `db:"city"`

	ReceptionID        *string    `db:"reception_id"`
	ReceptionDateTime  *time.Time `db:"reception_date_time"`
	ReceptionStatus    *string    `db:"reception_status"`
	ReceptionCreatedAt *time.Time `db:"reception_created_at"`
	ReceptionClosedAt  *time.Time `db:"reception_closed_at"`

	ProductID          *string    `db:"product_id"`
	ProductDateTime    *time.Time `db:"product_date_time"`
	ProductType        *string    `db:"product_type"`
	ProductReceptionID *string    `db:"product_reception_id"`
	ProductCreatedAt   *time.Time `db:"product_created_at"`
	ProductUpdatedAt   *time.Time `db:"product_updated_at"`
}

func (r *PVZPostgres) GetPVZWithDetails(startDate, endDate *time.Time, page, limit int) ([]entity.PVZResponse, error) {
	query := `
        SELECT 
            pvz.id,
            pvz.registration_date,
            pvz.city,
            reception.id AS reception_id,
            reception.date_time AS reception_date_time,
            reception.status AS reception_status,
            reception.created_at AS reception_created_at,
            reception.closed_at AS reception_closed_at,
            product.id AS product_id,
            product.date_time AS product_date_time,
            product.type AS product_type,
            product.reception_id AS product_reception_id,
            product.created_at AS product_created_at,
            product.updated_at AS product_updated_at
        FROM pvzs pvz
        JOIN receptions reception ON reception.pvz_id = pvz.id
        LEFT JOIN products product ON product.reception_id = reception.id
        WHERE 1=1
    `
	var args []interface{}
	argIdx := 1

	if startDate != nil {
		query += fmt.Sprintf(" AND reception.date_time >= $%d", argIdx)
		args = append(args, *startDate)
		argIdx++
	}
	if endDate != nil {
		query += fmt.Sprintf(" AND reception.date_time <= $%d", argIdx)
		args = append(args, *endDate)
		argIdx++
	}

	query += fmt.Sprintf(" ORDER BY pvz.registration_date ASC LIMIT $%d OFFSET $%d", argIdx, argIdx+1)
	args = append(args, limit, (page-1)*limit)

	var rows []pvzFlatRow
	if err := r.db.Select(&rows, query, args...); err != nil {
		return nil, fmt.Errorf("failed to get PVZ details: %v", err)
	}

	pvzMap := make(map[string]*entity.PVZResponse)
	for _, row := range rows {
		if _, ok := pvzMap[row.ID]; !ok {
			pvzMap[row.ID] = &entity.PVZResponse{
				PVZ: entity.PVZ{
					ID:               row.ID,
					RegistrationDate: row.RegistrationDate,
					City:             row.City,
				},
				Receptions: []entity.ReceptionResponse{},
			}
		}
		currentPVZ := pvzMap[row.ID]

		if row.ReceptionID == nil {
			continue
		}

		var recResp *entity.ReceptionResponse
		for i, rcp := range currentPVZ.Receptions {
			if rcp.Reception.ID == *row.ReceptionID {
				recResp = &currentPVZ.Receptions[i]
				break
			}
		}
		if recResp == nil {
			rec := entity.Reception{
				ID:        *row.ReceptionID,
				DateTime:  *row.ReceptionDateTime,
				PvzID:     row.ID,
				Status:    *row.ReceptionStatus,
				CreatedAt: *row.ReceptionCreatedAt,
				ClosedAt:  row.ReceptionClosedAt,
			}
			newRecResp := entity.ReceptionResponse{
				Reception: rec,
				Products:  []entity.Product{},
			}
			currentPVZ.Receptions = append(currentPVZ.Receptions, newRecResp)
			recResp = &currentPVZ.Receptions[len(currentPVZ.Receptions)-1]
		}

		if row.ProductID != nil {
			prod := entity.Product{
				ID:          *row.ProductID,
				DateTime:    *row.ProductDateTime,
				Type:        *row.ProductType,
				ReceptionID: *row.ProductReceptionID,
				CreatedAt:   *row.ProductCreatedAt,
				UpdatedAt:   *row.ProductUpdatedAt,
			}
			recResp.Products = append(recResp.Products, prod)
		}
	}

	result := make([]entity.PVZResponse, 0, len(pvzMap))
	for _, pvz := range pvzMap {
		result = append(result, *pvz)
	}

	return result, nil
}

func (r *PVZPostgres) GetByID(pvzID string) (*entity.PVZ, error) {
	var pvz entity.PVZ
	query := `SELECT id, registration_date, city FROM pvzs WHERE id = $1`
	err := r.db.Get(&pvz, query, pvzID)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get pvz: %v", err)
	}
	return &pvz, nil
}
