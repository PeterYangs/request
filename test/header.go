package main

import (
	"fmt"
	"github.com/PeterYangs/request"
)

func main() {

	//全局设置header
	client := request.NewClient().Header(map[string]string{"test": "demo"})

	//请求设置header
	content, err := client.R().Header(map[string]string{"test": "demo2"}).GetToContent("http://list.com/demo/header.php")

	if err != nil {

		fmt.Println(err)

		return
	}

	fmt.Println(content.ToString())

}
