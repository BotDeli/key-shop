package handlers

import (
	"github.com/gin-gonic/gin"
	"key-shop/internal/database/noSql/redis"
	"key-shop/internal/database/sql/postgres"
	"net/http"
)

func handlerGetMyItems(sessia redis.SessionCache, display postgres.PageDisplay) gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionKey, err := getSessionKey(c)
		if err != nil {
			return
		}

		login, err := getLogin(c, sessia, sessionKey)
		if err != nil {
			return
		}

		page, err := display.GetPageMyItems(login)
		if err != nil {
			c.Status(http.StatusNotFound)
			return
		}

		c.JSON(http.StatusOK, page)
	}
}
