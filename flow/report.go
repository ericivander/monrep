package flow

import "github.com/ericivander/monrep/entity"

// ReportFlow holds contracts for reporting functions
type ReportFlow interface {
	GetMonthlyReport(month int, year int) (*entity.Report, error)
}

// ReportProvider is struct to holds dependencies needed for reporting functions
type ReportProvider struct {
	categoryRepo CategoryRepository
	incomeRepo   IncomeRepository
	expenseRepo  ExpenseRepository
}

// NewReportProvider construct report provider given dependencies
func NewReportProvider(cat CategoryRepository, inc IncomeRepository, exp ExpenseRepository) *ReportProvider {
	return &ReportProvider{
		categoryRepo: cat,
		incomeRepo:   inc,
		expenseRepo:  exp,
	}
}

// GetMonthlyReport return report on specific month and year
func (r *ReportProvider) GetMonthlyReport(month int, year int) (*entity.Report, error) {
	categories, err := r.categoryRepo.GetCategories()
	if err != nil {
		return nil, err
	}

	categoriesName := make(map[int64]string, 0)
	for _, category := range categories {
		categoriesName[category.ID] = category.Name
	}

	incomes, err := r.incomeRepo.GetIncomesByMonthYear(month, year)
	if err != nil {
		return nil, err
	}

	lastTotalIncome, err := r.incomeRepo.GetTotalIncomeBeforeMonthYear(month, year)
	if err != nil {
		return nil, err
	}

	totalIncome := lastTotalIncome
	for _, income := range incomes {
		totalIncome += income.AmountIDR
		income.CategoryName = categoriesName[income.CategoryID]
	}

	expenses, err := r.expenseRepo.GetExpensesByMonthYear(month, year)
	if err != nil {
		return nil, err
	}

	lastTotalExpense, err := r.expenseRepo.GetTotalExpenseBeforeMonthYear(month, year)
	if err != nil {
		return nil, err
	}

	totalExpense := lastTotalExpense
	for _, expense := range expenses {
		totalExpense += expense.AmountIDR
		expense.CategoryName = categoriesName[expense.CategoryID]
	}

	return &entity.Report{
		Month:        month,
		Year:         year,
		Incomes:      incomes,
		Expenses:     expenses,
		StartBalance: lastTotalIncome - lastTotalExpense,
		EndBalance:   totalIncome - totalExpense,
	}, nil
}
