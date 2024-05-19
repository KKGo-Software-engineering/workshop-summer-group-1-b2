package transaction

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/KKGo-Software-engineering/workshop-summer/api/config"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetAllExpenses(t *testing.T) {
	t.Run("get all expenses succesfully", func(t *testing.T) {
		e := echo.New()
		defer e.Close()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		defer db.Close()

		rows := sqlmock.NewRows([]string{"id", "date", "amount", "category", "transaction_type", "note", "image_url", "spender_id"}).
			AddRow(1, "2021-01-01", 100.0, "Food", "expense", "Lunch", "image_url", 1).
			AddRow(2, "2021-01-02", 150.0, "Transport", "expense", "Taxi", "image_url", 2)
		mock.ExpectQuery(`SELECT  id, date, amount, category, transaction_type, note, image_url, spender_id FROM transaction WHERE transaction_type = 'expense'"`).WillReturnRows(rows)

		h := New(config.FeatureFlag{}, db)
		err := h.GetAllTransaction(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, `[{"id":1,"date":"2021-01-01","amount":100,"catergory":"Food","transaction_type":"expense","note":"Lunch","image_url":"image_url","spender_id":1},{"id":2,"date":"2021-01-02","amount":150,"catergory":"Transport","transaction_type":"expense","note":"Taxi","image_url":"image_url","spender_id":2}]`, rec.Body.String())
	})

	t.Run("get all expenses failed on database", func(t *testing.T) {
		e := echo.New()
		defer e.Close()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		defer db.Close()

		mock.ExpectQuery(`SELECT  id, date, amount, category, transaction_type, note, image_url, spender_id FROM transaction WHERE transaction_type = 'expense'"`).WillReturnError(assert.AnError)

		h := New(config.FeatureFlag{}, db)
		err := h.GetAllTransaction(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})
}
