package entity

import "time"

// Income is struct for income
type Income struct {
	ID           int       `json:"id"`
	ReceivedAt   time.Time `json:"received_at"`
	AmountIDR    int64     `json:"amount_idr"`
	Description  string    `json:"description"`
	CategoryID   int       `json:"-"`
	CategoryName string    `json:"category_name"`
}
