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

	Product2_Max_Num int
	Product2_Query_Name string
	Product2_Query_String string

	Product3_Max_Num int
	Product3_Query_Name string
	Product3_Query_String string
)

func Init() {
	myConfig := new(util.Config)
	path,_ := os.Getwd()
	//fmt.Println(path)
	path += "/conf/properties"
	//fmt.Println(path)
	myConfig.InitConfig(path)
	//fmt.Println(myConfig.Read("common", "ip"))
	//fmt.Printf("%v", myConfig.Mymap)
	Ip = myConfig.Read("common", "ip")
	Port = myConfig.Read("common", "port")
	Product1_Max_Num, _ = strconv.Atoi(myConfig.Read("product_1", "max_num"))
	Product1_Query_Name = myConfig.Read("product_1", "query_name")
	Product1_Query_String = myConfig.Read("product_1", "total_query_name")

	Product2_Max_Num, _ = strconv.Atoi(myConfig.Read("product_2", "max_num"))
	Product2_Query_Name = myConfig.Read("product_2", "query_name")
	Product2_Query_String = myConfig.Read("product_2", "total_query_name")

	Product3_Max_Num, _ = strconv.Atoi(myConfig.Read("product_3", "max_num"))
	Product3_Query_Name = myConfig.Read("product_3", "query_name")
	Product3_Query_String = myConfig.Read("product_3", "total_query_name")
}