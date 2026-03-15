package svc

import (
	"context"

	"github.com/HappyLadySauce/Beehive-M/services/user/internal/config"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB	*gorm.DB

	Redis *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	// Connect to PostgreSQL
	db, err := gorm.Open(postgres.Open(c.PostgreSQLDSN), &gorm.Config{})
	if err != nil {
		panic("failed to connect to PostgreSQL: " + err.Error())
	}

	// Connect to Redis
	redis := redis.NewClient(&redis.Options{
		Addr:     c.RedisHost,
		Password: c.RedisPass,
		DB:       c.RedisDB,
	})
	// Ping Redis to check if the connection is successful
	_, err = redis.Ping(context.Background()).Result()
	if err != nil {
		panic("failed to connect to Redis: " + err.Error())
	}

	// Return the ServiceContext
	return &ServiceContext{
		Config: c,
		DB:     db,
		Redis:  redis,
	}
}