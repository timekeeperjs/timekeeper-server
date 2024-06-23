package health

import (
	"net/http"
	"timekeeper-backend/internal/models"

	"github.com/gin-gonic/gin"
)

// HealthCheck godoc
// @Summary Health check endpoint
// @Description Returns the status of the server
// @Tags health
// @Produce  json
// @Success 200 {object} models.SuccessResponse
// @Router /health-check [get]
func HealthCheck(c *gin.Context) {
	response := models.SuccessResponse{
		Status: "OK",
	}
	c.JSON(http.StatusOK, response)
}
