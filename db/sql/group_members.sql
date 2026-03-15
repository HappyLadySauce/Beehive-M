CREATE TABLE group_members (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    group_id BIGINT NOT NULL COMMENT '群ID',
    user_id BIGINT NOT NULL COMMENT '用户ID',
    role TINYINT DEFAULT 0 COMMENT '角色: 0-普通成员 1-管理员 2-群主',
    nickname_in_group VARCHAR(100) COMMENT '群内昵称',
    join_time DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '加入时间',
    last_read_msg_id BIGINT DEFAULT 0 COMMENT '最后已读消息ID',
    is_muted TINYINT DEFAULT 0 COMMENT '是否禁言: 0-否 1-是',
    is_special TINYINT DEFAULT 0 COMMENT '特别关注: 0-否 1-是',
    UNIQUE KEY uk_group_user(group_id, user_id),
    INDEX idx_group_id(group_id),
    INDEX idx_user_id(user_id),
    FOREIGN KEY (group_id) REFERENCES groups(group_id),
    FOREIGN KEY (user_id) REFERENCES users(user_id)
) COMMENT '群成员表';