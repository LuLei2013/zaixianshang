package main
import (
	"fmt"
	"GoRedisService"
	"strconv"
)

func main () {
	defer GoRedisService.CloseRedis();
	fmt.Println("listen the world!")
	GoRedisService.OpenRedis("localhost","6379")


	for  {

		if GoRedisService.HGetValue("Product1") == "100" {
			break
		}

		popValue := GoRedisService.RPopValue("list")
		if popValue != "" {
			fmt.Print("popValue:"+popValue)
			//fmt.Print("\n")
			if GoRedisService.HGetValue("Product1") == "" {
				GoRedisService.HSetValue("Product1","1")
				//fmt.Print("if here")
			}else {
				//fmt.Print("else here")
				tmp :=GoRedisService.HGetValue("Product1")
				fmt.Print("  "+tmp)
				fmt.Print("\n")
				b,_ := strconv.Atoi(tmp);
				GoRedisService.HSetValue("Product1",strconv.Itoa(b+1))
			}
			//fmt.Print(popValue + "\n")
		}

	}


}