package transaction

import (
	"net/http"

	"github.com/KKGo-Software-engineering/workshop-summer/api/mlog"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func (h handler) GetSummaryExpensesHandler(c echo.Context) error {
	logger := mlog.L(c)
	ctx := c.Request().Context()

	rows, err := h.db.QueryContext(ctx, `SELECT * FROM transaction WHERE transaction_type = $1`, "expense")
	if err != nil {
		logger.Error("query error", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	defer rows.Close()

	var ts []Transaction
	for rows.Next() {
		var t Transaction

		err := rows.Scan(&t.Id, &t.Date, &t.Amount, &t.Catergory, &t.TransactionType, &t.Note, &t.ImageUrl, &t.SpenderId)
		if err != nil {
			logger.Error("query error", zap.Error(err))
			return c.JSON(http.StatusInternalServerError, ResponseMsg{
				Message: QueryErrorMsg,
			})
		}

		ts = append(ts, t)
	}

	if len(ts) == 0 {
		return c.JSON(http.StatusOK, ResponseMsg{
			Message: NoTransactionMsg,
		})
	}

	s := GetSummary(ts)

	return c.JSON(http.StatusOK, SummaryExpenses{
		TotalAmountSpent:     s.Total,
		AvgAmountSpentPerDay: s.Average,
		TotalNumberSpent:     s.TotalNumber,
	})
}

func (h handler) GetSummaryIncomeHandler(c echo.Context) error {
	s := GetSummary(getTransection())

	return c.JSON(http.StatusOK, SummaryIncome{
		TotalAmountEarned:     s.Total,
		AvgAmountEarnedPerDay: s.Average,
		TotalNumberEarned:     s.TotalNumber,
	})
}

func (h handler) GetSummaryBalanceHandler(c echo.Context) error {

	s := GetBalance(getTransection())

	return c.JSON(http.StatusOK, Balance{
		TotalAmountEarned: s.TotalAmountEarned,
		TotalAmountSpent:  s.TotalAmountSpent,
		TotalAmountSaved:  s.TotalAmountSaved,
	})

}
