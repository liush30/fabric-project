package errorcode

type ServiceError struct {
	errmsg string
	code   int32
}

const (
	ErrCodeCommon      = 1000 // 操作失败
	ErrCodeInvalidArgs = 1001 // 无效参数
	ErrCodeRequestBusy = 1002 // 请求频繁
	ErrCodeServerBusy  = 1010 // 服务繁忙

	ErrCodeNotAuthenticated      = 1011 // 未授权
	ErrCodeAuthenticationExpired = 1012 // 授权过期
	ErrCodeWrongSig              = 1013 // 签名错误

	ErrCodeTextFormat   = 1020 // 无效格式或内容
	ErrCodeRenameNum    = 1021 // 改名次数不足
	ErrCodeNameRepeated = 1022 // 名称已存在

	ErrCodeUserNotFound = 1035 //用户不存在
)

var (
	ErrServerBusy    = NewError(1010, "server busy")
	ErrOperateFailed = NewError(ErrCodeCommon)
)

func NewError(code int32, msg ...string) *ServiceError {
	if len(msg) > 0 {
		return &ServiceError{
			code:   code,
			errmsg: msg[0],
		}
	}

	return &ServiceError{
		code: code,
	}
}

func (s *ServiceError) Error() string {
	return s.errmsg
}

func (s *ServiceError) Code() int32 {
	return s.code
}
