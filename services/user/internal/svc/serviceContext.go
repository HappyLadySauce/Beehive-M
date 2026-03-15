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

	db, err := gorm.Open(postgres.Open(c.PostgreSQLDSN), &gorm.Config{})
	if err != nil {
		panic("failed to connect to PostgreSQL: " + err.Error())
	}

	redis := redis.NewClient(&redis.Options{
		Addr:     c.RedisHost,
		Password: c.RedisPass,
		DB:       c.RedisDB,
	})

	_, err = redis.Ping(context.Background()).Result()
	if err != nil {
		panic("failed to connect to Redis: " + err.Error())
	}

	return &ServiceContext{
		Config: c,
		DB:     db,
		Redis:  redis,
	}
}
