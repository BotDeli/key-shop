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
	router.POST("/items", handlerGetPageItems(storage))
	router.GET("/count_pages", handlerGetCountPages(storage))
	router.POST("/add_item", handlerAddItem(sessia, storage))
	router.GET("/get_login", handlerGetLogin(sessia))
	router.GET("/account", showAccountPage(sessia))
	router.GET("/my_items", handlerGetMyItems(sessia, storage))
}

// TODO: Сделать подсказки
// TODO: Организовать доставку описания к уведомлениям
