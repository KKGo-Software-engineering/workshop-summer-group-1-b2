package transaction

import "testing"

var mockSpecificTypeTransaction = []Transaction{
	{
		Id:              1,
		Date:            "2024-04-30T09:00:00.000Z",
		Amount:          1000,
		Catergory:       "Food",
		TransactionType: "income",
		Note:            "Lunch",
		ImageUrl:        "https://example.com/image1.jpg",
		SpenderId:       1,
	},
	{
		Id:              2,
		Date:            "2024-04-29T19:00:00.000Z",
		Amount:          2000,
		Catergory:       "Transport",
		TransactionType: "income",
		Note:            "Salary",
		ImageUrl:        "https://example.com/image2.jpg",
		SpenderId:       1,
	},
	{
		Id:              3,
		Date:            "2024-04-28T19:00:00.000Z",
		Amount:          2000,
		Catergory:       "Transport",
		TransactionType: "income",
		Note:            "Salary",
		ImageUrl:        "https://example.com/image3.jpg",
		SpenderId:       1,
	},
}

func TestGetSummary(t *testing.T) {
	s := GetSummary(mockSpecificTypeTransaction)

	if s.Total != 5000 {
		t.Errorf("Expected total amount spent is 1000, but got %v", s.Total)
	}

	if s.Average != 1666.67 {
		t.Errorf("Expected average amount spent per day is 1000, but got %v", s.Average)
	}

	if s.TotalNumber != 3 {
		t.Errorf("Expected total number spent is 1, but got %v", s.TotalNumber)
	}
}

func TestGetBalance(t *testing.T) {
	s := GetBalance(getTransection())

	if s.TotalAmountEarned != 4000 {
		t.Errorf("Expected total amount spent is 1000, but got %v", s.TotalAmountEarned)
	}

	if s.TotalAmountSpent != 1000 {
		t.Errorf("Expected average amount spent per day is 1000, but got %v", s.TotalAmountSpent)
	}

	if s.TotalAmountSaved != 3000 {
		t.Errorf("Expected total number spent is 1, but got %v", s.TotalAmountSaved)
	}
}
