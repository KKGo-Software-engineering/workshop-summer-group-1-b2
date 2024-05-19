package transaction

import (
	_ "database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
)

type RequestData struct {
	Date            string  `json:"date"`
	Amount          float64 `json:"amount"`
	Category       string  `json:"category"`
	TransactionType string  `json:"transaction_type"`
	Note            string  `json:"note"`
	ImageUrl        string  `json:"image_url"`
	SpenderId       int64   `json:"spender_id"`
}

type ResponseData struct {
	ID int64 `json:"id"`
}

type FeatureFlagTransaction struct {
	EnableCreateTransaction bool `env:"ENABLE_CREATE_TRANSACTION"`
}

const (
	stmt = `INSERT INTO transaction (date, amount, category, transaction_type, note, image_url, spender_id) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id;`
)


func (h handler) CreateTransaction(c echo.Context) error {
	var request RequestData

	ctx := c.Request().Context()

	// Unmarshal the request body into the request data
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// Execute the query
	var id int64
	err := h.db.QueryRowContext(ctx, stmt, request.Date, request.Amount, request.Category, request.TransactionType, request.Note, request.ImageUrl, request.SpenderId).Scan(&id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, ResponseData{
		ID: id,
	})
}
