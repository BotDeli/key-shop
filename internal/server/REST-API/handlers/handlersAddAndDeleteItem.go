package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"key-shop/internal/database/noSql/redis"
	"key-shop/internal/database/sql/postgres"
	"net/http"
)

var (
	ErrorGetSessionKey = errors.New("error getting session key")
	ErrorGetLogin      = errors.New("error getting login")
	ErrorGetItem       = errors.New("error getting item")
)

func handlerAddItem(sessia redis.SessionCache, display postgres.ItemsDisplay) gin.HandlerFunc {
	return func(c *gin.Context) {
		login, item, err := getLoginAndItem(c, sessia)
		if err != nil {
			return
		}

		if display.ExistsItem(item) {
			c.Status(http.StatusConflict)
			return
		}

		err = display.AddItem(login, item)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		c.Status(http.StatusCreated)
	}
}

func getLoginAndItem(c *gin.Context, sessia redis.SessionCache) (login string, item postgres.Item, err error) {
	sessionKey, err := getSessionKey(c)
	if err != nil {
		return
	}

	login, err = getLogin(c, sessia, sessionKey)
	if err != nil {
		return
	}

	item, err = getItem(c)
	return
}

func getSessionKey(c *gin.Context) (string, error) {
	sessionKey, err := c.Cookie("sessia")
	if err != nil || sessionKey == "" {
		c.Status(http.StatusUnauthorized)
		return "", ErrorGetSessionKey
	}
	return sessionKey, nil
}

func getLogin(c *gin.Context, sessia redis.SessionCache, sessionKey string) (string, error) {
	login, err := sessia.GetLogin(sessionKey)
	if err != nil || login == "" {
		c.Status(http.StatusUnauthorized)
		return "", ErrorGetLogin
	}
	return login, nil
}

func getItem(c *gin.Context) (postgres.Item, error) {
	var item postgres.Item
	err := c.ShouldBindJSON(&item)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return item, ErrorGetItem
	}
	return item, nil
}

func handlerDeleteItem(sessia redis.SessionCache, display postgres.ItemsDisplay) gin.HandlerFunc {
	return func(c *gin.Context) {
		login, item, err := getLoginAndItem(c, sessia)
		if err != nil {
			return
		}

		err = display.DeleteItem(login, item)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}

		c.Status(http.StatusAccepted)
	}
}
