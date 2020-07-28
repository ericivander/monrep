package flow

import "github.com/ericivander/monrep/entity"

// CategoryRepository holds contracts for category database functions
type CategoryRepository interface {
	GetCategories() ([]*entity.Category, error)
}

// IncomeRepository holds contracts for income database functions
type IncomeRepository interface {
	GetIncomesByMonthYear(month int, year int) ([]*entity.Income, error)
	GetTotalIncomeBeforeMonthYear(month int, year int) (int64, error)
}

// ExpenseRepository holds contracts for expense database functions
type ExpenseRepository interface {
	GetExpensesByMonthYear(month int, year int) ([]*entity.Expense, error)
	GetTotalExpenseBeforeMonthYear(month int, year int) (int64, error)
}
