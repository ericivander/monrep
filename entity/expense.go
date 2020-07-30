package entity

import "time"

// Expense is struct for income
type Expense struct {
	ID           int64     `json:"id"`
	PaidAt       time.Time `json:"paid_at"`
	AmountIDR    int64     `json:"amount_idr"`
	Description  string    `json:"description"`
	CategoryID   int64     `json:"-"`
	CategoryName string    `json:"category_name"`
}
