package client

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	simplejson "github.com/bitly/go-simplejson"
)

var (
	AppKey    string = ""
	AppSecret string = ""
	Router    string = ""
)

// 执行API接口
func Execute(method string, params map[string]string) (res *simplejson.Json, err error) {
	err = checkConfig()
	if err != nil {
		return
	}
	params["method"] = method
	var req *http.Request
	req, err = http.NewRequest("POST", Router, strings.NewReader(getRequestData(params)))
	if err != nil {
		return
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")
	httpclient := &http.Client{}
	httpclient.Timeout = time.Second * 3
	var response *http.Response
	response, err = httpclient.Do(req)
	if err != nil {
		return
	}

	if response.StatusCode != 200 {
		err = fmt.Errorf("请求错误:%d", response.StatusCode)
		return
	}
	var body []byte
	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	res, err = simplejson.NewJson(body)
	if err != nil {
		return
	}
	if responseError, ok := res.CheckGet("error_response"); ok {
		errorBytes, _ := responseError.Encode()
		err = errors.New("执行错误:" + string(errorBytes))
	}
	return
}

// 检查配置
func checkConfig() error {
	if AppKey == "" {
		return errors.New("AppKey 不能为空")
	}
	if AppSecret == "" {
		return errors.New("AppSecret 不能为空")
	}
	if Router == "" {
		return errors.New("Router 不能为空")
	}
	return nil
}

// 获取请求数据
func getRequestData(params map[string]string) string {
	// 公共参数
	args := url.Values{}
	hh, _ := time.ParseDuration("8h")
	loc := time.Now().UTC().Add(hh)
	args.Add("timestamp", strconv.FormatInt(loc.Unix(), 10))
	args.Add("format", "json")
	args.Add("app_key", AppKey)
	args.Add("v", "2.0")
	args.Add("sign_method", "md5")
	args.Add("partner_id", "Undesoft")
	// 请求参数
	for key, val := range params {
		args.Set(key, val)
	}
	// 设置签名
	args.Add("sign", getSign(args))
	return args.Encode()
}

// 获取签名
func getSign(args url.Values) string {
	// 获取Key
	keys := []string{}
	for k := range args {
		keys = append(keys, k)
	}
	// 排序asc
	sort.Strings(keys)
	// 把所有参数名和参数值串在一起
	query := AppSecret
	for _, k := range keys {
		query += k + args.Get(k)
	}
	query += AppSecret
	// 使用MD5加密
	signBytes := md5.Sum([]byte(query))
	// 把二进制转化为大写的十六进制
	return strings.ToUpper(hex.EncodeToString(signBytes[:]))
}
