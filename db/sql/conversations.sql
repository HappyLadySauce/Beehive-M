CREATE TABLE conversations (
    conversation_id BIGINT PRIMARY KEY AUTO_INCREMENT,
    user_id BIGINT NOT NULL COMMENT '用户ID',
    target_type TINYINT NOT NULL COMMENT '目标类型: 1-用户 2-群组',
    target_id BIGINT NOT NULL COMMENT '目标ID(用户ID或群ID)',
    last_msg_id BIGINT COMMENT '最后一条消息ID',
    last_msg_preview VARCHAR(500) COMMENT '最后消息预览',
    unread_count INT DEFAULT 0 COMMENT '未读消息数',
    is_pinned TINYINT DEFAULT 0 COMMENT '是否置顶',
    is_muted TINYINT DEFAULT 0 COMMENT '是否免打扰',
    last_interact_time DATETIME DEFAULT CURRENT_TIMESTAMP,
    create_time DATETIME DEFAULT CURRENT_TIMESTAMP,
    UNIQUE KEY uk_user_target(user_id, target_type, target_id),
    INDEX idx_user_interact_time(user_id, last_interact_time),
    INDEX idx_user_unread(user_id, unread_count),
    FOREIGN KEY (user_id) REFERENCES users(user_id)
) COMMENT '会话表';