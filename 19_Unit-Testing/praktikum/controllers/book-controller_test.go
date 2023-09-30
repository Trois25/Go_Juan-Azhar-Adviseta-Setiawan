package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"praktikum/config"
	"praktikum/models"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var mockDataBook = models.Book{
	Judul:    "Alta",
	Penulis:  "wawan",
	Penerbit: "upn",
}

func InsertDataBookForGetBooks() error {
	config.InitDBTest()
	book := models.Book{
		Judul:    "Alta",
		Penulis:  "wawan",
		Penerbit: "upn",
	}

	var err error
	if err = config.DB.Save(&book).Error; err != nil {
		return err
	}
	return nil
}

func TestGetBooksController(t *testing.T) {

	e := InitEchoTestAPI()
	InsertDataBookForGetBooks()

	type TestCase []struct {
		name       string
		path       string
		expectCode int
		sizeData   int
	}

	acceptedcase := TestCase{
		{"get all book", "/books", http.StatusOK, 1},
	}

	for _, testCase := range acceptedcase {
		req := httptest.NewRequest(http.MethodGet, testCase.path, nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, GetBooksController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
			body := rec.Body.String()

			// Parse JSON response
			var bookResponse struct {
				Message string        `json:"message"`
				Data    []models.Book `json:"data"`
			}
			err := json.Unmarshal([]byte(body), &bookResponse)

			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, testCase.sizeData, len(bookResponse.Data))
		}
	}
	t.Run("failed get book", func(t *testing.T) {
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
			path:       "/books",
			expectCode: http.StatusBadRequest,
		}

		req := httptest.NewRequest(http.MethodGet, "/books", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		config.DB.Migrator().DropTable(&models.Book{})

		GetBooksController(c)

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

func TestGetSpecificBookController(t *testing.T) {

	e := InitEchoTestAPI()
	InsertDataBookForGetBooks()

	var testCases = []struct {
		name       string
		path       string
		expectCode int
		sizeData   int
	}{
		{
			name:       "book found",
			path:       "/books",
			expectCode: http.StatusOK,
			sizeData:   1,
		},
	}

	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodGet, testCase.path, nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath("/books/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")

		if assert.NoError(t, GetBookController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
			body := rec.Body.String()

			// Parse JSON response
			var bookResponse struct {
				Message string      `json:"message"`
				Data    models.Book `json:"data"`
			}
			err := json.Unmarshal([]byte(body), &bookResponse)

			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, "Alta", bookResponse.Data.Judul)
		}
	}

	t.Run("failed get book data", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/books/10", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := GetBookController(c)

		c.SetPath("/users/:id")
		c.SetParamNames("id")
		c.SetParamValues("10")

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

}

// Delete user
func TestDeleteSpecificBookController(t *testing.T) {

	e := InitEchoTestAPI()

	var testCases = []struct {
		name       string
		path       string
		expectCode int
		sizeData   int
	}{
		{
			name:       "book deleted success",
			path:       "/books",
			expectCode: http.StatusOK,
			sizeData:   1,
		},
	}

	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodGet, testCase.path, nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath("/books/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")

		if assert.NoError(t, DeleteBookController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
			body := rec.Body.String()

			// Parse JSON response
			var bookResponse struct {
				Message string      `json:"message"`
				Data    models.Book `json:"data"`
			}
			err := json.Unmarshal([]byte(body), &bookResponse)

			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, "", bookResponse.Data.Judul)
		}
	}

	t.Run("failed delete book data", func(t *testing.T) {

		config.DB.Migrator().DropTable(&models.Book{})

		req := httptest.NewRequest(http.MethodGet, "/books/10", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := DeleteBookController(c)

		c.SetPath("/users/:id")
		c.SetParamNames("id")
		c.SetParamValues("10")

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

}

func TestCreateBookController(t *testing.T) {

	e := InitEchoTestAPI()
	InsertDataBookForGetBooks()

	type bookResponse struct {
		Message string              `json:"message"`
		Data    models.Book `json:"data"`
	}

	type bookResponseFail struct {
		Message string
		Statis  string
	}

	var testCases = struct {
		name       string
		path       string
		expectCode int
	}{

		name:       "success to create book",
		path:       "/books",
		expectCode: http.StatusOK,
	}

	body, err := json.Marshal(mockDataBook)
	if err != nil {
		t.Error(t, err, "error insert mock data")
	}

	req := httptest.NewRequest(http.MethodPost, "/books", bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if assert.NoError(t, CreateBookController(c)) {
		bodyResponses := rec.Body.String()
		var book bookResponse

		err := json.Unmarshal([]byte(bodyResponses), &book)
		if err != nil {
			assert.Error(t, err, "error")
		}

		assert.Equal(t, testCases.expectCode, rec.Code)
		assert.Equal(t, "Alta", book.Data.Judul)
		assert.Equal(t, "wawan", book.Data.Penulis)
		assert.Equal(t, "upn", book.Data.Penerbit)
		
	}

	t.Run("failed create user", func(t *testing.T) {

		config.DB.Migrator().DropTable(&models.Book{})

		req := httptest.NewRequest(http.MethodPost, "/books", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, CreateBookController(c)) {
			bodyResponses := rec.Body.String()
			var book bookResponseFail

			err := json.Unmarshal([]byte(bodyResponses), &book)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, "Failed create book", book.Message)

		}
	})

}

func TestUpdateBookController(t *testing.T) {

	type bookResponseFail struct {
		Message string
		Statis  string
	}
	e := InitEchoTestAPI()
	InsertDataBookForGetBooks()

	t.Run("book updated", func(t *testing.T) {
        // Persiapkan data pembaruan yang berbeda
        mockUpdateData := models.Book{
            Judul:    "Alta updated",
			Penulis:  "wawan updated",
			Penerbit: "upn updated",
        }
        requestBody, fail := json.Marshal(mockUpdateData)
		if fail != nil {
			t.Error(t, fail, "error update book data")
		}

        // Buat permintaan PUT dengan parameter id yang benar
        req := httptest.NewRequest(http.MethodPut, "/books/1", bytes.NewBuffer(requestBody))
        req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
        rec := httptest.NewRecorder()
		

        c := e.NewContext(req, rec)

		c.SetPath("/books/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")

        err := UpdateBookController(c)

        // Memeriksa hasil respons
        assert.NoError(t, err)
        assert.Equal(t, http.StatusOK, rec.Code)

        // Memeriksa apakah data pengguna telah diperbarui di dalam database
        updateBook := models.Book{}
        err = config.DB.First(&updateBook, 1).Error
        assert.NoError(t, err)
		
        assert.Equal(t, "Alta updated", updateBook.Judul)
        assert.Equal(t, "wawan updated", updateBook.Penulis)
        assert.Equal(t, "upn updated", updateBook.Penerbit)
    })

	t.Run("book not found", func(t *testing.T) {

		// config.DB.Migrator().DropTable(&models.User{})

		req := httptest.NewRequest(http.MethodPost, "/books/100", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, UpdateBookController(c)) {
			bodyResponses := rec.Body.String()
			var user bookResponseFail

			err := json.Unmarshal([]byte(bodyResponses), &user)
			if err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, "Failed update book", user.Message)

		}
	})

}
