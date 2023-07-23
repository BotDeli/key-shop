package handlers

import (
	"github.com/gin-gonic/gin"
	"key-shop/internal/database/noSql/redis"
)

func showMainPage(sessia redis.SessionCache) gin.HandlerFunc {
	return func(c *gin.Context) {
		if authorized(c, sessia) {
			c.Redirect(301, "/authorized")
		} else {
			c.HTML(200, "main.html", nil)
		}
	}
}

func authorized(c *gin.Context, sessia redis.SessionCache) bool {
	sessionKey, err := c.Cookie("sessia")
	return err == nil && len(sessionKey) == 16 && sessia.ExistsSessionKey(sessionKey)
}

func showAuthorizedPage(sessia redis.SessionCache) gin.HandlerFunc {
	return func(c *gin.Context) {
		if authorized(c, sessia) {
			c.HTML(200, "authorized.html", nil)
		} else {
			c.Redirect(301, "/")
		}
	}
}

func showAccountPage(sessia redis.SessionCache) gin.HandlerFunc {
	return func(c *gin.Context) {
		if authorized(c, sessia) {
			c.HTML(200, "account.html", nil)
		} else {
			c.Redirect(301, "/")
		}
	}
}
