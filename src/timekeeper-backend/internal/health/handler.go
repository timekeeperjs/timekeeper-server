package health

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type SuccessResponse struct {
	Status string `json:"status"`
}

// HealthCheck godoc
// @Summary Health check endpoint
// @Description Returns the status of the server
// @Tags health
// @Produce  json
// @Success 200 {object} SuccessResponse
// @Router /health-check [get]
func HealthCheck(c *gin.Context) {
	response := SuccessResponse{
		Status: "OK",
	}
	c.JSON(http.StatusOK, response)
}
