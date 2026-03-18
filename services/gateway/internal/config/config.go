// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf

	RedisHost     string `json:"RedisHost"`
	RedisPass     string `json:"RedisPass"`
	RedisDB       int    `json:"RedisDB"`
}
