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

func TestWriteShortLink(t *testing.T) {
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("failed to setup test database: %v", err)
	}
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.POST("/auth/login", controller.AuthUser(db))

	protected := router.Group("/")
	protected.Use(router2.JWTAuthMiddleware())
	{
		protected.POST("/link/writeShortLink", controller.WriteShortLink(db))
	}

	t.Run("successful write short link", func(t *testing.T) {
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
			ShortUrl:    "testshort",
		}
		jsonData, _ = json.Marshal(url)

		req, _ = http.NewRequest("POST", "/link/writeShortLink", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+token)

		w = httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("missing fields in write short link", func(t *testing.T) {
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
			OriginalUrl: "",
			ShortUrl:    "",
		}
		jsonData, _ = json.Marshal(url)

		req, _ = http.NewRequest("POST", "/link/writeShortLink", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+token)

		w = httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, `{"originalUrl":"original or short url is empty"}`, w.Body.String())
	})

	t.Run("invalid URL in write short link", func(t *testing.T) {
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
			OriginalUrl: "invalid_url",
			ShortUrl:    "testshort",
		}
		jsonData, _ = json.Marshal(url)

		req, _ = http.NewRequest("POST", "/link/writeShortLink", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+token)

		w = httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, `{"originalUrl":"URL is not validated"}`, w.Body.String())
	})

	t.Run("existing short URL in write short link", func(t *testing.T) {
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
			ShortUrl:    "testshort",
		}
		jsonData, _ = json.Marshal(url)

		req, _ = http.NewRequest("POST", "/link/writeShortLink", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+token)

		w = httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		req, _ = http.NewRequest("POST", "/link/writeShortLink", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+token)

		w = httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, `{"originalUrl":"short url is exist"}`, w.Body.String())
	})
}
