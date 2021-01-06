package repository

import (
	"fmt"
	"time"

	"github.com/spie/moneymy/expenses"
	"github.com/spie/moneymy/reader"
)

type Repository interface {
	GetCategories() ([]string, error)
	GetYears() ([]*expenses.Expense, error)
	GetMonths(year int) ([]*expenses.Expense, error)
	GetDays(year int, month int) ([]*expenses.Expense, error)
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

func (r csvRepository) GetYears() ([]*expenses.Expense, error) {
	entries, err := r.reader.GetAll()
	if err != nil {
		return []*expenses.Expense{}, err
	}

	comulatedMap := map[int]map[string]*expenses.Expense{}
	for _, entry := range entries {
		year := entry.Date.Year()
		if _, exists := comulatedMap[year]; !exists {
			comulatedMap[year] = map[string]*expenses.Expense{}
		}
		if _, exists := comulatedMap[year][entry.Category]; !exists {
			date, err := time.Parse("2006-01-02", fmt.Sprintf("%d-01-01", year))
			if err != nil {
				return []*expenses.Expense{}, err
			}

			comulatedMap[year][entry.Category] = &expenses.Expense{Date: date, Category: entry.Category, Amount: 0.00}
		}

		comulatedMap[year][entry.Category].Add(*entry)
	}

	return createComulatedSlice(comulatedMap), nil
}

func (r csvRepository) GetMonths(year int) ([]*expenses.Expense, error) {
	entries, err := r.reader.GetAll()
	if err != nil {
		return []*expenses.Expense{}, err
	}

	comulatedMap := map[int]map[string]*expenses.Expense{}
	for _, entry := range entries {
		if entry.Date.Year() != year {
			continue
		}

		month := int(entry.Date.Month())
		if _, exists := comulatedMap[month]; !exists {
			comulatedMap[month] = map[string]*expenses.Expense{}
		}
		if _, exists := comulatedMap[month][entry.Category]; !exists {
			date, err := time.Parse("2006-01-02", fmt.Sprintf("%d-%02d-01", year, month))
			if err != nil {
				return []*expenses.Expense{}, err
			}

			comulatedMap[month][entry.Category] = &expenses.Expense{Date: date, Category: entry.Category, Amount: 0.00}
		}

		comulatedMap[month][entry.Category].Add(*entry)
	}

	return createComulatedSlice(comulatedMap), nil
}

func (r csvRepository) GetDays(year int, month int) ([]*expenses.Expense, error) {
	// entries, err := r.reader.GetAll()
	// if err != nil {
	// 	return []*expenses.Expense{}, err
	// }

	firstDay, err := time.Parse("2006-01-02", fmt.Sprintf("%d-%02d-01", year, month))
	if err != nil {
		return []*expenses.Expense{}, err
	}

	categories, err := r.GetCategories()
	if err != nil {
		return []*expenses.Expense{}, err
	}

	comulatedMap := map[int]map[string]*expenses.Expense{}
	for int(firstDay.Month()) == month {
		comulatedMap[firstDay.Day()] = map[string]*expenses.Expense{}
		for _, category := range categories {
			comulatedMap[firstDay.Day()][category] = &expenses.Expense{
				Date:     firstDay,
				Category: category,
				Amount:   0.00,
			}
		}

		firstDay = firstDay.Add(time.Hour * 24)
	}

	return createComulatedSlice(comulatedMap), nil
}

func createComulatedSlice(comulatedMap map[int]map[string]*expenses.Expense) []*expenses.Expense {
	comulated := []*expenses.Expense{}
	for _, categoriesInYear := range comulatedMap {
		for _, exp := range categoriesInYear {
			comulated = append(comulated, exp)
		}
	}

	return comulated
}
