package handlers

import (
	"github.com/gin-gonic/gin"
	"key-shop/internal/database/noSql/redis"
)

func handlerShowMainPage(sessia redis.SessionCache) gin.HandlerFunc {
	return func(c *gin.Context) {
		showDontAuthorizedPage(c, sessia, "main.html", "/authorized")
	}
}

func showDontAuthorizedPage(c *gin.Context, sessia redis.SessionCache, htmlName, location string) {
	if authorized(c, sessia) {
		c.Redirect(301, location)
	} else {
		c.HTML(200, htmlName, nil)
	}
}

func authorized(c *gin.Context, sessia redis.SessionCache) bool {
	sessionKey, err := c.Cookie("sessia")
	return err == nil && len(sessionKey) == 16 && sessia.ExistsSessionKey(sessionKey)
}

func handlerShowAuthorizedPage(sessia redis.SessionCache) gin.HandlerFunc {
	return func(c *gin.Context) {
		showAuthorizedPage(c, sessia, "authorized.html")
	}
}

func showAuthorizedPage(c *gin.Context, sessia redis.SessionCache, htmlName string) {
	if authorized(c, sessia) {
		c.HTML(200, htmlName, nil)
	} else {
		c.Redirect(301, "/")
	}
}

func handlerShowAccountPage(sessia redis.SessionCache) gin.HandlerFunc {
	return func(c *gin.Context) {
		showAuthorizedPage(c, sessia, "account.html")
	}
}
