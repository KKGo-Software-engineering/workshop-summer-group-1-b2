package transaction

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/KKGo-Software-engineering/workshop-summer/api/config"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestUpdateTransaction(t *testing.T) {
	t.Run("can update transaction", func(t *testing.T) {

		e := echo.New()
		defer e.Close()

		req := httptest.NewRequest(http.MethodPut, "/transactions", strings.NewReader(`{"date": "2021-08-02", "amount": 2000.0, "category": "food", "image_url": "https://example.com/image.jpg"}`))

		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")

		db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		defer db.Close()

		row := sqlmock.NewRows([]string{"id","date","amount","category","image_url"}).AddRow(1,"2021-08-02",2000.0,"food" ,"https://example.com/image.jpg")
		mock.ExpectQuery(stmtUpdate).WithArgs("2021-08-02", 2000.0, "food", "https://example.com/image.jpg",1).WillReturnRows(row)
		cfg := config.FeatureFlag{EnableCreateSpender: true}

		h := New(cfg, db)
		err := h.UpdateTransaction(c)
		if err != nil {
			t.Errorf("error should be nil, got %v", err)
		}

		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.JSONEq(t, `{"id": 1, "date": "2021-08-02", "amount": 2000.0, "category": "food", "image_url": "https://example.com/image.jpg"}`, rec.Body.String())
	})
}
