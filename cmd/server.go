package cmd

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/widcha/armada/configs"
	"github.com/widcha/armada/internal/app"
	"github.com/widcha/armada/internal/app/delivery/mqtt"
	"github.com/widcha/armada/internal/app/delivery/rest"
	"github.com/widcha/armada/internal/pkg"
	"github.com/widcha/armada/internal/pkg/migration"
	"github.com/widcha/armada/worker/geofence"
)

var serverCmd = &cobra.Command{
	Use:   "start",
	Short: "Runs the server",
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

var workerCmd = &cobra.Command{
	Use:   "worker",
	Short: "Start geofence alert worker",
	Run: func(cmd *cobra.Command, args []string) {
		err := geofence.StartGeofenceWorker(configs.Get().RabbitMQ)
		if err != nil {
			log.Fatal("Worker Geofence error:", err)
		}
	},
}

func run() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	datasource := pkg.NewDataSource()
	migration.CreateVehicleLocationsTable(datasource.Postgre)
	container := app.NewContainer(datasource)

	subscriber := mqtt.NewSubscriber(datasource.Postgre)
	go subscriber.Start()

	router := rest.NewRouter(r, datasource, container)
	router.RegisterRouter()

	r.Run()
}
