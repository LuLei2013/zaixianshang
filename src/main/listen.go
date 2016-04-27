package main
import (
	"fmt"
	"GoRedisService"
	"strconv"
	"encoding/json"
	"vo"
)

func main () {
	defer GoRedisService.CloseRedis();
	fmt.Println("listen the world!")
	GoRedisService.OpenRedis("localhost","6379")

	for  {
		//超过商品总数,停止监听
		//todo 添加配置文件,尝试多商品
		if count,_ := strconv.Atoi(GoRedisService.HGetValue("Product1")); count>= 100 {
			fmt.Println("listen finish!")
			break
		}

		popValue := GoRedisService.RPopValue("list")
		if popValue != "" {
			fmt.Print("popValue:"+popValue)
			//fmt.Print("\n")
			var  qe vo.QueueEntry
			var userId string
			if err := json.Unmarshal([]byte(popValue), &qe); err == nil {
				fmt.Println(qe.Userid)
				userId = qe.Userid
			}else {
				continue
			}

			if checkValid(userId) {
				incCount("Product1",userId)//写入redis
			}
			//fmt.Print(popValue + "\n")
		}
	}
}

func checkValid(userId string ) bool{
	if GoRedisService.HGetValue(userId) != "" {
		return false
	}
	return true
}

func incCount(productId string,userId string){
	//商品数量计数器加一
	goodsid := 1;
	if GoRedisService.HGetValue(productId) == "" {
		GoRedisService.HSetValue(productId,"1")
		//fmt.Print("if here")
	}else {
		//fmt.Print("else here")
		tmp :=GoRedisService.HGetValue(productId)
		//fmt.Print("  "+tmp)
		//fmt.Print("\n")
		b,_ := strconv.Atoi(tmp);
		goodsid = b+1
		GoRedisService.HSetValue(productId,strconv.Itoa(goodsid))
	}

	GoRedisService.HSetValue(userId,strconv.Itoa(goodsid))
	fmt.Print(userId+"****"+strconv.Itoa(goodsid))
	fmt.Print("\n")
	newValue := userId+"*"+strconv.Itoa(goodsid)
	fmt.Print(newValue)
	fmt.Print("\n")
	GoRedisService.LPushValue("sum"+productId,newValue)
}
