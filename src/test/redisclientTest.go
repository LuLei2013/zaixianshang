package main

import (
	"fmt"
	"redis"
	"time"
)

func main() {
	//defer GoRedisService.CloseRedis();
	fmt.Println("Hello World")
	//GoRedisService.OpenRedis("localhost","6379")
	//GoRedisService.LPushValue("list","1 ")
	//GoRedisService.LPushValue("list","2 ")
	//GoRedisService.LPushValue("list","3 ")


	//popValue := GoRedisService.RPopValue("list")
	//for  popValue != ""{
	//	fmt.Print(popValue+"\n")
	//	popValue = GoRedisService.RPopValue("list")
	//}
	Rediss := redis.NewRedisClient("127.0.0.1:6379", 100, 100, 100, 100, 100, 100)
	Rediss.Set("123", "2123")
	fmt.Println(Rediss.Get("123"))
	fmt.Println(Rediss.Exists("123"))
	time.Sleep(10 * time.Second)

}
