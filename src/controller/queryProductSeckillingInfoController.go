package controller

import (
	"net/http"
	"fmt"
	"strings"
	"github.com/garyburd/redigo/redis"
	"vo"
	"encoding/json"
)

func QueryProductSeckillingInfo(resp http.ResponseWriter, req *http.Request) {
	req.ParseForm()  //解析参数，默认是不会解析的
	productid := ""
	//for k, v := range req.Form {
	//	fmt.Println("key:", k)
	//	fmt.Fprintf(resp, k)
	//	fmt.Println("val:", strings.Join(v, ""))
	//	fmt.Fprintf(resp, ":")
	//	fmt.Fprintf(resp, strings.Join(v, ""))
	//	fmt.Fprintf(resp, "\n")
	//}
	for key, value := range req.Form {
		if key == "productid" {
			productid = strings.Join(value, "")
		}
	}
	if productid == "" {
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
	productInfo, _ := redis.Strings(conn.Do("LRANGE", "sumProduct1", 0, -1))
	returnMsg := &vo.ResultProductMsg{0, nil}
	if productInfo == nil {
		returnMsg.SetErrno(1)
		returnMsg.SetList(nil)
	} else {
		goodsList := []vo.KillEntry{}
		for _, entry := range productInfo {
			tmp := strings.Split(entry, "*")
			userid := tmp[0]
			goodsid := tmp[1]
			killEntry := vo.KillEntry{userid, goodsid}
			goodsList = append(goodsList, killEntry)
			fmt.Println(goodsList)
		}
		returnMsg.SetErrno(0)
		returnMsg.SetList(goodsList)
	}
	if jsonstr, jsonerr := json.Marshal(returnMsg); jsonerr == nil {
		fmt.Fprintf(resp, string(jsonstr))
		return
	}
	fmt.Fprintf(resp, "error")
}