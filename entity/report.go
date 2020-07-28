package entity

// Report is struct for monthly report
type Report struct {
	Month        int        `json:"month"`
	Year         int        `json:"year"`
	Incomes      []*Income  `json:"incomes"`
	Expenses     []*Expense `json:"expenses"`
	StartBalance int64      `json:"start_balance"`
	EndBalance   int64      `json:"end_balance"`
}
