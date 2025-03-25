package middleware

import (
	"fmt"
	"miner/common/rsp"
	"miner/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type Limiter struct {
	redisCClient *redis.ClusterClient
	limit        int
	duration     time.Duration
}

func NewLimiter(limit int, duration time.Duration) *Limiter {
	return &Limiter{
		redisCClient: utils.RDB.Client,
		limit:        limit,
		duration:     duration,
	}
}

func (m *Limiter) Limit() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID := ctx.GetInt("user_id")
		url := ctx.Request.RequestURI
		method := ctx.Request.Method
		key := fmt.Sprintf("limit:%s:%s:%d", method, url, userID)
		redisCtx := ctx.Request.Context()
		now := time.Now().UnixMilli()

		minScore := strconv.FormatInt(now-int64(m.duration.Milliseconds()), 10)
		m.redisCClient.ZRemRangeByScore(redisCtx, key, "0", minScore)

		// 统计当前时间区间内的请求数
		count, err := m.redisCClient.ZCard(redisCtx, key).Result()
		if err != nil {
			rsp.Error(ctx, http.StatusInternalServerError, "", err.Error())
			ctx.Abort()
			return
		}

		// 检查是否超过限制
		if count >= int64(m.limit) {
			rsp.Error(ctx, http.StatusTooManyRequests, "request is too frequent", "")
			ctx.Abort()
			return
		}

		// 记录本次请求
		member := redis.Z{
			Score:  float64(now),
			Member: now,
		}
		_, err = m.redisCClient.ZAdd(redisCtx, key, member).Result()
		if err != nil {
			rsp.Error(ctx, http.StatusInternalServerError, "", err.Error())
			ctx.Abort()
			return
		}

		// 设置过期时间
		m.redisCClient.Expire(redisCtx, key, m.duration)

		ctx.Next()
	}
}
