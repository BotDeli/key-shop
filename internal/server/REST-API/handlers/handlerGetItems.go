package handlers

import (
	"github.com/gin-gonic/gin"
	"key-shop/internal/database/sql/postgres"
	"net/http"
)

func handlerGetPageItems(display postgres.PageDisplay) gin.HandlerFunc {
	return func(c *gin.Context) {
		numberPage, err := getNumberPage(c)

		page, err := display.GetPageAllItems(numberPage)

		if err != nil {
			c.Status(http.StatusNotFound)
			return
		}

		c.JSON(http.StatusOK, page)
	}
}

func getNumberPage(c *gin.Context) (int, error) {
	var numberPage struct {
		Page int `json:"page"`
	}
	err := c.ShouldBindJSON(&numberPage)
	return numberPage.Page, err
}
