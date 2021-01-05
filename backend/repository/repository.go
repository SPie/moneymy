package repository

import (
	"time"

	"github.com/spie/moneymy/expenses"
	"github.com/spie/moneymy/reader"
)

type Repository interface {
	GetCategories() ([]string, error)
	GetYears(first int, last int) ([]expenses.Expense, error)
}

type csvRepository struct {
	reader reader.Reader
}

func NewRepository(reader reader.Reader) Repository {
	return csvRepository{reader: reader}
}

func (r csvRepository) GetCategories() ([]string, error) {
	entries, err := r.reader.GetAll()
	if err != nil {
		return []string{}, err
	}

	availableCategories := map[string]bool{}
	categories := []string{}
	for _, entry := range entries {
		if _, exists := availableCategories[entry.Category]; !exists {
			availableCategories[entry.Category] = true
			categories = append(categories, entry.Category)
		}
	}

	return categories, nil
}

func (r csvRepository) GetYears(first int, last int) ([]expenses.Expense, error) {
	entries, err := r.reader.GetAll()
	if err != nil {
		return []expenses.Expense{}, err
	}

	comulatedMap := map[int]map[string]expenses.Expense{}

	categories, err := r.GetCategories()
	if err != nil {
		return []expenses.Expense{}, err
	}

	for _, category := range categories {
		for i := first; i <= last; i++ {
			comulatedMap[i][category] = expenses.Expense{Date: time.Time{}, Category: category, Amount: 0.00}
		}
	}

	comulated := []expenses.Expense{}
	for _, entry := range entries {
		year := entry.Date.Year()
		_, ok := comulatedMap[year]
		if !ok {
			// TODO
		}
	}

	return comulated, nil
}
