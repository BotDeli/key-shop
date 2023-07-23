package redis

import (
	"github.com/go-redis/redis"
	"key-shop/internal/config"
	"log"
)

type SessionCache interface {
	DisconnectSessionCache()
	GetSessionKey(string) (string, error)
	ExistsSessionKey(string) bool
	DeleteSessionKey(string) error
	GetLogin(string) (string, error)
}

type Redis struct {
	Client *redis.Client
}

func StartSessionCache(cfg config.RedisConfig) SessionCache {
	options := &redis.Options{
		Addr: cfg.Address,
	}
	client := redis.NewClient(options)
	pingClient(client)
	return Redis{Client: client}
}

func pingClient(client *redis.Client) {
	status := client.Ping()
	if status.Err() != nil {
		log.Fatal(status.Err())
	}
}

func (r Redis) DisconnectSessionCache() {
	err := r.Client.Close()
	if err != nil {
		log.Println(err)
	}
}
