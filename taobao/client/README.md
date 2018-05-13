# Example 🌰
```go
package main

import (
	"fmt"

	tbclient "github.com/Undesoft/pkg/taobao/client"
)

func init() {
	tbclient.AppKey = ""
	tbclient.AppSecret = ""
	tbclient.Router = "http://gw.api.taobao.com/router/rest"
}

func main() {
	res, err := tbclient.Execute("taobao.tbk.item.get", map[string]string{
		"fields": "num_iid,title,pict_url,small_images,reserve_price,zk_final_price,user_type,provcity,item_url,seller_id,volume,nick",
		"q":      "女装",
		"cat":    "16,18",
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("商品数量:", res.Get("tbk_item_get_response").Get("total_results").MustInt())
	var imtes []interface{}
	imtes, _ = res.Get("tbk_item_get_response").Get("results").Get("n_tbk_item").Array()
	for _, v := range imtes {
		fmt.Println("======")
		item := v.(map[string]interface{})
		fmt.Println("商品名称:", item["title"])
		fmt.Println("商品价格:", item["reserve_price"])
		fmt.Println("商品链接:", item["item_url"])
	}
}

```