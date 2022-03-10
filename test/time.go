package main

import (
	"fmt"
	"github.com/PeterYangs/request/v2"
)

func main() {

	client := request.NewClient()

	content, err := client.R().GetToContent("https://www.baidu.com")

	if err != nil {

		fmt.Println(err)

		return
	}

	fmt.Println(content.ToString())

	fmt.Println(content.Time().Seconds())

}
