package code

// 业务响应/错误码常量（通用服务 1xxx）
// 后三位：0=成功，1xx=请求/参数，2xx=认证/授权，3xx=业务逻辑，5xx=服务端
const (
	// 成功
	CodeSuccess = 200
	// 请求/参数错误
	CodeInvalidParam = 300
	// 服务端错误
	CodeInternal = 500
)