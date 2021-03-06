package main

import (
	"fmt"
	"github.com/PeterYangs/request/v2"
)

func main() {

	client := request.NewClient()

	r, err := client.R().Query(map[string]interface{}{
		"list":  "123",
		"array": []string{"123", "456"},
		"form": map[string]interface{}{
			"name":  "yy",
			"email": "904801074@qq.com",
			"age":   10,
			"order": []string{"1", "2", "3"},
		},
	}).Get("http://list.com/demo/get.php")

	if err != nil {

		fmt.Println(err)

		return
	}

	//响应头部
	//fmt.Println(r.Header())

	content, err := r.Content()

	if err != nil {

		fmt.Println(err)

		return
	}

	fmt.Println(content.ToString())

}
