package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/spie/moneymy/expenses"
	"github.com/spie/moneymy/repository"
)

type ExpensesHandler struct {
	repo repository.Repository
}

func NewExpensesHandler(repo repository.Repository) ExpensesHandler {
	return ExpensesHandler{repo: repo}
}

func (eh ExpensesHandler) GetCategories() gin.HandlerFunc {
	return func(c *gin.Context) {
		categories, err := eh.repo.GetCategories()
		if err != nil {
			c.AbortWithError(400, err)
			return
		}

		c.JSON(200, map[string][]string{"categories": categories})
	}
}

func (eh ExpensesHandler) GetYears() gin.HandlerFunc {
	return func(c *gin.Context) {
		years, err := eh.repo.GetYears()
		if err != nil {
			c.AbortWithError(400, err)
			return
		}

		c.JSON(200, map[string][]*expenses.Expense{"expenses": years})
	}
}

func (eh ExpensesHandler) GetMonths() gin.HandlerFunc {
	return func(c *gin.Context) {
		year, err := strconv.Atoi(c.Query("year"))
		if err != nil {
			c.AbortWithError(400, err)
			return
		}

		months, err := eh.repo.GetMonths(year)
		if err != nil {
			c.AbortWithError(400, err)
			return
		}

		c.JSON(200, map[string][]*expenses.Expense{"expenses": months})
	}
}

func (eh ExpensesHandler) GetDays() gin.HandlerFunc {
	return func(c *gin.Context) {
		year, err := strconv.Atoi(c.Query("year"))
		if err != nil {
			c.AbortWithError(400, err)
			return
		}
		month, err := strconv.Atoi(c.Query("month"))
		if err != nil {
			c.AbortWithError(400, err)
			return
		}

		days, err := eh.repo.GetDays(year, month)
		if err != nil {
			c.AbortWithError(400, err)
			return
		}

		c.JSON(200, map[string][]*expenses.Expense{"expenses": days})
	}
}
