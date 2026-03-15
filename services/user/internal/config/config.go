package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf

	// PostgreSQL DSN
	PostgreSQLDSN	string `json:""`

	// Redis connction
	RedisHost		string `json:""`
	RedisPass		string `json:""`
	RedisDB			int `json:""`
}
