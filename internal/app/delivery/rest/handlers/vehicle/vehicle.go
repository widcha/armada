package vehicle

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/widcha/armada/internal/app/usecase/vehicle"
)

// GetLatestLocationHandler godoc
// @Summary      Get latest vehicle location
// @Tags         vehicle
// @Produce      json
// @Param        vehicle_id  path string true "Vehicle ID"
// @Success      200 {object} vehicle.LocationResponse
// @Router       /vehicles/{vehicle_id}/location [get]
func GetLatestLocationHandler(inport vehicle.Inport) gin.HandlerFunc {
	return func(c *gin.Context) {
		vehicleID := c.Param("vehicle_id")
		log.Println(vehicleID)
		resp, err := inport.GetLatestLocation(c.Request.Context(), vehicleID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, resp)
	}
}

// GetHistoryHandler godoc
// @Summary      Get vehicle location history
// @Tags         vehicle
// @Produce      json
// @Param        vehicle_id  path string true "Vehicle ID"
// @Param        start       query int64 true "Start timestamp"
// @Param        end         query int64 true "End timestamp"
// @Success      200 {array} vehicle.LocationResponse
// @Router       /vehicles/{vehicle_id}/history [get]
func GetHistoryHandler(inport vehicle.Inport) gin.HandlerFunc {
	return func(c *gin.Context) {
		vehicleID := c.Param("vehicle_id")
		start, err1 := strconv.ParseInt(c.Query("start"), 10, 64)
		end, err2 := strconv.ParseInt(c.Query("end"), 10, 64)

		if err1 != nil || err2 != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid query parameter"})
			return
		}

		resp, err := inport.GetLocationHistory(c.Request.Context(), vehicleID, start, end)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, resp)
	}
}
