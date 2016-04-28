package main

import (
	"fmt"
	"github.com/Terry-Mao/goconf"
)

func main() {
	conf := goconf.New()
	if err := conf.Parse("../redis/redisConfig.txt"); err != nil {
		panic(err)
	}
	ip := conf.String("redis","ip")

	fmt.Println(ip)

}