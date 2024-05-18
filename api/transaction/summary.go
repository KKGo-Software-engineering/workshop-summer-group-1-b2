package transaction

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

type Balance struct {
	TotalAmountEarned float64 `json:"total_amount_earned"`
	TotalAmountSpent  float64 `json:"total_amount_spent"`
	TotalAmountSaved  float64 `json:"total_amount_saved"`
}
