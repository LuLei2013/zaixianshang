package vo

type ResultProductMsg struct {
	Errno int	`json:"errno"`
	List []KillEntry	`json:"list"`
}

func (rpm *ResultProductMsg) SetErrno (errno int){
	rpm.Errno = errno
}

func (rpm *ResultProductMsg) SetList (list []KillEntry){
	rpm.List = list
}

func (rpm *ResultProductMsg) GetErrno () (res int){
	return rpm.Errno
}

func (rpm *ResultProductMsg) GetList () (list []KillEntry){
	return rpm.List
}