package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"service"
)

//返回商品秒杀结果，首先解析请求参数，再调用ServiceQueryProductSeckillingInfo获得商品的秒杀结果
func QueryProductSeckillingInfo(resp http.ResponseWriter, req *http.Request) {
	req.ParseForm()

	info := service.ServiceQueryProductSeckillingInfo(req)
	if jsonstr, jsonerr := json.Marshal(info); jsonerr == nil {
		fmt.Fprintf(resp, string(jsonstr))
		return
	}
	fmt.Fprintf(resp, "error")
}
