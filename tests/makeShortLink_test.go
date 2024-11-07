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
	router2 "urlShortener/internal/interface/router"
)

func TestMakeShortLink(t *testing.T) {
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("failed to setup test database: %v", err)
	}
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.POST("/auth/login", controller.AuthUser(db))
	router.POST("/link/makeShortLink", controller.MakeShortLink(db))
	router.Use(router2.JWTAuthMiddleware())

	t.Run("successful make short link", func(t *testing.T) {
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
		token := response["token"]

		url := entity.Urls{
			OriginalUrl: "https://github.com/stretchr/testify",
		}
		jsonData, _ = json.Marshal(url)

		req, _ = http.NewRequest("POST", "/link/makeShortLink", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+token)

		w = httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("missing make short link", func(t *testing.T) {
		url := entity.Urls{
			OriginalUrl: "",
		}
		jsonData, _ := json.Marshal(url)

		req, _ := http.NewRequest("POST", "/link/makeShortLink", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, `{"originalUrl":"String is empty"}`, w.Body.String())
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
