package errs

import (
	"context"
	"git.fogcdn.top/ecf/common-lib/middleware/i18n"
	"strconv"
)

// NewI18nError 给前端返回的时候用这个，支持错误信息的翻译
// 注意code必须在对应项目的i18n目录下配置中英文，否则会报错
func NewI18nError(ctx context.Context, code int) error {
	msg := i18n.Translate(ctx, strconv.Itoa(code))
	s := &Error{
		Code: int32(code),
		Msg:  msg,
	}
	return s
}
