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

func (r *RedisService) HSet(ctx context.Context, key string, field string, value interface{}) error {
	return r.Client.HSet(ctx, key, field, value).Err()
}

func (r *RedisService) GSet(ctx context.Context, key string, field string) (string, error) {
	return r.Client.HGet(ctx, key, field).Result()
}

func (r *RedisService) Scan(ctx context.Context, pattern string) ([]string, error) {
	var keys []string
	var cursor uint64
	for {
		var err error
		var keys2 []string
		keys2, cursor, err = r.Client.Scan(ctx, cursor, pattern, 100).Result()
		if err != nil {
			return nil, err
		}
		keys = append(keys, keys2...)

		if cursor == 0 {
			break
		}
	}
	return keys, nil
}

func (r *RedisService) Close(ctx context.Context) error {
	return r.Client.Close()
}

// func (r *RedisService) Exists(ctx context.Context, key string) bool {
// 	exists := r.Client.Exists(ctx, key).Err()
// 	return exists
// }
