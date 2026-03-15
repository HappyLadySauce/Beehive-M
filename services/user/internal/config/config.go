package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf

	// PostgreSQL
	PostgreSQLDSN string `json:"PostgreSQLDSN"`

	// Redis
	RedisHost     string `json:"RedisHost"`
	RedisPass     string `json:"RedisPass"`
	RedisDB       int    `json:"RedisDB"`
}
