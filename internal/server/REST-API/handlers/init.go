package handlers

import (
	"github.com/gin-gonic/gin"
	"key-shop/internal/database/noSql/redis"
	"key-shop/internal/database/sql/postgres"
)

const (
	path = "keyShop/internal/server/REST-API/handlers/"
)

func InitHandlers(router *gin.Engine, sessia redis.SessionCache, storage postgres.Storage) {
	router.GET("/", showMainPage(sessia))
	router.GET("/authorized", showAuthorizedPage(sessia))
	router.POST("/login", handlerLogin(sessia, storage))
	router.POST("/registration", handlerRegistration(sessia, storage))
	router.POST("/exit", handlerExitUser(sessia))
	router.GET("/items", getItems(storage))
	router.GET("/count_pages", getCountPages())
	router.POST("/add_item", handlerAddItem(sessia, storage))
}

// TODO: Сделать таблицу
// TODO: Сделать подсказки

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

func getItems(display postgres.PageDisplay) gin.HandlerFunc {
	return func(c *gin.Context) {
		page, err := display.GetPage(1)
		if err != nil {
			c.Status(500)
			return
		}
		c.JSON(200, page)
	}
}

func getCountPages() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, struct {
			Pages int `json:"pages"`
		}{Pages: 10})
	}
}
