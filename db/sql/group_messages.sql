CREATE TABLE group_messages (
    msg_id BIGINT PRIMARY KEY AUTO_INCREMENT,
    group_id BIGINT NOT NULL COMMENT '群ID',
    sender_id BIGINT NOT NULL COMMENT '发送者ID',
    msg_type TINYINT NOT NULL COMMENT '消息类型: 1-文本 2-图片 3-文件 4-表情',
    content TEXT COMMENT '消息内容',
    file_url VARCHAR(500) COMMENT '文件URL',
    file_size INT COMMENT '文件大小',
    file_name VARCHAR(255) COMMENT '文件名',
    is_recalled TINYINT DEFAULT 0 COMMENT '是否撤回',
    at_user_ids JSON COMMENT '@的用户ID列表',
    send_time DATETIME DEFAULT CURRENT_TIMESTAMP,
    recall_time DATETIME COMMENT '撤回时间',
    INDEX idx_group_send_time(group_id, send_time),
    INDEX idx_sender_group(sender_id, group_id),
    INDEX idx_group_msg_time(group_id, msg_id),
    FOREIGN KEY (group_id) REFERENCES groups(group_id),
    FOREIGN KEY (sender_id) REFERENCES users(user_id)
) COMMENT '群聊消息表';