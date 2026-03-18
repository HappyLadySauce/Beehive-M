package ws

import (
	"context"
	"fmt"
	"sync"

	"github.com/HappyLadySauce/Beehive-M/services/gateway/internal/config"
	"github.com/HappyLadySauce/errors"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	"github.com/redis/go-redis/v9"
)

// 定义 Connection 表示单条 WebSocket 连接，用于管理单条连接的上下文信息
type Connection struct {
	// 上下文
	ctx context.Context
	// 连接ID
	ConnID string
	// 用户ID
	UserID string
	// 设备ID
	DeviceID string
	// 网关ID
	GatewayID string
	// 连接
	conn *websocket.Conn
	// 写锁
	writeMu sync.Mutex
}

// 定义 Hub 用于管理当前进程所有 WebSocket 连接，用于分配 ConnID 与 Connection 的映射关系
type Hub struct {
	// 上下文
	ctx context.Context
	// 网关ID
	gatewayID string
	// 连接
	conns map[string]*Connection // ConnID 与 Connection 的映射关系
	// 读写锁
	mu sync.RWMutex
	// Redis 客户端
	rdb redis.Client
	// 分布式锁
	rs *redsync.Redsync
}

// NewHub 创建连接管理器，gatewayID 用于多实例部署时区分本实例。
func NewHub(gatewayID string, cfg config.Config) *Hub {
	// 生成网关ID
	if gatewayID == "" {
		gatewayID = "gateway-" + uuid.Must(uuid.NewRandom()).String()
	}

	// 创建 redis 客户端
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisHost,
		Password: cfg.RedisPass,
		DB:       cfg.RedisDB,
	})
	// 测试连接
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		panic("failed to connect to Redis: " + err.Error())
	}

	pool := goredis.NewPool(client)
	rs := redsync.New(pool)

	return &Hub{
		ctx:       context.Background(),
		gatewayID:  gatewayID,
		conns:      make(map[string]*Connection),
		rs:  rs,
	}
}

// Close 关闭连接管理器
func (h *Hub) Close() []error {
	h.mu.Lock()
	defer h.mu.Unlock()

	var closeErrs []error

	// 关闭所有连接
	for connID, conn := range h.conns {
		if err := conn.Close(); err != nil {
			closeErrs = append(closeErrs, errors.Wrap(err, "failed to close connection"))
		}
		delete(h.conns, connID)
	}

	// 关闭 Redis 客户端
	if err := h.rdb.Close(); err != nil {
		closeErrs = append(closeErrs, errors.Wrap(err, "failed to close Redis client"))
	}

	return closeErrs
}

// Close 关闭连接
func (c *Connection) Close() error {
	// 连接信息
	info := fmt.Sprintf("connID: %s, userID: %s, deviceID: %s, gatewayID: %s",
		c.ConnID, c.UserID, c.DeviceID, c.GatewayID)
	c.writeMu.Lock()
	defer c.writeMu.Unlock()
	if c.conn != nil {
		err := c.conn.Close()
		if err != nil {
			return errors.Wrap(err, "failed to close WebSocket connection: "+info)
		}
		c.conn = nil
	}
	return nil
}

// Register 将已升级的 WebSocket 注册到 Hub，返回带 ConnID 的 Connection。
func (h *Hub) Register(conn *websocket.Conn, userID, deviceID string) *Connection {
	connID := "connID:" + userID + "-" + deviceID
	newConn := &Connection{
		ctx:       h.ctx,
		ConnID:    connID,
		UserID:    userID,
		DeviceID:  deviceID,
		GatewayID: h.gatewayID,
		conn:      conn,
	}

	return newConn
}




