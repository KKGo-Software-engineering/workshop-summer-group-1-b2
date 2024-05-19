package transaction

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/KKGo-Software-engineering/workshop-summer/api/config"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var expectSummaryExpesnse = SummaryExpenses{
	TotalAmountSpent:     2800,
	AvgAmountSpentPerDay: 1400,
	TotalNumberSpent:     2,
}

var expectSummaryIncome = SummaryIncome{
	TotalAmountEarned:     2800,
	AvgAmountEarnedPerDay: 1400,
	TotalNumberEarned:     2,
}

var expectBalance = Balance{
	TotalAmountEarned: 4000,
	TotalAmountSpent:  1000,
	TotalAmountSaved:  3000,
}

func TestGetSummaryExpenses(t *testing.T) {

	t.Run("have error with query should return error", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/api/v1/expenses/summary", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		defer db.Close()

		sqlmock.NewRows([]string{"id", "date", "amount", "category", "transaction_type", "note", "image_url", "spender_id"})
		mock.ExpectQuery(`SELECT FROM transaction WHERE transaction_type = $1`).WithArgs("expense").WillReturnError(assert.AnError)
		cfg := config.FeatureFlag{EnableCreateSpender: true}

		p := New(cfg, db)
		err := p.GetSummaryExpensesHandler(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})

	t.Run("get all summary expenses should return list", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/api/v1/expenses/summary", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		defer db.Close()

		rows := sqlmock.NewRows([]string{"id", "date", "amount", "category", "transaction_type", "note", "image_url", "spender_id"}).
			AddRow(1, "2024-04-30 09:00:00+00", 1000, "Food", "expense", "Lunch", "https://example.com/image1.jpg", 1).
			AddRow(2, "2024-04-27 19:00:00+00", 1800, "Bills", "expense", "Electricity bill", "https://example.com/image9.jpg", 1)

		mock.ExpectQuery(`SELECT * FROM transaction WHERE transaction_type = $1`).WithArgs("expense").WillReturnRows(rows)
		cfg := config.FeatureFlag{EnableCreateSpender: true}

		p := New(cfg, db)
		err := p.GetSummaryExpensesHandler(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)

		m, _ := json.Marshal(expectSummaryExpesnse)
		assert.JSONEq(t, string(m), rec.Body.String())
	})

	t.Run("no transaction should return message transaction not found", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/api/v1/expenses/summary", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		defer db.Close()

		rows := sqlmock.NewRows([]string{"id", "date", "amount", "category", "transaction_type", "note", "image_url", "spender_id"})
		mock.ExpectQuery(`SELECT * FROM transaction WHERE transaction_type = $1`).WithArgs("expense").WillReturnRows(rows)
		cfg := config.FeatureFlag{EnableCreateSpender: true}

		p := New(cfg, db)
		err := p.GetSummaryExpensesHandler(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)

		m, _ := json.Marshal(ResponseMsg{Message: NoTransactionMsg})
		assert.JSONEq(t, string(m), rec.Body.String())
	})

}

func TestGetSummaryIncome(t *testing.T) {

	t.Run("have error with query should return error", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/api/v1/incomes/summary", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		defer db.Close()

		sqlmock.NewRows([]string{"id", "date", "amount", "category", "transaction_type", "note", "image_url", "spender_id"})
		mock.ExpectQuery(`SELECT FROM transaction WHERE transaction_type = $1`).WithArgs("income").WillReturnError(assert.AnError)
		cfg := config.FeatureFlag{EnableCreateSpender: true}

		p := New(cfg, db)
		err := p.GetSummaryIncomeHandler(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})

	t.Run("get all summary income should return list", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/api/v1/incomes/summary", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		defer db.Close()

		rows := sqlmock.NewRows([]string{"id", "date", "amount", "category", "transaction_type", "note", "image_url", "spender_id"}).
			AddRow(1, "2024-04-30 09:00:00+00", 1000, "Food", "income", "Lunch", "https://example.com/image1.jpg", 1).
			AddRow(2, "2024-04-27 19:00:00+00", 1800, "Bills", "income", "Electricity bill", "https://example.com/image9.jpg", 1)

		mock.ExpectQuery(`SELECT * FROM transaction WHERE transaction_type = $1`).WithArgs("income").WillReturnRows(rows)
		cfg := config.FeatureFlag{EnableCreateSpender: true}

		p := New(cfg, db)
		err := p.GetSummaryIncomeHandler(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)

		m, _ := json.Marshal(expectSummaryIncome)
		assert.JSONEq(t, string(m), rec.Body.String())
	})

	t.Run("no transaction should return message transaction not found", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/api/v1/incomes/summary", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		defer db.Close()

		rows := sqlmock.NewRows([]string{"id", "date", "amount", "category", "transaction_type", "note", "image_url", "spender_id"})
		mock.ExpectQuery(`SELECT * FROM transaction WHERE transaction_type = $1`).WithArgs("income").WillReturnRows(rows)
		cfg := config.FeatureFlag{EnableCreateSpender: true}

		p := New(cfg, db)
		err := p.GetSummaryIncomeHandler(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)

		m, _ := json.Marshal(ResponseMsg{Message: NoTransactionMsg})
		assert.JSONEq(t, string(m), rec.Body.String())
	})

}

func TestGetBalanceHandler(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/balance", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	cfg := config.FeatureFlag{EnableCreateSpender: true}

	p := New(cfg, nil)
	err := p.GetSummaryBalanceHandler(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	m, _ := json.Marshal(expectBalance)

	assert.JSONEq(t, string(m), rec.Body.String())
}
