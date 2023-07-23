package handlers

import (
	"github.com/gin-gonic/gin"
	"key-shop/internal/database/noSql/redis"
	"net/http"
)

func handlerGetLogin(sessia redis.SessionCache) gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionKey, err := c.Cookie("sessia")
		if err != nil {
			c.Status(http.StatusBadRequest)
		}

		login, err := sessia.GetLogin(sessionKey)
		if err != nil {
			c.Status(http.StatusBadRequest)
		}

		c.JSON(http.StatusOK, struct {
			Login string `json:"login"`
		}{
			Login: login,
		})
	}
}
