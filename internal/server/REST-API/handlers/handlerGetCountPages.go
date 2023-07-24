package handlers

import (
	"github.com/gin-gonic/gin"
	"key-shop/internal/database/sql/postgres"
	"net/http"
)

func handlerGetCountPages(display postgres.PageDisplay) gin.HandlerFunc {
	return func(c *gin.Context) {
		pages, err := display.GetCountPages()
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}

		c.JSON(http.StatusOK, struct {
			Pages int `json:"pages"`
		}{
			Pages: pages,
		})
	}
}
