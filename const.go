package errs

const (
	EsportsSrv = iota + 10
	EsportsWebsocketSrv
	EsportsCronSrv
)

// Err 框架错误值
const (
	OK       = 0     // 成功
	AUserErr = 10000 // A类错误，用户类错误

	BParamErr = 20000 // B类错误，参数错误

	CCallErr      = 30000 // C类错误，调用第三方服务错误
	CCallDBErr    = 31000 // 数据库类错误
	CCallRedisErr = 32000 //redis类错误

	DInnerErr = 40000 // D类错误，内部错误
)

var ErrInfo = map[int]*Error{
	AUserErr: New(AUserErr, "用户类错误"),

	BParamErr: New(BParamErr, "参数错误"),

	CCallErr:      New(CCallErr, "调用第三方服务错误"),
	CCallDBErr:    New(CCallDBErr, "数据库错误"),
	CCallRedisErr: New(CCallRedisErr, "Redis错误"),

	DInnerErr: New(DInnerErr, "内部错误"),
}
