package transaction

import (
	"net/http"

	"github.com/KKGo-Software-engineering/workshop-summer/api/config"
	"github.com/kkgo-software-engineering/workshop/mlog"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func (h handler) GetAllTransaction(c echo.Context) error {
	logger := mlog.L(c)
	ctx := c.Request().Context()

	rows, err := h.db.QueryContext(ctx, `SELECT  id, date, amount, category, transaction_type, note, image_url, spender_id FROM transaction WHERE transaction_type = 'expense'"`)
	if err != nil {
		logger.Error("query error", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer rows.Close()

	var trans []Transaction
	for rows.Next() {
		var tran Transaction
		err := rows.Scan(&tran.Id, &tran.Date, &tran.Amount, &tran.Catergory, &tran.TransactionType, &tran.Note, &tran.ImageUrl, &tran.SpenderId)
		if err != nil {
			logger.Error("scan error", zap.Error(err))
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		trans = append(trans, tran)
	}

	return c.JSON(http.StatusOK, trans)
}

// func GetAllTransaction(db *sql.DB) func(c echo.Context) error {
// 	return func(c echo.Context) error {
// 		logger := mlog.L(c)
// 		ctx := c.Request().Context()

// 		rows, err := db.QueryContext(ctx, `SELECT  id, date, amount, category, transaction_type, note, image_url, spender_id FROM transaction WHERE transaction_type = 'expense'"`)

// 		if err != nil {
// 			logger.Error("query error", zap.Error(err))
// 			return c.JSON(http.StatusInternalServerError, err.Error())
// 		}
// 		defer rows.Close()

// 		var trans []Transaction
// 		for rows.Next() {
// 			var tran Transaction
// 			err := rows.Scan(&tran.Id, &tran.Date, &tran.Amount, &tran.Catergory, &tran.TransectionType, &tran.Note, &tran.ImageUrl, &tran.SpenderId)
// 			if err != nil {
// 				logger.Error("scan error", zap.Error(err))
// 				return c.JSON(http.StatusInternalServerError, err.Error())
// 			}
// 			trans = append(trans, tran)
// 		}

// 		return c.JSON(http.StatusOK, trans)

// 	}

// }
