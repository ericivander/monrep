package types

import (
	"encoding/json"
	"errors"
)

// CategoryType is type for category type
type CategoryType int32

const (
	// IncomeType is enum value of category type 'income'
	IncomeType CategoryType = iota
	// ExpenseType is enum value of category type 'expense'
	ExpenseType
)

// Scan override scan on sqlx
func (c *CategoryType) Scan(value interface{}) error {
	val := CategoryType(value.(int32))
	if val < 0 || val > 1 {
		return errors.New("invalid value")
	}

	*c = val
	return nil
}

var toString = map[CategoryType]string{
	IncomeType:  "Income",
	ExpenseType: "Expense",
}

var toID = map[string]CategoryType{
	"Income":  IncomeType,
	"Expense": ExpenseType,
}

func (c CategoryType) String() string {
	return toString[c]
}

// MarshalJSON marshals the enum as a quoted json string
func (c CategoryType) MarshalJSON() ([]byte, error) {
	return json.Marshal(toString[c])
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (c *CategoryType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	// Note that if the string cannot be found then it will be set to the zero value, 'IncomeType' in this case.
	*c = toID[s]
	return nil
}
