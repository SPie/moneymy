package main

import (
	"encoding/csv"
	"os"
	"regexp"
	"strings"

	"github.com/spie/moneymy/reader"
)

func main() {
	if len(os.Args) < 2 {
		panic("file name is required")
	}
	if len(os.Args) < 3 {
		panic("new file name is required")
	}

	records, err := reader.NewCsvReader(os.Args[1]).GetAllRaw()
	if err != nil {
		panic(err)
	}

	file, err := os.Create(os.Args[2])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, record := range records {
		if record[1] != "Bargeld" && record[1] != "Freizeit" && record[1] != "Haushalt" && record[1] != "Konto" {
			continue
		}

		category := record[2]
		if matched, _ := regexp.MatchString(`From\ `, category); matched {
			continue
		}
		if matched, _ := regexp.MatchString(`To\ `, category); matched {
			continue
		}

		amount := record[5]
		if matched := regexp.MustCompile(`\.`).FindAllString(amount, -1); len(matched) > 1 {
			amount = strings.Replace(amount, ".", "", len(matched)-1)
		}

		converted := []string{record[0], category, amount, record[6]}
		err := writer.Write(converted)
		if err != nil {
			panic(err)
		}
	}
}
