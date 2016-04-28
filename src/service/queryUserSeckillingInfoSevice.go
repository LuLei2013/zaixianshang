package service

import (
	"redis"
	"strings"
	"vo"
	"net/http"
	"strconv"
)
//
// errrno : 1
//         status :2 参数错误
// errno  :0
//         status :0  秒杀还未开始 ,商品没有被卖了
//         status :1  秒杀成功    , 成功秒杀到，redis中查询到goodsId
//         status :2  秒杀失败    , 没有秒杀到，redis中未查询到goodsId，且商品已经被卖完
//         status :3  在秒杀中    , 没有秒杀到，redis中未查询到goodsId，但是商品还未卖完
//
//
func QueryUserSeckillingInfo(req *http.Request)  *vo.ResultPersonMsg {
	req.ParseForm() //解析参数，默认是不会解析的
	productid := ""
	userid := ""
	retMessage := &vo.ResultPersonMsg{0, "", ""}
	for key, value := range req.Form {
		if key == "userid" {
			userid = strings.Join(value, "")
		} else if key == "productid" {
			productid = strings.Join(value, "")
		}
	}
	if userid == "" || productid == "" {
		retMessage.SetErrno(1)
		retMessage.SetStatus("2")
		retMessage.SetGoodsId("秒杀失败，秒杀参数错误")
		return retMessage
	}
	value, _ := redis.RedisPoolOne.Get(vo.Product1_Query_Name)
	countGoodsSold, _ := strconv.Atoi(string(value)); 
	if 0 == countGoodsSold{
		retMessage.SetErrno(0)
		retMessage.SetStatus("0")
		retMessage.SetGoodsId("秒杀未开始...")
		return retMessage;
	}

	goodsid, _ := redis.RedisPoolOne.Get(userid)
	if goodsid != ""{   // 秒杀成功
		retMessage.SetErrno(0)
		retMessage.SetStatus("1")
		retMessage.SetGoodsId(goodsid)
	} else { // 秒杀失败
		if  countGoodsSold< vo.Product1_Max_Num {
			retMessage.SetErrno(0)
			retMessage.SetStatus("3")
			retMessage.SetGoodsId("在秒杀中，请稍后查询...")
			return retMessage;
		}
		retMessage.SetErrno(0)
		retMessage.SetStatus("2")
		retMessage.SetGoodsId("秒杀失败,未秒杀到商品")
	}

	return retMessage
}
