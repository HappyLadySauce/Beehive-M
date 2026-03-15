CREATE TABLE sync_status (
    user_id BIGINT NOT NULL COMMENT '用户ID',
    device_id VARCHAR(100) NOT NULL COMMENT '设备标识',
    last_sync_msg_id BIGINT DEFAULT 0 COMMENT '最后同步消息ID',
    last_sync_time DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    sync_offset BIGINT DEFAULT 0 COMMENT '同步偏移量',
    PRIMARY KEY (user_id, device_id),
    INDEX idx_user_last_sync(user_id, last_sync_time),
    FOREIGN KEY (user_id) REFERENCES users(user_id)
) COMMENT '消息同步状态表';