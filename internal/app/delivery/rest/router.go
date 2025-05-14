package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/widcha/armada/internal/app"
	"github.com/widcha/armada/internal/app/delivery/rest/handlers/healthcheck"
	"github.com/widcha/armada/internal/app/delivery/rest/handlers/vehicle"
	"github.com/widcha/armada/internal/pkg"
)

type Router struct {
	router     gin.IRouter
	datasource *pkg.DataSource
	container  *app.Container
}

func NewRouter(router gin.IRouter, datasource *pkg.DataSource, container *app.Container) *Router {
	return &Router{
		router:     router,
		datasource: datasource,
		container:  container,
	}
}

func (h *Router) RegisterRouter() {
	h.router.GET("/health", healthcheck.HealthCheckHandler(h.container.HealthCheckInport))

	h.router.GET("/vehicles/:vehicle_id/location", vehicle.GetLatestLocationHandler(h.container.VehicleInport))
	h.router.GET("/vehicles/:vehicle_id/history", vehicle.GetHistoryHandler(h.container.VehicleInport))
}
