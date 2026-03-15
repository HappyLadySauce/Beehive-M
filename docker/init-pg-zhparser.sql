-- 首次启动时在 beehive 库中启用 zhparser 与 chinese 全文配置
-- 仅当数据目录为空时执行；若已有数据请手动执行本文件内容

\connect beehive

CREATE EXTENSION IF NOT EXISTS zhparser;

CREATE TEXT SEARCH CONFIGURATION chinese (PARSER = zhparser);
ALTER TEXT SEARCH CONFIGURATION chinese ADD MAPPING FOR n,v,a,i,e,l,j WITH simple;
