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

	categories, err := r.GetCategories()
	if err != nil {
		return []*expenses.Expense{}, err
	}

	firstYear := 0
	for _, entry := range entries {
		if entry.Date.Year() < firstYear || firstYear == 0 {
			firstYear = entry.Date.Year()
		}
	}

	periods := []int{}
	comulatedMap := map[int]map[string]*expenses.Expense{}
	for i := firstYear; i <= time.Now().Year(); i++ {
		periods = append(periods, i)
		comulatedMap[i] = map[string]*expenses.Expense{}

		for _, category := range categories {
			date, err := time.Parse("2006-01-02", fmt.Sprintf("%d-01-01", i))
			if err != nil {
				return []*expenses.Expense{}, err
			}

			comulatedMap[i][category] = &expenses.Expense{
				Date: date,
				Category: category,
				Amount: 0.00,
			}
		}
	}

	for _, entry := range entries {
		if expense, exists := comulatedMap[entry.Date.Year()][entry.Category]; exists {
			expense.Add(*entry)
		}
	}

	return createComulatedSlice(comulatedMap, periods, categories), nil
}

func (r csvRepository) GetMonths(year int) ([]*expenses.Expense, error) {
	entries, err := r.reader.GetAll()
	if err != nil {
		return []*expenses.Expense{}, err
	}

	categories, err := r.GetCategories()
	if err != nil {
		return []*expenses.Expense{}, err
	}

	periods := []int{}
	comulatedMap := map[int]map[string]*expenses.Expense{}
	for i:= 1; i <= 12; i++ {
		periods = append(periods, i)
		comulatedMap[i] = map[string]*expenses.Expense{}

		for _, category := range categories {
			date, err := time.Parse("2006-01-02", fmt.Sprintf("%d-%02d-01", year, i))
			if err != nil {
				return []*expenses.Expense{}, err
			}
			comulatedMap[i][category] = &expenses.Expense{
				Date: date,
				Category: category,
				Amount: 0.00,
			}
		}
	}

	for _, entry := range entries {
		if entry.Date.Year() != year {
			continue
		}

		month := int(entry.Date.Month())
		if expense, exists := comulatedMap[month][entry.Category]; exists {
			expense.Add(*entry)
		}
	}

	return createComulatedSlice(comulatedMap, periods, categories), nil
}

func (r csvRepository) GetDays(year int, month int) ([]*expenses.Expense, error) {
	entries, err := r.reader.GetAll()
	if err != nil {
		return []*expenses.Expense{}, err
	}

	firstDay, err := time.Parse("2006-01-02", fmt.Sprintf("%d-%02d-01", year, month))
	if err != nil {
		return []*expenses.Expense{}, err
	}

	categories, err := r.GetCategories()
	if err != nil {
		return []*expenses.Expense{}, err
	}

	periods := []int{}
	comulatedMap := map[int]map[string]*expenses.Expense{}
	for int(firstDay.Month()) == month {
		periods = append(periods, firstDay.Day())
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
	for _, entry := range entries {
		if expense, exists := comulatedMap[entry.Date.Day()][entry.Category]; exists {
			expense.Add(*entry)
		}
	}

	return createComulatedSlice(comulatedMap, periods, categories), nil
}

func createComulatedSlice(comulatedMap map[int]map[string]*expenses.Expense, periods []int, categories []string) []*expenses.Expense {
	comulated := []*expenses.Expense{}
	for _, period := range periods {
		for _, category := range categories {
			if expense, exists := comulatedMap[period][category]; exists {
				comulated = append(comulated, expense)
			}
		}
	}

	return comulated
}
