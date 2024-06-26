package transaction

import (
	"database/sql"

	"github.com/KKGo-Software-engineering/workshop-summer/api/config"
)

type Transaction struct {
	Id              int64   `json:"id"`
	Date            string  `json:"date"`
	Amount          float64 `json:"amount"`
	Catergory       string  `json:"catergory"`
	TransactionType string  `json:"transaction_type"`
	Note            string  `json:"note"`
	ImageUrl        string  `json:"image_url"`
	SpenderId       int64   `json:"spender_id"`
}

type ResponseMsg struct {
	Message string `json:"message"`
}

const (
	NoTransactionMsg = "transaction not found"
	QueryErrorMsg    = "something went wrong with query"
)

type handler struct {
	flag config.FeatureFlag
	db   *sql.DB
}

func New(cfg config.FeatureFlag, db *sql.DB) *handler {
	return &handler{cfg, db}
}
