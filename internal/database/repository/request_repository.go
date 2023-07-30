package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/AntoniKania/nbp-api-wrapper-go/internal/api/model"
)

type RequestRepository interface {
	SaveRequest(ctx context.Context, request *model.Request) (int64, error)
}

type requestRepository struct {
	db *sql.DB
}

func NewRequestRepository(db *sql.DB) RequestRepository {
	return &requestRepository{
		db: db,
	}
}

func (r *requestRepository) SaveRequest(ctx context.Context, request *model.Request) (int64, error) {
	var id int64
	err := r.db.QueryRow(`INSERT INTO request(currency, average, start_date, end_date, request_date)
	VALUES($1, $2, $3, $4, $5) RETURNING id`, request.Currency, request.Average, request.StartDate, request.EndDate, request.RequestDate).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("SaveRequest: %v", err)
	}

	return id, nil
}
