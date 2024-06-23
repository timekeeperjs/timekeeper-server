package remote

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"timekeeper-backend/internal/db"
	"timekeeper-backend/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func setupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.GET("/get-remote", GetRemoteHandler(db))
	r.POST("/push-remote", PushRemoteHandler(db))
	return r
}

func TestGetRemoteHandler(t *testing.T) {
	// Initialize the database
	db := db.Init()
	defer db.Close()

	// Seed the database
	db.Create(&models.Remote{RemoteName: "testRemote", Version: "1.0", RemoteURL: "http://example.com"})

	router := setupRouter(db)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/get-remote?remoteName=testRemote&version=1.0", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var response models.RemoteResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "testRemote", response.RemoteName)
	assert.Equal(t, "1.0", response.Version)
	assert.Equal(t, "http://example.com", response.RemoteURL)
}

func TestPushRemoteHandler(t *testing.T) {
	// Initialize the database
	db := db.Init()
	defer db.Close()

	router := setupRouter(db)

	newRemote := PushRemoteRequest{
		RemoteName: "newRemote",
		Version:    "1.0",
		BaseURL:    "http://newexample.com",
	}
	jsonValue, _ := json.Marshal(newRemote)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/push-remote", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var response models.RemoteResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "newRemote", response.RemoteName)
	assert.Equal(t, "1.0", response.Version)
	assert.Equal(t, "http://newexample.com", response.RemoteURL)
}
