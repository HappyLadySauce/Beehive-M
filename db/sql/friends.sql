CREATE TABLE friends (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    user_id BIGINT NOT NULL COMMENT '用户ID',
    friend_id BIGINT NOT NULL COMMENT '好友ID',
    remark VARCHAR(100) COMMENT '好友备注',
    source VARCHAR(50) COMMENT '来源:手机/QQ号/群聊等',
    is_blocked TINYINT DEFAULT 0 COMMENT '是否拉黑: 0-否 1-是',
    is_special TINYINT DEFAULT 0 COMMENT '特别关心: 0-否 1-是',
    is_muted TINYINT DEFAULT 0 COMMENT '消息免打扰: 0-否 1-是',
    add_time DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
    last_interact_time DATETIME COMMENT '最后互动时间',
    UNIQUE KEY uk_user_friend(user_id, friend_id),
    INDEX idx_friend_id(friend_id),
    FOREIGN KEY (user_id) REFERENCES users(user_id),
    FOREIGN KEY (friend_id) REFERENCES users(user_id)
) COMMENT '好友关系表';