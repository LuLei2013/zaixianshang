package vo

import (
	"util"
	"os"
	"strconv"
)

var (
	Ip string
	Port string
	Product_Pre string

	Product1_Max_Num int
	Product1_Query_Name string
	Product1_Query_String string

	Flag bool

	Product2_Max_Num int
	Product2_Query_Name string
	Product2_Query_String string

	Product3_Max_Num int
	Product3_Query_Name string
	Product3_Query_String string
)

func init() {
	myConfig := new(util.Config)
	path,_ := os.Getwd()
	path += "/conf/properties"
	myConfig.InitConfig(path)

	Ip = myConfig.Read("common", "ip")
	Port = myConfig.Read("common", "port")
	Product_Pre = myConfig.Read("common", "query_prefix")

	Product1_Max_Num, _ = strconv.Atoi(myConfig.Read("product_1", "max_num"))
	Product1_Query_Name = myConfig.Read("product_1", "query_name")
	Product1_Query_String = myConfig.Read("product_1", "total_query_name")

	Product2_Max_Num, _ = strconv.Atoi(myConfig.Read("product_2", "max_num"))
	Product2_Query_Name = myConfig.Read("product_2", "query_name")
	Product2_Query_String = myConfig.Read("product_2", "total_query_name")

	Product3_Max_Num, _ = strconv.Atoi(myConfig.Read("product_3", "max_num"))
	Product3_Query_Name = myConfig.Read("product_3", "query_name")
	Product3_Query_String = myConfig.Read("product_3", "total_query_name")

	Flag = true
}