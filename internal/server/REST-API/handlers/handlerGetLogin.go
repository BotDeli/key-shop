package handlers

import (
	"github.com/gin-gonic/gin"
	"key-shop/internal/database/noSql/redis"
	"net/http"
)

func handlerGetLogin(sessia redis.SessionCache) gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionKey, err := getSessionKey(c)
		if err != nil {
			return
		}

		login, err := getLogin(c, sessia, sessionKey)
		if err != nil {
			return
		}

		c.JSON(http.StatusOK, gin.H{"login": login})
	}
}
