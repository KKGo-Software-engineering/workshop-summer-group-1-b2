package transaction

import (
	_ "database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

const (
	stmtUpdate = `UPDATE transaction SET date = $1, amount = $2, category = $3, image_url = $4 WHERE id = $5 RETURNING id,date,amount,category,image_url;`
)

type RequestDataUpdate struct {
	Date     string  `json:"date"`
	Amount   float64 `json:"amount"`
	Category string  `json:"category"`
	ImageUrl string  `json:"image_url"`
}

type ResponseDataUpdate struct {
	ID       int64   `json:"id"`
	Date     string  `json:"date"`
	Amount   float64 `json:"amount"`
	Category string  `json:"category"`
	ImageUrl string  `json:"image_url"`
}

func (h handler) UpdateTransaction(c echo.Context) error {
	var request RequestDataUpdate

	ctx := c.Request().Context()

	// Unmarshal the request body into the request data
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// Execute the query
	fmt.Println("-", c.Param("id"))

	idRaw, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	id := int64(idRaw)

	err = h.db.QueryRowContext(ctx, stmtUpdate, request.Date, request.Amount, request.Category, request.ImageUrl, id).Scan(&id, &request.Date, &request.Amount, &request.Category, &request.ImageUrl)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, ResponseDataUpdate{
		ID:       id,
		Date:     request.Date,
		Amount:   request.Amount,
		Category: request.Category,
		ImageUrl: request.ImageUrl,
	})
}
