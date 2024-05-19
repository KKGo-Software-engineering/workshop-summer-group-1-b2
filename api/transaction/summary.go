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

func (h handler) GetSummaryIncomeHandler(c echo.Context) error {
	s := GetSummary(getTransection())

	return c.JSON(http.StatusOK, SummaryIncome{
		TotalAmountEarned:     s.Total,
		AvgAmountEarnedPerDay: s.Average,
		TotalNumberEarned:     s.TotalNumber,
	})
}

func (h handler) GetSummaryBalanceHandler(c echo.Context) error {

	var mockTransaction = []Transaction{
		{
			Id:              1,
			Date:            "2024-10-30T09:00:00.000Z",
			Amount:          1000,
			Catergory:       "Food",
			TransactionType: "income",
			Note:            "Lunch",
			ImageUrl:        "https://example.com/image21.jpg",
			SpenderId:       1,
		},
		{
			Id:              2,
			Date:            "2024-10-29T19:00:00.000Z",
			Amount:          2000,
			Catergory:       "Transport",
			TransactionType: "income",
			Note:            "Salary",
			ImageUrl:        "https://example.com/image22.jpg",
			SpenderId:       1,
		},
		{
			Id:              3,
			Date:            "2024-10-29T19:00:00.000Z",
			Amount:          2000,
			Catergory:       "Transport",
			TransactionType: "income",
			Note:            "Salary",
			ImageUrl:        "https://example.com/image23.jpg",
			SpenderId:       1,
		},
		{
			Id:              4,
			Date:            "2024-10-29T19:00:00.000Z",
			Amount:          2000,
			Catergory:       "Transport",
			TransactionType: "expense",
			Note:            "Salary",
			ImageUrl:        "https://example.com/image24.jpg",
			SpenderId:       1,
		},
	}

	s := GetBalance(mockTransaction)

	return c.JSON(http.StatusOK, Balance{
		TotalAmountEarned: s.TotalAmountEarned,
		TotalAmountSpent:  s.TotalAmountSpent,
		TotalAmountSaved:  s.TotalAmountSaved,
	})

}
