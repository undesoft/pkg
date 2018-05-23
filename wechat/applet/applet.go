package applet

import (
	"errors"

	"github.com/Undesoft/pkg/net/http"
)

var (
	AppId     = ""
	AppSecret = ""
)

func Execute(url string, args map[string]string) (result string, err error) {
	err = checkConfig()
	if err != nil {
		return
	}
	args["appid"] = AppId
	args["secret"] = AppSecret
	return http.Get(url, args)
}

func checkConfig() error {
	if AppId == "" {
		return errors.New("AppKey 不能为空")
	}
	if AppSecret == "" {
		return errors.New("AppSecret 不能为空")
	}
	return nil
}
