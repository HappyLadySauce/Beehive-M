CREATE TABLE friend_recommendations (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    user_id BIGINT NOT NULL COMMENT '用户ID',
    recommended_user_id BIGINT NOT NULL COMMENT '推荐用户ID',
    reason VARCHAR(200) COMMENT '推荐原因',
    score DECIMAL(5,4) COMMENT '推荐分数',
    source VARCHAR(50) COMMENT '推荐来源',
    status TINYINT DEFAULT 0 COMMENT '0-未处理 1-已添加 2-忽略',
    create_time DATETIME DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_user_status(user_id, status),
    INDEX idx_score(score DESC),
    FOREIGN KEY (user_id) REFERENCES users(user_id),
    FOREIGN KEY (recommended_user_id) REFERENCES users(user_id)
) COMMENT '好友推荐表';