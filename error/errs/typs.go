package errs

// 错误级别
type Level int32

const (
	Success     Level = 0     // 成功
	ServerLevel       = -5000 // 服务级别错误
	ClientLevel       = -1000 // 客户端请求错误
	UserLevel         = -3000 // 自定义级别错误
)
