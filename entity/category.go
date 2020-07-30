package entity

import "github.com/ericivander/monrep/entity/types"

// Category is struct for category
type Category struct {
	ID   int64              `json:"id"`
	Name string             `json:"name"`
	Type types.CategoryType `json:"type"`
}
