CREATE TABLE group_invitations (
    invitation_id BIGINT PRIMARY KEY AUTO_INCREMENT,
    group_id BIGINT NOT NULL COMMENT '群ID',
    inviter_id BIGINT NOT NULL COMMENT '邀请人ID',
    invitee_id BIGINT NOT NULL COMMENT '被邀请人ID',
    status TINYINT DEFAULT 0 COMMENT '状态: 0-待处理 1-同意 2-拒绝 3-过期',
    expire_time DATETIME COMMENT '过期时间',
    create_time DATETIME DEFAULT CURRENT_TIMESTAMP,
    handle_time DATETIME COMMENT '处理时间',
    UNIQUE KEY uk_group_invitee(group_id, invitee_id),
    INDEX idx_invitee_status(invitee_id, status),
    FOREIGN KEY (group_id) REFERENCES groups(group_id),
    FOREIGN KEY (inviter_id) REFERENCES users(user_id),
    FOREIGN KEY (invitee_id) REFERENCES users(user_id)
) COMMENT '群组邀请表';