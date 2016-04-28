package test

import (
	"testing"
	"net/http"
	"net/url"
)

func Test_QueryUserSeckillingInfo_1(t *testing.T){
	u, _ := url.Parse("http://192.168.2.114:9090/zaixianshang/queryUserSeckillingInfo")
	q := u.Query()
	q.Set("userid", "")
	q.Set("productid", "")
	u.RawQuery = q.Encode()
	t.Log("****aladlgsadg****")
	if res, err := http.Get(u.String()); res != nil || err != nil {
		t.Error("用例1:参数为空校验测试失败!")
	} else {
		t.Log("用例1:参数为空校验测试通过!")
	}
	t.Log("teglsdlgjadgjasdlg;al")
}
