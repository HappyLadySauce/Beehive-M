CREATE TABLE user_emoticons (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    user_id BIGINT NOT NULL COMMENT '用户ID',
    emoticon_id VARCHAR(100) NOT NULL COMMENT '表情标识',
    emoticon_url VARCHAR(500) NOT NULL COMMENT '表情URL',
    emoticon_type TINYINT DEFAULT 1 COMMENT '1-系统 2-自定义',
    category VARCHAR(50) DEFAULT 'default' COMMENT '分类',
    use_count INT DEFAULT 0 COMMENT '使用次数',
    last_use_time DATETIME COMMENT '最后使用时间',
    create_time DATETIME DEFAULT CURRENT_TIMESTAMP,
    UNIQUE KEY uk_user_emoticon(user_id, emoticon_id),
    INDEX idx_user_category(user_id, category),
    FOREIGN KEY (user_id) REFERENCES users(user_id)
) COMMENT '用户表情包表';