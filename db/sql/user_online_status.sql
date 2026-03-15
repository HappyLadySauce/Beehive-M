CREATE TABLE user_online_status (
    user_id BIGINT PRIMARY KEY,
    status TINYINT DEFAULT 0 COMMENT '在线状态: 0-离线 1-在线 2-忙碌 3-离开',
    device_type VARCHAR(50) COMMENT '设备类型: web/android/ios',
    last_heartbeat DATETIME COMMENT '最后心跳时间',
    client_ip VARCHAR(50) COMMENT '客户端IP',
    update_time DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(user_id)
) COMMENT '用户在线状态表';