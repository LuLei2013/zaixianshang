package service

import (
	"redis"
	"strings"
	"vo"
)

func ServiceQueryProductSeckillingInfo(productid string) *vo.ResultProductMsg {
	productInfo, _ := redis.RedisPoolOne.LRange(vo.Product1_Query_String)
	returnMsg := &vo.ResultProductMsg{0, nil}
	if productInfo == nil {
		returnMsg.SetErrno(1)
		returnMsg.SetList(nil)
		return returnMsg
	}
	goodsList := []vo.KillEntry{}
	for _, entry := range productInfo {
		tmp := strings.Split(entry, "*")
		userid := tmp[0]
		goodsid := tmp[1]
		killEntry := vo.KillEntry{userid, goodsid}
		goodsList = append(goodsList, killEntry)
	}
	returnMsg.SetErrno(0)
	returnMsg.SetList(goodsList)
	return returnMsg

}
