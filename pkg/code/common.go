package code

// 业务响应/错误码常量
// 0=成功，1xx=请求/参数，2xx=认证/授权，3xx=业务逻辑，5xx=服务端
const (
	// 成功
	CodeSuccess = 1000 + iota
)

const (
	// 请求/参数错误
	CodeInvalidParam = 1100 + iota
)

const (
	// 服务端错误
	CodeInternal = 1500 + iota
)