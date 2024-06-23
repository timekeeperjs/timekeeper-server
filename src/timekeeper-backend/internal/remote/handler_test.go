package remote

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"timekeeper-backend/internal/db"
	"timekeeper-backend/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

const testDBFilePath = "test.db"

func setupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.GET("/get-remote", GetRemoteHandler(db))
	r.POST("/push-remote", PushRemoteHandler(db))
	r.GET("/remote-names", GetRemoteNamesHandler(db))
	r.GET("/dashboard", DashboardHandler(db))
	return r
}

func TestMain(m *testing.M) {
	// Run tests
	code := m.Run()

	// Clean up
	os.Remove(testDBFilePath)

	os.Exit(code)
}

func TestGetRemoteHandler(t *testing.T) {
	// Initialize the database
	database := db.Init(testDBFilePath)
	defer database.Close()

	// Seed the database
	database.Create(&models.Remote{RemoteName: "testRemote", Version: "1.0", RemoteURL: "http://example.com"})

	router := setupRouter(database)

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

	// Clear the database
	err = db.ClearDatabase(testDBFilePath)
	assert.NoError(t, err)
}

func TestGetRemoteHandler_NotFound(t *testing.T) {
	// Initialize the database
	database := db.Init(testDBFilePath)
	defer database.Close()

	router := setupRouter(database)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/get-remote?remoteName=nonexistent&version=1.0", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)

	// Clear the database
	err := db.ClearDatabase(testDBFilePath)
	assert.NoError(t, err)
}

func TestPushRemoteHandler(t *testing.T) {
	// Initialize the database
	database := db.Init(testDBFilePath)
	defer database.Close()

	router := setupRouter(database)

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

	// Clear the database
	err = db.ClearDatabase(testDBFilePath)
	assert.NoError(t, err)
}

// need to investigate
func TestPushRemoteHandler_BadRequest(t *testing.T) {
	// Initialize the database
	database := db.Init(testDBFilePath)
	defer database.Close()

	router := setupRouter(database)

	invalidRemote := map[string]string{
		"invalidField": "value",
	}
	jsonValue, _ := json.Marshal(invalidRemote)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/push-remote", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Clear the database
	err := db.ClearDatabase(testDBFilePath)
	assert.NoError(t, err)
}

func TestGetRemoteNamesHandler(t *testing.T) {
	// Initialize the database
	database := db.Init(testDBFilePath)
	defer database.Close()

	// Seed the database
	database.Create(&models.Remote{RemoteName: "testRemote1", Version: "1.0", RemoteURL: "http://example1.com"})
	database.Create(&models.Remote{RemoteName: "testRemote2", Version: "1.0", RemoteURL: "http://example2.com"})
	database.Create(&models.Remote{RemoteName: "testRemote1", Version: "2.0", RemoteURL: "http://example3.com"})

	router := setupRouter(database)

	// Test without remoteName query parameter
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/dashboard", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var response []models.Remote
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Len(t, response, 3)

	// Test with remoteName query parameter
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/dashboard?remoteName=testRemote1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Len(t, response, 2)
	assert.Equal(t, "testRemote1", response[0].RemoteName)
	assert.Equal(t, "testRemote1", response[1].RemoteName)

	// Clear the database
	err = db.ClearDatabase(testDBFilePath)
	assert.NoError(t, err)
}

func TestDashboardHandler_Empty(t *testing.T) {
	// Initialize the database
	database := db.Init(testDBFilePath)
	defer database.Close()

	router := setupRouter(database)

	// Test without remoteName query parameter
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/dashboard", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var response []models.Remote
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Len(t, response, 0)

	// Test with remoteName query parameter
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/dashboard?remoteName=nonexistent", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Len(t, response, 0)

	// Clear the database
	err = db.ClearDatabase(testDBFilePath)
	assert.NoError(t, err)
}
