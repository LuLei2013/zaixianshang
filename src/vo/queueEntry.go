package vo

type QueueEntry struct {
	Userid string	`json:"userid"`
	Productid string	`json:"productid"`
	Time string	`json:"time"`
}

func (queueEntry *QueueEntry) GetUserid () (str string) {
	return queueEntry.Userid
}
func (queueEntry *QueueEntry) GetProductid () (str string) {
	return queueEntry.Productid
}
func (queueEntry *QueueEntry) GetTime () (str string) {
	return queueEntry.Time
}

func (queueEntry *QueueEntry) SetUserid (userid string) {
	queueEntry.Userid = userid
}
func (queueEntry *QueueEntry) SetProductid (productid string)  {
	queueEntry.Productid = productid
}
func (queueEntry *QueueEntry) SetTime (time string) {
	queueEntry.Time = time
}