package main

import (
	"fmt"
	"github.com/PeterYangs/request/v2"
)

func main() {

	c := request.NewClient()

	content, err := c.R().GetToContent("https://www.baidu.com")

	if err != nil {

		fmt.Println(err)

		return
	}

	fmt.Println(content.ToString())

}
