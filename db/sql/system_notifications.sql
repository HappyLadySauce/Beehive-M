CREATE TABLE system_notifications (
    notification_id BIGINT PRIMARY KEY AUTO_INCREMENT,
    user_id BIGINT NOT NULL COMMENT '接收用户ID',
    type TINYINT NOT NULL COMMENT '通知类型: 1-好友申请 2-群申请 3-入群邀请 4-系统公告',
    title VARCHAR(200) COMMENT '通知标题',
    content TEXT NOT NULL COMMENT '通知内容',
    related_id BIGINT COMMENT '关联ID(申请ID/群ID等)',
    is_read TINYINT DEFAULT 0 COMMENT '是否已读',
    extra_data JSON COMMENT '额外数据',
    create_time DATETIME DEFAULT CURRENT_TIMESTAMP,
    read_time DATETIME COMMENT '已读时间',
    INDEX idx_user_type_status(user_id, type, is_read),
    INDEX idx_create_time(create_time),
    FOREIGN KEY (user_id) REFERENCES users(user_id)
) COMMENT '系统通知表';