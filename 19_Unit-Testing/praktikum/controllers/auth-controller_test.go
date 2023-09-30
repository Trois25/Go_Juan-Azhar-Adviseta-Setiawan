package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"praktikum/config"
	"praktikum/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	e := InitEchoTestAPI()
	InsertDataUserForGetUsers()
	t.Run("Login success", func(t *testing.T) {

		var check = models.User{
			Email:    "alta@gmail.com",
			Password: "123",
		}

		body, err := json.Marshal(check)
		if err != nil {
			t.Error(t, err, "error insert mock data")
		}

		// assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/users/login", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)

		errors := LoginController(c)

		assert.NoError(t, errors)
		assert.Equal(t, http.StatusOK, rec.Code)

		CheckLogin := models.User{}
		err = config.DB.First(&CheckLogin, 1).Error
		assert.NoError(t, err)
		assert.Equal(t, "alta@gmail.com", CheckLogin.Email)
		assert.Equal(t, "123", CheckLogin.Password)
	})

	t.Run("data not found",func(t *testing.T) {
		type userLoginFail struct {
			Message string
			Statis  string
		}

		req := httptest.NewRequest(http.MethodPost, "/users/login", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, LoginController(c)) {
			bodyResponses := rec.Body.String()
			var user userLoginFail

			err := json.Unmarshal([]byte(bodyResponses), &user)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusInternalServerError, rec.Code)
			assert.Equal(t, "error login", user.Message)

		}
	})
}
