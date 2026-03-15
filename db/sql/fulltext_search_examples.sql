-- =============================================================================
-- 全文检索示例查询（执行 fulltext_search.sql 后再使用）
-- =============================================================================

-- -----------------------------------------------------------------------------
-- 1. 搜索用户（按关键词，排除已冻结，按相关度排序）
-- -----------------------------------------------------------------------------
-- 示例：关键词 "小明"
/*
SELECT user_id, account, nickname, avatar, signature,
       ts_rank(search_vector, plainto_tsquery('simple', '小明')) AS rank
FROM users
WHERE status = 1
  AND search_vector @@ plainto_tsquery('simple', '小明')
ORDER BY rank DESC
LIMIT 20;
*/

-- 短语或分词搜索（支持 AND/OR）
/*
SELECT user_id, nickname,
       ts_rank(search_vector, query) AS rank
FROM users,
     to_tsquery('simple', '小明 & 北京') AS query
WHERE status = 1 AND search_vector @@ query
ORDER BY rank DESC
LIMIT 20;
*/

-- -----------------------------------------------------------------------------
-- 2. 搜索群组（按关键词，仅正常群）
-- -----------------------------------------------------------------------------
/*
SELECT group_id, group_number, group_name, avatar, description,
       ts_rank(search_vector, plainto_tsquery('simple', '技术')) AS rank
FROM groups
WHERE status = 1
  AND search_vector @@ plainto_tsquery('simple', '技术')
ORDER BY rank DESC
LIMIT 20;
*/

-- -----------------------------------------------------------------------------
-- 3. 搜索私聊记录（某两人之间的消息，按时间倒序）
-- -----------------------------------------------------------------------------
-- 示例：sender_id=1 与 receiver_id=2 的会话中搜 "会议"
/*
SELECT msg_id, sender_id, receiver_id, content, send_time,
       ts_headline('simple', content, plainto_tsquery('simple', '会议'),
                   'MaxWords=35, MinWords=15, StartSel=<<, StopSel=>>') AS headline
FROM private_messages
WHERE ((sender_id = 1 AND receiver_id = 2) OR (sender_id = 2 AND receiver_id = 1))
  AND search_vector @@ plainto_tsquery('simple', '会议')
ORDER BY send_time DESC
LIMIT 50;
*/

-- -----------------------------------------------------------------------------
-- 4. 搜索群聊记录（某群内消息）
-- -----------------------------------------------------------------------------
-- 示例：group_id=100 内搜 "公告"
/*
SELECT msg_id, group_id, sender_id, content, send_time,
       ts_headline('simple', content, plainto_tsquery('simple', '公告'),
                   'MaxWords=35, MinWords=15, StartSel=<<, StopSel=>>') AS headline
FROM group_messages
WHERE group_id = 100
  AND search_vector @@ plainto_tsquery('simple', '公告')
ORDER BY send_time DESC
LIMIT 50;
*/

-- -----------------------------------------------------------------------------
-- 5. 应用层使用方式（Go 等）
-- 将用户输入做转义后传入，避免 SQL 注入；中文可先装 zhparser 再改用 'chinese'
-- plainto_tsquery 会自动把空格变成 AND，适合简单关键词
-- to_tsquery 可写 '词1 & 词2' 或 '词1 | 词2'，适合高级搜索
-- -----------------------------------------------------------------------------
