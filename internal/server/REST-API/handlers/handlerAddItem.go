package handlers

import (
	"github.com/gin-gonic/gin"
	"key-shop/internal/database/noSql/redis"
	"key-shop/internal/database/sql/postgres"
	"net/http"
)

func handlerAddItem(sessia redis.SessionCache, display postgres.ItemsDisplay) gin.HandlerFunc {
	return func(c *gin.Context) {
		var item postgres.Item
		err := c.ShouldBindJSON(&item)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		sessionKey, err := c.Cookie("sessia")
		if err != nil || sessionKey == "" {
			c.Status(http.StatusUnauthorized)
			return
		}

		login, err := sessia.GetLogin(sessionKey)
		if err != nil || login == "" {
			c.Status(http.StatusUnauthorized)
			return
		}

		if !display.ExistsItem(item) {
			err = display.AddItem(login, item)
			if err != nil {
				c.Status(http.StatusBadRequest)
				return
			}
		}

		c.Status(http.StatusCreated)
	}
}
