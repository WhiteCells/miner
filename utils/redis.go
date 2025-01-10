package utils

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	Client *redis.ClusterClient
}

var (
	RDB          *RedisClient
	onceRDB      sync.Once
	initRDBError error
)

func InitRDB() error {
	onceRDB.Do(func() {
		addrs := fmt.Sprintf("%s:%d", Config.Redis.Host, Config.Redis.Port)
		client := redis.NewClusterClient(&redis.ClusterOptions{
			// Addr:     fmt.Sprintf("%s:%d", Config.Redis.Host, Config.Redis.Port),
			Addrs: []string{
				addrs,
				// "127.0.0.1:7000",
				// "127.0.0.1:7001",
				// "127.0.0.1:7002",
				// "127.0.0.1:7003",
				// "127.0.0.1:7004",
				// "127.0.0.1:7005",
			},
			Password: Config.Redis.Password,
			// DB:       Config.Redis.DB,
			// PoolSize: Config.Redis.PoolSize,
		})

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if _, err := client.Ping(ctx).Result(); err != nil {
			fmt.Println("failed to connect redis")
			initRDBError = err
			return
		}

		fmt.Println("redis connect successfully")

		RDB = &RedisClient{Client: client}
	})
	return initRDBError
}

func (r *RedisClient) RPush(ctx context.Context, key string, value string) error {
	return r.Client.RPush(ctx, key, value).Err()
}

func (r *RedisClient) LPop(ctx context.Context, key string) (string, error) {
	return r.Client.LPop(ctx, key).Result()
}

func (r *RedisClient) LLen(ctx context.Context, key string) (int64, error) {
	return r.Client.LLen(ctx, key).Result()
}

func (r *RedisClient) HSet(ctx context.Context, field string, key string, value string) error {
	return r.Client.HSet(ctx, field, key, value).Err()
}

func (r *RedisClient) HGet(ctx context.Context, field string, key string) (string, error) {
	return r.Client.HGet(ctx, field, key).Result()
}

func (r *RedisClient) HGetAll(ctx context.Context, field string) (map[string]string, error) {
	return r.Client.HGetAll(ctx, field).Result()
}

func (r *RedisClient) HDel(ctx context.Context, field string, key string) error {
	return r.Client.HDel(ctx, field, key).Err()
}

func (r *RedisClient) Set(ctx context.Context, key string, value interface{}, opts ...time.Duration) error {
	exp := time.Duration(0)
	if len(opts) > 0 {
		exp = opts[0]
	}
	return r.Client.Set(ctx, key, value, exp).Err()
}

func (r *RedisClient) Get(ctx context.Context, key string) (string, error) {
	return r.Client.Get(ctx, key).Result()
}

func (r *RedisClient) Del(ctx context.Context, key ...string) error {
	return r.Client.Del(ctx, key...).Err()
}

func (r *RedisClient) Scan(ctx context.Context, pattern string) ([]string, error) {
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

func (r *RedisClient) Exists(ctx context.Context, key string) bool {
	cnt, err := r.Client.Exists(ctx, key).Result()
	if err != nil {
		return false
	}
	return cnt > 0
}

func (r *RedisClient) Close(ctx context.Context) error {
	return r.Client.Close()
}
