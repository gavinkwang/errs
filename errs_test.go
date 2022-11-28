package errs_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.ctyuncdn.cn/wangzhk/errs"
)

// go test -v -coverprofile=cover.out
// go tool cover -func=cover.out

func TestErrs(t *testing.T) {

	var err *errs.Error
	str := err.Error()
	assert.Contains(t, str, "success")

	e := errs.New(111, "inner fail")
	assert.NotNil(t, e)

	assert.Equal(t, 111, errs.Code(e))
	assert.Equal(t, "inner fail", errs.Msg(e))

	str = e.Error()
	assert.Contains(t, str, "code")
	assert.Contains(t, str, "msg")

	e = errs.Newf(111, "inner fail")
	assert.NotNil(t, e)

	assert.Equal(t, 111, errs.Code(e))
	assert.Equal(t, "inner fail", errs.Msg(e))

	str = err.Error()

	assert.Equal(t, 0, errs.Code(nil))
	assert.Equal(t, "success", errs.Msg(nil))

	assert.Equal(t, 0, errs.Code((*errs.Error)(nil)))
	assert.Equal(t, "success", errs.Msg((*errs.Error)(nil)))

	eo := errors.New("unknown error")
	assert.Equal(t, errs.RetUnknown, errs.Code(eo))
	assert.Equal(t, "unknown error", errs.Msg(eo))
}

func TestService(t *testing.T) {
	errs.Init(errs.EsportsSrv)
	assert.Equal(t, errs.EsportsSrv*100000+errs.AUserErr, errs.Code(errs.ErrInfo[errs.AUserErr]))
}

func TestEqual(t *testing.T) {
	errs.Init(errs.EsportsSrv)

	e := errs.New(10000, "test")

	assert.Equal(t, true, e.Equal(errs.ErrInfo[errs.AUserErr]))
}
