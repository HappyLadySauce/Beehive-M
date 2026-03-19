package ws

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/HappyLadySauce/Beehive-M/pkg/code"
	"github.com/HappyLadySauce/Beehive-M/services/gateway/internal/config"

	"github.com/HappyLadySauce/errors"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logx"
)

const (
	lockKeyPrefix = "ws:lock:"
	connKeyPrefix = "ws:conn:"
	userDevKeyPrefix = "ws:u:d:"
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

// redis 中存储的连接信息
type ConnectionInfo struct {
	ConnID    string `json:"connId"`
	UserID    string `json:"userId"`
	DeviceID  string `json:"deviceId"`
	GatewayID string `json:"gatewayId"`
	Status    string `json:"status"` // "online" 在线 "closing" 已关闭
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
	// 日志记录器
	logx.Logger
}

// NewHub 创建连接管理器，gatewayID 用于多实例部署时区分本实例。
func NewHub(gatewayID string, cfg config.Config) (*Hub, error) {
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
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}
	// 创建 redsync 实例
	pool := goredis.NewPool(client)
	rs := redsync.New(pool)

	return &Hub{
		ctx:       context.Background(),
		gatewayID:  gatewayID,
		conns:      make(map[string]*Connection),
		rs:  rs,
		Logger:	logx.WithContext(context.Background()),
	}, nil
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
func (h *Hub) Register(conn *websocket.Conn, userID, deviceID string) (*Connection, error) {
	connID := "connID:" + userID + "-" + deviceID

	// 获取分布式锁
	mutex := h.rs.NewMutex(lockKeyPrefix + connID)
	
	if err := mutex.Lock(); err != nil {
		h.Logger.Errorf("failed to lock connection: %v", connID)
		return nil, errors.WithCode(code.CodeDistributedLockGetFailed, "failed to lock connection")
	}
	defer mutex.Unlock()

	var oldConnInfo ConnectionInfo
	// 获取 Redis 中旧连接信息 并将其关闭
	if val, err := h.rdb.Get(h.ctx, connKeyPrefix + connID).Result(); err == nil && val != "" {
		// 如果连接存在 删除旧连接
		if err := json.Unmarshal([]byte(val), &oldConnInfo); err != nil {
			if oldConnInfo.GatewayID != h.gatewayID {
				// 通过消息队列向 gateway 频道发送关闭消息
			} else {
				// 如果旧连接在当前网关 直接关闭
				h.mu.RLock()
				if oldConn, exists := h.conns[connID]; exists {
					oldConn.Close()
				}
				h.mu.RUnlock()
			}
			// 删除旧连接
			h.rdb.Del(h.ctx, connKeyPrefix + connID)

			// 删除用户 - 设备映射
			h.rdb.Del(h.ctx, userDevKeyPrefix + userID + deviceID)
		} else {
			h.Logger.Errorf("nmarshal connctionInfo from redis failed: %v", err)
			return nil, errors.WithCode(code.CodeUnmarshalFailed, "unmarshal failed")
		}
	} else if err == nil && val == "" {
		h.Logger.Info("connctionInfo: %v from redis is nil", connID)
		// 删除旧连接
		h.rdb.Del(h.ctx, connKeyPrefix + connID)
		// 删除用户 - 设备映射
		h.rdb.Del(h.ctx, userDevKeyPrefix + userID + deviceID)
	} else {
		h.Logger.Errorf("get connctionInfo from redis failed: %v", err)
		return nil, errors.WithCode(code.CodeCacheGetFailed, "get connctionInfo from redis failed")
	}

	// 定义连接信息
	connInfo := ConnectionInfo{
		ConnID:    connID,
		UserID:    userID,
		DeviceID:  deviceID,
		GatewayID: h.gatewayID,
		Status:    "online",
	}

	// 转换连接信息 并存入 redis
	if data, err := json.Marshal(connInfo); err == nil {
		// 将 connectionInfo 存入 redis
		if err := h.rdb.SetNX(h.ctx, connKeyPrefix + connID, data, 24 * time.Hour).Err(); err != nil {
			h.Logger.Errorf("set connectionInfo to redis failed: %v", err)
			return nil, errors.WithCode(code.CodeCacheSetFailed, "set connectionInfo to redis failed")
		}

		// 将 user-dev 映射存入
		if err := h.rdb.SetNX(h.ctx, userDevKeyPrefix + userID + deviceID, connID, 24 * time.Hour).Err(); err != nil {
			h.Logger.Errorf("set user-dev to redis failed: %v", err)
			return nil, errors.WithCode(code.CodeCacheSetFailed, "set user-dev to redis failed")
		}
	}

	// 创建新连接
	newConn := &Connection{
		ctx:       h.ctx,
		ConnID:    connID,
		UserID:    userID,
		DeviceID:  deviceID,
		GatewayID: h.gatewayID,
		conn:      conn,
		writeMu:   sync.Mutex{},
	}
	
	// 保存到本地连接池
	h.mu.Lock()
	h.conns[connID] = newConn
	h.mu.Unlock()

	// 启动 gateway 消息队列

	// 返回连接
	return newConn, nil
}

// 注销 Websocket 连接
func (h *Hub) Deregister(connID string) error {
	h.mu.Lock()
	delete(h.conns, connID)
	h.mu.Unlock()

	var connInfo ConnectionInfo
	// 从 redis 删除 connectionInfo
	if val, err := h.rdb.Get(h.ctx, connKeyPrefix + connID).Result(); err != nil {
		if err := json.Unmarshal([]byte(val), &connInfo); err == nil {
			// 删除 redis connectionInfo
			h.rdb.Del(h.ctx, connKeyPrefix + connID)
			// 删除 user-dev 映射
			h.rdb.Del(h.ctx, userDevKeyPrefix + connInfo.UserID + connInfo.DeviceID)
		} else {
			h.Logger.Errorf("unmarshal connectionInfo from redis failed: %v", err)
			return errors.WithCode(code.CodeUnmarshalFailed, "unmarshal failed")
		}
		
	}

	return nil
}


