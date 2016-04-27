package vo

type KillEntry struct {
	Userid string	`json:"userid"`
	Productid string	`json:"productid"`
}

func (entry *KillEntry) SetUserid (userid string){
	entry.Userid = userid
}

func (entry *KillEntry) SetProductid (productid string){
	entry.Productid = productid
}

func (entry *KillEntry) GetProductid () (res string){
	return entry.Productid
}

func (entry *KillEntry) GetUserid () (res string){
	return entry.Userid
}