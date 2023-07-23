package redis

import (
	"errors"
	"key-shop/pkg/UUIDGenerator"
	"key-shop/pkg/errorHandle"
	"log"
	"time"
)

const (
	path = "keyShop/internal/database/noSql/redis/"
)

var (
	ErrorGetSessionKey          = errors.New("error getting session key")
	ErrorCreateUniqueSessionKey = errors.New("error creating unique session key")
	ErrorDeleteSessionKey       = errors.New("error deleting session key")
	ErrorUserDontExist          = errors.New("error user does not exist")
)

func (r Redis) GetSessionKey(login string) (string, error) {
	sessionKey, err := getUniqueSessionKey(r, 1)
	if err != nil {
		return "", err
	}
	status := r.Client.Set(sessionKey, login, 24*time.Hour)
	if status.Err() != nil {
		log.Println(errorHandle.ErrorFormat(path, "sessia.go", "AddSessionKey", status.Err()))
		return "", ErrorGetSessionKey
	}
	return sessionKey, nil
}

func getUniqueSessionKey(r Redis, n int) (string, error) {
	if n == 5 {
		return "", ErrorCreateUniqueSessionKey
	}
	sessionKey := UUIDGenerator.NewUUID()
	if r.ExistsSessionKey(sessionKey) {
		return getUniqueSessionKey(r, n+1)
	}
	return sessionKey, nil
}

func (r Redis) ExistsSessionKey(sessionKey string) bool {
	status := r.Client.Exists(sessionKey)
	if status.Err() != nil {
		return false
	}
	return status.Val() == 1
}

func (r Redis) DeleteSessionKey(sessionKey string) error {
	status := r.Client.Del(sessionKey)
	if status.Err() != nil {
		log.Print(errorHandle.ErrorFormat(path, "sessia.go", "DeleteSessionKey", status.Err()))
		return ErrorDeleteSessionKey
	}
	return nil
}

func (r Redis) GetLogin(sessionKey string) (string, error) {
	status := r.Client.Get(sessionKey)
	if status.Err() != nil {
		return "", ErrorUserDontExist
	}
	return status.Val(), nil
}
