package reader

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/spie/moneymy/expenses"
)

type Reader interface {
	GetAll() ([]*expenses.Expense, error)
	GetAllRaw() ([][]string, error)
}

type csvReader struct {
	fileName string
}

func NewCsvReader(fileName string) Reader {
	return csvReader{fileName: fileName}
}

func (r csvReader) GetAll() ([]*expenses.Expense, error) {
	file, err := os.Open(r.fileName)
	if err != nil {
		return []*expenses.Expense{}, err
	}
	defer file.Close()

	csv := csv.NewReader(file)

	allExpenses := []*expenses.Expense{}
	for {
		record, err := csv.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return []*expenses.Expense{}, err
		}

		amount, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			return []*expenses.Expense{}, err
		}

		date, err := time.Parse("02/01/2006", record[0])
		if err != nil {
			return []*expenses.Expense{}, err
		}

		allExpenses = append(allExpenses, &expenses.Expense{
			Date:     date,
			Amount:   amount,
			Currency: record[3],
			Category: record[1],
		})
	}

	return allExpenses, nil
}

func (r csvReader) GetAllRaw() ([][]string, error) {
	file, err := os.Open(r.fileName)
	if err != nil {
		return [][]string{}, err
	}
	defer file.Close()

	csv := csv.NewReader(file)

	return csv.ReadAll()
}
