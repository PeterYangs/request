package main

import (
	"fmt"
	"github.com/PeterYangs/request"
)

func main() {

	client := request.NewClient()

	c, err := client.R().Params(map[string]interface{}{"list": []string{"1", "2", "3"}}).DeleteToContent("http://www.cms.com/put")

	if err != nil {

		fmt.Println(err)

		return

	}

	fmt.Println(c.ToString())

}
