package transaction

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h handler) GetSummaryExpensesHandler(c echo.Context) error {

	// logger := mlog.L(c)
	// ctx := c.Request().Context()

	// rows, err := h.db.QueryContext(ctx, `SELECT * FROM transaction`)
	// if err != nil {
	// 	logger.Error("query error", zap.Error(err))
	// 	return c.JSON(http.StatusInternalServerError, err.Error())

	// }

	// defer rows.Close()

	// var transactions []Transaction

	// for rows.Next() {
	// 	var t Transaction
	// 	err := rows.Scan(&t.Id, &t.Date, &t.Amount, &t.Catergory, &t.TransectionType, &t.Note, &t.ImageUrl, &t.SpenderId)
	// 	if err != nil {
	// 		logger.Error("scan error", zap.Error(err))
	// 		return c.JSON(http.StatusInternalServerError, err.Error())
	// 	}

	// 	transactions = append(transactions, t)
	// }

	// print(transactions)

	// summaryExpenses := GetSummaryExpenses(getTransection())
	s := GetSummary(getTransection())

	return c.JSON(http.StatusOK, SummaryExpenses{
		TotalAmountSpent:     s.Total,
		AvgAmountSpentPerDay: s.Average,
		TotalNumberSpent:     s.TotalNumber,
	})
}
