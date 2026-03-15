CREATE TABLE resources (
    resource_id VARCHAR(64) PRIMARY KEY COMMENT '资源ID(MD5/雪花ID)',
    uploader_id BIGINT NOT NULL COMMENT '上传者ID',
    file_name VARCHAR(255) NOT NULL COMMENT '原始文件名',
    file_type VARCHAR(50) NOT NULL COMMENT '文件类型: image/jpeg',
    file_size BIGINT NOT NULL COMMENT '文件大小(字节)',
    file_hash VARCHAR(64) COMMENT '文件哈希(去重)',
    storage_path VARCHAR(500) NOT NULL COMMENT '存储路径',
    width INT COMMENT '图片/视频宽度',
    height INT COMMENT '图片/视频高度',
    upload_time DATETIME DEFAULT CURRENT_TIMESTAMP,
    status TINYINT DEFAULT 1 COMMENT '1-正常 2-已删除',
    is_temp TINYINT DEFAULT 0 COMMENT '是否为临时文件',
    expire_time DATETIME COMMENT '过期时间(临时文件)',
    INDEX idx_uploader(uploader_id),
    INDEX idx_file_hash(file_hash),
    INDEX idx_expire_time(expire_time),
    FOREIGN KEY (uploader_id) REFERENCES users(user_id)
) COMMENT '资源文件表';