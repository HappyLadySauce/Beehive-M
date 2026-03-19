// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"github.com/HappyLadySauce/Beehive-M/services/gateway/internal/config"
	"github.com/HappyLadySauce/Beehive-M/services/gateway/internal/ws"
	"github.com/HappyLadySauce/Beehive-M/services/gateway/internal/mq"
)

type ServiceContext struct {
	Config	config.Config

	Hub		*ws.Hub
	MQ		*mq.MQ
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 创建连接管理器
	hub, err := ws.NewHub("", c)
	if err != nil {
		panic(err)
	}

	// 创建 MQ 网关消息队列
	mq, err := mq.NewMQ(c)
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config: c,
		Hub:	hub,
		MQ:		mq,
	}
}

func (s *ServiceContext) Close() {
	s.Hub.Close()
	s.MQ.Close()
}