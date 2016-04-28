package service

import (
	"dao"
	"net/http"
	"strings"
	"vo"
	"fmt"
)

//根据商品id查询成功秒杀的所有用户id和用户购买商品的具体编号
func ServiceQueryProductSeckillingInfo(req *http.Request) *vo.ResultProductMsg {
	req.ParseForm()
	productid := ""
	for key, value := range req.Form {
		if key == "productid" {
			productid = strings.Join(value, "")
		}
	}
	returnMsg := &vo.ResultProductMsg{0, nil}
	//fmt.Print(productid)
	if (productid != vo.Product1_Query_Name) {
		fmt.Println("errMsg:", "productid不存在")
		panic("productid不存在")
	}

	productInfo, _ := redis.RedisPoolOne.LRange(vo.Product1_Query_String)

	if productInfo == nil {
		fmt.Println("errMsg:", "无法查询到结果")
		panic("无法查询到结果")
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
