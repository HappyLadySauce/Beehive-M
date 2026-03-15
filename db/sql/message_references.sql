CREATE TABLE message_references (
    reference_id BIGINT PRIMARY KEY AUTO_INCREMENT,
    msg_id BIGINT NOT NULL COMMENT '消息ID',
    msg_scene_type TINYINT NOT NULL COMMENT '消息场景: 1-私聊 2-群聊',
    ref_msg_id BIGINT NOT NULL COMMENT '被引用消息ID',
    ref_msg_scene_type TINYINT NOT NULL COMMENT '被引用消息场景: 1-私聊 2-群聊',
    ref_sender_id BIGINT COMMENT '被引用发送者ID',
    ref_content_preview VARCHAR(500) COMMENT '被引用内容预览',
    create_time DATETIME DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_msg_scene_id(msg_scene_type, msg_id),
    INDEX idx_ref_msg_scene_id(ref_msg_scene_type, ref_msg_id)
) COMMENT '消息引用关系表';