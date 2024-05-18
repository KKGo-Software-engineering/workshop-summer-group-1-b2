package transaction

type Transaction struct {
	Id              int64   `json:"id"`
	Date            string  `json:"date"`
	Amount          float64 `json:"amount"`
	Catergory       string  `json:"catergory"`
	TransectionType string  `json:"transection_type"`
	Note            string  `json:"note"`
	ImageUrl        string  `json:"image_url"`
	SpenderId       int64   `json:"spender_id"`
}
