package postgres

import (
	"time"

	"github.com/ericivander/monrep/entity"
	dbEntity "github.com/ericivander/monrep/repository/postgres/entity"
	"github.com/jmoiron/sqlx"
)

// ExpenseProvider holds database connection
type ExpenseProvider struct {
	db *sqlx.DB
}

// NewExpense initiate ExpenseProvider
func NewExpense(db *sqlx.DB) *ExpenseProvider {
	return &ExpenseProvider{db: db}
}

func (p *ExpenseProvider) fetch(query string, args ...interface{}) ([]*entity.Expense, error) {
	rows, err := p.db.Queryx(query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	result := make([]*entity.Expense, 0)

	for rows.Next() {
		tmpExpense := dbEntity.Expense{}
		if err := rows.StructScan(&tmpExpense); err != nil {
			return nil, err
		}

		result = append(result, tmpExpense.ToEntity())
	}

	return result, nil
}

// GetExpensesByMonthYear return Expenses on database by month and year
func (p *ExpenseProvider) GetExpensesByMonthYear(month int, year int) ([]*entity.Expense, error) {
	query := "SELECT id, paid_at, amount_idr, description, category_id FROM expenses WHERE paid_at >= $1 AND paid_at < $2"

	startOfMonth := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	endOfMonth := startOfMonth.AddDate(0, 1, 0)

	return p.fetch(query, startOfMonth, endOfMonth)
}

// GetTotalExpenseBeforeMonthYear return total Expenses before month and year
func (p *ExpenseProvider) GetTotalExpenseBeforeMonthYear(month int, year int) (int64, error) {
	query := "SELECT SUM(amount_id) FROM expenses WHERE paid_at < $1"

	startOfMonth := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)

	var totalExpense int64
	err := p.db.QueryRowx(query, startOfMonth).Scan(&totalExpense)

	return totalExpense, err
}
