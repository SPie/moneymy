package api

import (
	"github.com/gin-gonic/gin"
	"github.com/spie/moneymy/handlers"
)

func SetUp(expensesHandler handlers.ExpensesHandler) *gin.Engine {
	api := gin.Default()

	api.GET("categories", expensesHandler.GetCategories())
	api.GET("years", expensesHandler.GetYears())
	api.GET("months", expensesHandler.GetMonths())
	api.GET("days", expensesHandler.GetDays())

	return api
}
