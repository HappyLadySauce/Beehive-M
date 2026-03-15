CREATE TABLE group_announcements (
    announcement_id BIGINT PRIMARY KEY AUTO_INCREMENT,
    group_id BIGINT NOT NULL COMMENT '群ID',
    author_id BIGINT NOT NULL COMMENT '发布者ID',
    title VARCHAR(200) NOT NULL COMMENT '公告标题',
    content TEXT NOT NULL COMMENT '公告内容',
    is_pinned TINYINT DEFAULT 0 COMMENT '是否置顶',
    is_important TINYINT DEFAULT 0 COMMENT '是否重要',
    publish_time DATETIME DEFAULT CURRENT_TIMESTAMP,
    expire_time DATETIME COMMENT '过期时间',
    read_count INT DEFAULT 0 COMMENT '阅读数',
    status TINYINT DEFAULT 1 COMMENT '1-正常 2-删除',
    INDEX idx_group_time(group_id, publish_time),
    INDEX idx_group_pinned_time(group_id, is_pinned DESC, publish_time DESC),
    FOREIGN KEY (group_id) REFERENCES groups(group_id),
    FOREIGN KEY (author_id) REFERENCES users(user_id)
) COMMENT '群公告表';

CREATE TABLE group_announcement_reads (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    announcement_id BIGINT NOT NULL COMMENT '公告ID',
    user_id BIGINT NOT NULL COMMENT '用户ID',
    read_time DATETIME DEFAULT CURRENT_TIMESTAMP,
    UNIQUE KEY uk_announcement_user(announcement_id, user_id),
    INDEX idx_user_id(user_id),
    FOREIGN KEY (announcement_id) REFERENCES group_announcements(announcement_id),
    FOREIGN KEY (user_id) REFERENCES users(user_id)
) COMMENT '群公告阅读记录表';