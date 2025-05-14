package vehicle

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type interactor struct {
	db *sqlx.DB
}

func NewUsecase(db *sqlx.DB) Inport {
	return interactor{db: db}
}

func (i interactor) GetLatestLocation(ctx context.Context, vehicleID string) (LocationResponse, error) {
	query := `
		SELECT vehicle_id, latitude, longitude, EXTRACT(EPOCH FROM timestamp)::bigint AS timestamp
		FROM vehicle_locations
		WHERE vehicle_id = $1
		ORDER BY timestamp DESC
		LIMIT 1
	`

	var loc LocationResponse
	err := i.db.GetContext(ctx, &loc, query, vehicleID)
	return loc, err
}

func (i interactor) GetLocationHistory(ctx context.Context, vehicleID string, start, end int64) ([]LocationResponse, error) {
	query := `
		SELECT vehicle_id, latitude, longitude, EXTRACT(EPOCH FROM timestamp)::bigint AS timestamp
		FROM vehicle_locations
		WHERE vehicle_id = $1
		AND EXTRACT(EPOCH FROM timestamp)::bigint BETWEEN $2 AND $3
		ORDER BY timestamp ASC
	`

	var history []LocationResponse
	err := i.db.SelectContext(ctx, &history, query, vehicleID, start, end)
	return history, err
}
