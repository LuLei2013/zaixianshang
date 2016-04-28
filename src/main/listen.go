package main

import (
	"GoRedisService"
	"encoding/json"
	"fmt"
	"strconv"
	"vo"
)

func main() {
	defer GoRedisService.CloseRedis()
	defer func(){ // 必须要先声明defer，否则不能捕获到panic异常
		if err:=recover();err!=nil{
			fmt.Println(err)
		}
	}()

	fmt.Println("listen the world!")
	//vo.Init()//初始化配置文件
	GoRedisService.OpenRedis(vo.Ip, vo.Port)
	for {
		//超过商品总数,停止监听
		//todo 添加配置文件,尝试多商品
		if count,err := strconv.Atoi(GoRedisService.HGetValue(vo.Product1_Query_Name)); count>= vo.Product1_Max_Num {
			if err != nil {
				panic(err)
			}
			fmt.Println("listen finish!")
			break
		}
		popValue := GoRedisService.RPopValue("list") //消费队列
		if popValue != "" {
			fmt.Print("popValue:" + popValue)
			var qe vo.QueueEntry
			var userId string
			//解压数据
			if err := json.Unmarshal([]byte(popValue), &qe); err == nil {
				fmt.Println(qe.Userid)
				userId = qe.Userid
			} else {
				continue//解压失败则淘汰该用户,这里不抛出异常
			}
			if checkValid(userId) {
				incCount(vo.Product1_Query_Name,userId)//写入redis
			}
			//fmt.Print(popValue + "\n")
		}
	}
}

/**
 *重入判断
 *@param userId string
 *@return bool
 */
func checkValid(userId string) bool {
	if GoRedisService.HGetValue(userId) != "" {
		return false
	}
	return true
}

//商品计数器加1,并存储,异常均在异步方法中打印
func incCount(productId string, userId string) {
	//商品数量计数器加一
	goodsid := 1
	tmp := GoRedisService.HGetValue(productId)
	//fmt.Print("  "+tmp)
	//fmt.Print("\n")
	b, _ := strconv.Atoi(tmp)
	goodsid = b + 1
	GoRedisService.HSetValue(productId, strconv.Itoa(goodsid))


	//存储用户和其购买的商品关系
	GoRedisService.HSetValue(userId, strconv.Itoa(goodsid))
	fmt.Print(userId + "****" + strconv.Itoa(goodsid))
	fmt.Print("\n")
	newValue := userId + "*" + strconv.Itoa(goodsid)
	fmt.Print(newValue)
	fmt.Print("\n")
	//向全部购买信息中push数据
	GoRedisService.LPushValue(vo.Product1_Query_String,newValue)
}
