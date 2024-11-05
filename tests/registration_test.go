package tests

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"urlShortener/internal/domain/entity"
	"urlShortener/internal/interface/controller"
)

func TestRegUser(t *testing.T) {
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("failed to setup test database: %v", err)
	}

	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.POST("/auth/login", controller.RegUser(db))

	t.Run("successful registration", func(t *testing.T) {
		user := entity.User{
			Login:    "testuser",
			Password: "testpassword",
			Name:     "Test User",
		}
		jsonData, _ := json.Marshal(user)

		req, _ := http.NewRequest("POST", "/auth/login", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, `{"status":true}`, w.Body.String())
	})

	t.Run("missing fields registration", func(t *testing.T) {
		user := entity.User{
			Login:    "testuser",
			Password: "",
			Name:     "Test User",
		}
		jsonData, _ := json.Marshal(user)

		req, _ := http.NewRequest("POST", "/auth/login", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, `{"error":"login or password or name is empty"}`, w.Body.String())
	})

	t.Run("invalid JSON", func(*testing.T) {
		req, _ := http.NewRequest("POST", "/auth/login", bytes.NewBufferString("invalid json"))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "error")
	})
}
