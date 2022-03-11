package main

import (
	"fmt"
	"github.com/PeterYangs/request/v2"
)

func main() {

	c := request.NewClient()

	rsp, err := c.R().Get("https://www.baidu.com")

	if err != nil {

		fmt.Println(err)

		return
	}

	content, err := rsp.Content()

	if err != nil {

		fmt.Println(err)

		return
	}

	//fmt.Println(content.ToString())
	fmt.Println(content.Header().Get("Connection"))
	fmt.Println(content.StatusCode())
	fmt.Println(content.Proto())

}
