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

	// GET

	router.GET("/", handlerShowMainPage(sessia))
	router.GET("/authorized", handlerShowAuthorizedPage(sessia))
	router.GET("/count_pages", handlerGetCountPages(storage))
	router.GET("/get_login", handlerGetLogin(sessia))
	router.GET("/account", handlerShowAccountPage(sessia))
	router.GET("/my_items", handlerGetMyItems(sessia, storage))

	// POST

	router.POST("/login", handlerLogin(sessia, storage))
	router.POST("/registration", handlerRegistration(sessia, storage))
	router.POST("/exit", handlerExitUser(sessia))
	router.POST("/items", handlerGetPageItems(storage))
	router.POST("/add_item", handlerAddItem(sessia, storage))

	// DELETE

	router.DELETE("/delete_item", handlerDeleteItem(sessia, storage))
}
