package vo

type ResultPersonMsg struct {
	Errno int   `json:"errno"`
	Status string `json:"status"`
	GoodsId string `json:"goodsid"`
}

func (msg *ResultPersonMsg) SetErrno (no int) {
	msg.Errno = no
}

func (msg *ResultPersonMsg) GetErrno () (res int) {
	return msg.Errno
}

func (msg *ResultPersonMsg) SetStatus (status string) {
	msg.Status = status
}

func (msg *ResultPersonMsg) GetStatus () (status string) {
	return msg.Status
}

func (msg *ResultPersonMsg) SetGoodsId (goodsId string) {
	msg.GoodsId = goodsId
}

func (msg *ResultPersonMsg) GetGoodsId () (goodsId string) {
	return msg.GoodsId
}

