CREATE TABLE friend_requests (
    request_id BIGINT PRIMARY KEY AUTO_INCREMENT,
    applicant_id BIGINT NOT NULL COMMENT '申请人ID',
    target_id BIGINT NOT NULL COMMENT '目标用户ID',
    reason VARCHAR(500) COMMENT '申请理由',
    status TINYINT DEFAULT 0 COMMENT '状态: 0-待处理 1-同意 2-拒绝 3-忽略',
    handle_time DATETIME COMMENT '处理时间',
    create_time DATETIME DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_applicant(applicant_id),
    INDEX idx_target(target_id),
    INDEX idx_target_status(target_id, status),
    INDEX idx_status(status),
    FOREIGN KEY (applicant_id) REFERENCES users(user_id),
    FOREIGN KEY (target_id) REFERENCES users(user_id)
) COMMENT '好友申请表';