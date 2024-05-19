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

func TestCreateTransaction(t *testing.T) {
	t.Run("can update transaction", func(t *testing.T) {
		e := echo.New()
		defer e.Close()

		req := httptest.NewRequest(http.MethodPost, "/transactions", strings.NewReader(`{"date": "2021-08-01", "amount": 1000.0, "category": "food", "transaction_type": "expense", "note": "lunch", "image_url": "https://example.com/image.jpg", "spender_id": 1}`))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		defer db.Close()

		row := sqlmock.NewRows([]string{"id"}).AddRow(1)
		mock.ExpectQuery(stmt).WithArgs("2021-08-01", 1000.0, "food", "expense", "lunch", "https://example.com/image.jpg", 1).WillReturnRows(row)
		cfg := config.FeatureFlag{EnableCreateSpender: true}

		h := New(cfg, db)
		err := h.CreateTransaction(c)
		if err != nil {
			t.Errorf("error should be nil, got %v", err)
		}

		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.JSONEq(t, `{"id": 1}`, rec.Body.String())
	})
}
