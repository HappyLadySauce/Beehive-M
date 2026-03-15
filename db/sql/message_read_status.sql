CREATE TABLE message_read_status (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    msg_id BIGINT NOT NULL COMMENT '消息ID',
    user_id BIGINT NOT NULL COMMENT '用户ID',
    scene_type TINYINT NOT NULL COMMENT '消息场景: 1-私聊 2-群聊',
    is_read TINYINT DEFAULT 0 COMMENT '是否已读',
    read_time DATETIME COMMENT '已读时间',
    UNIQUE KEY uk_msg_user_scene(msg_id, user_id, scene_type),
    INDEX idx_user_scene(user_id, scene_type),
    INDEX idx_user_scene_read(user_id, scene_type, is_read),
    INDEX idx_scene_msg(scene_type, msg_id)
) COMMENT '消息已读状态表';