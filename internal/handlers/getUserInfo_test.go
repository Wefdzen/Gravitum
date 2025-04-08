package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetUserInfo(t *testing.T) {
	r := gin.Default()
	r.GET("/users/", GetUserInfo())

	// Тест на успешное получение пользователя
	t.Run("return 200 status code when user exists", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/users/?username=testuser", nil) // Используем query параметр
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assertStatus(t, w.Code, http.StatusOK)
	})

	// Тест на ситуацию, когда пользователь не найден
	t.Run("return 404 status code when user does not exist", func(t *testing.T) {

		req, _ := http.NewRequest(http.MethodGet, "/users/?username=ll", nil)
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assertStatus(t, w.Code, http.StatusNotFound)
	})
}
