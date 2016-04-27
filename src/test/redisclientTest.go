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
	redisone.HSet("1232", "1", "one")
	redisone.HSet("1232", "2", "two")
	fmt.Println(redisone.HGet("1232","1"))
	//fmt.Println(redisone.HGetall("123"))


}
