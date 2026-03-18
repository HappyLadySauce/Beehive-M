package code

// 业务响应/错误码常量（通用服务 1xxx）
// 后三位：0=成功，1xx=请求/参数，2xx=认证/授权，3xx=业务逻辑，5xx=服务端
const (
	// 成功
	CodeSuccess = 200
	// 请求/参数错误
	CodeInvalidParam = 300
	// 解析失败
	CodeUnmarshalFailed = 301
	// 数据库错误
	CodeDBQueryFailed = 302
	// 缓存错误
	CodeCacheSetFailed = 303
	// 缓存获取失败
	CodeCacheGetFailed = 304
	// 服务端错误
	CodeInternal = 500
)

// 分布式锁相关错误
const (
	// 分布式锁获取失败
	CodeDistributedLockGetFailed = 305
	// 分布式锁释放失败
	CodeDistributedLockUnlockFailed = 306
	// 分布式锁已存在
	CodeDistributedLockAlreadyExists = 307
	// 分布式锁获取超时
	CodeDistributedLockGetTimeout = 308
)