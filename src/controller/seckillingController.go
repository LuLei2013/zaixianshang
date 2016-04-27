package controller
import (
	"net/http"
	"fmt"
	"strconv"
	"vo"
	"encoding/json"
	"strings"
	"github.com/garyburd/redigo/redis"
)

var counter = 0
func Seckilling(resp http.ResponseWriter, req *http.Request){
	req.ParseForm()  //解析参数，默认是不会解析的
	//fmt.Println(r.Form)  //这些信息是输出到服务器端的打印信息
	//fmt.Println("path", r.URL.Path)
	//fmt.Println("scheme", r.URL.Scheme)
	//fmt.Println(r.Form["url_long"])
	counter += 1
	fmt.Println(counter)
	for k, v := range req.Form {
		fmt.Println("key:", k)
		fmt.Fprintf(resp, k)
		fmt.Println("val:", strings.Join(v, ""))
		fmt.Fprintf(resp, ":")
		fmt.Fprintf(resp, strings.Join(v, ""))
		fmt.Fprintf(resp, "\n")
	}
	message := &vo.ReturnMsg{0, ""}
	//test dw

	IPAndPort := "192.168.2.165:6379"
	conn, _ := redis.Dial("tcp", IPAndPort)
	if conn == nil {
		fmt.Printf("redis连接失败\n")
	}
	defer conn.Close()
	//GoRedisService.OpenRedis("192.168.2.165","6379")
	//defer GoRedisService.CloseRedis();
	value, _ := redis.String(conn.Do("GET", "Product1"))
	if count, _ := strconv.Atoi(string(value)); count >= 100 {
		message.SetErrno(1)
		message.SetErrMsg("秒杀失败")
		if jsonstr, jsonerr := json.Marshal(message); jsonerr == nil {
			fmt.Fprintf(resp, string(jsonstr))
		}
		//fmt.Fprintf(resp, "失败")
		//return message
	} else {
		entry := &vo.QueueEntry{"", "", ""}
		for key, value := range req.Form {
			switch key {
			case "userid" :
				entry.SetUserid(strings.Join(value, ""))
			case "productid" :
				entry.SetProductid(strings.Join(value, ""))
			default:
				fmt.Fprintf(resp, "出错了")
				return
			}
		}

		if str, err := json.Marshal(entry); err == nil {
			_, err := conn.Do("lpush", "list", string(str))
			if err != nil {
				fmt.Println("errMsg:", err)
			}
			//GoRedisService.LPushValue("list", string(str))
			message.SetErrno(0)
			message.SetErrMsg("秒杀中")
			if jsonstr, jsonerr := json.Marshal(message); jsonerr == nil {
				fmt.Fprintf(resp, string(jsonstr))
			}
		} else {
			message.SetErrno(1)
			message.SetErrMsg("参数出错")
			if jsonstr, jsonerr := json.Marshal(message); jsonerr == nil {
				fmt.Fprintf(resp, string(jsonstr))
			}
		}
		//return message
	}
	//fmt.Fprintf(w, "Hello World!") //这个写入到w的是输出到客户端的
}