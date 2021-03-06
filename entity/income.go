package entity

import "time"

// Income is struct for income
type Income struct {
	ID           int64     `json:"id"`
	ReceivedAt   time.Time `json:"received_at"`
	AmountIDR    int64     `json:"amount_idr"`
	Description  string    `json:"description"`
	CategoryID   int64     `json:"-"`
	CategoryName string    `json:"category_name"`
}
