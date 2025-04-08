package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestUpdateUserInfo(t *testing.T) {
	r := gin.Default()
	r.PUT("/users", UpdateUserInfo())

	// Тест на успешное обновление данных пользователя
	t.Run("return 200 status code when user info is updated", func(t *testing.T) {
		updateData := UpdateData{
			UserName:    "testuser",
			NewNameUser: "newuser",
			NewPassword: "newpassword123",
		}
		jsonData, _ := json.Marshal(updateData)

		req, _ := http.NewRequest(http.MethodPut, "/users", bytes.NewBuffer(jsonData))
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assertStatus(t, w.Code, http.StatusOK)
	})

	// Тест на ситуацию, когда переданы некорректные данные
	t.Run("return 400 status code when update data is invalid", func(t *testing.T) {
		updateData := UpdateData{
			UserName:    "",
			NewNameUser: "",
			NewPassword: "",
		}
		jsonData, _ := json.Marshal(updateData)

		req, _ := http.NewRequest(http.MethodPut, "/users", bytes.NewBuffer(jsonData))
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assertStatus(t, w.Code, http.StatusBadRequest)
	})
}
