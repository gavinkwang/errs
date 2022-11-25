package errs

type serviceID struct {
	Sid int
}

type ServiceID interface {
	Init(id int)
	GetSid() int
}

var singleton ServiceID = &serviceID{}

// Init 初始化服务ID
func Init(sid int) {
	singleton.Init(sid)
}

func GetInstance() ServiceID {
	return singleton
}

func Reset(s ServiceID) {
	singleton = s
}

// Init 初始化服务ID
func (s *serviceID) Init(sid int) {
	s.Sid = sid
	for k, v := range ErrInfo {
		ErrInfo[k] = New(k, Msg(v))
	}
}

// GetSid 获取服务ID
func (s *serviceID) GetSid() int {
	return s.Sid
}
