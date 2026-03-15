CREATE TABLE recalled_messages (
    recall_id BIGINT PRIMARY KEY AUTO_INCREMENT,
    msg_id BIGINT NOT NULL COMMENT '原消息ID',
    msg_type TINYINT NOT NULL COMMENT '消息类型',
    operator_id BIGINT NOT NULL COMMENT '操作人ID',
    reason VARCHAR(200) COMMENT '撤回原因',
    original_content TEXT COMMENT '原始内容备份',
    original_file_info JSON COMMENT '原始文件信息',
    recall_time DATETIME DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_msg_type_id(msg_type, msg_id),
    INDEX idx_operator_time(operator_id, recall_time)
) COMMENT '撤回消息记录表';