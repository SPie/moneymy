package handlers

import (
	"github.com/gin-gonic/gin"
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
