package main

import (
	"fmt"
	"GoRedisService"
	"strconv"
	"time"
)

func main () {
	//defer GoRedisService.CloseRedis();
	fmt.Println("Hello World")
	//GoRedisService.OpenRedis("localhost","6379")
	//GoRedisService.LPushValue("list","1 ")
	//GoRedisService.LPushValue("list","2 ")
	//GoRedisService.LPushValue("list","3 ")

	for index := 100; index > 0; index-- {
		go func(i int) {
			defer GoRedisService.CloseRedis();
			GoRedisService.OpenRedis("localhost","6379")
			GoRedisService.LPushValue("list",strconv.Itoa(i))
		}(index)
		//GoRedisService.LPushValue("list",strconv.Itoa(index))
	}


	//popValue := GoRedisService.RPopValue("list")
	//for  popValue != ""{
	//	fmt.Print(popValue+"\n")
	//	popValue = GoRedisService.RPopValue("list")
	//}

	time.Sleep(10*time.Second)

}