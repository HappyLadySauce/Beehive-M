CREATE TABLE groups (
    group_id BIGINT PRIMARY KEY AUTO_INCREMENT,
    group_number VARCHAR(50) UNIQUE NOT NULL COMMENT '群号',
    group_name VARCHAR(200) NOT NULL COMMENT '群名称',
    avatar VARCHAR(500) COMMENT '群头像',
    owner_id BIGINT NOT NULL COMMENT '群主用户ID',
    announcement TEXT COMMENT '群公告',
    description VARCHAR(500) COMMENT '群描述',
    max_members INT DEFAULT 500 COMMENT '最大成员数',
    privacy TINYINT DEFAULT 0 COMMENT '隐私设置: 0-公开 1-私密',
    verification TINYINT DEFAULT 0 COMMENT '加群验证: 0-自由 1-验证 2-禁止',
    status TINYINT DEFAULT 1 COMMENT '状态: 1-正常 2-解散',
    create_time DATETIME DEFAULT CURRENT_TIMESTAMP,
    update_time DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_owner(owner_id),
    INDEX idx_group_number(group_number),
    FOREIGN KEY (owner_id) REFERENCES users(user_id)
) COMMENT '群组表';