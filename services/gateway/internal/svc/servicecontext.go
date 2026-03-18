// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"github.com/HappyLadySauce/Beehive-M/services/gateway/internal/config"
	"github.com/HappyLadySauce/Beehive-M/services/gateway/internal/ws"
)

type ServiceContext struct {
	Config config.Config

	Hub  *ws.Hub
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 创建 Redis 客户端


	// 创建连接管理器
	hub := ws.NewHub("", c)

	return &ServiceContext{
		Config: c,
		Hub:   hub,
	}
}

func (s *ServiceContext) Close() {
	s.Hub.Close()
}