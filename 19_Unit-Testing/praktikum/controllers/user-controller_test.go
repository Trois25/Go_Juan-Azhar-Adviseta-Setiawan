package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"praktikum/config"
	"praktikum/models"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var mockData = models.User{
	Name:     "banteng",
	Email:    "banteng@example.com",
	Password: "banteng123",
}

func InitEchoTestAPI() *echo.Echo {
	e := echo.New()
	return e
}

func InsertDataUserForGetUsers() error {
	config.InitDBTest()
	user := models.User{
		Name:     "Alta",
		Password: "123",
		Email:    "alta@gmail.com",
	}

	var err error
	if err = config.DB.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

func TestGetUsersController(t *testing.T) {

	e := InitEchoTestAPI()
	InsertDataUserForGetUsers()

	type TestCase []struct {
		name       string
		path       string
		expectCode int
		sizeData   int
	}

	acceptedcase := TestCase{
		{"get user normal", "/users", http.StatusOK, 1},
	}

	for _, testCase := range acceptedcase {
		req := httptest.NewRequest(http.MethodGet, testCase.path, nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, GetUsersController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
			body := rec.Body.String()

			// Parse JSON response
			var userResponse struct {
				Message string                `json:"message"`
				Data    []models.UserResponse `json:"data"`
			}
			err := json.Unmarshal([]byte(body), &userResponse)

			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, testCase.sizeData, len(userResponse.Data))
		}
	}

	t.Run("failed get data user", func(t *testing.T) {
		type ResponseFailed struct {
			Status  string
			Message string
		}

		var FailTestCase = struct {
			name       string
			path       string
			expectCode int
		}{
			name:       "failed to get data",
			path:       "/users",
			expectCode: http.StatusBadRequest,
		}

		req := httptest.NewRequest(http.MethodGet, "/users", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		config.DB.Migrator().DropTable(&models.User{})

		GetUsersController(c)

		body := rec.Body.String()

		var responses ResponseFailed
		err := json.Unmarshal([]byte(body), &responses)
		assert.Equal(t, FailTestCase.expectCode, rec.Code)
		if err != nil {
			assert.Error(t, err, "error")
		}
		assert.Equal(t, "failed", responses.Status)
		assert.Equal(t, "error get data", responses.Message)
	})
}

func TestGetSpecificUserController(t *testing.T) {

	e := InitEchoTestAPI()
	InsertDataUserForGetUsers()

	var testCases = []struct {
		name       string
		path       string
		expectCode int
		sizeData   int
	}{
		{
			name:       "user found",
			path:       "/users",
			expectCode: http.StatusOK,
			sizeData:   1,
		},
	}

	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodGet, testCase.path, nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath("/users/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")

		if assert.NoError(t, GetUserController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
			body := rec.Body.String()

			// Parse JSON response
			var userResponse struct {
				Message string              `json:"message"`
				Data    models.UserResponse `json:"data"`
			}
			err := json.Unmarshal([]byte(body), &userResponse)

			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, "Alta", userResponse.Data.Name)
		}
	}

	t.Run("failed get user data", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/users/10", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := GetUserController(c)

		c.SetPath("/users/:id")
		c.SetParamNames("id")
		c.SetParamValues("10")

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

}

// Delete user
func TestDeleteSpecificUserController(t *testing.T) {

	e := InitEchoTestAPI()

	var testCases = []struct {
		name       string
		path       string
		expectCode int
		sizeData   int
	}{
		{
			name:       "user deleted success",
			path:       "/users",
			expectCode: http.StatusOK,
			sizeData:   1,
		},
	}

	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodGet, testCase.path, nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath("/users/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")

		if assert.NoError(t, DeleteUserController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
			body := rec.Body.String()

			// Parse JSON response
			var userResponse struct {
				Message string              `json:"message"`
				Data    models.UserResponse `json:"data"`
			}
			err := json.Unmarshal([]byte(body), &userResponse)

			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, "", userResponse.Data.Name)
		}
	}

	t.Run("failed delete user data", func(t *testing.T) {

		config.DB.Migrator().DropTable(&models.User{})

		req := httptest.NewRequest(http.MethodGet, "/users/10", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := DeleteUserController(c)

		c.SetPath("/users/:id")
		c.SetParamNames("id")
		c.SetParamValues("10")

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

}

func TestCreateUserController(t *testing.T) {

	e := InitEchoTestAPI()
	InsertDataUserForGetUsers()

	type userResponse struct {
		Message string              `json:"message"`
		Data    models.UserResponse `json:"data"`
	}

	type userResponseFail struct {
		Message string
		Statis  string
	}

	var testCases = struct {
		name       string
		path       string
		expectCode int
	}{

		name:       "success to create user",
		path:       "/users",
		expectCode: http.StatusOK,
	}

	body, err := json.Marshal(mockData)
	if err != nil {
		t.Error(t, err, "error insert mock data")
	}

	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if assert.NoError(t, CreateUserController(c)) {
		bodyResponses := rec.Body.String()
		var user userResponse

		err := json.Unmarshal([]byte(bodyResponses), &user)
		if err != nil {
			assert.Error(t, err, "error")
		}

		assert.Equal(t, testCases.expectCode, rec.Code)
		assert.Equal(t, "banteng", user.Data.Name)
		assert.Equal(t, "banteng@example.com", user.Data.Email)
	}

	t.Run("failed create user", func(t *testing.T) {

		config.DB.Migrator().DropTable(&models.User{})

		req := httptest.NewRequest(http.MethodPost, "/users", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, CreateUserController(c)) {
			bodyResponses := rec.Body.String()
			var user userResponseFail

			err := json.Unmarshal([]byte(bodyResponses), &user)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, "Failed create user", user.Message)

		}
	})

}

func TestUpdateUserController(t *testing.T) {

	type userResponseFail struct {
		Message string
		Statis  string
	}
	e := InitEchoTestAPI()
	InsertDataUserForGetUsers()

	t.Run("user updated", func(t *testing.T) {
        // Persiapkan data pembaruan yang berbeda
        mockUpdateData := models.User{
            Name:     "Updated Name",
            Email:    "updated@example.com",
            Password: "updatedPassword",
        }
        requestBody, fail := json.Marshal(mockUpdateData)
		if fail != nil {
			t.Error(t, fail, "error update user data")
		}

        // Buat permintaan PUT dengan parameter id yang benar
        req := httptest.NewRequest(http.MethodPut, "/users/1", bytes.NewBuffer(requestBody))
        req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
        rec := httptest.NewRecorder()
		

        c := e.NewContext(req, rec)

		c.SetPath("/users/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")

        err := UpdateUserController(c)
		fmt.Println(err)

        // Memeriksa hasil respons
        assert.NoError(t, err)
        assert.Equal(t, http.StatusOK, rec.Code)

        // Memeriksa apakah data pengguna telah diperbarui di dalam database
        updatedUser := models.User{}
        err = config.DB.First(&updatedUser, 1).Error
        assert.NoError(t, err)
        assert.Equal(t, "Updated Name", updatedUser.Name)
        assert.Equal(t, "updated@example.com", updatedUser.Email)
        assert.Equal(t, "updatedPassword", updatedUser.Password)
    })

	t.Run("user not found", func(t *testing.T) {

		// config.DB.Migrator().DropTable(&models.User{})

		req := httptest.NewRequest(http.MethodPost, "/users/100", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, UpdateUserController(c)) {
			bodyResponses := rec.Body.String()
			var user userResponseFail

			err := json.Unmarshal([]byte(bodyResponses), &user)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, "Failed update user", user.Message)

		}
	})

}
