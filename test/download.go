package main

import (
	"fmt"
	"github.com/PeterYangs/request"
)

func main() {

	url := "http://list.com/demo/demo.zip"

	client := request.NewClient()

	err := client.R().Download(url, "123.zip")

	if err != nil {

		fmt.Println(err)
	}
}
