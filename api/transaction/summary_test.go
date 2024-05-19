package transaction

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/KKGo-Software-engineering/workshop-summer/api/config"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var mockSummaryExpesnse = SummaryExpenses{
	TotalAmountSpent:     5000,
	AvgAmountSpentPerDay: 1666.67,
	TotalNumberSpent:     3,
}

var mockSummaryIncome = SummaryIncome{
	TotalAmountEarned:     5000,
	AvgAmountEarnedPerDay: 1666.67,
	TotalNumberEarned:     3,
}

func TestGetSummaryExpenses(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/expenses/summary", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	cfg := config.FeatureFlag{EnableCreateSpender: true}

	p := New(cfg, nil)
	err := p.GetSummaryExpensesHandler(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	m, _ := json.Marshal(mockSummaryExpesnse)

	assert.JSONEq(t, string(m), rec.Body.String())

}

func TestGetSummaryIncome(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/expenses/summary", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	cfg := config.FeatureFlag{EnableCreateSpender: true}

	p := New(cfg, nil)
	err := p.GetSummaryIncomeHandler(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	m, _ := json.Marshal(mockSummaryIncome)

	assert.JSONEq(t, string(m), rec.Body.String())

}
