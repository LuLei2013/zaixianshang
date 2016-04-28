package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"service"
)

var counter = 0

func Seckilling(resp http.ResponseWriter, req *http.Request) {
	req.ParseForm() //解析参数，默认是不会解析的
	counter += 1
	fmt.Println(counter)
	for k, v := range req.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	message := service.ServiceSeckilling(req)
	if jsonstr, jsonerr := json.Marshal(message); jsonerr == nil {
		fmt.Fprintf(resp, string(jsonstr))
	}
	//test dw


	//value, _ := redis.RedisPoolOne.Get(vo.Product1_Query_Name)
	//if count, _ := strconv.Atoi(string(value)); count >= vo.Product1_Max_Num {
	//	message.SetErrno(1)
	//	message.SetErrMsg("秒杀失败")
	//	if jsonstr, jsonerr := json.Marshal(message); jsonerr == nil {
	//		fmt.Fprintf(resp, string(jsonstr))
	//	}
	//} else {
	//	entry := &vo.QueueEntry{"", "", ""}
	//	for key, value := range req.Form {
	//		switch key {
	//		case "userid":
	//			entry.SetUserid(strings.Join(value, ""))
	//		case "productid":
	//			entry.SetProductid(strings.Join(value, ""))
	//		default:
	//			fmt.Fprintf(resp, "出错了")
	//			return
	//		}
	//	}
	//
	//	if str, err := json.Marshal(entry); err == nil {
	//		err := redis.RedisPoolOne.RPush("list", string(str))
	//		if err != nil {
	//			fmt.Println("errMsg:", err)
	//			return
	//		}
	//		message.SetErrno(0)
	//		message.SetErrMsg("秒杀中")
	//		if jsonstr, jsonerr := json.Marshal(message); jsonerr == nil {
	//			fmt.Fprintf(resp, string(jsonstr))
	//		}
	//	} else {
	//		message.SetErrno(1)
	//		message.SetErrMsg("参数出错")
	//		if jsonstr, jsonerr := json.Marshal(message); jsonerr == nil {
	//			fmt.Fprintf(resp, string(jsonstr))
	//		}
	//	}
	//	//return message
	//}
	//fmt.Fprintf(w, "Hello World!") //这个写入到w的是输出到客户端的
}
