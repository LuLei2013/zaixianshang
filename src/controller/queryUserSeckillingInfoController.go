package controller

import (
	"fmt"
	"net/http"
	"redis"
	"strings"
    	"vo"
    	"encoding/json"
	"strconv"
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
	retMessage := &vo.ResultPersonMsg{0, "", ""}
	value, _ := redis.RedisPoolOne.Get(vo.Product1_Query_Name)
	if count, _ := strconv.Atoi(string(value)); count < vo.Product1_Max_Num {
		retMessage.SetErrno(2)
		retMessage.SetStatus("3")
		retMessage.SetGoodsId("还在秒杀中，请稍后查询")
		return
	}
	goodsid, _ := redis.RedisPoolOne.Get(userid)
	if goodsid != ""{   // 秒杀成功
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
