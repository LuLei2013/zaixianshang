package test

import (
	"service"
	"testing"
	"vo"
)

func Test_ServiceSeckilling_1(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Log("用例1测试通过！")
		} else {
			t.Error("用例1测试未通过！")
		}
	}()
	service.ServiceSeckilling("", "")
}

func Test_ServiceSeckilling_2(t *testing.T) {
	vo.Flag = false
	message := service.ServiceSeckilling("123", "111")
	if message.GetErrMsg() == "秒杀失败" && message.GetErrno() == 0 {
		t.Log("用例2测试通过！")
	} else {
		t.Error("用例2测试未通过！")
	}
}

func Test_ServiceSeckilling_3(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Log("用例3测试通过！")
		} else {
			t.Error("用例3测试未通过！")
		}
	}()
	service.ServiceSeckilling("123", "123")
}

func Test_ServiceSeckilling_4(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Log("用例4测试未通过！")
		} else {
			t.Error("用例4测试通过！")
		}
	}()
	service.ServiceSeckilling("123", "111")
}