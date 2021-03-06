package main

import (
	"fmt"
	"github.com/PeterYangs/request/v2"
)

func main() {

	client := request.NewClient().Proxy("http://127.0.0.1:4780")

	content, err := client.R().GetToContent("https://www.google.com.hk/")

	if err != nil {

		fmt.Println(err)

		return
	}

	fmt.Println(content.ToString())

}
