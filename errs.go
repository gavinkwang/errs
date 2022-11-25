// Package errs 错误码类型，里面包含errcode errmsg，多语言通用
package errs

import (
	"fmt"
)

// Success 成功提示字符串
const (
	Success    = "success"
	RetUnknown = -1
)

// Error 错误码结构 包含 错误码类型 错误码 错误信息
type Error struct {
	Code int32
	Msg  string
	Desc string
}

// Error 实现error接口，返回error描述
func (e *Error) Error() string {
	if e == nil {
		return Success
	}

	return fmt.Sprintf("code:%d, msg:%s", e.Code, e.Msg)
}

// GetCode 获取Error的Code
func (e *Error) GetCode() int {
	return int(e.Code)
}

// GetMsg 获取Error的msg
func (e *Error) GetMsg() string {
	return e.Msg
}

// Equal 比较两个Error是否相等(只比较Code,去除服务编号)
func (e *Error) Equal(new *Error) bool {
	return e.Code&0x11111 == new.Code&0x11111
}

// New 创建一个error
func New(code int, msg string) error {
	err := &Error{
		Code: int32(code + GetInstance().GetSid()*100000),
		Msg:  msg,
	}

	return err
}

// Newf 创建一个error，msg支持格式化字符串
func Newf(code int, format string, params ...interface{}) error {
	msg := fmt.Sprintf(format, params...)
	return New(code, msg)
}

// Code 通过error获取error code
func Code(e error) int {
	if e == nil {
		return 0
	}
	err, ok := e.(*Error)
	if !ok {
		return RetUnknown
	}
	if err == (*Error)(nil) {
		return 0
	}
	return int(err.Code)
}

// Msg 通过error获取error msg
func Msg(e error) string {
	if e == nil {
		return Success
	}
	err, ok := e.(*Error)
	if !ok {
		return e.Error()
	}
	if err == (*Error)(nil) {
		return Success
	}
	return err.Msg
}
