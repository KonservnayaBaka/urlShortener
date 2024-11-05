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

func TestAuthUser(t *testing.T) {
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("failed to setup test database: %v", err)
	}
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.POST("/auth/login", controller.AuthUser(db))

	t.Run("successful auth", func(t *testing.T) {
		user := entity.User{
			Login:    "testuser",
			Password: "testpassword",
		}
		jsonData, _ := json.Marshal(user)

		req, _ := http.NewRequest("POST", "/auth/login", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string]string
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.NotEmpty(t, response["token"])
	})

	t.Run("missing fields authorization", func(t *testing.T) {
		user := entity.User{
			Login:    "testuser",
			Password: "",
		}
		jsonData, _ := json.Marshal(user)

		req, _ := http.NewRequest("POST", "/auth/login", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.JSONEq(t, `{"error":"login or password is empty"}`, w.Body.String())
	})

	t.Run("invalid JSON", func(t *testing.T) {
		req, _ := http.NewRequest("POST", "/auth/login", nil)
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.JSONEq(t, `{"error":"invalid request"}`, w.Body.String())
	})
}
