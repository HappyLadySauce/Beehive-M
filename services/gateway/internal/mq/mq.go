package mq

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/HappyLadySauce/Beehive-M/pkg/code"
	"github.com/HappyLadySauce/Beehive-M/services/gateway/internal/config"
	"github.com/HappyLadySauce/errors"

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

// Publish 发布消息
func (m *MQ) Publish(routingKey string, msg *Message) error {
	data, err := json.Marshal(msg)
	if err != nil {
		m.logger.Errorf("failed to marshal message: %w", err)
		return errors.WithCode(code.CodeMarshalFailed, "failed to marshal message")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	if err := m.channel.PublishWithContext(
		ctx,
		"gateway_exchange",	// 交换机
		routingKey,			// 路由键
		false,				// 强制
		false,				// 立即
		amqp.Publishing{
			ContentType: "application/json",
			DeliveryMode: amqp.Persistent,	// 持久化消息
			Timestamp: time.Now(),
			Body: data,
		},
	); err != nil {
		m.logger.Errorf("failed to publish message: %w", err)
		return errors.WithCode(code.CodeMQPushFailed, "failed to publish message")
	}

	return nil
}

// Subscribe 订阅消息
func (m *MQ) Subscribe(queueName, routingKey string) (<- chan *Message, error) {
	// 确保队列存在
	if _, err := m.channel.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		m.logger.Errorf("failed to declare queue: %w", err)
		return nil, errors.WithCode(code.CodeMQDeclareFailed, "failed to declare queue")
	}

	// 绑定队列
	if err := m.channel.QueueBind(
		queueName,
		routingKey,
		"gateway_exchange",
		false,
		nil,
	); err != nil {
		m.logger.Errorf("failed to bind queue: %w", err)
		return nil, errors.WithCode(code.CodeMQBindQueueFailed, "failed to bind queue")	
	}

	// 设置 QoS
	if err := m.channel.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	); err != nil {
		m.logger.Errorf("failed to set Qos: %w", err)
		return nil, errors.WithCode(code.CodeMQSetQosFailed, "failed to set Qos")	
	}
	
	msgs, err := m.channel.Consume(
		queueName,
		"",
		false,	// auto-ack
		false,	// exclusive
		false,	// no-local
		false,	// no-wait
		nil,
	)
	if err != nil {
		m.logger.Errorf("failed to consume messages: %w", err)
		return nil, errors.WithCode(code.CodeMQConsumeFailed, "failed to consume messages")
	}

	messageChan := make(chan *Message, 100)

	go func() {
		defer close(messageChan)
		for d := range msgs {
			var msg Message
			if err := json.Unmarshal(d.Body, &msg); err != nil {
				m.logger.Errorf("failed to unmarshal message: %v", err)
				d.Nack(false, true)
				continue
			}

			select {
			case messageChan <- &msg:
				d.Ack(false)
			default:
				m.logger.Errorf("message channel full, dropping message: %s", msg.ConnID)
				d.Nack(false, false) // 丢弃消息		
			}
		}
	}()

	return messageChan, nil
}





























