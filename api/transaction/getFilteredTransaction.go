package transaction

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/kkgo-software-engineering/workshop/mlog"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func (h handler) GetFilteredTransaction(c echo.Context) error {
	logger := mlog.L(c)
	ctx := c.Request().Context()

	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	offset := (page - 1) * limit

	query := `SELECT id, date, amount, category, transaction_type, note, image_url, spender_id FROM transaction WHERE 1=1`
	var filters []string
	var args []interface{}

	if len(filters) > 0 {
		query += " AND " + strings.Join(filters, " AND ")
	}
	query += " ORDER BY date DESC OFFSET ? ROWS FETCH NEXT ? ROWS ONLY"
	args = append(args, offset, limit)

	rows, err := h.db.QueryContext(ctx, query, args...)
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
