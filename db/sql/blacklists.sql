CREATE TABLE blacklists (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    user_id BIGINT NOT NULL COMMENT '用户ID',
    blocked_user_id BIGINT NOT NULL COMMENT '被拉黑用户ID',
    blocked_type TINYINT DEFAULT 1 COMMENT '1-私聊 2-群聊 3-全部',
    remark VARCHAR(200) COMMENT '拉黑备注',
    block_time DATETIME DEFAULT CURRENT_TIMESTAMP,
    expire_time DATETIME COMMENT '过期时间(空为永久)',
    UNIQUE KEY uk_user_blocked(user_id, blocked_user_id),
    INDEX idx_expire_time(expire_time),
    FOREIGN KEY (user_id) REFERENCES users(user_id),
    FOREIGN KEY (blocked_user_id) REFERENCES users(user_id)
) COMMENT '黑名单表';