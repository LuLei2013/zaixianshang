package vo

type KillEntry struct {
	Userid string	`json:"userid"`
	Goodsid string	`json:"goodsid"`
}

func (entry *KillEntry) SetUserid (userid string){
	entry.Userid = userid
}

func (entry *KillEntry) SetGoodsid (goodsid string){
	entry.Goodsid = goodsid
}

func (entry *KillEntry) GetGoodsid () (res string){
	return entry.Goodsid
}

func (entry *KillEntry) GetUserid () (res string){
	return entry.Userid
}