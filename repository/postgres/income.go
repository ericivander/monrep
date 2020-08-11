package postgres

import (
	"time"

	"github.com/ericivander/monrep/entity"
	dbEntity "github.com/ericivander/monrep/repository/postgres/entity"
	"github.com/jmoiron/sqlx"
)

// IncomeProvider holds database connection
type IncomeProvider struct {
	db *sqlx.DB
}

// NewIncome initiate IncomeProvider
func NewIncome(db *sqlx.DB) *IncomeProvider {
	return &IncomeProvider{db: db}
}

func (p *IncomeProvider) fetch(query string, args ...interface{}) ([]*entity.Income, error) {
	rows, err := p.db.Queryx(query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	result := make([]*entity.Income, 0)

	for rows.Next() {
		tmpIncome := dbEntity.Income{}
		if err := rows.StructScan(&tmpIncome); err != nil {
			return nil, err
		}

		result = append(result, tmpIncome.ToEntity())
	}

	return result, nil
}

// GetIncomesByMonthYear return incomes on database by month and year
func (p *IncomeProvider) GetIncomesByMonthYear(month int, year int) ([]*entity.Income, error) {
	query := "SELECT id, received_at, amount_idr, description, category_id FROM incomes WHERE received_at >= $1 AND received_at < $2"

	startOfMonth := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	endOfMonth := startOfMonth.AddDate(0, 1, 0)

	return p.fetch(query, startOfMonth, endOfMonth)
}

// GetTotalIncomeBeforeMonthYear return total incomes before month and year
func (p *IncomeProvider) GetTotalIncomeBeforeMonthYear(month int, year int) (int64, error) {
	query := "SELECT SUM(amount_id) FROM incomes WHERE received_at < $1"

	startOfMonth := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)

	var totalIncome int64
	err := p.db.QueryRowx(query, startOfMonth).Scan(&totalIncome)

	return totalIncome, err
}
