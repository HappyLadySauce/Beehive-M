CREATE TABLE search_histories (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    user_id BIGINT NOT NULL COMMENT '用户ID',
    keyword VARCHAR(200) NOT NULL COMMENT '搜索关键词',
    search_type TINYINT COMMENT '搜索类型:1-用户 2-群组 3-聊天记录',
    result_count INT DEFAULT 0 COMMENT '结果数量',
    search_time DATETIME DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_user_keyword(user_id, keyword),
    INDEX idx_user_time(user_id, search_time DESC),
    FOREIGN KEY (user_id) REFERENCES users(user_id)
) COMMENT '搜索历史表';