package test

import (
	"testing"
)

func Test_Division_1(t *testing.T) {
	if i, e := Division(6, 2); i != 3 || e != nil { //try a unit test on function
		t.Error("用例1:除法测试功能测试失败!") // 如果不是如预期的那么就报错
	} else {
		t.Log("用例1:除法测试功能测试通过!") //记录一些你期望记录的信息
	}
}

func Test_Division_2(t *testing.T) {
	if _, e := Division(6, 0); e == nil { //try a unit test on function
		t.Error("用例2:除数不为0,异常测试失败!") // 如果不是如预期的那么就报错
	} else {
		t.Log("用例2:除数为0情况,校验通过!", e) //记录一些你期望记录的信息
	}
}
