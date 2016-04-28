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
		message.SetErrno(3)
		message.SetErrMsg("秒杀失败")
		return message
	}
	if value, err := redis.RedisPoolOne.Get(vo.Product1_Query_Name); err == nil {
		if count, _ := strconv.Atoi(string(value)); count >= vo.Product1_Max_Num {
			vo.Flag = false
			message.SetErrno(3)
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
						message.SetErrno(1)
						message.SetErrMsg("参数错误")
						return message
					}
				}
				default:
					continue
				}
			}
			if entry.GetUserid() == "" || entry.GetProductid() == "" {
				message.SetErrno(1)
				message.SetErrMsg("参数错误")
				return message
			}
			if str, err := json.Marshal(entry); err == nil {
				err := redis.RedisPoolOne.RPush("list", string(str))
				if err != nil {
					fmt.Println("errMsg:", err)
					message.SetErrno(2)
					message.SetErrMsg("Redis链接错误")
					return message
				}
				message.SetErrno(0)
				message.SetErrMsg("秒杀中，请稍后查询")
			} else {
				message.SetErrno(1)
				message.SetErrMsg("参数错误")
			}
		}
	} else {
		message.SetErrno(2)
		message.SetErrMsg("Redis链接错误")
	}
	return message
}