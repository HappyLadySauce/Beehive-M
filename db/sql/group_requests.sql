CREATE TABLE group_requests (
    request_id BIGINT PRIMARY KEY AUTO_INCREMENT,
    applicant_id BIGINT NOT NULL COMMENT '申请人ID',
    group_id BIGINT NOT NULL COMMENT '群ID',
    reason VARCHAR(500) COMMENT '申请理由',
    status TINYINT DEFAULT 0 COMMENT '状态: 0-待处理 1-同意 2-拒绝 3-忽略',
    handler_id BIGINT COMMENT '处理人ID',
    handle_time DATETIME COMMENT '处理时间',
    create_time DATETIME DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_applicant_group(applicant_id, group_id),
    INDEX idx_group_status(group_id, status),
    FOREIGN KEY (applicant_id) REFERENCES users(user_id),
    FOREIGN KEY (group_id) REFERENCES groups(group_id),
    FOREIGN KEY (handler_id) REFERENCES users(user_id)
) COMMENT '群申请表';