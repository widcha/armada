package migration

import (
	"context"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

func CreateVehicleLocationsTable(db *sqlx.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS vehicle_locations (
		id SERIAL PRIMARY KEY,
		vehicle_id TEXT NOT NULL,
		latitude DOUBLE PRECISION,
		longitude DOUBLE PRECISION,
		timestamp TIMESTAMP NOT NULL,
		created_at TIMESTAMP DEFAULT NOW()
	);
	`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := db.ExecContext(ctx, query)
	if err != nil {
		log.Fatalf("Failed to create vehicle_locations table: %v", err)
	}

	log.Println("Table vehicle_locations is ready")
}
