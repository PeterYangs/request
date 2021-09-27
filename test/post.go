package main

import (
	"fmt"
	"github.com/PeterYangs/request"
)

func main() {

	client := request.NewClient()

	r, err := client.R().Params(map[string]interface{}{
		"list":  "123",
		"array": []string{"123", "456"},
		"form": map[string]interface{}{
			"name":  "yy",
			"email": "904801074@qq.com",
			"age":   10,
			"order": []string{"1", "2", "3"},
		},
	}).Post("http://list.com/demo/post.php")

	if err != nil {

		fmt.Println(err)

		return
	}

	content, err := r.Body().Content()

	if err != nil {

		fmt.Println(err)

		return
	}

	fmt.Println(content.ToString())

}
