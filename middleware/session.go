package middleware

import (
	"miner/utils"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

// 初始化 Session
func InitSession(router *gin.Engine) error {
	conf := utils.Config.Session
	store, err := redis.NewStore(
		10,
		"tcp",
		utils.Config.Redis.Host+":"+strconv.Itoa(utils.Config.Redis.Port),
		utils.Config.Redis.Password,
		[]byte(conf.Secret),
	)
	if err != nil {
		return err
	}

	store.Options(sessions.Options{
		Path:     "/",
		MaxAge:   conf.MaxAge,
		Secure:   true,
		HttpOnly: true,
	})

	router.Use(sessions.Sessions(conf.Name, store))
	return nil
}
