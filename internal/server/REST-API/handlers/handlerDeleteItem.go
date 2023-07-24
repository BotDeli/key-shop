package handlers

import (
	"github.com/gin-gonic/gin"
	"key-shop/internal/database/noSql/redis"
	"key-shop/internal/database/sql/postgres"
	"net/http"
)

func handlerDeleteItem(sessia redis.SessionCache, display postgres.ItemsDisplay) gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionKey, err := c.Cookie("sessia")
		if err != nil {
			c.Status(http.StatusUnauthorized)
			return
		}

		login, err := sessia.GetLogin(sessionKey)
		if err != nil {
			c.Status(http.StatusUnauthorized)
			return
		}

		item, err := getItem(c)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		err = display.DeleteItem(login, item)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}

		c.Status(http.StatusAccepted)
	}
}

func getItem(c *gin.Context) (postgres.Item, error) {
	var item postgres.Item
	err := c.ShouldBindJSON(&item)
	return item, err
}
