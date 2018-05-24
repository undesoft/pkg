package applet

import (
	"fmt"
	"testing"
)

func init() {
	AppId = "123"
	AppSecret = "abc"
}

func TestExecute(t *testing.T) {
	// 小程序登录
	url := "https://api.weixin.qq.com/sns/jscode2session"
	args := map[string]string{
		"js_code":    "登录时获取的 code",
		"grant_type": "authorization_code",
	}
	json, err := Execute(url, args)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(json)
}
