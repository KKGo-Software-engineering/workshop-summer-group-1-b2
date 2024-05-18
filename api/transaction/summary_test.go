package transaction

import "testing"

func TestGetSummaryExpenses(t *testing.T) {
	summaryExpenses := GetSummaryExpenses()

	if summaryExpenses.TotalAmountSpent != 1000 {
		t.Errorf("Expected total amount spent is 1000, but got %v", summaryExpenses.TotalAmountSpent)
	}

	if summaryExpenses.AvgAmountSpentPerDay != 1000 {
		t.Errorf("Expected average amount spent per day is 1000, but got %v", summaryExpenses.AvgAmountSpentPerDay)
	}

	if summaryExpenses.TotalNumberSpent != 1 {
		t.Errorf("Expected total number spent is 1, but got %v", summaryExpenses.TotalNumberSpent)
	}
}
