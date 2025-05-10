package healthcheck

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/widcha/armada/configs"
	"github.com/widcha/armada/internal/app/usecase/healthcheck"
)

// HealthCheckHandler godoc
// @Summary      Health Check
// @Tags         health-check
// @Accept       json
// @Produce      json
// @Success 	 200 {object} healthcheck.InportResponse "Result"
// @Router       / [get]
func HealthCheckHandler(inport healthcheck.Inport) gin.HandlerFunc {
	return func(c *gin.Context) {
		resp, err := inport.Execute(c.Copy().Request.Context())
		if err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"message": "failed",
				"status":  resp.Status,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
			"name":    configs.Get().AppName,
			"version": configs.Get().AppVersion,
			"status":  resp.Status,
		})
	}
}
