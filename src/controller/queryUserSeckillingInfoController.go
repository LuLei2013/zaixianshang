package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"redis"
	"strings"
	"vo"
)

func QueryUserSeckillingInfo(resp http.ResponseWriter, req *http.Request) {
	req.ParseForm() //解析参数，默认是不会解析的
	productid := ""
	userid := ""
	for key, value := range req.Form {
		if key == "userid" {
			userid = strings.Join(value, "")
		} else if key == "productid" {
			productid = strings.Join(value, "")
		}
	}
	if userid == "" || productid == "" {
		fmt.Fprintf(resp, "error")
		return
	}
	goodsid, _ := redis.RedisPoolOne.Get(userid)
	retMessage := &vo.ResultPersonMsg{0, "", ""}
	if goodsid != "" { // 秒杀成功
		retMessage.SetErrno(0)
		retMessage.SetStatus("1")
		retMessage.SetGoodsId(goodsid)
	} else { // 秒杀失败
		retMessage.SetErrno(1)
		retMessage.SetStatus("2")
		retMessage.SetGoodsId("sorry,you miss the goods")
	}
	// 返回请求结果
	if jsonstr, jsonerr := json.Marshal(retMessage); jsonerr == nil {
		fmt.Fprintf(resp, string(jsonstr))
		return
	}
	fmt.Fprintf(resp, "error")
}
