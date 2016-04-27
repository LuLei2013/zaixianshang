package vo

type ReturnMsg struct {
	Errno int   `json:"errno"`
	ErrMsg string `json:"errMsg"`
}

func (msg *ReturnMsg) SetErrno (no int) {
	msg.Errno = no
}

func (msg *ReturnMsg) GetErrno (no int) (res int) {
	return msg.Errno
}

func (msg *ReturnMsg) SetErrMsg (errMsg string) {
	msg.ErrMsg = errMsg
}

func (msg *ReturnMsg) GetErrMsg (no int) (res string) {
	return msg.ErrMsg
}