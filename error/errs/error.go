package errs

import "fmt"

type Error interface {
	Code() int32
	Error() string
}

func newError(level Level, mark int32, msg string) Error {
	if 0 < mark && mark < 1000 {
		code := int32(level) - mark
		return &errorStatus{
			code: code,
			msg:  msg,
		}
	}
	panic(fmt.Sprintf("mark must be less than 1000 and greater than 0"))
}

func Succes() Error {
	return &errorStatus{
		code: 0,
		msg:  "成功",
	}
}

// 系统错误
func SystemError(mark int32, msg string) Error {
	return newError(ServerLevel, 500+mark, msg)
}

// 客户端错误
func ClientError(mark int32, msg string) Error {
	return newError(ClientLevel, 500+mark, msg)
}

// 业务错误
func UserError(mark int32, msg string) Error {
	return newError(UserLevel, 500+mark, msg)
}

/////////////////////////////////////////
type errorStatus struct {
	code int32
	msg  string
}

func (p errorStatus) Code() int32 {
	return p.code
}

func (p errorStatus) Error() string {
	return p.msg
}
