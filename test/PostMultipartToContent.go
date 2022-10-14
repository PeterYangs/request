package main

import (
	"fmt"
	"github.com/PeterYangs/request/v2"
	"os"
)

func main() {

	f, e := os.OpenFile("README.md", os.O_CREATE|os.O_RDWR, 0644)

	if e != nil {

		return
	}

	client := request.NewClient()

	ct, err := client.R().Params(map[string]interface{}{"name": "yy", "file": f}).PostMultipartToContent("http://www.yycms.com/post_test")

	if err != nil {

		fmt.Println(err)

		return
	}

	fmt.Println(ct.ToString())

}
