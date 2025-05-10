package cmd

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/widcha/armada/internal/app"
	"github.com/widcha/armada/internal/app/delivery/rest"
	"github.com/widcha/armada/internal/pkg"
)

var serverCmd = &cobra.Command{
	Use:   "start",
	Short: "Runs the server",
	Run: func(cmd *cobra.Command, args []string) {
		run()
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
	container := app.NewContainer(datasource)

	router := rest.NewRouter(r, datasource, container)
	router.RegisterRouter()

	r.Run()
}
