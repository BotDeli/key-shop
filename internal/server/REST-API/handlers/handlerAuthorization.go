package handlers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"key-shop/internal/database/noSql/redis"
	"key-shop/internal/database/sql/postgres"
	"key-shop/pkg/errorHandle"
	"log"
	"net/http"
	"time"
)

var (
	ErrorGetDataRequest = errors.New("error get data from request")
)

func handlerRegistration(sessia redis.SessionCache, a postgres.Authorization) gin.HandlerFunc {
	return handlerAuthorization(sessia, a.Registration, "handlerAuthorization.go", "handlerRegistration")
}

func handlerLogin(sessia redis.SessionCache, a postgres.Authorization) gin.HandlerFunc {
	return handlerAuthorization(sessia, a.Login, "handlerAuthorization.go", "handlerLogin")
}

func handlerAuthorization(sessia redis.SessionCache, method func(postgres.User) error, fileName, functionName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user postgres.User
		err := c.ShouldBindJSON(&user)
		if err != nil {
			log.Println(errorHandle.ErrorFormat(path, fileName, functionName, err))
			sendResponse(c, http.StatusBadRequest, ErrorGetDataRequest)
			return
		}
		err = method(user)
		if err != nil {
			sendResponse(c, http.StatusConflict, err)
			return
		}
		sessionKey, err := sessia.GetSessionKey(user.Login)
		setCookieSessia(c, sessionKey, int(24*time.Hour))
		sendResponse(c, http.StatusAccepted, err)
	}
}

func setCookieSessia(c *gin.Context, sessionKey string, maxAge int) {
	cookie := &http.Cookie{
		Name:     "sessia",
		Value:    sessionKey,
		Path:     "/",
		Domain:   "localhost",
		MaxAge:   maxAge,
		HttpOnly: true,
	}
	http.SetCookie(c.Writer, cookie)
}

func sendResponse(c *gin.Context, statusCode int, err error) {
	c.JSON(statusCode, gin.H{
		"ERR": fmt.Sprint(err),
	})
}
