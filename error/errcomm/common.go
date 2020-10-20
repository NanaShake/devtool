package errcomm

import "github.com/NanaShake/devtool/error/errs"

// 业务错误
var (
	ErrMoneyNotEnough = errs.UserError(1, "游戏币不足") // -3001
	ErrGoldNotEnough  = errs.UserError(2, "银币不足")  // -3002
)

// 客户端请求错误
var (
	ErrClientInvalid = errs.ClientError(1, "操作失败，条件不满足") // -1001
	ErrClientParam   = errs.ClientError(2, "请求参数错误")     // -1002
	ErrClientPerm    = errs.ClientError(3, "操作失败，权限不足")  // -1003

)

// 服务器错误
var (
	ErrServerConfig       = errs.SystemError(1, "服务器配置错误，请稍后再试") // -5001
	ErrServerLogic        = errs.SystemError(2, "服务器异常，请稍后再试")   // -5002
	ErrServerDb           = errs.SystemError(3, "服务器异常，请稍后再试")   // -5003
	ErrServerDataNotFound = errs.SystemError(4, "数据不存在")         // -5004

)
