package api

import (
	"github.com/gin-gonic/gin"
	"github.com/spie/moneymy/handlers"
)

func SetUp(expensesHandler handlers.ExpensesHandler) *gin.Engine {
	api := gin.Default()

	api.GET("/categories", expensesHandler.GetCategories())

	return api
}
