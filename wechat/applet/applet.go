package applet

import (
	"errors"

	"github.com/Undesoft/pkg/net/http"
	simplejson "github.com/bitly/go-simplejson"
)

var (
	AppId     = ""
	AppSecret = ""
)

func Execute(url string, args map[string]string) (json *simplejson.Json, err error) {
	err = checkConfig()
	if err != nil {
		return
	}
	args["appid"] = AppId
	args["secret"] = AppSecret

	result, err := http.Get(url, args)
	if err != nil {
		return
	}
	json, err = simplejson.NewJson([]byte(result))
	if err != nil {
		return
	}
	if errmsg, ok := json.CheckGet("errcode"); ok {
		bytes, _ := errmsg.Encode()
		err = errors.New(string(bytes))
		return
	}
	return
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
