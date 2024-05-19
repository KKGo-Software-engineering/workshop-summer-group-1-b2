package transaction

import (
	"fmt"
	"math"
)

type Balance struct {
	TotalAmountEarned float64 `json:"total_amount_earned"`
	TotalAmountSpent  float64 `json:"total_amount_spent"`
	TotalAmountSaved  float64 `json:"total_amount_saved"`
}

type SummaryExpenses struct {
	TotalAmountSpent     float64 `json:"total_amount_spent"`
	AvgAmountSpentPerDay float64 `json:"avg_amount_spent_per_day"`
	TotalNumberSpent     int     `json:"total_number_spent"`
}

type SummaryIncome struct {
	TotalAmountEarned     float64 `json:"total_amount_earned"`
	AvgAmountEarnedPerDay float64 `json:"avg_amount_earned_per_day"`
	TotalNumberEarned     int     `json:"total_number_earned"`
}

func getTransection() []Transaction {
	mockTransaction := []Transaction{
		{
			Id:              1,
			Date:            "2024-04-30T09:00:00.000Z",
			Amount:          1000,
			Catergory:       "Food",
			TransectionType: "expense",
			Note:            "Lunch",
			ImageUrl:        "https://example.com/image1.jpg",
			SpenderId:       1,
		},
		{
			Id:              2,
			Date:            "2024-04-29T19:00:00.000Z",
			Amount:          2000,
			Catergory:       "Transport",
			TransectionType: "income",
			Note:            "Salary",
			ImageUrl:        "https://example.com/image2.jpg",
			SpenderId:       1,
		},
		{
			Id:              3,
			Date:            "2024-04-29T19:00:00.000Z",
			Amount:          2000,
			Catergory:       "Transport",
			TransectionType: "income",
			Note:            "Salary",
			ImageUrl:        "https://example.com/image2.jpg",
			SpenderId:       1,
		},
	}

	return mockTransaction
}

type Summary struct {
	Total       float64
	Average     float64
	TotalNumber int
}

func roundedNumber(n float64) float64 {
	fmt.Println(n)
	fmt.Println(n * 100)
	return math.Round(n*100) / 100
}

func GetSummary(ts []Transaction) Summary {
	total := 0.0
	totalNumber := 0

	for _, t := range ts {
		total += t.Amount
		totalNumber++
	}

	return Summary{
		Total:       roundedNumber(total),
		Average:     roundedNumber(total / float64(totalNumber)),
		TotalNumber: totalNumber,
	}
}

// func GetSummaryExpenses(t []Transaction) SummaryExpenses {
// 	totalAmountSpent := 0.0
// 	totalNumberSpent := 0

// 	for _, transaction := range t {
// 		if transaction.TransectionType == "expense" {
// 			totalAmountSpent += transaction.Amount
// 			totalNumberSpent++
// 		}
// 	}

// 	return SummaryExpenses{
// 		TotalAmountSpent:     totalAmountSpent,
// 		AvgAmountSpentPerDay: totalAmountSpent / float64(totalNumberSpent),
// 		TotalNumberSpent:     totalNumberSpent,
// 	}

// }
