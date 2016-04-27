package main

import (
	"fmt"
	"redis"

)

func main() {

	redisone := redis.GetRedisInstance()
	//redis的数据输入输出
	redisone.Set("123", "2123")
	fmt.Println(redisone.Get("123"))
	fmt.Println(redisone.Exists("123"))
	//redis list数据输入输出
	redisone.RPush("1243", "1")
	redisone.RPush("1243", "2")
	fmt.Println(redisone.LPop("1243"))

	//fmt.Println(redisone.HGetall("123"))


}
