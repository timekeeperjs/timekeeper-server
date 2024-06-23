package health

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	// Set up Gin router
	router := gin.Default()
	router.GET("/health-check", HealthCheck)

	// Create a new HTTP request
	req, _ := http.NewRequest("GET", "/health-check", nil)
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"status": "OK"}`, w.Body.String())
}
