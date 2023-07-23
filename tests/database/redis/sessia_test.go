package redis

import (
	"key-shop/internal/config"
	"key-shop/internal/database/noSql/redis"
	"testing"
)

const (
	address     = "localhost:6379"
	testedLogin = "testLogin"
)

func TestSessia(t *testing.T) {
	cache := connectSessionCache(t)
	defer cache.DisconnectSessionCache()

	functionalTest(t, cache)
}

func functionalTest(t *testing.T, cache redis.SessionCache) {
	sessionKey := getSessionKey(t, cache)
	checkSessionKey(t, sessionKey)
	checkExistsSessionKey(t, cache, sessionKey)
	checkGetLogin(t, cache, sessionKey)
	checkRemoveSessionKey(t, cache, sessionKey)
}

func connectSessionCache(t *testing.T) redis.SessionCache {
	cache := redis.StartSessionCache(config.RedisConfig{
		Address: address,
	})
	if cache == nil {
		t.Fail()
	}
	return cache
}

func getSessionKey(t *testing.T, cache redis.SessionCache) string {
	sessionKey, err := cache.GetSessionKey(testedLogin)
	if err != nil {
		t.Error(err)
	}
	return sessionKey
}

func checkSessionKey(t *testing.T, sessionKey string) {
	if sessionKey == testedLogin {
		t.Error("session key equals tested login")
	}

	if sessionKey == "" {
		t.Error("session key is empty")
	}
}

func checkExistsSessionKey(t *testing.T, cache redis.SessionCache, sessionKey string) {
	if !cache.ExistsSessionKey(sessionKey) {
		t.Error("session key does not exist")
	}
}

func checkGetLogin(t *testing.T, cache redis.SessionCache, sessionKey string) {
	login, err := cache.GetLogin(sessionKey)
	if err != nil {
		t.Error(err)
	}

	if login != testedLogin {
		t.Error("login dont equals tested login")
	}

	if login == "" {
		t.Error("login is empty")
	}
}

func checkRemoveSessionKey(t *testing.T, cache redis.SessionCache, sessionKey string) {
	err := cache.DeleteSessionKey(sessionKey)
	if err != nil {
		t.Error(err)
	}
	if cache.ExistsSessionKey(sessionKey) {
		t.Error("session key dont removed")
	}
}
