package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type RedisClient struct {
	Redis *redis.Client
}

func NewRedisClient() ICache {
	var client = &redis.Client{}
	client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", "localhost", "6379"),
		Password: "",
		DB:       0,
	})

	fmt.Println("Redis client created")

	pong, err := client.Ping(ctx).Result()
	if err != nil {
		fmt.Println("Redis ping failed", err)
	} else {
		fmt.Println("Redis ping success", pong)
	}

	redisClient := &RedisClient{
		Redis: client,
	}
	return redisClient
}

func (r *RedisClient) Set(key string, value interface{}) error {
	err := r.Redis.Set(ctx, key, value, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisClient) SetWithDuration(key string, value interface{}, duration int) error {
	err := r.Redis.Set(ctx, key, value, time.Duration(duration)).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisClient) Get(key string) (interface{}, error) {
	val, err := r.Redis.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	return val, nil
}

func (r *RedisClient) Delete(key string) error {
	err := r.Redis.Del(ctx, key).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisClient) Update(key string, value interface{}) error {
	_ = r.Delete(key)

	err := r.Set(key, value)
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisClient) Close() {
	r.Redis.Close()
}
