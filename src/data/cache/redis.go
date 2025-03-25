package cache

import (
	"golang-web-api/config"
	"time"
	"golang-web-api/pkg/logging"

	"github.com/go-redis/redis"
)

var redisClient *redis.Client


var logger = logging.NewLogger(&config.GetConfig().Logger)


func InitRedis(cfg *config.Redis) error {

	redisClient = redis.NewClient(&redis.Options{
		Addr:               cfg.Host + ":" + cfg.Port,
		Password:           cfg.Password,
		DB:                 cfg.DB,
		PoolSize:           cfg.PoolSize,
		PoolTimeout:        cfg.PoolTimeout * time.Second,
		IdleTimeout:        cfg.IdleTimeout * time.Second,
		ReadTimeout:        cfg.ReadTimeout * time.Second,
		WriteTimeout:       cfg.WriteTimeout * time.Second,
		IdleCheckFrequency: cfg.IdleCheckFrequency * time.Second,	
	})

	_, err := redisClient.Ping().Result()
	if err != nil {
		return err
	}

	logger.Info(logging.Redis , logging.StartUp , "Connected to Redis", nil)

	return nil
}

func GetRedisClient() *redis.Client {
	return redisClient
}

func CloseRedis() {
	redisClient.Close()
}
