package main

import (
	"fmt"
	"github.com/PeterYangs/request"
)

func main() {

	client := request.NewClient()

	content, err := client.R().Params(map[string]interface{}{
		"list":  "123",
		"array": []string{"123", "456"},
		"form": map[string]interface{}{
			"name":  "yy",
			"email": "904801074@qq.com",
			"age":   10,
			"order": []string{"1", "2", "3"},
		},
	}).GetToContent("http://list.com/demo/get.php")

	if err != nil {

		fmt.Println(err)

		return

	}

	fmt.Println(content.ToString())

	//fmt.Println(content.ToJsonMap())

}
