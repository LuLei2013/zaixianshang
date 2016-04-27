package main
import (
	"fmt"
	"redis"
	"strconv"
)

func main () {

	fmt.Println("listen the world!")



	for  {

		if RedisPoolOne.Get("Product1") == "100" {
			break
		}

		popValue := RedisPoolOne.LPop("list111")
		if popValue != "" {
			fmt.Print("popValue:"+popValue)
			//fmt.Print("\n")
			if RedisPoolOne.Get("Product1") == "" {
				RedisPoolOne.Set("Product1","1")
				//fmt.Print("if here")
			}else {
				//fmt.Print("else here")
				tmp :=RedisPoolOne.Get("Product1")
				fmt.Print("  "+tmp)
				fmt.Print("\n")
				b,_ := strconv.Atoi(tmp);
				RedisPoolOne.Set("Product1",strconv.Itoa(b+1))
			}
			//fmt.Print(popValue + "\n")
		}

	}


}