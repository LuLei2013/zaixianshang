package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"service"
	"vo"
	"strings"
)

var counter = 0

func Seckilling(resp http.ResponseWriter, req *http.Request) {
	req.ParseForm() //解析参数，默认是不会解析的
	counter += 1
	fmt.Println(counter)
	message := &vo.ReturnMsg{0, ""}
	defer func() {//异常处理
		if err := recover(); err != nil {
			message.SetErrno(1)
			message.SetErrMsg("出错了！")
		}
		if jsonstr, jsonerr := json.Marshal(message); jsonerr == nil {
			fmt.Fprintf(resp, string(jsonstr))
		}
	}()
	for k, v := range req.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	message = service.ServiceSeckilling(req)
	if jsonstr, jsonerr := json.Marshal(message); jsonerr == nil {
		fmt.Fprintf(resp, string(jsonstr))
	}
}
