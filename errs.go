// Package errs 错误码类型，里面包含errcode errmsg，多语言通用
package errs

import (
	"fmt"
	pbcommon "gitlab.ctyuncdn.cn/wangzhk/errs/protos/goout/status"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strconv"
)

// Success 成功提示字符串
const (
	Success    = "success"
	RetUnknown = 9999999
)

// Error 错误码结构 包含 错误码类型 错误码 错误信息
type Error struct {
	Code   int32  // 错误码
	Msg    string // 错误信息(展示给用户看, 默认为空，在i18n配置里定义)
	Detail string // 错误信息(展示给开发人员看)
}

// Error 实现error接口，返回error描述
func (e *Error) Error() string {
	if e == nil {
		return Success
	}

	return fmt.Sprintf("code:%d, detail:%s", e.GetCode(), e.GetDetail())
}

// GetCode 获取Error的Code
func (e *Error) GetCode() int {
	return int(e.Code)
}

// GetMsg 获取Error的msg
func (e *Error) GetMsg() string {
	return e.Msg
}

// GetDetail 获取Error的detail
func (e *Error) GetDetail() string {
	return e.Detail
}

// Equal 比较两个Error是否相等(只比较Code,去除服务编号)
func (e *Error) Equal(new *Error) bool {
	return e.Code&0x11111 == new.Code&0x11111
}

// GRPCStatus 兼容GRPC错误码
func (e *Error) GRPCStatus() *status.Status {
	msg := e.Msg
	if msg == "" {
		msg = "请求出错，请稍后重试"
	}
	rs := &pbcommon.NewResponseStatus{
		Code:    strconv.Itoa(int(e.Code)),
		Message: msg,
		Detail:  e.Detail,
	}
	st, _ := status.New(codes.Code(e.Code), msg).WithDetails(rs)
	return st
}

// AppendErr 在原有Error基础上追加err信息，常用于拼接外部的错误信息
func (e *Error) AppendErr(err error) error {
	detail := fmt.Sprintf("%s(%s)", e.Detail, err.Error())
	return &Error{
		Code:   e.Code,
		Detail: detail,
	}
}

// New 创建一个error
func New(code int, detail string) *Error {
	return &Error{
		Code:   int32(GetInstance().GetSid()*100000 + code),
		Detail: detail,
	}
}

// Newf 创建一个error，msg支持格式化字符串
func Newf(code int, format string, params ...interface{}) *Error {
	detail := fmt.Sprintf(format, params...)
	return New(code, detail)
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

// Detail 通过error获取error Detail
func Detail(e error) string {
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
	return err.Detail
}
