package datastore

import (
	"context"
	"fmt"
	"go-ibooking/internal/config"
	"go-ibooking/internal/infrastructure/logger"
	"go-ibooking/internal/utils"
	"log"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisDatastore struct {
	Client redis.UniversalClient
}

func NewRedisDatastore(cfg *config.Config) *RedisDatastore {
	isClusterEnabled := cfg.Get("Redis")["ClusterEnabled"].(bool)

	var client redis.UniversalClient

	if isClusterEnabled {
		client = connectRedisCluster(cfg)
	} else {
		client = connectRedisStandalone(cfg)
	}

	return &RedisDatastore{Client: client}
}

func (r *RedisDatastore) Close() error {
	err := r.Client.Close()

	if err != nil {
		return err
	}

	return nil
}

func Ping(r *RedisDatastore) error {
	var ctx = context.Background()

	return r.Client.Ping(ctx).Err()
}

func connectRedisStandalone(cfg *config.Config) *redis.Client {
	host := cfg.Get("Redis")["Host"].(string)
	port := cfg.Get("Redis")["Port"].(string)
	password := cfg.Get("Redis")["Password"].(string)
	db := cfg.Get("Redis")["DB"].(string)

	options := &redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: password,
		DB:       utils.StringToInt(db),
	}

	client := redis.NewClient(options)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := client.Ping(ctx).Result()
	if err != nil {
		logger.Log.Err(err).Msg("failed to connect to Redis (Standalone)")
		log.Printf("Failed to connect to Redis (Standalone): %v", err)
	}

	logger.Write.Info().Msg("Successfully connected to Redis (Standalone)")
	return client
}

func connectRedisCluster(cfg *config.Config) *redis.ClusterClient {
	nodes := cfg.Get("Redis")["ClusterNodes"].(string)
	addrs := strings.Split(nodes, ",")
	password := cfg.Get("Redis")["Password"].(string)

	options := &redis.ClusterOptions{
		Addrs:    addrs,
		Password: password,
	}

	client := redis.NewClusterClient(options)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := client.Ping(ctx).Result()
	if err != nil {
		logger.Log.Err(err).Msg("failed to connect to Redis Cluster")
		log.Printf("Failed to connect to Redis Cluster: %v", err)
	}

	logger.Write.Info().Msg("Successfully connected to Redis Cluster")
	return client
}
