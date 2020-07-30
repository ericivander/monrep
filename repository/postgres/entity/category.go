package entity

import (
	"github.com/ericivander/monrep/entity"
	"github.com/ericivander/monrep/entity/types"
)

// Category is struct for category on database
type Category struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
	Type int16  `db:"type"`
}

// ToEntity convert database object to entity
func (c *Category) ToEntity() *entity.Category {
	return &entity.Category{
		ID:   c.ID,
		Name: c.Name,
		Type: types.CategoryType(c.Type),
	}
}
