package controller

import (
	"net/http"
	"fmt"
	"strings"
    	"vo"
    	"encoding/json"
	"github.com/garyburd/redigo/redis"
)

func QueryUserSeckillingInfo(resp http.ResponseWriter, req *http.Request) {
	req.ParseForm()  //解析参数，默认是不会解析的
	productid := ""
	userid := ""
	for key, value := range req.Form {
		if key == "userid"{
			userid = strings.Join(value, "")
		}else if key == "productid" {
			productid = strings.Join(value, "")
		}
	}
	if userid == "" || productid == ""{
		fmt.Fprintf(resp, "error")
		return 
	}
	IPAndPort := "192.168.2.165:6379"
	conn, _ := redis.Dial("tcp", IPAndPort)
	if conn == nil {
		fmt.Printf("redis连接失败\n")
		return
	}
	defer conn.Close()
	goodsid, _ :=	redis.String(conn.Do("GET", userid))
	retMessage := &vo.ResultPersonMsg{0, "", ""}
	if goodsid != ""{   // 秒杀成功
		retMessage.SetErrno(0)
		retMessage.SetStatus("1")
		retMessage.SetGoodsId(goodsid)
	}else{              // 秒杀失败
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
