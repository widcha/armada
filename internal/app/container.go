package app

import (
	"github.com/widcha/armada/internal/app/usecase/healthcheck"
	"github.com/widcha/armada/internal/app/usecase/vehicle"
	"github.com/widcha/armada/internal/pkg"
)

type Container struct {
	HealthCheckInport healthcheck.Inport
	VehicleInport     vehicle.Inport
}

func NewContainer(datasource *pkg.DataSource) *Container {
	return &Container{
		HealthCheckInport: healthcheck.NewUsecase(datasource.Postgre),
		VehicleInport:     vehicle.NewUsecase(datasource.Postgre),
	}
}
