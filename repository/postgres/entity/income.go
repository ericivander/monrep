package entity

import (
	"time"

	"github.com/ericivander/monrep/entity"
)

// Income is struct for income database structure
type Income struct {
	ID          int64     `db:"id"`
	ReceivedAt  time.Time `db:"received_at"`
	AmountIDR   int64     `db:"amount_idr"`
	Description string    `db:"description"`
	CategoryID  int64     `db:"category_id"`
}

// ToEntity convert income on database entity to income on monrep entity
func (i *Income) ToEntity() *entity.Income {
	return &entity.Income{
		ID:          i.ID,
		ReceivedAt:  i.ReceivedAt,
		AmountIDR:   i.AmountIDR,
		Description: i.Description,
		CategoryID:  i.CategoryID,
	}
}
