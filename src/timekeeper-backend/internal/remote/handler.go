package remote

import (
	"net/http"
	"timekeeper-backend/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type PushRemoteRequest struct {
	RemoteName string `json:"remoteName"`
	Version    string `json:"version"`
	BaseURL    string `json:"baseUrl"`
}

// GetRemoteHandler godoc
// @Summary Get a remote by name and version
// @Description Get a remote by name and version
// @Tags remote
// @Accept  json
// @Produce  json
// @Param remoteName query string true "Remote Name"
// @Param version query string true "Version"
// @Success 200 {object} models.RemoteResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /get-remote [get]
func GetRemoteHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		remoteName := c.Query("remoteName")
		version := c.Query("version")

		var remote models.Remote
		if err := db.Where("remote_name = ? AND version = ?", remoteName, version).First(&remote).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Remote not found"})
			return
		}

		c.JSON(http.StatusOK, remote)
	}
}

// PushRemoteHandler godoc
// @Summary Push a new remote
// @Description Push a new remote
// @Tags remote
// @Accept  json
// @Produce  json
// @Param remote body PushRemoteRequest true "Remote"
// @Success 200 {object} models.RemoteResponse
// @Failure 400 {object} models.ErrorResponse
// @Router /push-remote [post]
func PushRemoteHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req PushRemoteRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		remote := models.Remote{
			RemoteName: req.RemoteName,
			Version:    req.Version,
			RemoteURL:  req.BaseURL,
		}

		db.Create(&remote)
		c.JSON(http.StatusOK, remote)
	}
}
