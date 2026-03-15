package code

// 用户相关错误码常量
// 业务响应/错误码常量
// 0=成功，1xx=请求/参数，2xx=认证/授权，3xx=业务逻辑，5xx=服务端
const (
	// 用户不存在
	CodeUserNotFound = 2100 + iota
)

const (
	// 用户未授权
	CodeUserNotAuthorized = 2200 + iota
)

const(
	// 用户已存在
	CodeUserAlreadyExists = 2300 + iota
)
