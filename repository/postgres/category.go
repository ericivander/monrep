package postgres

import (
	"github.com/ericivander/monrep/entity"
	dbEntity "github.com/ericivander/monrep/repository/postgres/entity"
	"github.com/jmoiron/sqlx"
)

// CategoryProvider holds database connection
type CategoryProvider struct {
	db *sqlx.DB
}

// NewCategory initiate CategoryProvider
func NewCategory(db *sqlx.DB) *CategoryProvider {
	return &CategoryProvider{db: db}
}

func (p *CategoryProvider) fetch(query string, args ...interface{}) ([]*entity.Category, error) {
	rows, err := p.db.Queryx(query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	result := make([]*entity.Category, 0)

	for rows.Next() {
		tmpCategory := dbEntity.Category{}
		if err := rows.StructScan(&tmpCategory); err != nil {
			return nil, err
		}

		result = append(result, tmpCategory.ToEntity())
	}

	return result, nil
}

// GetCategories return all categories on database
func (p *CategoryProvider) GetCategories() ([]*entity.Category, error) {
	return p.fetch("SELECT id, name, type FROM categories")
}
