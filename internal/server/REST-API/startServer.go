package REST_API

import (
	"github.com/gin-gonic/gin"
	"key-shop/internal/config"
	"key-shop/internal/database/noSql/redis"
	"key-shop/internal/database/sql/postgres"
	"key-shop/internal/server/REST-API/handlers"
	"log"
	"net/http"
)

func StartServer(cfg config.HTTPServerConfig, sessia redis.SessionCache, storage postgres.Storage) {
	router := getRouter()
	loadFiles(router)
	server := getHttpServer(cfg, router)
	handlers.InitHandlers(router, sessia, storage)
	log.Fatal(server.ListenAndServe())
}

func getRouter() *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	return router
}

func loadFiles(router *gin.Engine) {
	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*")
}

func getHttpServer(cfg config.HTTPServerConfig, router *gin.Engine) *http.Server {
	return &http.Server{
		Addr:              cfg.Address,
		Handler:           router,
		ReadHeaderTimeout: cfg.ReadHeaderTimeout,
		IdleTimeout:       cfg.IdleTimeout,
	}
}
