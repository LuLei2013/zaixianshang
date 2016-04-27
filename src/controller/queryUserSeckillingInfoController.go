package controller

import (
	"net/http"
	"fmt"
	"strings"
)

func QueryUserSeckillingInfo(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()  //解析参数，默认是不会解析的
	//fmt.Println(r.Form)  //这些信息是输出到服务器端的打印信息
	//fmt.Println("path", r.URL.Path)
	//fmt.Println("scheme", r.URL.Scheme)
	//fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Fprintf(w, k)
		fmt.Println("val:", strings.Join(v, ""))
		fmt.Fprintf(w, ":")
		fmt.Fprintf(w, strings.Join(v, ""))
		fmt.Fprintf(w, "\n")
	}
	//fmt.Fprintf(w, "Hello World!") //这个写入到w的是输出到客户端的
}