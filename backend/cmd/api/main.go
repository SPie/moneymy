package main

import (
	"os"

	"github.com/spie/moneymy/api"
	"github.com/spie/moneymy/handlers"
	"github.com/spie/moneymy/reader"
	"github.com/spie/moneymy/repository"
)

func main() {
	if len(os.Args) < 2 {
		panic("Argument for file name is required")
	}

	r := reader.NewCsvReader(os.Args[1])

	repo := repository.NewRepository(r)
	
	expensesHandler := handlers.NewExpensesHandler(repo)

	api := api.SetUp(expensesHandler)

	err := api.Run()
	if err != nil {
		panic(err)
	}
}
