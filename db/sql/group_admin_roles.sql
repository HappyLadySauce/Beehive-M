CREATE TABLE group_admin_roles (
    role_id BIGINT PRIMARY KEY AUTO_INCREMENT,
    group_id BIGINT NOT NULL COMMENT '群ID',
    role_name VARCHAR(50) NOT NULL COMMENT '角色名',
    permissions JSON NOT NULL COMMENT '权限配置',
    is_default TINYINT DEFAULT 0 COMMENT '是否默认角色',
    create_time DATETIME DEFAULT CURRENT_TIMESTAMP,
    UNIQUE KEY uk_group_role(group_id, role_name),
    INDEX idx_group_id(group_id),
    FOREIGN KEY (group_id) REFERENCES groups(group_id)
) COMMENT '群管理员角色表';

CREATE TABLE group_admin_permissions (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    group_id BIGINT NOT NULL COMMENT '群ID',
    user_id BIGINT NOT NULL COMMENT '用户ID',
    role_id BIGINT NOT NULL COMMENT '角色ID',
    custom_permissions JSON COMMENT '自定义权限',
    expire_time DATETIME COMMENT '权限过期时间',
    grant_time DATETIME DEFAULT CURRENT_TIMESTAMP,
    grantor_id BIGINT COMMENT '授权人ID',
    UNIQUE KEY uk_group_user(group_id, user_id),
    INDEX idx_user_id(user_id),
    FOREIGN KEY (group_id) REFERENCES groups(group_id),
    FOREIGN KEY (user_id) REFERENCES users(user_id),
    FOREIGN KEY (role_id) REFERENCES group_admin_roles(role_id),
    FOREIGN KEY (grantor_id) REFERENCES users(user_id)
) COMMENT '群管理员权限分配表';