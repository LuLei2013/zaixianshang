package controller
import (
	"net/http"
	"fmt"
	"strings"
	"GoRedisService"
)

func Seckilling(resp http.ResponseWriter, req *http.Request) {
	req.ParseForm()  //解析参数，默认是不会解析的
	//fmt.Println(r.Form)  //这些信息是输出到服务器端的打印信息
	//fmt.Println("path", r.URL.Path)
	//fmt.Println("scheme", r.URL.Scheme)
	//fmt.Println(r.Form["url_long"])


	//test dw
	defer GoRedisService.CloseRedis();
	GoRedisService.OpenRedis("localhost","6379")
	GoRedisService.LPushValue("list","1")

	for k, v := range req.Form {
		fmt.Println("key:", k)
		fmt.Fprintf(resp, k)
		fmt.Println("val:", strings.Join(v, ""))
		fmt.Fprintf(resp, ":")
		fmt.Fprintf(resp, strings.Join(v, ""))
		fmt.Fprintf(resp, "\n")
	}
	//fmt.Fprintf(w, "Hello World!") //这个写入到w的是输出到客户端的
}