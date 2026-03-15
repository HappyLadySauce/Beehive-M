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
		// 关闭 DB 连接
		sqlDB, dbErr := db.DB()
		if dbErr == nil && sqlDB != nil {
			_ = sqlDB.Close()
		}
		// 关闭 Redis 连接
		_ = redis.Close()
		panic("failed to connect to Redis: " + err.Error())
	}

	return &ServiceContext{
		Config: c,
		DB:     db,
		Redis:  redis,
	}
}

// Close 释放 DB 与 Redis 连接，应在服务退出时调用（如 main 中 defer）
func (s *ServiceContext) Close() error {
	var err error
	if s.DB != nil {
		var sqlDB interface{ Close() error }
		sqlDB, err = s.DB.DB()
		if err == nil && sqlDB != nil {
			err = sqlDB.Close()
		}
	}
	if s.Redis != nil {
		if closeErr := s.Redis.Close(); closeErr != nil && err == nil {
			err = closeErr
		}
	}
	return err
}
