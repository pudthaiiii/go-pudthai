package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type CacheManager struct {
	client  redis.UniversalClient
	timeout time.Duration
}

func NewCacheManager(client redis.UniversalClient, timeout time.Duration) *CacheManager {
	return &CacheManager{
		client:  client,
		timeout: timeout,
	}
}

func (cm *CacheManager) Remember(ctx context.Context, key string, ttl time.Duration, fn func() (string, error)) (string, error) {
	val, err := cm.Get(ctx, key)
	if err == redis.Nil {
		val, err = fn()
		if err != nil {
			return "", err
		}

		err = cm.Set(ctx, key, val, ttl)
		if err != nil {
			return "", err
		}

		return val, nil
	} else if err != nil {
		return "", err
	}

	return val, nil
}

func (cm *CacheManager) Get(ctx context.Context, key string) (string, error) {
	return cm.client.Get(ctx, key).Result()
}

func (cm *CacheManager) Set(ctx context.Context, key string, value string, ttl time.Duration) error {
	return cm.client.Set(ctx, key, value, ttl).Err()
}

func (cm *CacheManager) Delete(ctx context.Context, key string) error {
	return cm.client.Del(ctx, key).Err()
}
