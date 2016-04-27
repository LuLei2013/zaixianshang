package controller
import (
	"net/http"
	"fmt"
	"GoRedisService"
	"strconv"
	"vo"
	"encoding/json"
)

func Seckilling(resp http.ResponseWriter, req *http.Request){
	req.ParseForm()  //解析参数，默认是不会解析的
	//fmt.Println(r.Form)  //这些信息是输出到服务器端的打印信息
	//fmt.Println("path", r.URL.Path)
	//fmt.Println("scheme", r.URL.Scheme)
	//fmt.Println(r.Form["url_long"])
	//for k, v := range req.Form {
	//	fmt.Println("key:", k)
	//	fmt.Fprintf(resp, k)
	//	fmt.Println("val:", strings.Join(v, ""))
	//	fmt.Fprintf(resp, ":")
	//	fmt.Fprintf(resp, strings.Join(v, ""))
	//	fmt.Fprintf(resp, "\n")
	//}
	message := &vo.ReturnMsg{0, ""}
	//test dw
	defer GoRedisService.CloseRedis();
	GoRedisService.OpenRedis("192.168.2.165","6379")
	if count, _ := strconv.Atoi(GoRedisService.HGetValue("Product1")); count >= 100 {
		message.SetErrno(1)
		message.SetErrMsg("秒杀失败")
		fmt.Fprintf(resp, "失败")
		//return message
	} else {
		entry := &vo.QueueEntry{"1", "2", "3"}
		if str, err := json.Marshal(entry); err == nil {
			GoRedisService.LPushValue("list", string(str))
			message.SetErrno(0)
			message.SetErrMsg("秒杀中")
			fmt.Fprintf(resp, "秒杀中")
		} else {
			message.SetErrno(1)
			message.SetErrMsg("参数出错")
			fmt.Fprintf(resp, "出错了！")
		}
		//return message
	}
	//fmt.Fprintf(w, "Hello World!") //这个写入到w的是输出到客户端的
}