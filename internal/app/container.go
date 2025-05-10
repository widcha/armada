package app

import (
	"github.com/widcha/armada/internal/app/usecase/healthcheck"
	"github.com/widcha/armada/internal/pkg"
)

type Container struct {
	HealthCheckInport healthcheck.Inport
}

func NewContainer(datasource *pkg.DataSource) *Container {
	return &Container{
		HealthCheckInport: healthcheck.NewUsecase(datasource.Postgre),
	}
}
