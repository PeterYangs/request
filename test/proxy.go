package main

import (
	"fmt"
	"github.com/PeterYangs/request"
)

func main() {

	client := request.NewClient().Proxy("http://127.0.0.1:4780")

	for i := 0; i < 10; i++ {

		content, err := client.R().GetToContent("https://www.peterdemo.net/api/admin/admin/info")

		if err != nil {

			fmt.Println(err)

			return
		}

		fmt.Println(content.ToString())

	}

}
