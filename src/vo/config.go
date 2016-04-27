package vo

import (
	"util"
	"os"
	"strconv"
)

var (
	Ip string
	Port string
	Product1_Max_Num int
	Product1_Query_Name string
	Product1_Query_String string
)

func Init() {
	myConfig := new(util.Config)
	path,_ := os.Getwd()
	//fmt.Println(path)
	path += "\\conf\\properties"
	//fmt.Println(path)
	myConfig.InitConfig(path)
	//fmt.Println(myConfig.Read("common", "ip"))
	//fmt.Printf("%v", myConfig.Mymap)
	Ip = myConfig.Read("common", "ip")
	Port = myConfig.Read("common", "port")
	Product1_Max_Num, _ = strconv.Atoi(myConfig.Read("product_1", "max_num"))
	Product1_Query_Name = myConfig.Read("product_1", "query_name")
	Product1_Query_String = myConfig.Read("product_1", "total_query_name")
}