package handlers

import (
	"github.com/gin-gonic/gin"
	"key-shop/internal/database/noSql/redis"
	"key-shop/pkg/errorHandle"
	"log"
	"net/http"
)

func handlerExitUser(sessia redis.SessionCache) gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionKey, err := getSessionKey(c)
		if err != nil {
			return
		}

		err = sessia.DeleteSessionKey(sessionKey)
		if err != nil {
			log.Println(errorHandle.ErrorFormat(path, "handlerExitUser", "handlerExitUser", err))
		}

		setCookieSessia(c, "", -1)
		c.Status(http.StatusAccepted)
	}
}
