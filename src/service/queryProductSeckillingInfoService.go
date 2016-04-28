package service

import (
	"dao"
	"net/http"
	"strings"
	"vo"
)

//根据商品id查询成功秒杀的所有用户id和用户购买商品的具体编号
func ServiceQueryProductSeckillingInfo(req *http.Request) *vo.ResultProductMsg {
	productid := ""
	for key, value := range req.Form {
		if key == "productid" {
			productid = strings.Join(value, "")
		}
	}
	returnMsg := &vo.ResultProductMsg{0, nil}
	if productid != vo.Product1_Query_Name {
		returnMsg.SetErrno(1)
		returnMsg.SetList(nil)
		return returnMsg
	}

	productInfo, _ := redis.RedisPoolOne.LRange(vo.Product1_Query_String)

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
