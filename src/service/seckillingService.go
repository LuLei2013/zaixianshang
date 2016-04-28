package service

import (
	"vo"
	"redis"
	"strconv"
	"net/http"
	"strings"
	"encoding/json"
	"fmt"
)

func ServiceSeckilling(req *http.Request) *vo.ReturnMsg {
	message := &vo.ReturnMsg{0, ""} //返回消息
	if !vo.Flag {
		message.SetErrno(0)
		message.SetErrMsg("秒杀失败")
		return message
	}
	if value, err := redis.RedisPoolOne.Get(vo.Product1_Query_Name); err == nil {
		if count, _ := strconv.Atoi(string(value)); count >= vo.Product1_Max_Num {
			vo.Flag = false
			message.SetErrno(0)
			message.SetErrMsg("秒杀失败")
		} else {
			req.ParseForm() //解析参数，默认是不会解析的
			entry := &vo.QueueEntry{"", "", ""}
			for key, value := range req.Form {
				strValue := strings.Join(value, "")
				switch key {
				case "userid":
					entry.SetUserid(strValue)
				case "productid": {
					if (strValue == vo.Product1_Query_Name || strValue == vo.Product2_Query_Name || strValue == vo.Product3_Query_Name) {
						entry.SetProductid(strValue)
					} else {
						fmt.Println("errMsg:", "productid不存在")
						panic("productid不存在")
					}
				}
				default:
					continue
				}
			}
			if entry.GetUserid() == "" || entry.GetProductid() == "" {
				fmt.Println("errMsg:", "参数错误")
				panic("参数错误")
			}
			if str, err := json.Marshal(entry); err == nil {
				redisError := redis.RedisPoolOne.RPush("list", string(str))
				if redisError != nil {
					fmt.Println("errMsg:", redisError)
					panic(redisError.Error())
				}
				message.SetErrno(0)
				message.SetErrMsg("秒杀中，请稍后查询")
			} else {
				fmt.Println("errMsg:", err)
				panic(err.Error())
			}
		}
	} else {
		fmt.Println("errMsg:", err)
		panic(err.Error())
	}
	return message
}