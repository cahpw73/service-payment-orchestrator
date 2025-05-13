package redis

import (
	"context"
	"log"
	"time"

	"service-payment-orchestrator/config"

	"github.com/redis/go-redis/v9"
)

var (
	ctx2 = context.Background()
)

type RedisServiceInterface interface {
	GetData(key string) (string, error)
	SetData(key, value string) error
	Close() error
}

type RedisService struct {
	client *redis.Client
}

func NewRedisService(addr string) RedisServiceInterface {
	redisClient := redis.NewClient(&redis.Options{
		Addr: addr,
		DB:   0,
	})

	_, err := redisClient.Ping(ctx2).Result()
	if err != nil {
		log.Fatalf("Unable to connect to Redis: %v", err)
	}

	return &RedisService{client: redisClient}
}

func (r *RedisService) GetData(key string) (string, error) {
	val, err := r.client.Get(ctx2, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func (r *RedisService) SetData(key, value string) error {
	expiration := time.Minute * time.Duration(config.RedisTTL)
	err := r.client.Set(ctx2, key, value, expiration).Err()
	return err
}

func (r *RedisService) Close() error {
	return r.client.Close()
}
