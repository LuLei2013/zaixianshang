package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"service"
	"strings"
)

func QueryProductSeckillingInfo(resp http.ResponseWriter, req *http.Request) {
	req.ParseForm() //解析参数，默认是不会解析的
	productid := ""

	for key, value := range req.Form {
		if key == "productid" {
			productid = strings.Join(value, "")
		}
	}
	if productid == "" {
		fmt.Fprintf(resp, "error")
		return
	}

	info := service.ServiceQueryProductSeckillingInfo()
	if jsonstr, jsonerr := json.Marshal(info); jsonerr == nil {
		fmt.Fprintf(resp, string(jsonstr))
		return
	}
	fmt.Fprintf(resp, "error")
}
