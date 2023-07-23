package app

import (
	"key-shop/internal/config"
	"key-shop/internal/database/noSql/redis"
	"key-shop/internal/database/sql/postgres"
	Rest "key-shop/internal/server/REST-API"
)

func StartApplication() {
	cfg := config.MustReadConfig()
	dataSourceName := cfg.Postgres.GetDataSourceName()
	storage := postgres.MustNewStorage(dataSourceName)
	defer storage.Disconnect()
	session := redis.StartSessionCache(cfg.Redis)
	defer session.DisconnectSessionCache()
	Rest.StartServer(cfg.HttpServer, session, storage)
}
