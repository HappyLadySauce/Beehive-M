# 错误码约定

## 码段分配（千位 = 服务）

| 千位 | 服务     | 说明     |
|-----|----------|----------|
| 1xxx | 通用服务 | 跨服务通用错误码 |
| 2xxx | User 服务 | 用户相关 |
| 3xxx | 下一服务 | 按项目扩展（如 Auth、Friend 等） |
| …   | …        | 以此类推 |

## 后三位分类（同一服务内）

| 区间   | 含义         | 说明                     |
|--------|--------------|--------------------------|
| x0xx / 0 | 成功         | 如 1000、2000 表示成功  |
| x1xx   | 请求/参数    | 参数缺失、非法、校验失败 |
| x2xx   | 认证/授权    | 未登录、无权限、Token 失效 |
| x3xx   | 业务逻辑     | 用户不存在、已存在等业务规则 |
| x5xx   | 服务端/未知  | 内部错误、未知异常       |

## 示例

- **通用**：`1000` 成功，`1101` 参数错误，`1501` 内部错误  
- **User**：`2100` 成功，`2101` 用户不存在，`2201` 未授权，`2501` 用户服务内部错误  

## 使用方式

Logic 层统一使用 `github.com/zeromicro/x/errors` 与 `pkg/code` 常量：

```go
import (
    "github.com/HappyLadySauce/Beehive-M/pkg/code"
    "github.com/zeromicro/x/errors"
)

return nil, errors.New(code.CodeInvalidParam, "user_id is required")
```

Gateway 对外统一响应格式为 `{"code": <码>, "msg": "<描述>", "data": ...}`，由 `xhttp.JsonBaseResponseCtx` 与 `httpx.SetErrorHandler` 保证。
