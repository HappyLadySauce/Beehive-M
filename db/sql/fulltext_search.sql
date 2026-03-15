-- =============================================================================
-- PostgreSQL 全文检索：用户 / 群组 / 私聊消息 / 群聊消息
-- 依赖：PostgreSQL 12+（GENERATED 列）
-- 中文分词（可选）：安装 zhparser 后，将下方 'simple' 改为 'chinese'
-- =============================================================================

-- 若使用中文分词，先执行（需先安装 zhparser 扩展）：
-- CREATE EXTENSION IF NOT EXISTS zhparser;
-- CREATE TEXT SEARCH CONFIGURATION chinese (PARSER = zhparser);
-- ALTER TEXT SEARCH CONFIGURATION chinese ADD MAPPING FOR n,v,a,i,e,l,j WITH simple;

-- -----------------------------------------------------------------------------
-- 1. 用户表：账号、昵称、个性签名（A 权重更高，B 次之）
-- -----------------------------------------------------------------------------
ALTER TABLE users
  ADD COLUMN IF NOT EXISTS search_vector tsvector
  GENERATED ALWAYS AS (
    setweight(to_tsvector('simple', coalesce(account, '')), 'A')
    || setweight(to_tsvector('simple', coalesce(nickname, '')), 'A')
    || setweight(to_tsvector('simple', coalesce(signature, '')), 'B')
  ) STORED;

CREATE INDEX IF NOT EXISTS idx_users_search ON users USING GIN (search_vector);

-- -----------------------------------------------------------------------------
-- 2. 群组表：群号、群名称、群描述
-- -----------------------------------------------------------------------------
ALTER TABLE groups
  ADD COLUMN IF NOT EXISTS search_vector tsvector
  GENERATED ALWAYS AS (
    setweight(to_tsvector('simple', coalesce(group_number, '')), 'A')
    || setweight(to_tsvector('simple', coalesce(group_name, '')), 'A')
    || setweight(to_tsvector('simple', coalesce(description, '')), 'B')
  ) STORED;

CREATE INDEX IF NOT EXISTS idx_groups_search ON groups USING GIN (search_vector);

-- -----------------------------------------------------------------------------
-- 3. 私聊消息表：仅文本消息且未撤回的内容参与检索（前 10000 字符）
-- -----------------------------------------------------------------------------
ALTER TABLE private_messages
  ADD COLUMN IF NOT EXISTS search_vector tsvector
  GENERATED ALWAYS AS (
    CASE
      WHEN msg_type = 1 AND content IS NOT NULL AND (is_recalled IS NULL OR is_recalled = 0)
      THEN to_tsvector('simple', left(content, 10000))
      ELSE ''::tsvector
    END
  ) STORED;

CREATE INDEX IF NOT EXISTS idx_private_messages_search ON private_messages USING GIN (search_vector);

-- -----------------------------------------------------------------------------
-- 4. 群聊消息表：仅文本消息且未撤回的内容参与检索
-- -----------------------------------------------------------------------------
ALTER TABLE group_messages
  ADD COLUMN IF NOT EXISTS search_vector tsvector
  GENERATED ALWAYS AS (
    CASE
      WHEN msg_type = 1 AND content IS NOT NULL AND (is_recalled IS NULL OR is_recalled = 0)
      THEN to_tsvector('simple', left(content, 10000))
      ELSE ''::tsvector
    END
  ) STORED;

CREATE INDEX IF NOT EXISTS idx_group_messages_search ON group_messages USING GIN (search_vector);
