package utils

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisService struct {
	Client *redis.Client
}

var (
	RDB          *RedisService
	onceRDB      sync.Once
	initRDBError error
)

func InitRDB() error {
	onceRDB.Do(func() {
		client := redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", Config.Redis.Host, Config.Redis.Port),
			Password: Config.Redis.Password,
			DB:       Config.Redis.DB,
			PoolSize: Config.Redis.PoolSize,
		})

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if _, err := client.Ping(ctx).Result(); err != nil {
			fmt.Println("failed to connect redis")
			initRDBError = err
			return
		}

		fmt.Println("redis connect successfully")

		RDB = &RedisService{Client: client}
	})
	return initRDBError
}

func (r *RedisService) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	return r.Client.Set(ctx, key, value, ttl).Err()
}

func (r *RedisService) Get(ctx context.Context, key string) (string, error) {
	return r.Client.Get(ctx, key).Result()
}

func (r *RedisService) Del(ctx context.Context, key string) error {
	return r.Client.Del(ctx, key).Err()
}

func (r *RedisService) HSet(ctx context.Context, key, field string, value interface{}) error {
	return r.Client.HSet(ctx, key, field, value).Err()
}

func (r *RedisService) GSet(ctx context.Context, key, field string) (string, error) {
	return r.Client.HGet(ctx, key, field).Result()
}
