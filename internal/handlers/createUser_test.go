package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Wefdzen/Gravitum/internal/db"
	"github.com/gin-gonic/gin"
)

func TestCreateUser(t *testing.T) {
	r := gin.Default()
	r.POST("/users", CreateUser()) // Регистрация маршрута

	// Тест на создание пользователя
	t.Run("return 201 status code when user is created", func(t *testing.T) {
		user := db.User{
			UserName: "testuser",
			Password: "password123",
		}
		jsonData, _ := json.Marshal(user)

		req, _ := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(jsonData))
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assertStatus(t, w.Code, http.StatusCreated)
	})

	// Тест на пустые поля username или password
	t.Run("return 400 status code when username or password is empty", func(t *testing.T) {
		user := db.User{
			UserName: "",
			Password: "",
		}
		jsonData, _ := json.Marshal(user)

		req, _ := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(jsonData))
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assertStatus(t, w.Code, http.StatusBadRequest)
	})
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}
