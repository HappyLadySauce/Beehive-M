CREATE TABLE user_activity_stats (
    stat_date DATE NOT NULL COMMENT '统计日期',
    user_id BIGINT NOT NULL COMMENT '用户ID',
    login_count INT DEFAULT 0 COMMENT '登录次数',
    msg_sent_count INT DEFAULT 0 COMMENT '发送消息数',
    msg_received_count INT DEFAULT 0 COMMENT '接收消息数',
    file_upload_count INT DEFAULT 0 COMMENT '文件上传数',
    file_download_count INT DEFAULT 0 COMMENT '文件下载数',
    online_duration INT DEFAULT 0 COMMENT '在线时长(分钟)',
    last_active_time DATETIME COMMENT '最后活跃时间',
    PRIMARY KEY (stat_date, user_id),
    INDEX idx_user_date(user_id, stat_date),
    FOREIGN KEY (user_id) REFERENCES users(user_id)
) COMMENT '用户活跃度统计表';