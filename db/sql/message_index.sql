CREATE TABLE message_index (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    msg_id BIGINT NOT NULL COMMENT '消息ID',
    msg_type TINYINT NOT NULL COMMENT '1-私聊 2-群聊',
    conversation_key VARCHAR(100) NOT NULL COMMENT '会话KEY: user1_user2 或 group_id',
    sender_id BIGINT NOT NULL COMMENT '发送者ID',
    send_time DATETIME NOT NULL COMMENT '发送时间',
    content_preview VARCHAR(200) COMMENT '消息预览',
    msg_data_type TINYINT COMMENT '数据类型',
    INDEX idx_conversation_time(conversation_key, send_time),
    INDEX idx_sender_time(sender_id, send_time),
    INDEX idx_msg_type_id(msg_type, msg_id)
) COMMENT '消息索引表';