package entity

import (
	"time"

	"github.com/ericivander/monrep/entity"
)

// Expense is struct for expense database structure
type Expense struct {
	ID          int64     `db:"id"`
	PaidAt      time.Time `db:"paid_at"`
	AmountIDR   int64     `db:"amount_idr"`
	Description string    `db:"description"`
	CategoryID  int64     `db:"category_id"`
}

// ToEntity convert expense on database entity to expense on monrep entity
func (e *Expense) ToEntity() *entity.Expense {
	return &entity.Expense{
		ID:          e.ID,
		PaidAt:      e.PaidAt,
		AmountIDR:   e.AmountIDR,
		Description: e.Description,
		CategoryID:  e.CategoryID,
	}
}
