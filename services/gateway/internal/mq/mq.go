package mq

import (
	"context"
	"fmt"

	"github.com/HappyLadySauce/Beehive-M/services/gateway/internal/config"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/zeromicro/go-zero/core/logx"
)

type Message struct {
	Action	string `json:"action"`
	ConnID    string                 `json:"connId"`
	UserID    string                 `json:"userId"`
	DeviceID  string                 `json:"deviceId"`
	GatewayID string                 `json:"gatewayId"`
	Reason    string                 `json:"reason,omitempty"`
	Timestamp int64                  `json:"timestamp"`
	Data      map[string]interface{} `json:"data,omitempty"`
}

// MQ RabbitMQ 客户端
type MQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	logger  logx.Logger
}

func NewMQ(cfg config.Config) (*MQ, error) {
	conn, err := amqp.Dial(cfg.RabbitMQURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}

	channel, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("failed to open channel: %w", err)
	}

	// 声明交换机
	if err := channel.ExchangeDeclare(
		"gateway_exchange",		// 交换机名称
		"direct",       // 类型
		true,           // 持久化
		false,          // 自动删除
		false,          // 内部
		false,          // 不等待
		nil,            // 参数
	); err != nil {
		channel.Close()
		conn.Close()
		return nil, fmt.Errorf("failed to declare exchange: %w", err)
	}

	// 声明队列
	if _, err := channel.QueueDeclare(
		"gateway_queue", // 队列名称
		true,      // 持久化
		false,     // 自动删除
		false,     // 排他
		false,     // 不等待
		nil,       // 参数
	); err != nil {
		channel.Close()
		conn.Close()
		return nil, fmt.Errorf("failed to declare queue: %w", err)
	}

	// 绑定队列到交换机
	if err := channel.QueueBind(
		"gateway_queue",			// 队列名称
		"gateway_event",	// 路由键
		"gateway_exchange",			// 交换机名称
		false,
		nil,
	); err != nil {
		channel.Close()
		conn.Close()
		return nil, fmt.Errorf("failed to bind queue: %w", err)
	}

	return &MQ{
		conn: conn,
		channel: channel,
		logger:	logx.WithContext(context.Background()),
	}, nil
}

func (m *MQ) Close() error {
	if m.channel != nil {
		if err := m.channel.Close(); err != nil {
			m.logger.Errorf("failed to close channel: %v", err)
		}
	}
	if m.conn != nil {
		if err := m.conn.Close(); err != nil {
			m.logger.Errorf("failed to close connection: %v", err)
		}
	}
	return nil
}