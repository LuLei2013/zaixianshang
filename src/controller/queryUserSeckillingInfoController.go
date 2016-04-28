package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"service"
)

func QueryUserSeckillingInfo(resp http.ResponseWriter, req *http.Request) {
	
    req.ParseForm() //解析参数，默认是不会解析的
	counter += 1
	fmt.Println("Server 已收到的请求总数 : ",counter)
	for k, v := range req.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	message := service.QueryUserSeckillingInfo(req)
	if jsonstr, jsonerr := json.Marshal(message); jsonerr == nil {
		fmt.Fprintf(resp, string(jsonstr))
	}

}