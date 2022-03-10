package main

import (
	"fmt"
	"github.com/PeterYangs/request/v2"
)

func main() {

	client := request.NewClient()

	content, err := client.R().ReTry(1).GetToContent("http://list.com/asdjk.php")

	if err != nil {

		fmt.Println(err)

		return
	}

	fmt.Println(content.ToString())

}
