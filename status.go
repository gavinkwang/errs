package errs

import (
	"context"
	"git.fogcdn.top/ecf/common-lib/middleware/i18n"
	"strconv"
)

// I18nTrans 给前端返回的时候用这个，支持错误信息的翻译
// 注意code必须在对应项目的i18n目录下配置中英文，否则会panic。中英文配置的code不带服务标识
func I18nTrans(ctx context.Context, err *Error) *Error {
	code := err.GetCode()
	msg := err.GetMsg()
	detail := err.GetDetail()

	if code < 100000 {
		// code 小于100000时，返回需带上服务号
		code = GetInstance().GetSid()*100000 + code
		msg = i18n.Translate(ctx, strconv.Itoa(code))
	} else {
		msg = i18n.Translate(ctx, strconv.Itoa(code%100000))
	}

	return &Error{
		Code:   int32(code),
		Msg:    msg,
		Detail: detail,
	}
}
