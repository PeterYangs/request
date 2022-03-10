package main

import (
	"fmt"
	"github.com/PeterYangs/request/v2"
)

func main() {

	client := request.NewClient()

	rsp, err := client.R().Params(map[string]interface{}{"list": []string{"1", "2", "3"}}).Put("http://www.cms.com/put")

	if err != nil {

		fmt.Println(err)

		return

	}

	c, err := rsp.Content()

	if err != nil {

		fmt.Println(err)

		return

	}

	fmt.Println(c.ToString())

}
