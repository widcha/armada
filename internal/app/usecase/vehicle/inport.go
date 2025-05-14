package vehicle

import (
	"context"
)

type Inport interface {
	GetLatestLocation(ctx context.Context, vehicleID string) (LocationResponse, error)
	GetLocationHistory(ctx context.Context, vehicleID string, start, end int64) ([]LocationResponse, error)
}

type LocationResponse struct {
	VehicleID string  `json:"vehicle_id" db:"vehicle_id"`
	Latitude  float64 `json:"latitude" db:"latitude"`
	Longitude float64 `json:"longitude" db:"longitude"`
	Timestamp int64   `json:"timestamp" db:"timestamp"`
}
