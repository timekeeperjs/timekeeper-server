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

// GetRemoteNamesHandler godoc
// @Summary Get all unique remote names
// @Description Get all unique remote names
// @Tags remote
// @Accept  json
// @Produce  json
// @Success 200 {array} string
// @Failure 500 {object} models.ErrorResponse
// @Router /remote-names [get]
func GetRemoteNamesHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var remoteNames []string
		if err := db.Model(&models.Remote{}).Pluck("DISTINCT remote_name", &remoteNames).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve remote names"})
			return
		}

		c.JSON(http.StatusOK, remoteNames)
	}
}

// DashboardHandler godoc
// @Summary Get all remotes or remotes by name
// @Description Get all remotes or remotes by name
// @Tags remote
// @Accept  json
// @Produce  json
// @Param remoteName query string false "Remote Name"
// @Success 200 {array} models.RemoteResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /dashboard [get]
func DashboardHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		remoteName := c.Query("remoteName")

		var remotes []models.Remote
		var err error

		if remoteName != "" {
			err = db.Where("remote_name = ?", remoteName).Find(&remotes).Error
		} else {
			err = db.Find(&remotes).Error
		}

		if err != nil && err != gorm.ErrRecordNotFound {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve remotes"})
			return
		}

		c.JSON(http.StatusOK, remotes)
	}
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
// @Failure 500 {object} models.ErrorResponse
// @Router /get-remote [get]
func GetRemoteHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		remoteName := c.Query("remoteName")
		version := c.Query("version")

		var remote models.Remote
		if err := db.Where("remote_name = ? AND version = ?", remoteName, version).First(&remote).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "Remote not found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
			}
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
// @Failure 500 {object} models.ErrorResponse
// @Router /push-remote [post]
func PushRemoteHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req PushRemoteRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Additional validation checks
		if req.RemoteName == "" || req.Version == "" || req.BaseURL == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields"})
			return
		}

		remote := models.Remote{
			RemoteName: req.RemoteName,
			Version:    req.Version,
			RemoteURL:  req.BaseURL,
		}

		if err := db.Create(&remote).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create remote"})
			return
		}

		c.JSON(http.StatusOK, remote)
	}
}
